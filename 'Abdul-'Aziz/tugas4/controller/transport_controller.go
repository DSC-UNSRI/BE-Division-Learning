package controller

import (
	"fmt"

	"tugas/tugas4/infrastructure/input"
	"tugas/tugas4/usecase"
)

type TransportController struct {
    transportUseCase *usecase.TransportUseCase
    inputReader      input.InputReader
}

func NewTransportController(
    transportUC *usecase.TransportUseCase,
    reader input.InputReader,
) *TransportController {
    return &TransportController{
        transportUseCase: transportUC,
        inputReader:      reader,
    }
}

func (c *TransportController) ShowTransportMenu() (string, error) {
    fmt.Println("\nPilih Kendaraan:")
    fmt.Println("1. Kendaraan Pribadi")
    fmt.Println("2. Bus Kaleng")
    fmt.Println("3. Nebeng")
    fmt.Println("4. Travel")
    fmt.Print("Pilihan: ")
    
    choice := c.inputReader.ReadInt(1, 4)
    options := []string{
        "Kendaraan Pribadi",
        "Bus Kaleng", 
        "Nebeng",
        "Travel",
    }
    
    vehicle := options[choice-1]
    err := c.transportUseCase.UpdateVehicle(vehicle)
    if err != nil {
        return "", fmt.Errorf("Gagal update kendaraan: %v", err)
    }
    
    return vehicle, nil
}