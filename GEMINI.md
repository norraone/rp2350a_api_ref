# GEMINI Project Overview

This project is a MicroPython-based WiFi signal strength monitor that uses an SSD1309 OLED display.

## Project Overview

The main application scans for a specified WiFi network and displays the signal strength (RSSI) on the OLED screen. The display shows the target SSID, the current RSSI value, a trend chart of the RSSI over time, and a live bar graph of the current signal strength. A set of LEDs also provide a visual indication of the signal strength.

The project is structured as follows:

*   `src/main.py`: The main application entry point. It contains the logic for scanning for the WiFi network, updating the display, and controlling the LEDs.
*   `src/ssd1309.py`: A driver for the SSD1309 OLED display. It provides low-level control of the display via I2C or SPI.
*   `src/oled_logger.py`: A helper class for displaying scrolling text on the OLED display.
*   `ref/`: Contains reference materials, including C drivers for the SSD1309, and other documentation.
*   `release/`: Contains a release script and a release version of the `ssd1309.py` driver.

## Tools Used

*   **MicroPython:** The project is written in MicroPython.
*   **ampy:** The `ampy` tool is used to transfer files to the MicroPython board. See `ref/ampy.txt` for usage instructions.

## Building and Running

This is a MicroPython project, so it's intended to be run on a microcontroller with MicroPython firmware.

To run the project, you would typically:

1.  **Set up your hardware:** Connect the SSD1309 OLED display and LEDs to your MicroPython board as defined in the `src/main.py` file.
2.  **Configure the WiFi network:** Modify the `TARGET_SSID` variable in `src/main.py` to match the WiFi network you want to monitor.
3.  **Upload the code:** Use the `ampy` tool to upload the files from the `src` directory to your MicroPython board. For example:

    ```bash
    ampy --port /dev/tty.SLAB_USBtoUART put src/main.py
    ampy --port /dev/tty.SLAB_USBtoUART put src/ssd1309.py
    ```

4.  **Run the code:** The `main.py` script will run automatically on boot, or you can run it manually from the REPL.

## Project History

The `ref/experiment_log.md` file contains a detailed log of the project's development. Key takeaways include:

*   The project initially used a generic `ssd1306.py` driver, which caused display issues. A custom `ssd1309.py` driver was created based on C-language reference files to resolve the problem.
*   Hardware instability issues (I2C timeouts, device not found errors) were traced to a logic level mismatch between the 5V OLED and the 3.3V ESP32. The issue was resolved by powering the OLED from the ESP32's 3.3V pin.
*   Once the hardware was stable, the I2C frequency was increased to 400kHz for better performance.

This historical context is valuable for understanding the design decisions and troubleshooting steps taken during the project's development.