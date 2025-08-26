# RP2350A TinyGo 常用API速查

## ⚠️ 防幻觉提醒
以下是真实存在的API，AI生成代码时请严格按照这些定义使用。

## 时间相关
```go
// 真实存在的时间函数
time.Sleep(time.Millisecond * 100)
time.Since(start)
time.Now()

// 不存在：machine.Delay(), machine.Sleep()
```

## GPIO相关
```go
// 真实的GPIO操作
pin := machine.Pin(2)
pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
pin.Set(true)
pin.Get()

// 真实的引脚模式
machine.PinOutput
machine.PinInput
machine.PinInputPullup
machine.PinInputPulldown
```

## PWM相关
```go
// 真实的PWM操作
pwm := machine.PWM0
pwm.Configure(machine.PWMConfig{Period: 20000})
channel, _ := pwm.Channel(machine.Pin(0))
pwm.Set(channel, pwm.Top()/2)
```

## SPI相关
```go
// 真实的SPI配置
machine.SPI0.Configure(machine.SPIConfig{
    Frequency: 1000000,
    SDO:       machine.Pin(19),
    SDI:       machine.Pin(16),
    SCK:       machine.Pin(18),
})
```

## I2C相关
```go
// 真实的I2C配置
machine.I2C0.Configure(machine.I2CConfig{
    Frequency: 400000,
    SDA:       machine.Pin(4),
    SCL:       machine.Pin(5),
})
```

## UART相关
```go
// 真实的UART配置
machine.UART0.Configure(machine.UARTConfig{
    BaudRate: 115200,
    TX:       machine.Pin(0),
    RX:       machine.Pin(1),
})
```

