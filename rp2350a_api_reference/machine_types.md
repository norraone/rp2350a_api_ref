# Machine 包类型定义

type ioType struct {
type irqCtrl struct {
type irqSummary struct {
type ioBank0Type struct {
type padsBank0Type struct {
type pinFunc uint8
type PinChange uint8
type xoscType struct {
type WatchdogConfig struct {
type watchdog interface {
type UART struct {
type clockIndex uint8
type clockType struct {
type fc struct {
type clock struct {
// Not all clocks have both types of mux.
type timerType struct {
typedef unsigned char uint8_t;
typedef unsigned short uint16_t;
typedef unsigned long uint32_t;
typedef unsigned long size_t;
typedef unsigned long uintptr_t;
typedef long int intptr_t;
typedef const volatile uint16_t io_ro_16;
typedef const volatile uint32_t io_ro_32;
typedef volatile uint16_t io_rw_16;
typedef volatile uint32_t io_rw_32;
typedef volatile uint32_t io_wo_32;
typedef int bool;
typedef void (*flash_exit_xip_fn)(void);
typedef void (*flash_flush_cache_fn)(void);
typedef void (*flash_connect_internal_fn)(void);
typedef void (*flash_range_erase_fn)(uint32_t, size_t, uint32_t, uint16_t);
typedef void (*flash_range_program_fn)(uint32_t, const uint8_t*, size_t);
typedef void *(*rom_table_lookup_fn)(uint32_t code, uint32_t mask);
typedef int (*rom_reboot_fn)(uint32_t flags, uint32_t delay_ms, uint32_t p0, uint32_t p1);
typedef struct {
typedef struct {
typedef struct {
typedef struct {
typedef struct {
typedef struct flash_rp2350_qmi_save_state {
type pll struct {
type pllSearch struct {
type SPIConfig struct {
type SPI struct {
type watchdogImpl struct {
// correctly implemented the methods on the I2C type. They must match
// the i2cController interface method signatures type to type perfectly.
// If not implementing the I2C type please remove your target from the build tags
