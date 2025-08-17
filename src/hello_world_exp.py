# 导入我们需要的库
# 'machine' 库用于控制硬件，比如引脚(Pin)
# 'time' 库用于添加延时
import machine
import time

# 定义板载LED所连接的引脚
# 对于许多常见的 ESP32 开发板，板载LED连接到 GPIO 2
# 如果这个引脚不起作用，你可能需要查阅你具体板子的引脚图
led_pin = 2
led = machine.Pin(led_pin, machine.Pin.OUT)

# 打印一条消息到 Thonny 的 Shell 窗口，告诉我们程序已经开始
print(f"程序开始... 闪烁 GPIO {led_pin} 上的LED！")

# 创建一个无限循环，让LED不停地闪烁
# 你可以使用 Thonny 的 "停止" 按钮来中断程序
while True:
    try:
        # 点亮LED
        # led.value(1) 设置引脚为高电平
        led.value(1)
        print("LED ON") # 在Shell中打印状态
        
        # 等待1秒
        time.sleep(1)
        
        # 熄灭LED
        # led.value(0) 设置引脚为低电平
        led.value(0)
        print("LED OFF") # 在Shell中打印状态
        
        # 再次等待1秒
        time.sleep(1)
        
    except KeyboardInterrupt:
        # 如果用户在 Thonny 中按下了停止按钮 (Ctrl+C)，
        # 程序会捕获到 KeyboardInterrupt 异常
        print("程序已停止。")
        # 在退出前确保LED是熄灭的
        led.value(0)
        break
