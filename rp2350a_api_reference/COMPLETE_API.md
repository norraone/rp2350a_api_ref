# RP2350A TinyGo Complete API Reference

这是从真实源码中提取的完整API参考，用于避免AI幻觉生成不存在的API。

## 常量 (Constants)

```go
A0 Pin
A1 Pin
A2 Pin
A3 Pin
A4 Pin
A5 Pin
ADC0 Pin
ADC1 Pin
ADC2 Pin
ADC3 Pin
ADC4 Pin
ADC5 Pin
ADC6 Pin
ADC7 Pin
BUTTON Pin
CKN Pin
CKP Pin
ClkHSTX 
D0N Pin
D0P Pin
D10 Pin
D11 Pin
D1N Pin
D1P Pin
D2 Pin
D22 Pin
D23 Pin
D26 Pin
D27 Pin
D2N Pin
D2P Pin
D3 Pin
D4 Pin
D5 Pin
D6 Pin
D7 Pin
D8 Pin
D9 Pin
GP0 Pin
GP1 Pin
GP10 Pin
GP11 Pin
GP12 Pin
GP13 Pin
GP14 Pin
GP15 Pin
GP16 Pin
GP17 Pin
GP18 Pin
GP19 Pin
GP2 Pin
GP20 Pin
GP21 Pin
GP22 Pin
GP23 Pin
GP24 Pin
GP25 Pin
GP26 Pin
GP27 Pin
GP28 Pin
GP29 Pin
GP3 Pin
GP30 
GP31 
GP32 Pin
GP33 Pin
GP34 Pin
GP35 Pin
GP36 Pin
GP37 
GP38 
GP39 
GP4 Pin
GP40 
GP41 
GP42 
GP43 
GP44 
GP45 
GP46 
GP47 
GP5 Pin
GP6 Pin
GP7 Pin
GP8 Pin
GP9 Pin
GPIO0 Pin
GPIO1 Pin
GPIO10 Pin
GPIO11 Pin
GPIO12 Pin
GPIO13 Pin
GPIO14 Pin
GPIO15 Pin
GPIO16 Pin
GPIO17 Pin
GPIO18 Pin
GPIO19 Pin
GPIO2 Pin
GPIO20 Pin
GPIO21 Pin
GPIO22 Pin
GPIO23 Pin
GPIO24 Pin
GPIO25 Pin
GPIO26 Pin
GPIO27 Pin
GPIO28 Pin
GPIO29 Pin
GPIO3 Pin
GPIO30 Pin
GPIO31 Pin
GPIO32 Pin
GPIO33 Pin
GPIO34 Pin
GPIO35 Pin
GPIO36 Pin
GPIO37 Pin
GPIO38 Pin
GPIO39 Pin
GPIO4 Pin
GPIO40 Pin
GPIO41 Pin
GPIO42 Pin
GPIO43 Pin
GPIO44 Pin
GPIO45 Pin
GPIO46 Pin
GPIO47 Pin
GPIO5 Pin
GPIO6 Pin
GPIO7 Pin
GPIO8 Pin
GPIO9 Pin
I2C0_SCL_PIN 
I2C0_SDA_PIN 
I2C1_SCL_PIN 
I2C1_SDA_PIN 
I2CFinish 
I2CModeController I2CMode
I2CModeTarget 
I2CReceive I2CTargetEvent
I2CRequest 
LED 
LED_BLUE Pin
LED_GREEN Pin
LED_RED Pin
MISO Pin
MOSI Pin
Mode0 
Mode1 
Mode2 
Mode3 
NEOPIXEL Pin
NumberOfUSBEndpoints 
ParityEven 
ParityNone UARTParity
ParityOdd 
PinAnalog 
PinFalling PinChange
PinI2C 
PinInput 
PinInputPulldown 
PinInputPullup 
PinOutput PinMode
PinPIO0 
PinPIO1 
PinPIO2 
PinPWM 
PinRising 
PinSPI 
PinToggle 
PinUART 
RESETS_RESET_Msk 
RX Pin
SCK Pin
SCL Pin
SDA Pin
SDIO_DATA1 
SDIO_DATA2 
SD_CARD_DETECT 
SD_CS 
SD_MISO 
SD_MOSI 
SD_SCK 
SPI0_SCK_PIN 
SPI0_SDI_PIN 
SPI0_SDO_PIN 
SPI1_SCK_PIN 
SPI1_SDI_PIN 
SPI1_SDO_PIN 
SYSINFO_BASE 
SYSINFO_CHIP_ID_OFFSET 
SYSINFO_CHIP_ID_REVISION_BITS 
SYSINFO_CHIP_ID_REVISION_LSB 
TWI_FREQ_100KHZ 
TWI_FREQ_400KHZ 
TX Pin
UART0_RX_PIN 
UART0_TX_PIN 
UART1_RX_PIN 
UART1_TX_PIN 
UART_RX_PIN 
UART_TX_PIN 
USB_HOST_5V_POWER Pin
USB_HOST_DATA_MINUS Pin
USB_HOST_DATA_PLUS Pin
WS2812 Pin
WatchdogMaxTimeout 
XOSC_STARTUP_DELAY_MULTIPLIER 
_ 
_NUMBANK0_GPIOS 
_NUMBANK0_IRQS 
_NUMIRQ 
_NUMSPINLOCKS 
_PICO_SPINLOCK_ID_IRQ 
_VREG_VOLTAGE_AUTO_ADJUST_DELAY 
ackTimeout 
bits 
cdcLineInfoSize 
cfbdiv 
cfref 
clkADC 
clkGPOUT0 clockIndex
clkGPOUT1 
clkGPOUT2 
clkGPOUT3 
clkPeri 
clkRef 
clkSys 
clkUSB 
compileTimeCheckPeriod uint64
compileTimeCheckSetPeriod uint64
cpd1 
cpd2 
cpuFreq 
crefd 
cvco 
cycles 
defaultBaud uint32
deviceName 
eraseBlockSizeValue 
fifoDepth 
fnGPCK pinFunc
fnHSTX pinFunc
fnI2C pinFunc
fnNULL pinFunc
fnPIO0 pinFunc
fnPIO1 pinFunc
fnPIO2 pinFunc
fnPWM pinFunc
fnQMI pinFunc
fnSIO pinFunc
fnSPI pinFunc
fnUART pinFunc
fnUARTAlt pinFunc
fnUSB pinFunc
foutpd 
initDontReset 
initUnreset 
is48Pin 
maxBaud uint32
maxCPUFreq 
maxPWMPins 
maxPeriod 
maxTop 
maxXoscMHz 
memoryStart 
mhz 
microsecond 
milliseconds 
minSleep 
nanosecond 
notimpl 
numClocks 
numTimers 
numberOfCycles 
padEnableMask 
rp2350ExtraReg 
sleepAlarm 
sleepAlarmIRQ 
spi0DMAChannel 
spi1DMAChannel 
thermADC 
timeout_us 
topStart 
usbBuf0CtrlAvail 
usbBuf0CtrlData0Pid 
usbBuf0CtrlData1Pid 
usbBuf0CtrlFull 
usbBuf0CtrlLast 
usbBuf0CtrlLenMask 
usbBuf0CtrlSel 
usbBuf0CtrlStall 
usbBuf1CtrlAvail 
usbBuf1CtrlData0Pid 
usbBuf1CtrlData1Pid 
usbBuf1CtrlFull 
usbBuf1CtrlLast 
usbBuf1CtrlLenMask 
usbBuf1CtrlSel 
usbBuf1CtrlStall 
usbBufferLen 
usbEpControlBufferAddress 
usbEpControlDoubleBuffered 
usbEpControlEnable 
usbEpControlEndpointType 
usbEpControlEndpointTypeBulk 
usbEpControlEndpointTypeControl 
usbEpControlEndpointTypeISO 
usbEpControlEndpointTypeInterrupt 
usbEpControlInterruptOnNak 
usbEpControlInterruptOnStall 
usbEpControlInterruptPerBuff 
usbEpControlInterruptPerDoubleBuff 
usb_STRING_MANUFACTURER 
usb_STRING_PRODUCT 
writeBlockSize 
xoscFreq 
```

## 变量 (Variables)

```go
var DefaultUART 
var ErrBadPeriod 
var ErrLSBNotSupported 
var ErrSPIBaud 
var ErrSPITimeout 
var ErrTxInvalidSliceSize 
var ErrUSBBytesRead 
var ErrUSBBytesWritten 
var ErrUSBReadTimeout 
var Flash flashBlockDevice
var I2C0 
var I2C1 
var PWM0 
var PWM1 
var PWM10 
var PWM11 
var PWM2 
var PWM3 
var PWM4 
var PWM5 
var PWM6 
var PWM7 
var PWM8 
var PWM9 
var SPI0 
var SPI1 
var UART0 
var UART1 
var USBCDC Serialer
var USBDev 
var Watchdog 
var _ 
var _I2C0 
var _I2C1 
var _UART0 
var _UART1 
var _usbDPSRAM 
var abortReason i2cAbortError
var adcAref uint32
var adcLock sync.Mutex
var b []byte
var base *irqCtrl
var bestFbdiv uint64
var bestFreq uint64
var bestMargin int64
var bestRefdiv uint8
var bestpd1 uint8
var bestpd2 uint8
var ch *dmaChannel
var clocks 
var configuredFreq []uint32
var deviceIDBuf []byte
var dmaChannels 
var dreq uint32
var endPoints 
var epXPIDReset []bool
var epXdata0 []bool
var errFlashCannotErasePage 
var errFlashCannotErasePastEOF 
var errFlashCannotReadPastEOF 
var errFlashCannotWriteData 
var errFlashCannotWritePastEOF 
var errFlashInvalidWriteLength 
var errFlashNotAllowedWriteData 
var errI2CAckExpected 
var errI2CAlreadyListening 
var errI2CBusError 
var errI2CBusReadyTimeout 
var errI2CDisable 
var errI2CGeneric 
var errI2CMultipleDevices 
var errI2CNoDevices 
var errI2CNotImplemented 
var errI2COverflow 
var errI2COverread 
var errI2CReadTimeout 
var errI2CSignalReadTimeout 
var errI2CSignalStartTimeout 
var errI2CSignalStopTimeout 
var errI2CUnderflow 
var errI2CWriteTimeout 
var errI2CWrongAddress 
var errI2CWrongMode 
var errInvalidI2CBaudrate 
var errInvalidI2CSCL 
var errInvalidI2CSDA 
var errInvalidTgtAddr 
var errSPIInvalidMachineConfig 
var errSPIInvalidSCK 
var errSPIInvalidSDI 
var errSPIInvalidSDO 
var errUARTBufferEmpty 
var errVCOOverflow 
var fbrd uint32
var flashDataEnd []byte
var flashDataStart []byte
var frac 
var freq 
var gpio Pin
var ioBank0 
var isEndpointHalt 
var isRemoteWakeUpEnabled 
var lo uint32
var mask uint32
var msk uint32
var nextHi uint32
var okSCK bool
var okSCL bool
var okSDA bool
var okSDI bool
var okSDO bool
var padsBank0 
var pdTable 
var pen uint8
var pev uint8
var phc 
var pinCallbacks [][]interface{}
var pllRst uint32
var pllSys 
var pllUSB 
var pllsysFB uint32
var pllsysPD1 uint32
var pllsysPD2 uint32
var poly uint8
var pos uint8
var postdiv uint32
var prescale uint32
var randomByte uint8
var resetVal uint32
var resets 
var rhs 
var rxleft 
var sendOnEP0DATADONE interface{}
var setInt [][]bool
var setupBytes []byte
var timer 
var txleft 
var udd_ep_control_cache_buffer []uint8
var udd_ep_in_cache_buffer [][]uint8
var udd_ep_out_cache_buffer [][]uint8
var usbConfiguration uint8
var usbDescriptor descriptor.Descriptor
var usbEndpointDescriptors []descriptor.Device
var usbRxHandler []interface{}
var usbSetInterface uint8
var usbSetupHandler []interface{}
var usbStallHandler []interface{}
var usbTxHandler []interface{}
var usb_PID uint16
var usb_VID uint16
var usb_trans_buffer []uint8
var val uint32
var whole 
var xosc 
```

## 类型 (Types)

```go
type ADCChannel uint8
type ADCConfig interface{}
type BlockDevice interface{}
type I2C interface{}
type I2CConfig interface{}
type I2CMode int
type I2CTargetEvent uint8
type PinChange uint8
type SPI interface{}
type SPIConfig interface{}
type Serialer interface{}
type UART interface{}
type UARTParity uint8
type USBDevice interface{}
type WatchdogConfig interface{}
type clock interface{}
type clockIndex uint8
type clockType interface{}
type clocksType interface{}
type dmaChannel interface{}
type fc interface{}
type flashBlockDevice interface{}
type i2cAbortError uint32
type ioBank0Type interface{}
type ioType interface{}
type irqCtrl interface{}
type irqSummary interface{}
type padsBank0Type interface{}
type pinFunc uint8
type pll interface{}
type pllSearch interface{}
type pwmGroup interface{}
type timerType interface{}
type usbBuffer interface{}
type usbBufferControlRegister interface{}
type usbDPSRAM interface{}
type usbEndpointControlRegister interface{}
type watchdog interface{}
type watchdogImpl interface{}
type xoscType interface{}
```

## 函数 (Functions)

```go
func AckUsbOutTransfer(uint32)
func CPUFrequency() uint32
func ChipVersion() uint8
func ConfigureUSBEndpoint(descriptor.Descriptor, []usb.EndpointConfig, []usb.SetupConfig)
func CurrentCore() int
func DeviceID() []byte
func EnableCDC(interface{}, interface{}, interface{})
func EnterBootloader()
func FlashDataEnd() uintptr
func FlashDataStart() uintptr
func GetRNG() (uint32, error)
func InitADC()
func NumCores() int
func PWMPeripheral(Pin) (uint8, error)
func ReadTemperature() int32
func ReceiveUSBControlPacket() ([]byte, error)
func SendUSBInPacket(uint32, []byte) bool
func SendZlp()
func abs(int64) int64
func adjustCoreVoltage() bool
func armEPZeroStall()
func boolToBit(bool) uint32
func calcClockDiv(uint32) uint32
func calcVCO(uint32) uint32
func doFlashCommand([]byte, []byte) error
func enterBootloader()
func eraseBlockSize() int64
func flashPad([]byte, int) []byte
func genTable()
func getIntChange(Pin, uint32) PinChange
func getPWMGroup(uintptr) *pwmGroup
func gpioHandleInterrupt(interrupt.Interrupt)
func handleEndpointRx(uint32) []byte
func handleStandardSetup(usb.Setup) bool
func handleUSBIRQ(interrupt.Interrupt)
func handleUSBSetAddress(usb.Setup) bool
func init()
func initEndpoint(uint32)
func initUART(*UART)
func irqSet(uint32, bool)
func isReservedI2CAddr(uint8) bool
func lightSleep(uint64)
func machineInit()
func pllFreqOutPostdiv(uint64, uint8) (uint64, error)
func pwmGPIOToChannel(Pin) uint8
func pwmGPIOToSlice(Pin) uint8
func readAddress(int64) uintptr
func resetBlock(uint32)
func roscRandByte() uint8
func sendDescriptor(usb.Setup)
func sendUSBPacket(uint32, []byte, uint16)
func sendViaEPIn(uint32, []byte, int)
func setEPDataPID(uint32, bool)
func strToUTF16LEDescriptor(string, []byte)
func ticks() uint64
func u32max(uint32) uint32
func unresetBlock(uint32)
func unresetBlockWait(uint32)
func usbManufacturer() string
func usbProduct() string
func usbProductID() uint16
func usbSerial() string
func usbVendorID() uint16
func waitForReady()
func writeAddress(int64) uintptr
```

## 方法 (Methods)

### *I2C Methods

```go
func (*I2C) Configure(I2CConfig) error
func (*I2C) Listen(uint16) error
func (*I2C) ReadRegister(uint8, uint8, []byte) error
func (*I2C) Reply([]byte) error
func (*I2C) SetBaudRate(uint32) error
func (*I2C) Tx(uint16, []byte) error
func (*I2C) WaitForEvent([]byte) (I2CTargetEvent, int, error)
func (*I2C) WriteRegister(uint8, uint8, []byte) error
func (*I2C) clearAbortReason()
func (*I2C) deinit() uint32
func (*I2C) disable() error
func (*I2C) enable()
func (*I2C) getAbortReason() i2cAbortError
func (*I2C) init(I2CConfig) error
func (*I2C) interrupted(uint32) bool
func (*I2C) listen(uint8) error
func (*I2C) readAvailable() uint32
func (*I2C) reset()
func (*I2C) tx(uint8, []byte) error
func (*I2C) writeAvailable() uint32
```

### *SPI Methods

```go
func (*SPI) Configure(SPIConfig) error
func (*SPI) GetBaudRate() uint32
func (*SPI) PrintRegs()
func (*SPI) SetBaudRate(uint32) error
func (*SPI) Transfer(byte) (byte, error)
func (*SPI) Tx([]byte) error
func (*SPI) deinit() uint32
func (*SPI) initSPI(SPIConfig) error
func (*SPI) isBusy() bool
func (*SPI) isReadable() bool
func (*SPI) isWritable() bool
func (*SPI) reset()
func (*SPI) rx([]byte, byte) error
func (*SPI) setFormat(uint8)
func (*SPI) tx([]byte) error
func (*SPI) txrx([]byte) error
```

### *UART Methods

```go
func (*UART) Buffered() int
func (*UART) Configure(UARTConfig) error
func (*UART) Read([]byte) (int, error)
func (*UART) ReadByte() (byte, error)
func (*UART) Receive(byte)
func (*UART) SetBaudRate(uint32)
func (*UART) SetFormat(uint8, UARTParity) error
func (*UART) Write([]byte) (int, error)
func (*UART) WriteByte(byte) error
func (*UART) flush()
func (*UART) handleInterrupt(interrupt.Interrupt)
func (*UART) writeByte(byte) error
```

### *USBDevice Methods

```go
func (*USBDevice) ClearStallEPIn(uint32)
func (*USBDevice) ClearStallEPOut(uint32)
func (*USBDevice) Configure(UARTConfig)
func (*USBDevice) SetStallEPIn(uint32)
func (*USBDevice) SetStallEPOut(uint32)
```

### *clock Methods

```go
func (*clock) configure(uint32)
func (*clock) hasGlitchlessMux() bool
```

### *clocksType Methods

```go
func (*clocksType) clock(clockIndex) clock
func (*clocksType) init()
func (*clocksType) initRTC()
func (*clocksType) initTicks()
```

### *pll Methods

```go
func (*pll) init(uint32)
```

### *pwmGroup Methods

```go
func (*pwmGroup) Channel(Pin) (uint8, error)
func (*pwmGroup) Configure(PWMConfig) error
func (*pwmGroup) Counter() uint32
func (*pwmGroup) Enable(bool)
func (*pwmGroup) Get(uint8) uint32
func (*pwmGroup) IsEnabled() bool
func (*pwmGroup) Period() uint64
func (*pwmGroup) Set(uint8, uint32)
func (*pwmGroup) SetCounter(uint32)
func (*pwmGroup) SetInverting(uint8, bool)
func (*pwmGroup) SetPeriod(uint64) error
func (*pwmGroup) SetTop(uint32)
func (*pwmGroup) Top() uint32
func (*pwmGroup) enable(bool)
func (*pwmGroup) getChanLevel(uint8) uint16
func (*pwmGroup) getClockDiv() uint8
func (*pwmGroup) getPhaseCorrect() uint32
func (*pwmGroup) getWrap() uint32
func (*pwmGroup) init(PWMConfig, bool) error
func (*pwmGroup) peripheral() uint8
func (*pwmGroup) setChanLevel(uint8, uint16)
func (*pwmGroup) setClockDiv(uint8)
func (*pwmGroup) setDivMode(uint32)
func (*pwmGroup) setInverting(uint8, bool)
func (*pwmGroup) setPeriod(uint64) error
func (*pwmGroup) setPhaseCorrect(bool)
func (*pwmGroup) setWrap(uint16)
```

### *timerType Methods

```go
func (*timerType) lightSleep(uint64)
func (*timerType) timeElapsed() uint64
```

### *usbDPSRAM Methods

```go
func (*usbDPSRAM) clear()
func (*usbDPSRAM) setupBytes() []byte
```

### *watchdogImpl Methods

```go
func (*watchdogImpl) Configure(WatchdogConfig) error
func (*watchdogImpl) Start() error
func (*watchdogImpl) Update()
func (*watchdogImpl) startTick(uint32)
```

### *xoscType Methods

```go
func (*xoscType) init()
```

### ADC Methods

```go
func (ADC) Configure(ADCConfig) error
func (ADC) Get() uint16
func (ADC) GetADCChannel() (ADCChannel, error)
```

### ADCChannel Methods

```go
func (ADCChannel) Configure(ADCConfig) error
func (ADCChannel) Pin() (Pin, error)
func (ADCChannel) getOnce() uint16
func (ADCChannel) getVoltage() uint32
```

### Pin Methods

```go
func (Pin) Configure(PinConfig)
func (Pin) Get() bool
func (Pin) PortMaskClear() (*uint32, uint32)
func (Pin) PortMaskSet() (*uint32, uint32)
func (Pin) Set(bool)
func (Pin) SetInterrupt(PinChange, interface{}) error
func (Pin) acknowledgeInterrupt(PinChange)
func (Pin) clr()
func (Pin) ctrlSetInterrupt(PinChange, bool, *irqCtrl)
func (Pin) get() bool
func (Pin) init()
func (Pin) ioCtrl() *volatile.Register32
func (Pin) ioIntBit(PinChange) uint32
func (Pin) padCtrl() *volatile.Register32
func (Pin) pulldown()
func (Pin) pulloff()
func (Pin) pullup()
func (Pin) set()
func (Pin) setFunc(pinFunc)
func (Pin) setInterrupt(PinChange, bool)
func (Pin) setSchmitt(bool)
func (Pin) setSlew(bool)
func (Pin) xor()
```

### PinChange Methods

```go
func (PinChange) events() uint32
```

### flashBlockDevice Methods

```go
func (flashBlockDevice) EraseBlockSize() int64
func (flashBlockDevice) EraseBlocks(int64) error
func (flashBlockDevice) ReadAt([]byte, int64) (int, error)
func (flashBlockDevice) Size() int64
func (flashBlockDevice) WriteAt([]byte, int64) (int, error)
func (flashBlockDevice) WriteBlockSize() int64
func (flashBlockDevice) eraseBlocks(int64) error
func (flashBlockDevice) writeAt([]byte, int64) (int, error)
```

### i2cAbortError Methods

```go
func (i2cAbortError) Error() string
func (i2cAbortError) Reasons() []string
```

### pllSearch Methods

```go
func (pllSearch) CalcDivs(uint64) (uint64, uint8, error)
```

