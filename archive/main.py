from PIL import ImageGrab
from time import sleep
import pydirectinput

EUCLIDEAN_THRESHOLD = [0,60,60]
FREQUENCY = 100   # Hz

GREEN_ICON = (100, 190, 101)
GREEN_ICON_LOCATION = (1736, 1033)

RED_ICON = ((227, 38, 52))
RED_ICON_LOCATION = (1799, 1024)

BLUE_ICON = (92, 203, 255)
BLUE_ICON_LOCATION = (1655, 1037)

call_count = 0

def screenshot():
    return ImageGrab.grab()


    
def euclidean_color_distance(c1, c2):
    r = (c1[0] - c2[0]) ** 2
    g = (c1[1] - c2[1]) ** 2
    b = (c1[2] - c2[2]) ** 2
    return (r + g + b) ** 0.5

def check_pixel_color(cords: list, check: list):
    r = []
    try:
        img = screenshot()
    except:
        print('Screenshot failed')
        return [False, False, False]
    print('')
    for i, cord in enumerate(cords):
        p = img.getpixel(cord)
        print(f'\x1B[48;2;{p[0]};{p[1]};{p[2]}m    \x1B[0m ', end='')

        r.append(euclidean_color_distance(p, check[i]) <= EUCLIDEAN_THRESHOLD[i])

    return r

def answer_call_and_decline():
    print(' \x1B[41mCall detected\x1B[0m', end='')
    pydirectinput.press('enter')
    pydirectinput.press('backspace')
    
def decline_call():
    print(' \x1B[41mCall detected\x1B[0m', end='')
    # sleep(1/4)
    pydirectinput.press('backspace')

while True:
    sleep(1/FREQUENCY)
    c = check_pixel_color([BLUE_ICON_LOCATION, GREEN_ICON_LOCATION, RED_ICON_LOCATION], [(0,0,0), GREEN_ICON, RED_ICON])

    print(''.join([str(int(i)) for i in c]), call_count, end='')

    # declinable call
    if c == [True, True, True]:
        call_count += 1
        decline_call()
    # forced call
    elif c == [True, True, False]:
        call_count += 1
        answer_call_and_decline()

