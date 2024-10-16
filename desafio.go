package main

import "fmt"

const ebulicaoKelvin = 373.15

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func main() {
	tempKelvin := ebulicaoKelvin

	celsius := kelvinToCelsius(tempKelvin)
	fmt.Printf("A temperatura em Celsius é: %.2f\n", celsius)
}
