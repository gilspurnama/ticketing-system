package service

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var parkingLotSpace int
var emptyParkingLot int
var parkingVehicle []string

func MainParkingService(filename string) {
	lines, _ := ReadFile(filename)
	for _, line := range lines {
		commandLine := strings.Split(line, " ")
		switch commandLine[0] {
		case "park":
			response := ParkVehicle(commandLine[1])
			fmt.Println(response)
		case "create_parking_lot":
			parkingLotReq, _ := strconv.Atoi(commandLine[1])
			CreateParkingLot(parkingLotReq)
		case "leave":
			totalHours, _ := strconv.Atoi(commandLine[2])
			response := LeaveVehicle(commandLine[1], totalHours)
			fmt.Println(response)
		case "status":
			response := ParkingLotStatus()
			fmt.Println(response)
		default:
			fmt.Printf("Command is not recognized %s", commandLine[0])
		}
	}
}

func ErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	ErrorCheck(err)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func CreateParkingLot(parkingLotReq int) {
	parkingLotSpace = parkingLotReq
	emptyParkingLot = parkingLotSpace
}

func ParkVehicle(vehicleNumber string) string {
	if emptyParkingLot == 0 {
		return "Sorry, parking lot is full\n"
	}
	vehicleIndex := slices.Index(parkingVehicle, vehicleNumber)
	if vehicleIndex == -1 {
		parkingSpace := len(parkingVehicle)
		if parkingSpace == 0 {
			parkingVehicle = append(parkingVehicle, vehicleNumber)
			vehicleIndex = 1
		} else {
			emptyParkingSpot := slices.Index(parkingVehicle, "0")
			if emptyParkingSpot == -1 {
				parkingVehicle = append(parkingVehicle, vehicleNumber)
				newVehicleIndex := slices.Index(parkingVehicle, vehicleNumber)
				vehicleIndex = newVehicleIndex + 1
			} else {
				parkingVehicle[emptyParkingSpot] = vehicleNumber
				vehicleIndex = emptyParkingSpot + 1
			}
		}
		emptyParkingLot -= 1
		s := fmt.Sprintf("Allocated slot number: %d", vehicleIndex)
		return s
	} else {
		s := fmt.Sprintf("Registration number %s already inside the parking lot\n", vehicleNumber)
		return s
	}
}

func LeaveVehicle(vehicleNumber string, totalHours int) string {
	vehicleIndex := slices.Index(parkingVehicle, vehicleNumber)
	if vehicleIndex == -1 {
		s := fmt.Sprintf("Registration number %s not found\n", vehicleNumber)
		return s
	}
	parkingVehicle[vehicleIndex] = "0"
	totalCharge := 0
	if totalHours <= 2 {
		totalCharge = 10
	} else {
		totalCharge = 10 + 10*(totalHours-2)
	}
	emptyParkingLot += 1
	s := fmt.Sprintf("Registration number %s with Slot Number %d is free with Charge $%d\n", vehicleNumber, vehicleIndex+1, totalCharge)
	return s
}

func ParkingLotStatus() string {
	var sb strings.Builder
	sb.WriteString("Slot No. Registration No.\n")
	for i, vehicle := range parkingVehicle {
		if vehicle != "0" {
			fmt.Fprintf(&sb, "%d %s\n", i+1, vehicle)
		}
	}
	return sb.String()
}
