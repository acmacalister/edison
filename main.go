package main

import (
	"log"

	"time"

	"github.com/tarm/serial"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyMFD1", Baud: 57600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.Write([]byte("4"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Started 9 degress of freedom connection.")

	go func() {
		for {
			buf := make([]byte, 128)
			n, err := s.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			f := parseFreedomData(buf)
			//fmt.Printf("%s", string(buf[:n]))
		}
	}()

	gbot := gobot.NewGobot()

	e := edison.NewEdisonAdaptor("edison")
	blueLed := gpio.NewLedDriver(e, "led", "3")
	redLed := gpio.NewLedDriver(e, "led", "5")
	greenLed := gpio.NewLedDriver(e, "led", "6")
	yellowLed := gpio.NewLedDriver(e, "led", "9")
	var level byte = 0

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			blueLed.Brightness(level)
			redLed.Brightness(level)
			greenLed.Brightness(level)
			yellowLed.Brightness(level)
			level++
			if level >= 168 {
				level = 0
			}
			//led.Toggle()
		})
	}

	robot := gobot.NewRobot("quad",
		[]gobot.Connection{e},
		[]gobot.Device{blueLed, redLed, greenLed, yellowLed},
		work,
	)

	gbot.AddRobot(robot)
	gbot.Start()
}
