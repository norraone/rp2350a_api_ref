# master.py - Step 1: Basic UART Communication
from machine import UART
import time

uart = UART(1, baudrate=9600, tx=17, rx=16, timeout=1000)

print("Master is ready and listening...")

while True:
    if uart.any():
        received_data = uart.readline()
        if received_data:
            message = received_data.decode('utf-8').strip()
            print(f"Received: '{message}'")
    
    time.sleep_ms(100)
