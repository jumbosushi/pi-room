# pi-room :partly_sunny:

Raspberry Pi setup to track a room's following conditions:
- Temperature
- Humidity
- Air pressure
- Light brightness

Inspiration based on [this blog post](https://life.craftz.dog/entry/record-my-room-consitions) by [@inkdrop_app](https://twitter.com/inkdrop_app)

## Requirements
- Go 1.13+
- Raspberry Pi (I used Model 3)
- [ANAVI Infrared pHAT](https://www.crowdsupply.com/anavi-technology/infrared-phat)
- Cloud Firestore (Firebase) 

## Installation

```
go get github.com/jumbosushi/pi-room
```

## Usage

### 1. Set up Raspberry Pi

I recommend [this tutorial](https://desertbot.io/blog/headless-raspberry-pi-3-bplus-ssh-wifi-setup)!

### 2. Set up Cloud Firestore

Follow Firebase's [official doc](https://desertbot.io/blog/headless-raspberry-pi-3-bplus-ssh-wifi-setup) to create a service account. Down load the account's key file in to your path and update the path near `TODO` in `main.go`

### 3. Download anavi scripts

Clone [anavi-examples](https://github.com/AnaviTechnology/anavi-examples) repo within the pi at the home directory (`/home/pi`). Run `make` in the following directories:
- `sensors/HTU21D/c`
- `sensors/BMP180/c`
- `sensors/BH1750/c`

### 4. Run the script

`pi-room` will now be able to send metrics to Firestore :tada:

I recommend running it as a cron job every couple minutes and monitoring it on [CronHub](https://cronhub.io/)