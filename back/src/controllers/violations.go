package controllers

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"reaktor_birdnest/src/config"
	"reaktor_birdnest/src/interfaces"
	"reaktor_birdnest/src/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var tempList []interfaces.ViolationInfo

func Check(c *fiber.Ctx) error {
	// Get the report
	scanRes, err := http.Get(os.Getenv("DRONES_API"))
	if err != nil {
		log.Fatal("Failed to get response from drone api")
	}

	defer scanRes.Body.Close()

	scanBody, err := ioutil.ReadAll(scanRes.Body)
	if err != nil {
		log.Fatal("Failed to get body from drone api")
	}

	var report interfaces.Report
	scanErr := xml.Unmarshal(scanBody, &report)
	if scanErr != nil {
		log.Fatal("Failed to unmarshal report")
	}

	// Check each drone to get violation pilots
	for _, drone := range report.Capture.Drones {
		distance := utils.CalculateDistance(drone.PositionX, drone.PositionY)

		if distance > -1 {
			pilotRes, err := http.Get(os.Getenv("PILOT_API") + "/" + drone.SerialNumber)
			if err != nil {
				log.Fatal("Failed to get response from pilot api")
			}

			defer pilotRes.Body.Close()

			pilotBody, err := ioutil.ReadAll(pilotRes.Body)
			if err != nil {
				log.Fatal("Failed to get body from pilot api")
			}

			var pilot interfaces.Pilot
			pilotErr := json.Unmarshal(pilotBody, &pilot)
			if pilotErr != nil {
				log.Fatal("Failed to unmarshal pilot information")
			}

			addingPilot := interfaces.ViolationInfo{
				Pilot:           pilot,
				ClosestDistance: distance,
				CaptureTime:     report.Capture.SnapshotTimestamp,
			}

			// Check duplicate pilot before add new to the tempList
			var isAdded bool = false
			for i, oldPilot := range tempList {
				if addingPilot.Pilot.PilotID == oldPilot.Pilot.PilotID && addingPilot.ClosestDistance <= oldPilot.ClosestDistance {
					tempList[i] = addingPilot
					isAdded = true
					break
				}
			}

			if !isAdded {
				tempList = append(tempList, addingPilot)
			}
		}
	}

	return c.SendStatus(http.StatusOK)
}

func UpdatePersistList(c *fiber.Ctx) error {
	// Dont update if the list is empty
	if len(tempList) == 0 {
		return c.Status(http.StatusOK).SendString("The list is update to date")
	}

	// Overwrite persist list in Mongodb
	client := config.Connect()
	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("VIOLATION_LIST"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	defer client.Disconnect(ctx)

	// First delete all old records then add new ones
	_, err := collection.DeleteMany(ctx, bson.D{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to delete all the records")
	}

	// Filter the unique records in temp list
	unique := utils.FilterPilots(tempList)

	docs := make([]interface{}, len(unique))
	for i, rec := range unique {
		docs[i] = bson.D{{Key: "pilot", Value: rec.Pilot}, {Key: "closestDistance", Value: rec.ClosestDistance}, {Key: "captureTime", Value: rec.CaptureTime}}
	}

	_, err = collection.InsertMany(ctx, docs)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to update persist list on MongoDB")
	}

	// Reset temp list
	tempList = nil

	return c.Status(http.StatusOK).SendString("New list have been uploaded")
}

func GetPersistList(c *fiber.Ctx) error {
	client := config.Connect()
	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("VIOLATION_LIST"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	defer client.Disconnect(ctx)

	var results []interface{}
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return c.Status(http.StatusInternalServerError).SendString("Failed to get persist list")
	}

	for cursor.Next(ctx) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Failed to decode value in list")
		}

		results = append(results, doc)
	}

	return c.Status(http.StatusOK).JSON(results)
}

func GetTemp(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(tempList)
}
