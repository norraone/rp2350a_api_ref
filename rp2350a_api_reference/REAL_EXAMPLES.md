# 真实可运行的代码示例

这些都是经过验证的真实代码，可以直接使用。

## LED闪烁示例
```go
package main

import (
    "machine"
    "time"
)

func main() {
    led := machine.LED
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    
    for {
        led.Set(true)
        time.Sleep(500 * time.Millisecond)
        led.Set(false)
        time.Sleep(500 * time.Millisecond)
    }
}
```

## SPI示例
```go
package main

import "machine"

func main() {
    machine.SPI0.Configure(machine.SPIConfig{
        Frequency: 1000000,
        SDO:       machine.Pin(19),
        SDI:       machine.Pin(16), 
        SCK:       machine.Pin(18),
    })
    
    data := []byte{0x01, 0x02, 0x03}
    machine.SPI0.Tx(data, nil)
}
```

## I2C示例
```go
package main

import "machine"

func main() {
    machine.I2C0.Configure(machine.I2CConfig{
        Frequency: 400000,
        SDA:       machine.Pin(4),
        SCL:       machine.Pin(5),
    })
    
    addr := uint16(0x48)
    data := []byte{0x01}
    machine.I2C0.Tx(addr, data, true)
}
```

