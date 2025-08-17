
# MicroPython SSD1309 OLED driver
#
# Based on the work by Adafruit and other community contributors,
# adapted for the SSD1309 controller.

from micropython import const
import framebuf

# Command constants from the provided C driver
SET_CONTRAST = const(0x81)
SET_ENTIRE_ON = const(0xA4)
SET_NORM_INV = const(0xA6)
SET_DISP = const(0xAE) # 0xAE=off, 0xAF=on
SET_MEM_ADDR = const(0x20)
SET_COL_ADDR = const(0x21)
SET_PAGE_ADDR = const(0x22)
SET_DISP_START_LINE = const(0x40)
SET_SEG_REMAP = const(0xA0) # 0xA0=normal, 0xA1=remapped
SET_MUX_RATIO = const(0xA8)
SET_COM_OUT_DIR = const(0xC0) # 0xC0=normal, 0xC8=remapped
SET_DISP_OFFSET = const(0xD3)
SET_COM_PIN_CFG = const(0xDA)
SET_DISP_CLK_DIV = const(0xD5)
SET_PRECHARGE = const(0xD9)
SET_VCOM_DESEL = const(0xDB)
SET_CHARGE_PUMP = const(0x8D)
CMD_LOCK = const(0xFD)

class SSD1309:
    """Base class for SSD1309 display driver"""
    def __init__(self, width, height, external_vcc):
        self.width = width
        self.height = height
        self.external_vcc = external_vcc
        self.pages = self.height // 8
        self.buffer = bytearray(self.pages * self.width)
        self.framebuf = framebuf.FrameBuffer(self.buffer, self.width, self.height, framebuf.MVLSB)
        self.poweron()
        self.init_display()

    def init_display(self):
        # Initialization sequence based on typical SSD1309 datasheets
        # and comparing with the provided C driver reference.
        for cmd in (
            CMD_LOCK, 0x12, # Unlock commands
            SET_DISP | 0x00,  # Display OFF (0xAE)
            SET_DISP_CLK_DIV, 0x80, # Set Clock Divide Ratio/Oscillator Freq
            SET_MUX_RATIO, self.height - 1, # Set Multiplex Ratio
            SET_DISP_OFFSET, 0x00, # Set Display Offset
            SET_DISP_START_LINE | 0x00, # Set Display Start Line
            SET_CHARGE_PUMP, 0x14, # Enable Charge Pump
            SET_MEM_ADDR, 0x00,  # Set Memory Addressing Mode (Horizontal)
            SET_SEG_REMAP | 0x01,  # Set Segment Re-map (col 127 -> SEG0)
            SET_COM_OUT_DIR | 0x08,  # Set COM Output Scan Direction (scan from COM[N-1] to COM0)
            SET_COM_PIN_CFG, 0x12, # Set COM Pins Hardware Configuration
            SET_CONTRAST, 0xFF, # Set Contrast Control
            SET_PRECHARGE, 0xF1, # Set Pre-charge Period
            SET_VCOM_DESEL, 0x40, # Set VCOMH Deselect Level
            SET_ENTIRE_ON, # Entire Display ON (resume to RAM content)
            SET_NORM_INV, # Set Normal Display
            SET_DISP | 0x01):  # Display ON (0xAF)
            self.write_cmd(cmd)
        self.fill(0)
        self.show()

    def poweroff(self):
        self.write_cmd(SET_DISP | 0x00)

    def poweron(self):
        self.write_cmd(SET_DISP | 0x01)

    def contrast(self, contrast):
        self.write_cmd(SET_CONTRAST)
        self.write_cmd(contrast)

    def invert(self, invert):
        self.write_cmd(SET_NORM_INV | (invert & 1))

    def show(self):
        self.write_cmd(SET_COL_ADDR)
        self.write_cmd(0)
        self.write_cmd(self.width - 1)
        self.write_cmd(SET_PAGE_ADDR)
        self.write_cmd(0)
        self.write_cmd(self.pages - 1)
        self.write_data(self.buffer)

    def fill(self, col):
        self.framebuf.fill(col)

    def pixel(self, x, y, col):
        self.framebuf.pixel(x, y, col)

    def scroll(self, dx, dy):
        self.framebuf.scroll(dx, dy)

    def text(self, string, x, y, col=1):
        self.framebuf.text(string, x, y, col)

    def line(self, x1, y1, x2, y2, c):
        self.framebuf.line(x1, y1, x2, y2, c)

    def hline(self, x, y, w, c):
        self.framebuf.hline(x, y, w, c)

    def vline(self, x, y, h, c):
        self.framebuf.vline(x, y, h, c)

    def rect(self, x, y, w, h, c, f=False):
        self.framebuf.rect(x, y, w, h, c, f)


class SSD1309_I2C(SSD1309):
    def __init__(self, width, height, i2c, addr=0x3c, external_vcc=False):
        self.i2c = i2c
        self.addr = addr
        self.temp = bytearray(2)
        self.write_list = [b'\x40', None]  # Co=0, D/C#=1
        super().__init__(width, height, external_vcc)

    def write_cmd(self, cmd):
        self.temp[0] = 0x00  # Co=0, D/C#=0
        self.temp[1] = cmd
        self.i2c.writeto(self.addr, self.temp)

    def write_data(self, buf):
        self.write_list[1] = buf
        self.i2c.writevto(self.addr, self.write_list)
