# RP2350A 开发相关文件集合

这个文件夹包含了所有与 RP2350A 微控制器开发相关的文件，从原始项目中提取出来方便独立使用。

## 文件夹结构

### tinygo/ - TinyGo 编译器支持文件
```
tinygo/
├── src/
│   ├── machine/                    # 机器层抽象
│   │   ├── machine_rp2_2350a.go   # RP2350A 专用支持 ⭐
│   │   ├── machine_rp2_2350.go    # RP2350 通用支持
│   │   ├── machine_rp2_2350b.go   # RP2350B 支持
│   │   ├── machine_rp2350_*.go    # RP2350 专用功能
│   │   ├── machine_rp2_*.go       # RP2 通用功能 (支持 RP2350)
│   │   ├── board_*2350*.go        # RP2350 开发板支持
│   │   └── *.go                   # 通用接口文件
│   ├── runtime/                   # 运行时支持
│   │   ├── runtime_rp2.go         # RP2 运行时
│   │   └── runtime_rp2350.go      # RP2350 专用运行时
│   └── device/rp/                 # 设备定义
│       └── rp2350-extra.go        # RP2350 扩展设备定义
├── builder/                       # 构建系统
│   └── build.go                   # 包含 RP2350 构建支持
├── targets/                       # 目标配置文件
│   ├── rp2350.json                # RP2350 基础配置
│   ├── rp2350b.json               # RP2350B 配置
│   ├── pico2.json                 # Pico 2 配置
│   ├── pico-plus2.json            # Pico Plus 2 配置
│   ├── metro-rp2350.json          # Metro RP2350 配置
│   ├── elecrow-rp2350.json        # Elecrow RP2350 配置
│   ├── pga2350.json               # PGA2350 配置
│   ├── tiny2350.json              # Tiny2350 配置
│   └── arm.ld                     # ARM 链接器脚本
├── GNUmakefile                    # 构建系统
└── CHANGELOG.md                   # 变更日志
```

### pico/ - Pico SDK 相关文件
```
pico/
├── p/                             # 外设寄存器定义
│   ├── accessctrl/rp2350.go       # 访问控制
│   ├── adc/rp2350.go              # ADC
│   ├── bootram/rp2350.go          # 启动 RAM
│   ├── clocks/rp2350.go           # 时钟系统
│   ├── dma/rp2350.go              # DMA 控制器
│   ├── trng/rp2350.go             # 真随机数生成器 (新增)
│   ├── sha/rp2350.go              # SHA 加密引擎 (新增)
│   ├── hstx*/rp2350.go            # 高速传输 (新增)
│   ├── powman/rp2350.go           # 电源管理 (增强)
│   └── */rp2350.go                # 其他外设
├── hal/                           # 硬件抽象层
│   ├── dma/                       # DMA HAL
│   └── irq/                       # 中断 HAL
└── devboard/                      # 开发板支持
    ├── pico2/                     # Pico 2 开发板
    ├── weacta10/                  # WeAct A10 开发板
    ├── weactb/                    # WeAct B 开发板
    └── common/                    # 通用组件
```

### pio/ - 可编程 IO 支持
```
pio/
└── rp2-pio/                       # RP2 PIO 库
    ├── piolib/                    # PIO 库函数
    ├── examples/                  # 示例程序
    └── *.go                       # PIO 核心支持
```

## 关键文件说明

### RP2350A 专用文件
- **`tinygo/src/machine/machine_rp2_2350a.go`** - RP2350A 芯片的专用定义和配置
- **`tinygo/src/machine/machine_rp2350_rom.go`** - RP2350 ROM 函数接口
- **`tinygo/src/runtime/runtime_rp2350.go`** - RP2350 运行时支持

### RP2350 新增外设
- **`pico/p/trng/rp2350.go`** - 硬件真随机数生成器
- **`pico/p/sha/rp2350.go`** - 硬件 SHA 加密加速器
- **`pico/p/hstx*/rp2350.go`** - 高速传输外设
- **`pico/p/powman/rp2350.go`** - 增强的电源管理

### 开发板配置
- **Pico 2** - Raspberry Pi 官方 RP2350 开发板
- **Metro RP2350** - Adafruit Metro RP2350 开发板
- **Elecrow RP2350** - Elecrow RP2350 开发板
- **WeAct 开发板** - 第三方 RP2350 开发板

## 文件统计
- 总计 **.go 文件**: 181 个
- **TinyGo 核心文件**: ~30 个
- **外设定义文件**: ~58 个
- **HAL 层文件**: ~20 个
- **示例和驱动**: ~70+ 个

## RP2350A vs RP2040 主要差异

### 新增功能
1. **双核 ARM Cortex-M33** (vs M0+)
2. **硬件安全特性** (TrustZone-M, 访问控制)
3. **真随机数生成器** (TRNG)
4. **硬件 SHA 加速器**
5. **增强的电源管理**
6. **高速传输接口** (HSTX)
7. **更大的内存和存储**

### 兼容性
- 大部分 RP2040 代码可以直接在 RP2350A 上运行
- 新功能通过专用的 API 访问
- 引脚兼容 (部分开发板)

## 使用说明

这些文件可以用于：
1. **TinyGo 编译器开发** - 添加 RP2350A 支持
2. **驱动程序开发** - 基于这些 HAL 层开发外设驱动
3. **固件开发** - 直接使用这些定义开发 RP2350A 固件
4. **学习和参考** - 了解 RP2350A 的内部结构和编程接口

## 注意事项
- 这些文件基于特定版本提取，使用时请注意版本兼容性
- 某些文件可能依赖其他未包含的文件，使用时需要注意依赖关系
- 建议配合官方文档使用

---
提取时间: 2024-08-26
源项目: tinygo_ssd1309
目标芯片: RP2350A (ARM Cortex-M33)
