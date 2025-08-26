package main

import (
	"machine"
	"time"

	pio "github.com/tinygo-org/pio/rp2-pio"
	"github.com/tinygo-org/pio/rp2-pio/piolib"
)

func main() {
	time.Sleep(time.Second)
	const pin = machine.LED
	sm, _ := pio.PIO0.ClaimStateMachine()
	pulsar, err := piolib.NewPulsar(sm, pin)
	if err != nil {
		panic(err.Error())
	}
	println("start pulsing")

	for {
		// Max period is 0.5ms. PIO state machines can run at minimum of 2kHz.
		for period := time.Microsecond; period < time.Millisecond/3; period *= 2 {
			err = pulsar.SetPeriod(period)
			if err != nil {
				panic(err.Error())
			}
			for i := uint32(10); i < 100; i *= 2 {
				pulsar.TryQueue(i)
				time.Sleep(time.Second / 2)
			}
		}
	}
}
