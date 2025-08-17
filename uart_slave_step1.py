# slave.py - Step 1: Basic UART Communication
from machine import UART
import time

uart = UART(1, baudrate=9600, tx=17, rx=16)

print("Slave is ready and sending...")

count = 0
while True:
    count += 1
    message = f"Hello from slave, count: {count}\n"
    
    uart.write(message.encode('utf-8'))
    print(f"Sent: '{message.strip()}'")
    
    time.sleep(2)

