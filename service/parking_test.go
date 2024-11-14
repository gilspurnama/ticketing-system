package service

import (
	"strings"
	"testing"
)

var testRegistrationNumber_1 = "XX-11-YY-1111"
var testRegistrationNumber_2 = "XX-22-YY-2222"
var testRegistrationNumber_3 = "XX-33-YY-3333"

func TestMainFunc(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		t.Cleanup(resetValue)
		MainParkingService("../ticket_file_test.txt")
	})

	t.Run("Panic", func(t *testing.T) {
		t.Cleanup(resetValue)
		defer func() {
			if r := recover(); r == nil {
				t.Error("function should be panic")
			}
		}()
		MainParkingService("../ticket_file123.txt")
	})
}

func TestReadFile(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		t.Cleanup(resetValue)
		_, err := ReadFile("../ticket_file_test.txt")
		if err != nil {
			t.Errorf("Result was incorrect")
		}
	})

	t.Run("Panic", func(t *testing.T) {
		t.Cleanup(resetValue)
		defer func() {
			if r := recover(); r == nil {
				t.Error("function should be panic")
			}
		}()
		ReadFile("../ticket_file123.txt")
	})

}

func TestCreateParkingSpace(t *testing.T) {
	t.Cleanup(resetValue)
	CreateParkingLot(5)
	if parkingLotSpace != 5 {
		t.Error("Result was incorrect")
	}
}

func TestParkVehicle(t *testing.T) {
	t.Run("Full Parking Lot", func(t *testing.T) {
		t.Cleanup(resetValue)
		emptyParkingLot = 0
		response := ParkVehicle(testRegistrationNumber_1)
		if !strings.Contains(response, "parking lot is full") {
			t.Error("Result was incorrect, parking should be full")
		}
	})

	t.Run("Vehicle Already Inside", func(t *testing.T) {
		t.Cleanup(resetValue)
		emptyParkingLot = 1
		parkingVehicle = append(parkingVehicle, testRegistrationNumber_1)
		response := ParkVehicle(testRegistrationNumber_1)
		if !strings.Contains(response, "already inside the parking lot") {
			t.Error("Result was incorrect, registered vehicle should already be inside")
		}
	})

	t.Run("Empty Parking Lot", func(t *testing.T) {
		t.Cleanup(resetValue)
		emptyParkingLot = 1
		parkingLotSpace = 1
		response := ParkVehicle(testRegistrationNumber_1)
		if !strings.Contains(response, "Allocated slot number: 1") {
			t.Errorf("Result was incorrect, got: %s, want: Allocated slot number: 1", response)
		}
	})

	t.Run("Park at slot 2", func(t *testing.T) {
		t.Cleanup(resetValue)
		emptyParkingLot = 1
		parkingLotSpace = 2
		parkingVehicle = append(parkingVehicle, testRegistrationNumber_1)
		response := ParkVehicle(testRegistrationNumber_2)
		if !strings.Contains(response, "Allocated slot number: 2") {
			t.Errorf("Result was incorrect, got: %s, want: Allocated slot number: 2", response)
		}
	})

	t.Run("Empty On Nearest Exit", func(t *testing.T) {
		t.Cleanup(resetValue)
		emptyParkingLot = 1
		parkingLotSpace = 2
		parkingVehicle = append(parkingVehicle, testRegistrationNumber_1, "0", testRegistrationNumber_2)
		response := ParkVehicle(testRegistrationNumber_3)
		if !strings.Contains(response, "Allocated slot number: 2") {
			t.Errorf("Result was incorrect, got: %s, want: Allocated slot number: 2", response)
		}
	})
}

func TestLeaveVehicle(t *testing.T) {
	t.Run("Vehicle Not Found", func(t *testing.T) {
		t.Cleanup(resetValue)
		parkingVehicle = append(parkingVehicle, testRegistrationNumber_1)
		response := LeaveVehicle(testRegistrationNumber_2, 4)
		if !strings.Contains(response, "not found") {
			t.Errorf("Result was incorrect, got: %s, want: Registration number %s not found\n", response, testRegistrationNumber_2)
		}
	})

	t.Run("Leave Before 2 Hours", func(t *testing.T) {
		t.Cleanup(resetValue)
		parkingVehicle = append(parkingVehicle, testRegistrationNumber_1)
		response := LeaveVehicle(testRegistrationNumber_1, 1)
		if !strings.Contains(response, "$10") {
			t.Errorf("Result was incorrect, got: %s, want: Registration number %s with Slot Number %d is free with Charge $%d\n", response, testRegistrationNumber_1, 1, 10)
		}
	})

	t.Run("Leave After 2 Hours", func(t *testing.T) {
		t.Cleanup(resetValue)
		parkingVehicle = append(parkingVehicle, testRegistrationNumber_1)
		response := LeaveVehicle(testRegistrationNumber_1, 3)
		if !strings.Contains(response, "$20") {
			t.Errorf("Result was incorrect, got: %s, want: Registration number %s with Slot Number %d is free with Charge $%d\n", response, testRegistrationNumber_1, 1, 20)
		}
	})
}

func TestStatus(t *testing.T) {
	t.Run("Check Status", func(t *testing.T) {
		t.Cleanup(resetValue)
		parkingVehicle = append(parkingVehicle, testRegistrationNumber_1, testRegistrationNumber_2)
		response := ParkingLotStatus()
		lineCount := len(strings.Split(response, "\n"))
		if lineCount != 4 {
			t.Errorf("Result was incorrect, got: %d, want: 4", lineCount)
		}
	})
}

func resetValue() {
	parkingLotSpace = 0
	emptyParkingLot = 0
	parkingVehicle = parkingVehicle[:0]
}
