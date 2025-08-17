# A simple logger class to handle printing scrolling text to an OLED display.

class OLEDLogger:
    def __init__(self, display, font_width=8, font_height=8, inverted=False):
        self.display = display
        self.font_width = font_width
        self.font_height = font_height
        self.inverted = inverted
        
        self.width = self.display.width
        self.height = self.display.height
        
        self.max_lines = self.height // self.font_height
        self.max_chars = self.width // self.font_width
        
        self.buffer = []
        
        # Set colors based on inverted flag
        if self.inverted:
            self.text_color = 0
            self.bg_color = 1
            self.display.invert(True)
        else:
            self.text_color = 1
            self.bg_color = 0
            self.display.invert(False)
            
        self.display.fill(self.bg_color)
        self.display.show()

    def log(self, message):
        """Adds a new message to the log, automatically handling scrolling."""
        # Word wrap the message
        words = message.split(' ')
        lines = []
        current_line = ""
        for word in words:
            if len(current_line) + len(word) + 1 > self.max_chars:
                lines.append(current_line)
                current_line = word
            else:
                if current_line:
                    current_line += " "
                current_line += word
        lines.append(current_line)
        
        # Add new lines to the buffer
        for line in lines:
            self.buffer.append(line)
            if len(self.buffer) > self.max_lines:
                self.buffer.pop(0) # Remove the oldest line
        
        self._redraw()

    def _redraw(self):
        """Internal method to redraw the entire screen from the buffer."""
        self.display.fill(self.bg_color)
        y = 0
        for line in self.buffer:
            self.display.text(line, 0, y, self.text_color)
            y += self.font_height
        self.display.show()

from machine import Pin, I2C
import network
import time
import ssd1309

# --- Config ---
TARGET_SSID = "PCLJ-HOTEL"
LED_PINS = [4, 16, 17, 5]

# --- Hardware Setup ---
i2c = I2C(0, scl=Pin(22), sda=Pin(21), freq=400000)
display = ssd1309.SSD1309_I2C(128, 64, i2c)
leds = [Pin(p, Pin.OUT) for p in LED_PINS]
wlan = network.WLAN(network.STA_IF)
wlan.active(True)

# --- Invert Display ---
display.invert(True)

# --- Global State ---
# Create a list to store the history of RSSI values for the trend chart
rssi_history = [-100] * 128 # Start with no signal

def map_rssi(rssi, in_min, in_max, out_min, out_max):
    """General purpose mapping function."""
    if rssi < in_min: rssi = in_min
    if rssi > in_max: rssi = in_max
    return int((rssi - in_min) * (out_max - out_min) / (in_max - in_min) + out_min)

def update_leds(rssi):
    for led in leds: led.off()
    if rssi > -30: num_leds = 4
    elif rssi > -50: num_leds = 3
    elif rssi > -70: num_leds = 2
    elif rssi > -90: num_leds = 1
    else: num_leds = 0
    for i in range(num_leds): leds[i].on()

def redraw_display(rssi, status_text):
    display.fill(1)
    
    # 1. Top Status Line
    status_line = "{}: {}dBm".format(TARGET_SSID, rssi if status_text=="FOUND" else "--")
    display.text(status_line, 0, 0, 0)
    display.hline(0, 10, 128, 0)

    # 2. Middle Trend Chart
    chart_y_origin = 45
    chart_height = 30
    for x in range(len(rssi_history) - 1):
        y1 = chart_y_origin - map_rssi(rssi_history[x], -90, -30, 0, chart_height)
        y2 = chart_y_origin - map_rssi(rssi_history[x+1], -90, -30, 0, chart_height)
        display.line(x, y1, x + 1, y2, 0)

    # 3. Bottom Live Bar
    bar_width = map_rssi(rssi, -90, -30, 0, 128)
    display.rect(0, 50, bar_width, 10, 0, True)
    display.rect(0, 50, 128, 10, 0)
    
    display.show()

# --- Main Application Loop ---
print("Starting trend chart scanner...")

while True:
    networks = wlan.scan()
    found_network = None
    for ssid, bssid, channel, rssi, authmode, hidden in networks:
        if ssid.decode('utf-8') == TARGET_SSID:
            found_network = {"rssi": rssi}
            break

    status = "FOUND"
    if found_network:
        current_rssi = found_network["rssi"]
    else:
        current_rssi = -100
        status = "MISSING"

    # Update history for the trend chart
    rssi_history.append(current_rssi)
    rssi_history.pop(0) # Keep the list at 128 items

    update_leds(current_rssi)
    redraw_display(current_rssi, status)
    
    # A short delay to keep the loop from running too fast and making the chart unreadable
    time.sleep_ms(250)
