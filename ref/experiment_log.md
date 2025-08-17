# Experiment Log: ESP32 OLED Display Setup

**Date:** 2025-08-14

## Part 1: Initial Connection & Basic Test

### 1.1. Objective

The initial goal was to set up and test an ESP32 microcontroller with an external I2C OLED display (128x64, SSD1309 driver) using MicroPython.

### 1.2. Initial Setup

*   **Microcontroller:** ESP32
*   **Display:** 128x64 OLED, SSD1309 Driver
*   **Interface:** I2C (SDA: GPIO21, SCL: GPIO22)

### 1.3. Initial Success & Problems

After basic connectivity tests, a generic `ssd1306.py` driver was used. This led to a scrambled, "snowy" image on the screen. The issue was resolved by creating a new, dedicated `ssd1309.py` driver based on C-language reference files provided by the user. This highlighted the importance of using hardware-specific drivers over supposedly compatible ones.

--- 

## Part 2: Comprehensive Functionality & Stability Testing

### 2.1. Objective

After establishing a basic connection, the goal was to fully test all display features (graphics, text, inversion, scrolling) and optimize for a high refresh rate.

### 2.2. Debugging Hardware Instability

A comprehensive test script was created to cycle through various graphics and display mode commands. This immediately revealed significant instability.

*   **Problem 1: Script Crashing.** The script would fail and the screen would go blank when trying to run the graphics test. 
    *   **Analysis:** A `SyntaxError` was found in the diagnostic script due to an f-string, which was fixed. However, the core problem persisted.

*   **Problem 2: I2C Timeouts (`OSError: [Errno 116] ETIMEDOUT`).** When running a sequence of drawing commands, the I2C communication would time out. 
    *   **Analysis:** This indicated the display was not responding in time, likely because the I2C bus speed was too fast for the hardware setup.
    *   **Action:** The I2C frequency was lowered to a stable 100kHz. 

*   **Problem 3: No Device Found (`OSError: [Errno 19] ENODEV`).** This was the critical clue. Even at a slow speed, the display would sometimes fail to initialize or would disappear from the I2C bus entirely during operation.
    *   **Analysis:** This error proves the issue is not in the software, but is a fundamental hardware communication failure. The software is correctly reporting that the device has disconnected.

### 2.3. Final Solution: Hardware Power Correction

The combination of timeout and disconnection errors pointed to an unstable power supply or incorrect logic levels for the I2C communication lines (SDA/SCL).

*   **Root Cause:** The user, a hardware engineer, diagnosed that the OLED module was likely a 5V device, while the ESP32 operates at a 3.3V logic level. This mismatch was causing the I2C communication to be unreliable.
*   **Resolution:** The user connected the OLED's VCC pin to the ESP32's **3.3V** power supply. This aligned the logic levels and stabilized the hardware completely.

### 2.4. Optimization: High Refresh Rate

Once the hardware was stable, we could safely optimize for performance.

1.  **Increased I2C Speed:** The I2C frequency was increased from the safe 100kHz to a fast **400kHz**.
2.  **Optimized Animation Loop:** Delays were removed from the scrolling loop to allow it to run at the maximum speed.

This resulted in a final script that was both fully functional and highly performant.

### 2.5. Key Takeaways

*   `ETIMEDOUT` and `ENODEV` errors in I2C communication almost always point to underlying hardware issues (power, wiring, or speed), not software bugs.
*   Mismatched logic levels (e.g., 5V vs 3.3V) are a common cause of I2C instability.
*   Always stabilize the hardware before attempting to optimize software for performance.
*   Adding a decoupling capacitor across VCC and GND is a standard best practice to prevent power dips during current-intensive operations like a screen refresh.