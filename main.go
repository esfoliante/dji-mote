package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	// ? This will simply initiate the tello driver and
	// ? connect our program to port 8888 of the drone
	drone := tello.NewDriver("8888")

	// ? This will be our main function and will take our drone
	// ? to the skies
	work := func() {
		// ? We will start with a simple take-off
		drone.TakeOff()

		// ? And after 5 seconds we land our drone
		// ? (which, btw, didn't happen in my first try
		// ? and ended up going to a wall)
		// ? Maybe you should test drive in a safe environment
		gobot.After(5*time.Second, func() {
			drone.Land()
		})
	}

	// ? Then we just setup our robot (in this case a drone)
	// ? and we pass down our work function and drone driver
	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	// ? Finishing off, we simply start our drone
	robot.Start()
}
