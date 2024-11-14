package main

import (
	"ticketing-system/service"
)

func main() {
	service.MainParkingService("./ticket_file.txt")
}
