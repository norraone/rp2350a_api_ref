# RP2350A AI 代码生成使用指南

## 🎯 目标
通过提供完整、准确的 API 参考来消除 AI 生成代码中的"幻觉"问题，确保生成的代码接近 100% 无虚假 API。

## 📁 文件结构

```
rp2350a_api_reference/
├── AI_USAGE_GUIDE.md          # 本文件 - AI 使用指南
├── API_INDEX.md               # 总索引和说明
├── COMPLETE_API.md            # 从源码提取的完整API (⭐ 最重要)
├── QUICK_REFERENCE.md         # 常用 API 速查表
├── AVOID_HALLUCINATIONS.md    # 常见幻觉API列表
├── REAL_EXAMPLES.md           # 真实可运行的代码示例
├── machine_constants.md       # Machine 包常量
├── machine_types.md           # Machine 包类型定义
├── machine_functions.md       # Machine 包函数
└── board_pins.md              # 开发板引脚定义
```

## 🚀 AI 使用方法

### 1. 优先级顺序
当让 AI 生成 RP2350A 代码时，请按以下优先级参考文档：

1. **COMPLETE_API.md** ⭐ - 最权威，从源码直接提取
2. **AVOID_HALLUCINATIONS.md** - 避免常见错误
3. **QUICK_REFERENCE.md** - 常用 API 模式
4. **REAL_EXAMPLES.md** - 真实示例代码

### 2. 提示词模板

```
你是一个RP2350A TinyGo代码生成助手。请严格按照以下API参考生成代码：

[粘贴相关API文档内容]

要求：
- 只使用上述文档中明确列出的API
- 如果API不在文档中，绝对不要使用
- 优先参考COMPLETE_API.md中的定义
- 避免使用AVOID_HALLUCINATIONS.md中列出的虚假API
- 参考REAL_EXAMPLES.md中的代码模式

请生成: [你的具体需求]
```

### 3. 验证检查清单

生成代码后，请检查：
- ✅ 所有使用的常量都在 `COMPLETE_API.md` 中存在
- ✅ 所有函数调用都在 `COMPLETE_API.md` 中有定义  
- ✅ 没有使用 `AVOID_HALLUCINATIONS.md` 中的虚假API
- ✅ 代码结构参考了 `REAL_EXAMPLES.md` 中的模式

## 📊 API 统计信息

从真实源码中提取的API数量：
- **319 个常量** - 包括引脚定义、模式常量等
- **154 个变量** - 包括硬件实例、默认配置等  
- **40 个类型** - 包括配置结构体、接口等
- **65 个函数** - 包括初始化、配置函数等
- **多个方法** - 按类型分组的方法定义

## ⚠️ 常见陷阱避免

### 时间相关
❌ **不要使用**: `machine.Delay()`, `machine.KiloHertz`, `machine.MegaHertz`
✅ **正确使用**: `time.Sleep()`, 直接使用数字 `1000000`

### GPIO 操作  
❌ **不要使用**: `pin.SetHigh()`, `pin.SetLow()`, `pin.Toggle()`
✅ **正确使用**: `pin.Set(true)`, `pin.Set(false)`, 手动实现toggle

### PWM 操作
❌ **不要使用**: `pwm.SetDutyCycle()`, `pwm.SetFrequency()`  
✅ **正确使用**: `pwm.Set(channel, value)`, 在Configure中设置Period

## 🎯 最佳实践

### 1. 分步验证
```
1. 先检查API是否在COMPLETE_API.md中存在
2. 参考REAL_EXAMPLES.md中的使用模式
3. 避免AVOID_HALLUCINATIONS.md中的错误
4. 生成代码并进行最终检查
```

### 2. 常用代码模式

**GPIO 控制**:
```go
pin := machine.Pin(2)
pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
pin.Set(true)
```

**SPI 配置**:
```go
machine.SPI0.Configure(machine.SPIConfig{
    Frequency: 1000000,
    SDO: machine.Pin(19),
    SDI: machine.Pin(16),
    SCK: machine.Pin(18),
})
```

**定时器延时**:
```go
time.Sleep(100 * time.Millisecond)
```

## 🔧 开发板特定信息

### Pico 2 (RP2350A)
- LED: `machine.LED`
- 标准引脚: `machine.Pin(0)` 到 `machine.Pin(28)`
- SPI: `machine.SPI0`, `machine.SPI1`  
- I2C: `machine.I2C0`, `machine.I2C1`
- UART: `machine.UART0`, `machine.UART1`

### 其他开发板
参考 `board_pins.md` 获取特定开发板的引脚定义。

## 📝 使用示例

**提示AI生成LED闪烁代码**:
```
基于COMPLETE_API.md中的真实API，请生成RP2350A的LED闪烁代码。
要求：
- 使用machine.LED
- 使用time.Sleep()进行延时  
- 使用pin.Set()控制高低电平
- 不要使用任何AVOID_HALLUCINATIONS.md中的虚假API
```

**AI会生成**:
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

## 🎉 效果预期

使用这套API参考，AI生成的代码应该能够：
- ✅ **零幻觉** - 不包含虚假的API调用
- ✅ **直接编译** - 无需修改即可编译通过
- ✅ **正确运行** - 生成的代码逻辑正确
- ✅ **符合规范** - 遵循TinyGo的编码模式

---

**记住**: 这些API都是从真实源码中提取的，是100%准确的。当AI说某个API不存在时，请以这份文档为准！
