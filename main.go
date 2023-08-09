package main

import (
	"fmt"
	"image/color"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

var (
	// EUCLIDEAN_THRESHOLD is the euclidean distance between a pixel and a color
	// to be considered the same.
	EUCLIDEAN_THRESHOLD = []float64{0, 60, 60}
	// FREQUENCY is the frequency at which the program will check for calls.
	FREQUENCY = 10.0 // Hz

	// GREEN_ICON is the color of the green icon.
	GREEN_ICON = color.RGBA{99, 189, 102, 255} //rgb(99, 189, 102)
	// GREEN_ICON_LOCATION is the location of the green icon.
	GREEN_ICON_LOCATION = robotgo.Point{1736, 1033}

	// RED_ICON is the color of the red icon.
	RED_ICON = color.RGBA{227, 38, 52, 255}
	// RED_ICON_LOCATION is the location of the red icon.
	RED_ICON_LOCATION = robotgo.Point{1799, 1024}

	// BLUE_ICON is the color of the blue icon.
	BLUE_ICON = color.RGBA{92, 203, 255, 255}
	// BLUE_ICON_LOCATION is the location of the blue icon.
	BLUE_ICON_LOCATION = robotgo.Point{1655, 1037}
)

var callCount = 0

func euclideanColorDistance(c1, c2 color.RGBA) float64 {
	r := math.Pow(float64(c1.R-c2.R), 2)
	g := math.Pow(float64(c1.G-c2.G), 2)
	b := math.Pow(float64(c1.B-c2.B), 2)
	return math.Pow(r+g+b, 0.5)
}

func hexToRGB(hexColor string) color.RGBA {
	hexColor = strings.TrimPrefix(hexColor, "#")
	red, _ := strconv.ParseInt(hexColor[0:2], 16, 0)
	green, _ := strconv.ParseInt(hexColor[2:4], 16, 0)
	blue, _ := strconv.ParseInt(hexColor[4:6], 16, 0)
	rgbaColor := color.RGBA{uint8(red), uint8(green), uint8(blue), 255}
	return rgbaColor
}

func boolArrayToString(arr []bool) string {
	result := ""
	for _, value := range arr {
		if value {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}

func checkPixelColor(cords []robotgo.Point, check []color.RGBA) []bool {
	var re []bool
	fmt.Printf("\n")
	for i, cord := range cords {
		p := hexToRGB(robotgo.GetPixelColor(cord.X, cord.Y))

		// convert p from hex to rgb

		fmt.Printf("\x1B[48;2;%d;%d;%dm    \x1B[0m ", p.R, p.G, p.B)
		re = append(re, euclideanColorDistance(p, check[i]) <= EUCLIDEAN_THRESHOLD[i])
	}
	return re
}

func answerCallAndDecline() {
	fmt.Print(" \x1B[41mCall detected\x1B[0m")
	robotgo.KeyTap("enter")
	time.Sleep(50 * time.Millisecond)
	robotgo.KeyTap("backspace")
}

func declineCall() {
	fmt.Print(" \x1B[41mCall detected\x1B[0m")
	time.Sleep(50 * time.Millisecond)
	robotgo.KeyTap("backspace")
}

func main() {
	for {
		start_time := time.Now()
		time.Sleep(time.Duration(1/FREQUENCY) * time.Second)
		c := checkPixelColor([]robotgo.Point{BLUE_ICON_LOCATION, GREEN_ICON_LOCATION, RED_ICON_LOCATION}, []color.RGBA{{0, 0, 0, 255}, GREEN_ICON, RED_ICON})
		fmt.Printf(" %s %d %dms", boolArrayToString(c), callCount, time.Since(start_time).Milliseconds())
		if c[0] && c[1] && c[2] {
			callCount++
			declineCall()
		} else if c[0] && c[1] && !c[2] {
			callCount++
			answerCallAndDecline()
		}
	}
}
