# AI代码生成常见幻觉避免列表

## ❌ 常见幻觉API（不存在）

### 时间相关幻觉
- `machine.Delay()` - 不存在，使用 `time.Sleep()`
- `machine.DelayMs()` - 不存在，使用 `time.Sleep(time.Millisecond * n)`
- `machine.KiloHertz` - 不存在，直接使用数字如 `1000`
- `machine.MegaHertz` - 不存在，直接使用数字如 `1000000`

### GPIO相关幻觉
- `pin.SetHigh()` - 不存在，使用 `pin.Set(true)`
- `pin.SetLow()` - 不存在，使用 `pin.Set(false)`
- `pin.Toggle()` - 不存在，需要手动实现
- `machine.PinModeOutput` - 不存在，使用 `machine.PinOutput`

### PWM相关幻觉
- `machine.PWM.SetDutyCycle()` - 不存在，使用 `pwm.Set(channel, value)`
- `machine.PWM.SetFrequency()` - 不存在，在 Configure 中设置 Period

### 其他常见幻觉
- `machine.EnableInterrupts()` - 不存在
- `machine.DisableInterrupts()` - 不存在
- `machine.WaitForInterrupt()` - 不存在

## ✅ 正确的替代方案

### 时间操作
```go
// 正确的延时
time.Sleep(100 * time.Millisecond)
time.Sleep(1 * time.Second)

// 正确的频率设置
spi.Configure(machine.SPIConfig{Frequency: 1000000}) // 1MHz
```

### GPIO操作
```go
// 正确的GPIO设置
pin.Set(true)   // 设置为高电平
pin.Set(false)  // 设置为低电平
state := pin.Get() // 读取状态

// 手动实现Toggle
if pin.Get() {
    pin.Set(false)
} else {
    pin.Set(true)
}
```

### PWM操作
```go
// 正确的PWM设置
pwm.Configure(machine.PWMConfig{Period: 20000}) // 设置周期
channel, _ := pwm.Channel(pin)
pwm.Set(channel, pwm.Top()/2) // 50%占空比
```

