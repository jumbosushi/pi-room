package sensors

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// GetTemp returns the temperature
func GetTemp() (temp float64, err error) {
	out, err := exec.Command("/home/pi/anavi-examples/sensors/BMP180/c/BMP180").Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)
	lineTemp := strings.Split(s, "\n")[1]
	tempStr := strings.Fields(lineTemp)[1]

	temp, err = strconv.ParseFloat(tempStr, 32)
	if err != nil {
		log.Fatal(err)
	}

	return
}

// GetHumidity returns the humidity
func GetHumidity() (humidity float64, err error) {
	out, err := exec.Command("/home/pi/anavi-examples/sensors/HTU21D/c/HTU21D").Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)
	lineHum := strings.Split(s, "\n")[2]
	humidityStr := lineHum[:len(lineHum)-3]
	humidity, err = strconv.ParseFloat(humidityStr, 32)
	if err != nil {
		log.Fatal(err)
	}

	return
}

// GetPressure returns the pressure
func GetPressure() (pressure float64, err error) {
	out, err := exec.Command("/home/pi/anavi-examples/sensors/BMP180/c/BMP180").Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)
	linePress := strings.Split(s, "\n")[2]
	pressureStr := strings.Fields(linePress)[1]

	pressure, err = strconv.ParseFloat(pressureStr, 32)
	if err != nil {
		log.Fatal(err)
	}

	return
}

// GetLight returns the light
func GetLight() (light float64, err error) {
	out, err := exec.Command("/home/pi/anavi-examples/sensors/BH1750/c/BH1750").Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)
	lineLight := strings.Split(s, "\n")[1]
	lightStr := strings.Fields(lineLight)[1]

	light, err = strconv.ParseFloat(lightStr, 32)
	if err != nil {
		log.Fatal(err)
	}

	return
}
