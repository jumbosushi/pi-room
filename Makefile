.PHONY: bin
setup:
	if [ ! -d "/home/pi/anavi-examples" ] ; then
		git clone https://github.com/AnaviTechnology/anavi-examples.git
	fi

bin: main.go sensors/*.go
	go build -o bin/pi-room main.go
