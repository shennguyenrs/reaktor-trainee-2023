package utils

import (
	"sort"

	"reaktor_birdnest/src/interfaces"
)

func FilterPilots(pilots []interfaces.ViolationInfo) []interfaces.ViolationInfo {
	var unique []interfaces.ViolationInfo
	seen := make(map[string]bool)

	// Sort the list by the closest distance
	sort.Slice(pilots, func(i, j int) bool {
		return pilots[i].ClosestDistance < pilots[j].ClosestDistance
	})

	// Adding unique values
	for _, rec := range pilots {
		if !seen[rec.Pilot.PilotID] {
			unique = append(unique, rec)
			seen[rec.Pilot.PilotID] = true
		}
	}

	return unique
}
