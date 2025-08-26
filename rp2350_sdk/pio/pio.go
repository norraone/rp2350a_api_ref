package pio

import (
	"device/rp"
	"errors"
	"fmt"
	"math"
)

// PIO represents a single PIO instance (pio0 or pio1).
// We use a wrapper struct to make it easier to manage state.
type PIO struct {
	HW         *rp.PIO
	num        uint
	claimedSMs uint8 // Bitmask of claimed state machines.
}

// Program defines a PIO program.
type Program struct {
	Instructions []uint16
	Origin       int8 // -1 for auto-placement
}

// Config is an in-memory representation of the configuration for a state machine.
type Config struct {
	ClkDiv    uint32
	ExecCtrl  uint32
	ShiftCtrl uint32
	PinCtrl   uint32
}

// FifoJoin specifies how the TX and RX FIFOs are coupled.
type FifoJoin uint8

const (
	FifoJoinNone FifoJoin = 0 // TX and RX FIFOs are separate.
	FifoJoinTX   FifoJoin = 1 // TX FIFO is 8 words deep, RX is disabled.
	FifoJoinRX   FifoJoin = 2 // RX FIFO is 8 words deep, TX is disabled.
)

const (
	// RP2350 has 3 PIO instances, but we start with 2 for RP2040 compatibility.
	numPIOs          = 2
	instructionCount = 32
	numStateMachines = 4

	instrBitsJmp uint16 = 0x0000
)

var (
	// usedInstructionSpace is a bitmask for each PIO instance to track used instruction slots.
	usedInstructionSpace [numPIOs]uint32

	// PIO0 and PIO1 are the two PIO instances.
	PIO0 = &PIO{HW: rp.PIO0, num: 0}
	PIO1 = &PIO{HW: rp.PIO1, num: 1}
)

// DefaultConfig returns a default configuration for a state machine.
func DefaultConfig() Config {
	c := Config{}
	// From C SDK: sm_config_set_clkdiv_int_frac8(&c, 1, 0)
	c.ClkDiv = (1 << rp.PIO_SM0_CLKDIV_INT_Pos)
	// From C SDK: sm_config_set_wrap(&c, 0, 31)
	c.ExecCtrl = (31 << rp.PIO_SM0_EXECCTRL_WRAP_TOP_Pos) | (0 << rp.PIO_SM0_EXECCTRL_WRAP_BOTTOM_Pos)
	// From C SDK: sm_config_set_in_shift(&c, true, false, 32)
	c.ShiftCtrl = (1 << rp.PIO_SM0_SHIFTCTRL_IN_SHIFTDIR_Pos) | (31 << rp.PIO_SM0_SHIFTCTRL_PUSH_THRESH_Pos)
	// From C SDK: sm_config_set_out_shift(&c, true, false, 32)
	c.ShiftCtrl |= (1 << rp.PIO_SM0_SHIFTCTRL_OUT_SHIFTDIR_Pos) | (31 << rp.PIO_SM0_SHIFTCTRL_PULL_THRESH_Pos)
	return c
}

// SetOutShift configures the output shift register (OSR) behavior.
func (c *Config) SetOutShift(shiftRight, autopull bool, pullThreshold uint) {
	val := uint32(0)
	if shiftRight {
		val |= (1 << rp.PIO_SM0_SHIFTCTRL_OUT_SHIFTDIR_Pos)
	}
	if autopull {
		val |= (1 << rp.PIO_SM0_SHIFTCTRL_AUTOPULL_Pos)
	}
	val |= (uint32(pullThreshold&0x1f) << rp.PIO_SM0_SHIFTCTRL_PULL_THRESH_Pos)

	mask := uint32(rp.PIO_SM0_SHIFTCTRL_OUT_SHIFTDIR_Msk | rp.PIO_SM0_SHIFTCTRL_AUTOPULL_Msk | rp.PIO_SM0_SHIFTCTRL_PULL_THRESH_Msk)

	c.ShiftCtrl = (c.ShiftCtrl &^ mask) | val
}

// SetFIFOJoin configures the FIFO coupling.
func (c *Config) SetFIFOJoin(join FifoJoin) {
	mask := uint32(rp.PIO_SM0_SHIFTCTRL_FJOIN_TX_Msk | rp.PIO_SM0_SHIFTCTRL_FJOIN_RX_Msk)
	val := (uint32(join) << rp.PIO_SM0_SHIFTCTRL_FJOIN_TX_Pos)
	c.ShiftCtrl = (c.ShiftCtrl &^ mask) | val
}

// SetClkDiv sets the clock divider for the state machine.
func (c *Config) SetClkDiv(div float32) {
	var divInt uint32
	var divFrac uint8
	if div < 1 {
		divInt = 0
		divFrac = 0
	} else if div > 65536 {
		divInt = 65536
		divFrac = 0
	} else {
		divInt = uint32(div)
		divFrac = uint8((div - float32(divInt)) * 256)
	}
	c.ClkDiv = (uint32(divFrac) << rp.PIO_SM0_CLKDIV_FRAC_Pos) | (divInt << rp.PIO_SM0_CLKDIV_INT_Pos)
}

// Init applies the configuration to a state machine and sets it up to run.
package pio

import (
	"device/rp"
	"errors"
	"fmt"
	"math"
)

// PIO represents a single PIO instance (pio0 or pio1).
// We use a wrapper struct to make it easier to manage state.
type PIO struct {
	HW         *rp.PIO
	num        uint
	claimedSMs uint8 // Bitmask of claimed state machines.
}

// Program defines a PIO program.
type Program struct {
	Instructions []uint16
	Origin       int8 // -1 for auto-placement
}

// Config is an in-memory representation of the configuration for a state machine.
type Config struct {
	ClkDiv    uint32
	ExecCtrl  uint32
	ShiftCtrl uint32
	PinCtrl   uint32
}

// FifoJoin specifies how the TX and RX FIFOs are coupled.
type FifoJoin uint8

const (
	FifoJoinNone FifoJoin = 0 // TX and RX FIFOs are separate.
	FifoJoinTX   FifoJoin = 1 // TX FIFO is 8 words deep, RX is disabled.
	FifoJoinRX   FifoJoin = 2 // RX FIFO is 8 words deep, TX is disabled.
)

const (
	// RP2350 has 3 PIO instances, but we start with 2 for RP2040 compatibility.
	numPIOs          = 2
	instructionCount = 32
	numStateMachines = 4

	instrBitsJmp uint16 = 0x0000
)

var (
	// usedInstructionSpace is a bitmask for each PIO instance to track used instruction slots.
	usedInstructionSpace [numPIOs]uint32

	// PIO0 and PIO1 are the two PIO instances.
	PIO0 = &PIO{HW: rp.PIO0, num: 0}
	PIO1 = &PIO{HW: rp.PIO1, num: 1}
)

// DefaultConfig returns a default configuration for a state machine.
func DefaultConfig() Config {
	c := Config{}
	// From C SDK: sm_config_set_clkdiv_int_frac8(&c, 1, 0)
	c.ClkDiv = (1 << rp.PIO_SM0_CLKDIV_INT_Pos)
	// From C SDK: sm_config_set_wrap(&c, 0, 31)
	c.ExecCtrl = (31 << rp.PIO_SM0_EXECCTRL_WRAP_TOP_Pos) | (0 << rp.PIO_SM0_EXECCTRL_WRAP_BOTTOM_Pos)
	// From C SDK: sm_config_set_in_shift(&c, true, false, 32)
	c.ShiftCtrl = (1 << rp.PIO_SM0_SHIFTCTRL_IN_SHIFTDIR_Pos) | (31 << rp.PIO_SM0_SHIFTCTRL_PUSH_THRESH_Pos)
	// From C SDK: sm_config_set_out_shift(&c, true, false, 32)
	c.ShiftCtrl |= (1 << rp.PIO_SM0_SHIFTCTRL_OUT_SHIFTDIR_Pos) | (31 << rp.PIO_SM0_SHIFTCTRL_PULL_THRESH_Pos)
	return c
}

// SetOutShift configures the output shift register (OSR) behavior.
func (c *Config) SetOutShift(shiftRight, autopull bool, pullThreshold uint) {
	val := uint32(0)
	if shiftRight {
		val |= (1 << rp.PIO_SM0_SHIFTCTRL_OUT_SHIFTDIR_Pos)
	}
	if autopull {
		val |= (1 << rp.PIO_SM0_SHIFTCTRL_AUTOPULL_Pos)
	}
	val |= (uint32(pullThreshold&0x1f) << rp.PIO_SM0_SHIFTCTRL_PULL_THRESH_Pos)

	mask := uint32(rp.PIO_SM0_SHIFTCTRL_OUT_SHIFTDIR_Msk | rp.PIO_SM0_SHIFTCTRL_AUTOPULL_Msk | rp.PIO_SM0_SHIFTCTRL_PULL_THRESH_Msk)

	c.ShiftCtrl = (c.ShiftCtrl &^ mask) | val
}

// SetFIFOJoin configures the FIFO coupling.
func (c *Config) SetFIFOJoin(join FifoJoin) {
	mask := uint32(rp.PIO_SM0_SHIFTCTRL_FJOIN_TX_Msk | rp.PIO_SM0_SHIFTCTRL_FJOIN_RX_Msk)
	val := (uint32(join) << rp.PIO_SM0_SHIFTCTRL_FJOIN_TX_Pos)
	c.ShiftCtrl = (c.ShiftCtrl &^ mask) | val
}

// SetClkDiv sets the clock divider for the state machine.
func (c *Config) SetClkDiv(div float32) {
	var divInt uint32
	var divFrac uint8
	if div < 1 {
		divInt = 0
		divFrac = 0
	} else if div > 65536 {
		divInt = 65536
		divFrac = 0
	} else {
		divInt = uint32(div)
		divFrac = uint8((div - float32(divInt)) * 256)
	}
	c.ClkDiv = (uint32(divFrac) << rp.PIO_SM0_CLKDIV_FRAC_Pos) | (divInt << rp.PIO_SM0_CLKDIV_INT_Pos)
}

// Init applies the configuration to a state machine and sets it up to run.
func (p *PIO) Init(sm int, offset uint, cfg Config) {
	// Halt the state machine
	p.SetEnabled(sm, false)

	// Apply the configuration
	p.HW.SM[sm].CLKDIV.Set(cfg.ClkDiv)
	p.HW.SM[sm].EXECCTRL.Set(cfg.ExecCtrl)
	p.HW.SM[sm].SHIFTCTRL.Set(cfg.ShiftCtrl)
	p.HW.SM[sm].PINCTRL.Set(cfg.PinCtrl)

	// Clear FIFOs
	p.HW.SM[sm].SHIFTCTRL.SetBits(rp.PIO_SM0_SHIFTCTRL_FJOIN_RX_Msk)
	p.HW.SM[sm].SHIFTCTRL.ClearBits(rp.PIO_SM0_SHIFTCTRL_FJOIN_RX_Msk)

	// Restart the state machine and jump to the initial PC
	p.HW.CTRL.SetBits(1 << (rp.PIO_CTRL_SM_RESTART_Pos + sm))
	p.HW.CTRL.SetBits(1 << (rp.PIO_CTRL_CLKDIV_RESTART_Pos + sm))
	p.HW.SM[sm].INSTR.Set(encodeJmp(uint16(offset)))
}

// SetEnabled enables or disables a state machine.
func (p *PIO) SetEnabled(sm int, enabled bool) {
	if enabled {
		p.HW.CTRL.SetBits(1 << sm)
	} else {
		p.HW.CTRL.ClearBits(1 << sm)
	}
}

// PutBlocking writes a word to a state machine's TX FIFO, blocking if the FIFO is full.
func (p *PIO) PutBlocking(sm int, data uint32) {
	for p.IsTxFIFOFull(sm) {
		// wait
	}
	p.Put(sm, data)
}

// Put writes a word to a state machine's TX FIFO.
func (p *PIO) Put(sm int, data uint32) {
	p.HW.TXF[sm].Set(data)
}

// IsTxFIFOFull checks if the TX FIFO is full.
func (p *PIO) IsTxFIFOFull(sm int) bool {
	return p.HW.FSTAT.HasBits(1 << (rp.PIO_FSTAT_TXFULL_Pos + sm))
}

// ClaimUnusedSM finds and claims a free state machine on this PIO instance.
func (p *PIO) ClaimUnusedSM() (sm int, err error) {
	for i := 0; i < numStateMachines; i++ {
		mask := uint8(1 << i)
		if (p.claimedSMs&mask) == 0 {
			p.claimedSMs |= mask
			return i, nil
		}
	}
	return -1, fmt.Errorf("no free state machines on PIO%d", p.num)
}

// UnclaimSM marks a state machine as unused.
func (p *PIO) UnclaimSM(sm int) {
	if sm < 0 || sm >= numStateMachines {
		return
	}
	p.claimedSMs &^= (1 << sm)
}

// AddProgram finds a free spot in the PIO's instruction memory and loads the program.
func (p *PIO) AddProgram(prog *Program) (offset int, err error) {
	offset, err = p.findOffsetForProgram(prog)
	if err != nil {
		return -1, err
	}

	// Load the program into instruction memory.
	for i, instr := range prog.Instructions {
		finalInstr := instr
		// Check if the instruction is a JMP and adjust its address.
		if majorInstrBits(instr) == instrBitsJmp {
			finalInstr += uint16(offset)
		}
		p.HW.INSTR_MEM[offset+i].Set(finalInstr)
	}

	// Mark the instruction space as used.
	programMask := uint32(1<<len(prog.Instructions)) - 1
	usedInstructionSpace[p.num] |= (programMask << offset)

	return offset, nil
}

// findOffsetForProgram scans the instruction memory for a free slot.
func (p *PIO) findOffsetForProgram(prog *Program) (int, error) {
	if len(prog.Instructions) == 0 || len(prog.Instructions) > instructionCount {
		return -1, errors.New("invalid program length")
	}

	usedMask := usedInstructionSpace[p.num]
	programMask := uint32(1<<len(prog.Instructions)) - 1

	if prog.Origin >= 0 {
		// Specific origin requested.
		if prog.Origin+int8(len(prog.Instructions)) > instructionCount {
			return -1, errors.New("program does not fit at specified origin")
		}
		if (usedMask&(programMask<<prog.Origin)) != 0 {
			return -1, errors.New("space at specified origin is already used")
		}
		return int(prog.Origin), nil
	} else {
		// Find any free slot, starting from the top.
		for i := instructionCount - len(prog.Instructions); i >= 0; i-- {
			if (usedMask&(programMask<<i)) == 0 {
				return i, nil
			}
		}
		return -1, errors.New("not enough free space in instruction memory")
	}
}

// majorInstrBits extracts the 3 major instruction bits from a PIO instruction.
func majorInstrBits(instr uint16) uint16 {
	return instr & 0xe000
}

// encodeJmp creates a JMP instruction.
func encodeJmp(addr uint16) uint16 {
	return instrBitsJmp | (addr & 0x1f)
}

	// Halt the state machine
	p.SetEnabled(sm, false)

	// Apply the configuration
	p.HW.SM[sm].CLKDIV.Set(cfg.ClkDiv)
	p.HW.SM[sm].EXECCTRL.Set(cfg.ExecCtrl)
	p.HW.SM[sm].SHIFTCTRL.Set(cfg.ShiftCtrl)
	p.HW.SM[sm].PINCTRL.Set(cfg.PinCtrl)

	// Clear FIFOs
	p.HW.SM[sm].SHIFTCTRL.SetBits(rp.PIO_SM0_SHIFTCTRL_FJOIN_RX_Msk)
	p.HW.SM[sm].SHIFTCTRL.ClearBits(rp.PIO_SM0_SHIFTCTRL_FJOIN_RX_Msk)

	// Restart the state machine and jump to the initial PC
	p.HW.CTRL.SetBits(1 << (rp.PIO_CTRL_SM_RESTART_Pos + sm))
	p.HW.CTRL.SetBits(1 << (rp.PIO_CTRL_CLKDIV_RESTART_Pos + sm))
	p.HW.SM[sm].INSTR.Set(encodeJmp(uint16(offset)))
}

// SetEnabled enables or disables a state machine.
func (p *PIO) SetEnabled(sm int, enabled bool) {
	if enabled {
		p.HW.CTRL.SetBits(1 << sm)
	} else {
		p.HW.CTRL.ClearBits(1 << sm)
	}
}

// ClaimUnusedSM finds and claims a free state machine on this PIO instance.
func (p *PIO) ClaimUnusedSM() (sm int, err error) {
	for i := 0; i < numStateMachines; i++ {
		mask := uint8(1 << i)
		if (p.claimedSMs&mask) == 0 {
			p.claimedSMs |= mask
			return i, nil
		}
	}
	return -1, fmt.Errorf("no free state machines on PIO%d", p.num)
}

// UnclaimSM marks a state machine as unused.
func (p *PIO) UnclaimSM(sm int) {
	if sm < 0 || sm >= numStateMachines {
		return
	}
	p.claimedSMs &^= (1 << sm)
}

// AddProgram finds a free spot in the PIO's instruction memory and loads the program.
func (p *PIO) AddProgram(prog *Program) (offset int, err error) {
	offset, err = p.findOffsetForProgram(prog)
	if err != nil {
		return -1, err
	}

	// Load the program into instruction memory.
	for i, instr := range prog.Instructions {
		finalInstr := instr
		// Check if the instruction is a JMP and adjust its address.
		if majorInstrBits(instr) == instrBitsJmp {
			finalInstr += uint16(offset)
		}
		p.HW.INSTR_MEM[offset+i].Set(finalInstr)
	}

	// Mark the instruction space as used.
	programMask := uint32(1<<len(prog.Instructions)) - 1
	usedInstructionSpace[p.num] |= (programMask << offset)

	return offset, nil
}

// findOffsetForProgram scans the instruction memory for a free slot.
func (p *PIO) findOffsetForProgram(prog *Program) (int, error) {
	if len(prog.Instructions) == 0 || len(prog.Instructions) > instructionCount {
		return -1, errors.New("invalid program length")
	}

	usedMask := usedInstructionSpace[p.num]
	programMask := uint32(1<<len(prog.Instructions)) - 1

	if prog.Origin >= 0 {
		// Specific origin requested.
		if prog.Origin+int8(len(prog.Instructions)) > instructionCount {
			return -1, errors.New("program does not fit at specified origin")
		}
		if (usedMask&(programMask<<prog.Origin)) != 0 {
			return -1, errors.New("space at specified origin is already used")
		}
		return int(prog.Origin), nil
	} else {
		// Find any free slot, starting from the top.
		for i := instructionCount - len(prog.Instructions); i >= 0; i-- {
			if (usedMask&(programMask<<i)) == 0 {
				return i, nil
			}
		}
		return -1, errors.New("not enough free space in instruction memory")
	}
}

// majorInstrBits extracts the 3 major instruction bits from a PIO instruction.
func majorInstrBits(instr uint16) uint16 {
	return instr & 0xe000
}

// encodeJmp creates a JMP instruction.
func encodeJmp(addr uint16) uint16 {
	return instrBitsJmp | (addr & 0x1f)
}
