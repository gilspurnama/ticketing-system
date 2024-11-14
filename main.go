package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var parkingLot = 0
var emptyParkingLot = 0
var parkingVehicle []string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	check(err)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	fmt.Println("This is line testing")
	lines, err := readLines("./ticket_file.txt")
	if err != nil {
		fmt.Errorf("readLines: %s", err)
	}
	for _, line := range lines {
		commandLine := strings.Split(line, " ")
		switch commandLine[0] {
		case "park":
			park(commandLine[1])
		case "create_parking_lot":
			parkingLot, _ = strconv.Atoi(commandLine[1])
			emptyParkingLot = parkingLot
		case "leave":
			totalHours, _ := strconv.Atoi(commandLine[2])
			leave(commandLine[1], totalHours)
		case "status":
			status()
		default:
			fmt.Printf("Command is not recognized %s", commandLine[0])
		}
	}
}

func park(vehicleNumber string) {
	if emptyParkingLot == 0 {
		fmt.Println("Sorry, parking lot is full")
		return
	}
	vehicleIndex := slices.Index(parkingVehicle, vehicleNumber)
	if vehicleIndex == -1 {
		parkingSpace := len(parkingVehicle)
		if parkingSpace == 0 {
			parkingVehicle = append(parkingVehicle, vehicleNumber)
			fmt.Println("Allocated slot number: 1")
		} else {
			emptyParkingSpot := slices.Index(parkingVehicle, "0")
			if emptyParkingSpot == -1 {
				parkingVehicle = append(parkingVehicle, vehicleNumber)
				newVehicleIndex := slices.Index(parkingVehicle, vehicleNumber)
				fmt.Printf("Allocated slot number: %d\n", newVehicleIndex+1)
			} else {
				parkingVehicle[emptyParkingSpot] = vehicleNumber
				fmt.Printf("Allocated slot number: %d\n", emptyParkingSpot+1)
				fmt.Println()
			}
		}
		emptyParkingLot -= 1
	} else {
		fmt.Printf("Registration number %s not found\n", vehicleNumber)
	}

}

func leave(vehicleNumber string, totalHours int) {
	vehicleIndex := slices.Index(parkingVehicle, vehicleNumber)
	if vehicleIndex == -1 {
		fmt.Printf("Registration number %s not found\n", vehicleNumber)
		return
	}
	parkingVehicle[vehicleIndex] = "0"
	totalCharge := 0
	if totalHours <= 2 {
		totalCharge = 10
	} else {
		totalCharge = 10 + 10*(totalHours-2)
	}
	fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n", vehicleNumber, vehicleIndex+1, totalCharge)
	emptyParkingLot += 1
}

func status() {
	fmt.Println("Slot No. Registration No.")
	for i, vehicle := range parkingVehicle {
		if vehicle != "0" {
			fmt.Printf("%d %s\n", i+1, vehicle)
		}
	}
}
