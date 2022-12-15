#! /bin/sh
while true
do
  echo "Checking violations..."
  curl -X POST 'http://birdnest-api:3001/violations/temp/check'
  sleep 2
done
