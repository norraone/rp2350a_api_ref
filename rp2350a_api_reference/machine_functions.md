# Machine 包函数参考

// pinFunc represents a GPIO function.
// Each GPIO can have one function selected at a time.
func (p Pin) PortMaskSet() (*uint32, uint32) {
func (p Pin) set() {
func (p Pin) PortMaskClear() (*uint32, uint32) {
func (p Pin) clr() {
func (p Pin) xor() {
func (p Pin) get() bool {
func (p Pin) ioCtrl() *volatile.Register32 {
func (p Pin) padCtrl() *volatile.Register32 {
func (p Pin) pullup() {
func (p Pin) pulldown() {
func (p Pin) pulloff() {
func (p Pin) setSlew(sr bool) {
func (p Pin) setSchmitt(trigger bool) {
// setFunc will set pin function to fn.
func (p Pin) setFunc(fn pinFunc) {
func (p Pin) init() {
func (p Pin) Set(value bool) {
func (p Pin) Get() bool {
// on the RP2040. ORed PinChanges are valid input to most IRQ functions.
// nil func to unset the pin change interrupt. If you do so, the change
func (p Pin) SetInterrupt(change PinChange, callback func(Pin)) error {
func gpioHandleInterrupt(intr interrupt.Interrupt) {
func (change PinChange) events() uint32 {
func (p Pin) ioIntBit(change PinChange) uint32 {
func getIntChange(p Pin, status uint32) PinChange {
// This function will block until the crystal oscillator has stabilised.
func (osc *xoscType) init() {
// watchdog must be implemented by any platform supporting watchdog functionality
func (uart *UART) Configure(config UARTConfig) error {
func (uart *UART) SetBaudRate(br uint32) {
func (uart *UART) writeByte(c byte) error {
func (uart *UART) flush() {
func (uart *UART) SetFormat(databits, stopbits uint8, parity UARTParity) error {
func initUART(uart *UART) {
func (uart *UART) handleInterrupt(interrupt.Interrupt) {
func CPUFrequency() uint32 {
func (clks *clocksType) clock(cix clockIndex) clock {
func (clk *clock) hasGlitchlessMux() bool {
func (clk *clock) configure(src, auxsrc, srcFreq, freq uint32) {
// Note that the entire init function is computed at compile time
func init() {
// Must be called before any other clock function.
func (clks *clocksType) init() {
func (tmr *timerType) timeElapsed() (us uint64) {
// This function is a 'light' sleep and will return early if another
func (tmr *timerType) lightSleep(us uint64) {
// This is a RAM function because may be called during flash programming to enable save/restore of
// flash_exit_xip() ROM func, not the entirety of the QMI window state.
func enterBootloader() {
func doFlashCommand(tx []byte, rx []byte) error {
func (f flashBlockDevice) writeAt(p []byte, off int64) (n int, err error) {
func (f flashBlockDevice) eraseBlocks(start, length int64) error {
func (pll *pll) init(refdiv, fbdiv, postDiv1, postDiv2 uint32) {
func (ps pllSearch) CalcDivs(xoscRef, targetFreq, MHz uint64) (fbdiv uint64, refdiv, pd1, pd2 uint8, err error) {
func abs(a int64) int64 {
func pllFreqOutPostdiv(xosc, fbdiv, MHz uint64, refdiv, postdiv1, postdiv2 uint8) (foutpostdiv uint64, err error) {
func calcVCO(xoscFreq, fbdiv, refdiv uint32) uint32 {
func genTable() {
func (spi *SPI) Tx(w, r []byte) (err error) {
func (spi *SPI) Transfer(w byte) (byte, error) {
func (spi *SPI) SetBaudRate(br uint32) error {
func (spi *SPI) GetBaudRate() uint32 {
func (spi *SPI) Configure(config SPIConfig) error {
func (spi *SPI) initSPI(config SPIConfig) (err error) {
func (spi *SPI) setFormat(mode uint8) {
func (spi *SPI) reset() {
func (spi *SPI) deinit() (resetVal uint32) {
func (spi *SPI) isWritable() bool {
func (spi *SPI) isReadable() bool {
func (spi *SPI) PrintRegs() {
func (spi *SPI) isBusy() bool {
func (spi *SPI) tx(tx []byte) error {
func (spi *SPI) rx(rx []byte, txrepeat byte) error {
// Note this function is guaranteed to exit in a known amount of time (bits sent * time per bit)
func (spi *SPI) txrx(tx, rx []byte) error {
func (wd *watchdogImpl) Configure(config WatchdogConfig) error {
func (wd *watchdogImpl) Start() error {
func (wd *watchdogImpl) Update() {
func (i2c *I2C) WriteRegister(address uint8, register uint8, data []byte) error {
func (i2c *I2C) ReadRegister(address uint8, register uint8, data []byte) error {
func initEndpoint(ep, config uint32) {
func SendUSBInPacket(ep uint32, data []byte) bool {
func sendUSBPacket(ep uint32, data []byte, maxsize uint16) {
func ReceiveUSBControlPacket() ([cdcLineInfoSize]byte, error) {
func handleEndpointRx(ep uint32) []byte {
func AckUsbOutTransfer(ep uint32) {
func setEPDataPID(ep uint32, dataOne bool) {
func SendZlp() {
func sendViaEPIn(ep uint32, data []byte, count int) {
func (dev *USBDevice) SetStallEPIn(ep uint32) {
func (dev *USBDevice) SetStallEPOut(ep uint32) {
func (dev *USBDevice) ClearStallEPIn(ep uint32) {
func (dev *USBDevice) ClearStallEPOut(ep uint32) {
func (d *usbDPSRAM) setupBytes() []byte {
func (d *usbDPSRAM) clear() {
func InitADC() {
func (a ADC) Configure(config ADCConfig) error {
func (a ADC) Get() uint16 {
