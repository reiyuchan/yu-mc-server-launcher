package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

//go:generate goversioninfo
func main() {
	var maxMemory, minMemory string
	var max, min uint
	var noSize bool = false
	fmt.Println(`

___  _ _       _____ ____    _     ____  _     _      ____  _     _____ ____ 
\  \/// \ /\  /  __//  _ \  / \   /  _ \/ \ /\/ \  /|/   _\/ \ /|/  __//  __\
 \  / | | ||  | |  _| / \|  | |   | / \|| | ||| |\ |||  /  | |_|||  \  |  \/|
 / /  | \_/|  | |_//| \_/|  | |_/\| |-||| \_/|| | \|||  \_ | | |||  /_ |    /
/_/   \____/  \____\\____/  \____/\_/ \|\____/\_/  \|\____/\_/ \|\____\\_/\_\
                                                                             `)

	regex, _ := regexp.Compile("[a-zA-Z]")
	_, err := os.Stat("server.jar")
	if err != nil {
		fmt.Println("server file doesn't exist...")
		Pause()
	}
	fmt.Print("Enter the amount of max memory(ex: 4000M or 2G..G => Gigabytes..M => MegaBytes): ")
	fmt.Scanln(&maxMemory)
	fmt.Print("Enter the amount of min memory(ex: 4000M or 2G..G => Gigabytes..M => MegaBytes): ")
	fmt.Scanln(&minMemory)
	if strings.Contains(maxMemory, "G") || strings.Contains(minMemory, "G") {
		mxV, err := strconv.Atoi(strings.Split(maxMemory, "G")[0])
		if err != nil {
			fmt.Println("Not a valid number")
			Pause()
		}
		max = uint(math.Abs(float64(mxV)))
		msV, err := strconv.Atoi(strings.Split(minMemory, "G")[0])
		if err != nil {
			fmt.Println("Not a valid number")
			Pause()
		}
		min = uint(math.Abs(float64(msV)))
		if (max < 1 || min < 1) && !noSize {
			fmt.Println("min memory should be at least 1024M megabytes")
			Pause()
		}
	} else if strings.Contains(maxMemory, "g") || strings.Contains(minMemory, "g") {
		mxV, err := strconv.Atoi(strings.Split(maxMemory, "g")[0])
		if err != nil {
			fmt.Println("Not a valid number")
			Pause()
		}
		max = uint(math.Abs(float64(mxV)))
		msV, err := strconv.Atoi(strings.Split(minMemory, "g")[0])
		if err != nil {
			fmt.Println("Not a valid number")
			Pause()
		}
		min = uint(math.Abs(float64(msV)))
		if (max < 1 || min < 1) && !noSize {
			fmt.Println("min memory should be at least 1024M megabytes")
			Pause()
		}
	} else if strings.Contains(maxMemory, "M") || strings.Contains(minMemory, "M") {
		mxV, err := strconv.Atoi(strings.Split(maxMemory, "M")[0])
		if err != nil {
			fmt.Println("Not a valid number")
			Pause()
		}
		max = uint(math.Abs(float64(mxV)))
		msV, err := strconv.Atoi(strings.Split(minMemory, "M")[0])
		if err != nil {
			fmt.Println("Not a valid number")
			Pause()
		}
		min = uint(math.Abs(float64(msV)))
		if (max < 1024 || min < 1024) && !noSize {
			fmt.Println("min memory should be at least 1024M megabytes")
			Pause()
		}
	} else if strings.Contains(maxMemory, "m") || strings.Contains(minMemory, "m") {
		mxV, err := strconv.Atoi(strings.Split(maxMemory, "m")[0])
		if err != nil {
			fmt.Println("Not a valid number")
			Pause()
		}
		max = uint(math.Abs(float64(mxV)))
		msV, err := strconv.Atoi(strings.Split(minMemory, "m")[0])
		if err != nil {
			fmt.Println("Not a valid number")
			Pause()
		}
		min = uint(math.Abs(float64(msV)))
		if (max < 1024 || min < 1024) && !noSize {
			fmt.Println("min memory should be at least 1024M megabytes")
			Pause()
		}
	} else {
		if regex.MatchString(maxMemory) || regex.MatchString(minMemory) {
			fmt.Println("Not a valid number...")
		} else {
			fmt.Println("Values should have a size...(ex: M or G)")
		}
		noSize = true
		Pause()
	}

	cmd := exec.Command("java", "-Xmx"+maxMemory, "-Xms"+minMemory, "-jar", "server.jar", "nogui")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func Pause() {
	fmt.Print("Press enter to continue...")
	fmt.Scanln()
	os.Exit(1)
}
