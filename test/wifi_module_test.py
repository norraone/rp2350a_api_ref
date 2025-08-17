import network
import time

def scan():
    """Scans for WiFi networks and returns a list of results."""
    wlan = network.WLAN(network.STA_IF)
    wlan.active(True)
    
    # The scan can be slow, so we'll give it a moment
    time.sleep(2) 
    
    found_networks = wlan.scan()
    wlan.active(False) # Turn off WiFi radio to save power
    
    return found_networks