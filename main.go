package main

import (
	"github.com/jumbosushi/pi-room/sensors"

	"context"
	"log"
	"time"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// RoomData is a struct for storing room data
type RoomData struct {
	temp     float64
	humidity float64
	pressure float64
	light    float64
}

func recordData(roomData *RoomData) {
	ctx := context.Background()
	// TODO: Change the path to your key file
	opt := option.WithCredentialsFile("/home/pi/pi-room-key.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	collection := client.Collection("conditions")

	_, _, err = collection.Add(ctx, map[string]interface{}{
		"createdAt":   time.Now(),
		"temperature": roomData.temp,
		"humidity":    roomData.humidity,
		"pressure":    roomData.pressure,
		"light":       roomData.light,
	})
	if err != nil {
		log.Fatalf("Failed adding document: %v", err)
	}
}

func main() {
	log.Println("Starting pi-room ...")

	temp, _ := sensors.GetTemp()
	log.Printf("Temperature:  %f C\n", temp)

	humidity, _ := sensors.GetHumidity()
	log.Printf("Humidity:     %f %%rh\n", humidity)

	light, _ := sensors.GetLight()
	log.Printf("Light:        %f Lux\n", light)

	pressure, _ := sensors.GetPressure()
	log.Printf("Pressure:     %f %%rh\n", pressure)

	data := RoomData{
		temp:     temp,
		pressure: pressure,
		humidity: humidity,
		light:    light,
	}

	recordData(&data)
	log.Println("Successfully recorded data!")
}
