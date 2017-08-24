package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

var percentageRegex = regexp.MustCompile(`percentage:\ *(\d*)%`)

// getBatteryPercentage returns the current percentage for the given battery
func getBatteryPercentage(batteryName string) int {
	path := "/org/freedesktop/UPower/devices/battery_" + batteryName
	cmd := exec.Command("upower", "-i", path)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	result := percentageRegex.FindAllSubmatch(output, 1)

	if len(result) != 1 {
		log.Fatal("error while parsing percentage value")
	}

	if len(result[0]) <= 0 {
		log.Fatal("error while parsing percentage value")
	}

	strValue := string(result[0][1])

	if value, err := strconv.Atoi(strValue); err == nil {
		return value
	}

	return 0
}

func main() {

	b0 := getBatteryPercentage("BAT0")
	b1 := getBatteryPercentage("BAT1")

	fmt.Printf("%d%%/%d%%\n", b1, b0)
}
