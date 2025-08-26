package gpio

import "device/rp"

// Direction is the direction of a GPIO pin.
type Direction uint8

// Function is the function select for a GPIO pin.
type Function uint8

const (
	// Input is for reading the value of a GPIO pin.
	Input Direction = iota
	// Output is for writing a value to a GPIO pin.
	Output
)

const (
	FunctionSIO   Function = 5
	FunctionPIO0  Function = 6
	FunctionPIO1  Function = 7
)

// SetDir sets the direction of a GPIO pin.
func SetDir(pin uint, dir Direction) {
	mask := uint32(1) << pin
	if dir == Output {
		// Set the direction to output.
		rp.SIO.GPIO_OE_SET.Set(mask)
	} else {
		// Set the direction to input.
		rp.SIO.GPIO_OE_CLR.Set(mask)
	}
}

// Put sets the value of a GPIO pin.
func Put(pin uint, value bool) {
	mask := uint32(1) << pin
	if value {
		rp.SIO.GPIO_SET.Set(mask)
	} else {
		rp.SIO.GPIO_CLR.Set(mask)
	}
}

// Get reads the value of a GPIO pin.
func Get(pin uint) bool {
	mask := uint32(1) << pin
	return (rp.SIO.GPIO_IN.Get()&mask) != 0
}

// SetFunction selects the function for a GPIO pin.
func SetFunction(pin uint, fn Function) {
	// Set input enable on, output disable off
	rp.PADS_BANK0.GPIO[pin].SetBits(rp.PADS_BANK0_GPIO0_IE)
	rp.PADS_BANK0.GPIO[pin].ClearBits(rp.PADS_BANK0_GPIO0_OD)
	// Set the function select
	rp.IO_BANK0.GPIO[pin].CTRL.Set(uint32(fn) << rp.IO_BANK0_GPIO0_CTRL_FUNCSEL_Pos)
}
