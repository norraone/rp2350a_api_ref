from machine import Pin, I2C
import network
import ssd1309
import time
import _thread
import gc # 引入垃圾回收模块

# --- Config ---
# 已更新: 加入了新的4个LED引脚，现在总共8个
LED_PINS = [4, 16, 17, 5, 2, 12, 13, 27] 
# MAX_NETWORKS_TO_SHOW and DISPLAY_LINES are now the same since there's no scrolling.
DISPLAY_LINES = 5         
SMOOTHING_FACTOR = 0.1
RSSI_LOWER_BOUND = -95
UI_FRAME_RATE_MS = 33 # UI刷新率, 约等于30FPS

# --- Hardware Setup ---
i2c = I2C(0, scl=Pin(22), sda=Pin(21), freq=400000)
display = ssd1309.SSD1309_I2C(128, 64, i2c)
leds = [Pin(p, Pin.OUT) for p in LED_PINS]
wlan = network.WLAN(network.STA_IF)
wlan.active(True)

# --- Invert Display ---
display.invert(True)

# --- 全局共享状态 (多线程核心) ---
g_networks = [] 
network_list_lock = _thread.allocate_lock()

# --- 新增：内置5x7像素小字体 ---
# Font data for ASCII characters 32-126
FONT = [
    b'\x00\x00\x00\x00\x00',  # Space
    b'\x00\x00\x5f\x00\x00',  # !
    b'\x00\x07\x00\x07\x00',  # "
    b'\x14\x7f\x14\x7f\x14',  # #
    b'\x24\x2a\x7f\x2a\x12',  # $
    b'\x23\x13\x08\x64\x62',  # %
    b'\x36\x49\x55\x22\x50',  # &
    b'\x00\x05\x03\x00\x00',  # '
    b'\x00\x1c\x22\x41\x00',  # (
    b'\x00\x41\x22\x1c\x00',  # )
    b'\x14\x08\x3e\x08\x14',  # *
    b'\x08\x08\x3e\x08\x08',  # +
    b'\x00\x50\x30\x00\x00',  # ,
    b'\x08\x08\x08\x08\x08',  # -
    b'\x00\x60\x60\x00\x00',  # .
    b'\x20\x10\x08\x04\x02',  # /
    b'\x3e\x51\x49\x45\x3e',  # 0
    b'\x00\x42\x7f\x40\x00',  # 1
    b'\x42\x61\x51\x49\x46',  # 2
    b'\x21\x41\x45\x4b\x31',  # 3
    b'\x18\x14\x12\x7f\x10',  # 4
    b'\x27\x45\x45\x45\x39',  # 5
    b'\x3c\x4a\x49\x49\x30',  # 6
    b'\x01\x71\x09\x05\x03',  # 7
    b'\x36\x49\x49\x49\x36',  # 8
    b'\x06\x49\x49\x29\x1e',  # 9
    b'\x00\x36\x36\x00\x00',  # :
    b'\x00\x56\x36\x00\x00',  # ;
    b'\x08\x14\x22\x41\x00',  # <
    b'\x14\x14\x14\x14\x14',  # =
    b'\x00\x41\x22\x14\x08',  # >
    b'\x02\x01\x51\x09\x06',  # ?
    b'\x32\x49\x79\x41\x3e',  # @
    b'\x7e\x11\x11\x11\x7e',  # A
    b'\x7f\x49\x49\x49\x36',  # B
    b'\x3e\x41\x41\x41\x22',  # C
    b'\x7f\x41\x41\x22\x1c',  # D
    b'\x7f\x49\x49\x49\x41',  # E
    b'\x7f\x09\x09\x09\x01',  # F
    b'\x3e\x41\x49\x49\x7a',  # G
    b'\x7f\x08\x08\x08\x7f',  # H
    b'\x00\x41\x7f\x41\x00',  # I
    b'\x20\x40\x41\x3f\x01',  # J
    b'\x7f\x08\x14\x22\x41',  # K
    b'\x7f\x40\x40\x40\x40',  # L
    b'\x7f\x02\x04\x02\x7f',  # M
    b'\x7f\x04\x08\x10\x7f',  # N
    b'\x3e\x41\x41\x41\x3e',  # O
    b'\x7f\x09\x09\x09\x06',  # P
    b'\x3e\x41\x51\x21\x5e',  # Q
    b'\x7f\x09\x19\x29\x46',  # R
    b'\x46\x49\x49\x49\x31',  # S
    b'\x01\x01\x7f\x01\x01',  # T
    b'\x3f\x40\x40\x40\x3f',  # U
    b'\x1f\x20\x40\x20\x1f',  # V
    b'\x3f\x40\x38\x40\x3f',  # W
    b'\x63\x14\x08\x14\x63',  # X
    b'\x07\x08\x70\x08\x07',  # Y
    b'\x61\x51\x49\x45\x43',  # Z
    b'\x00\x7f\x41\x41\x00',  # [
    b'\x02\x04\x08\x10\x20',  # \
    b'\x00\x41\x41\x7f\x00',  # ]
    b'\x04\x02\x01\x02\x04',  # ^
    b'\x40\x40\x40\x40\x40',  # _
    b'\x00\x01\x02\x04\x00',  # `
    b'\x20\x54\x54\x54\x78',  # a
    b'\x7f\x48\x44\x44\x38',  # b
    b'\x38\x44\x44\x44\x20',  # c
    b'\x38\x44\x44\x48\x7f',  # d
    b'\x38\x54\x54\x54\x18',  # e
    b'\x08\x7e\x09\x01\x02',  # f
    b'\x0c\x52\x52\x52\x3e',  # g
    b'\x7f\x08\x04\x04\x78',  # h
    b'\x00\x44\x7d\x40\x00',  # i
    b'\x20\x40\x44\x3d\x00',  # j
    b'\x7f\x10\x28\x44\x00',  # k
    b'\x00\x41\x7f\x40\x00',  # l
    b'\x7c\x04\x18\x04\x78',  # m
    b'\x7c\x08\x04\x04\x78',  # n
    b'\x38\x44\x44\x44\x38',  # o
    b'\x7c\x14\x14\x14\x08',  # p
    b'\x08\x14\x14\x18\x7c',  # q
    b'\x7c\x08\x04\x04\x08',  # r
    b'\x48\x54\x54\x54\x20',  # s
    b'\x04\x3f\x44\x40\x20',  # t
    b'\x3c\x40\x40\x20\x7c',  # u
    b'\x1c\x20\x40\x20\x1c',  # v
    b'\x3c\x40\x30\x40\x3c',  # w
    b'\x44\x28\x10\x28\x44',  # x
    b'\x0c\x50\x50\x50\x3c',  # y
    b'\x44\x64\x54\x4c\x44',  # z
]

# --- 新增：自定义小字体绘制函数 ---
def draw_text_small(text, x, y, color):
    """
    使用内置的5x7字体逐像素绘制文本。
    """
    for i, char in enumerate(text):
        char_code = ord(char)
        if 32 <= char_code <= 126:
            font_char = FONT[char_code - 32]
            for col in range(5):
                for row in range(7):
                    if (font_char[col] >> row) & 1:
                        display.pixel(x + i * 6 + col, y + row, color)

# --- 后台扫描线程 ---
def scan_thread():
    global g_networks
    while True:
        try:
            scanned_results = wlan.scan()
            scanned_results.sort(key=lambda x: x[3], reverse=True)
            
            network_list_lock.acquire()
            try:
                g_networks = scanned_results
            finally:
                network_list_lock.release()
        
        except Exception as e:
            # 如果扫描出错，打印错误信息但不会让程序崩溃
            print("Scan thread error:", e)
        
        time.sleep_ms(500)

# --- 主线程的辅助函数 ---
def map_rssi(rssi, in_min, in_max, out_min, out_max):
    if rssi < in_min: rssi = in_min
    if rssi > in_max: rssi = in_max
    return int((rssi - in_min) * (out_max - out_min) / (in_max - in_min) + out_min)

def update_leds(current_rssi):
    """
    已更新: 改为使用固定的RSSI阈值来控制8个LED。
    """
    for led in leds: led.off()
    
    # 根据信号强度决定点亮多少个LED
    if current_rssi > -50: num_leds = 8
    elif current_rssi > -55: num_leds = 7
    elif current_rssi > -60: num_leds = 6
    elif current_rssi > -65: num_leds = 5
    elif current_rssi > -70: num_leds = 4
    elif current_rssi > -75: num_leds = 3
    elif current_rssi > -80: num_leds = 2
    elif current_rssi > -85: num_leds = 1
    else: num_leds = 0
            
    for i in range(num_leds): leds[i].on()

def redraw_display(networks_list, strongest_rssi_smooth):
    display.fill(1)
    # 使用新的小字体函数绘制标题
    draw_text_small("WiFi Scanner", 0, 0, 0)
    display.hline(0, 8, 128, 0)
    
    y_pos = 10
    if not networks_list:
        draw_text_small("Scanning...", 0, y_pos, 0)
    else:
        # 直接显示列表的前几项，不再滚动
        networks_to_display = networks_list[:DISPLAY_LINES]
        for net in networks_to_display:
            ssid_bytes, _, _, rssi, _, _ = net
            try:
                ssid = ssid_bytes.decode('utf-8', 'ignore')
            except:
                ssid = bytes(ssid_bytes).hex()
            
            # 小字体可以显示更多字符
            max_len = 12 
            display_ssid = ssid[:max_len] if len(ssid) > max_len else ssid
            line = "{:<12} {:>4}dBm".format(display_ssid, rssi)
            draw_text_small(line, 0, y_pos, 0)
            y_pos += 8 # 调整行高以适应5x7字体

    bar_width = map_rssi(strongest_rssi_smooth, RSSI_LOWER_BOUND, -30, 0, 128)
    display.rect(0, 54, bar_width, 10, 0, True)
    display.rect(0, 54, 128, 10, 0)
    display.show()

# --- 启动后台扫描线程 ---
_thread.start_new_thread(scan_thread, ())
print("Multi-Threaded Scanner Started with Small Font.")

# --- 主线程 (UI 循环) ---
smoothed_rssi = float(RSSI_LOWER_BOUND)

while True:
    # 手动触发垃圾回收，帮助系统保持稳定
    gc.collect()

    local_networks = []
    network_list_lock.acquire()
    try:
        local_networks = list(g_networks)
    finally:
        network_list_lock.release()

    strongest_current_rssi = RSSI_LOWER_BOUND
    if local_networks:
        strongest_current_rssi = local_networks[0][3]

    smoothed_rssi += (strongest_current_rssi - smoothed_rssi) * SMOOTHING_FACTOR

    # 更新LED的调用不再需要max_rssi_seen
    update_leds(strongest_current_rssi)
    # 调用 redraw_display 时不再需要传递滚动相关的参数
    redraw_display(local_networks, smoothed_rssi)
    
    time.sleep_ms(UI_FRAME_RATE_MS)
