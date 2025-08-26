package main

import (
	"time"

	"rp2350_sdk/gpio"
	"rp2350_sdk/pio"
)

// This is the compiled PIO program for WS2812 from the Raspberry Pi Pico SDK.
var ws2812Program = &pio.Program{
	Instructions: []uint16{
		0x80a0, //  0: pull   block
		0x6001, //  1: out    pins, 1
		0x0042, //  2: jmp    x--, 2
		0x8020, //  3: push   block
		0xa027, //  4: mov    x, osr
		0x0000, //  5: jmp    0
	},
	Origin: -1, // Auto-place the program
}

const (
	ws2812Pin    = 16
	ws2812Freq   = 800000 // 800kHz
	sysClockFreq = 125000000 // 125MHz
)

func main() {
	// 1. Choose a PIO instance
	p := pio.PIO0

	// 2. Load the PIO program
	offset, err := p.AddProgram(ws2812Program)
	if err != nil {
		panic(err.Error())
	}

	// 3. Claim a state machine
	sm, err := p.ClaimUnusedSM()
	if err != nil {
		panic(err.Error())
	}
	defer p.UnclaimSM(sm) // Make sure to release the SM when we're done

	// 4. Configure the GPIO pin for PIO
	gpio.SetFunction(ws2812Pin, gpio.FunctionPIO0)

	// 5. Create and customize the configuration
	cfg := pio.DefaultConfig()
	// The PIO program from pico-examples uses a JMP at the end, so we need to set the wrap.
	// The program is 6 instructions long, loaded at `offset`.
	cfg.ExecCtrl = (uint32(offset+5) << rp.PIO_SM0_EXECCTRL_WRAP_TOP_Pos) | (uint32(offset) << rp.PIO_SM0_EXECCTRL_WRAP_BOTTOM_Pos)
	cfg.SetOutShift(false, true, 24)
	cfg.SetFIFOJoin(pio.FifoJoinTX)

	// Calculate the clock divider
	// Formula from C SDK: clock_get_hz(clk_sys) / (freq * 8)
	div := float32(sysClockFreq) / (ws2812Freq * 8)
	cfg.SetClkDiv(div)

	// 6. Initialize and start the state machine
	p.Init(sm, uint(offset), cfg)
	p.SetEnabled(sm, true)

	// 7. Loop forever, sending colors
	colors := []uint32{0x00FF00, 0xFF0000, 0x0000FF, 0xFFFFFF, 0x000000} // Green, Red, Blue, White, Off
	i := 0
	for {
		color := colors[i]
		// WS2812 expects GRB format, but our data is RGB. Let's send it as is for this test.
		// The data needs to be left-shifted by 8 bits because the PIO program is for 24-bit RGB.
		p.PutBlocking(sm, color<<8)

		i = (i + 1) % len(colors)
		time.Sleep(500 * time.Millisecond)
	}
}
