package main

import (
	"errors"
	"fmt"
)

const (
	min        = "minimo"
	max        = "maximo"
	prom       = "promedio"
	categoriaA = "A"
	categoriaB = "B"
	categoriaC = "C"
)

func main() {
	fmt.Println("Inicio")

	// Calcular min
	min := calcularMin(2, 5, 6, 7, 8, 9, 2)

	max := calcularMax(2, 5, 6, 7, 8, 9, 2)

	prom := calcularProm(2, 5, 6, 7, 8, 9, 2)

	minFunc, err := funcionFactory("minimo")
	if err != nil {
		fmt.Printf(err.Error())
	}

	maxFunc, err := funcionFactory("maximo")
	if err != nil {
		fmt.Printf(err.Error())
	}

	promFunc, err := funcionFactory("promedio")
	if err != nil {
		fmt.Printf(err.Error())
	}
	promFactory := promFunc(2, 5, 6, 7, 8, 9, 2)
	minFactory := minFunc(2, 5, 6, 7, 8, 9, 2)
	maxFactory := maxFunc(2, 5, 6, 7, 8, 9, 2)

	fmt.Printf("El minimo sin factory es %d y el minimo con factory %d\n", min, minFactory)
	fmt.Printf("El minimo sin factory es %d y el minimo con factory %d\n", max, maxFactory)
	fmt.Printf("El minimo sin factory es %d y el minimo con factory %d\n", prom, promFactory)

	fmt.Println(calcularSalario(1000, "A"))
	fmt.Println(calcularSalario(1000, "B"))
	fmt.Println(calcularSalario(1000, "C"))

}

func calcularMin(nums ...int) int {
	// Programacion defensiva hacer validaciones
	if len(nums) == 0 {
		return 0
	}

	// Logica de negocio
	min := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	// Retorno
	return min
}

func calcularMax(nums ...int) int {
	// Programacion defensiva hacer validaciones
	if len(nums) == 0 {
		return 0
	}

	// Logica de negocio
	max := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	// Retorno
	return max
}

func calcularProm(nums ...int) int {
	// Programacion defensiva hacer validaciones
	if len(nums) == 0 {
		return 0
	}

	// Logica de negocio
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}

	// Retorno
	return sum / len(nums)
}

func funcionFactory(operacion string) (func(...int) int, error) {
	switch operacion {
	case min:
		return calcularMin, nil
	case max:
		return calcularMax, nil
	case prom:
		return calcularProm, nil
	default:
		return nil, errors.New("no existe la operacion")
	}

}

func calcularSalario(minutosTrabajados int, categoria string) float64 {
	switch categoria {
	case categoriaC:
		return float64(minutosTrabajados / 60 * 1000)
	case categoriaB:
		return float64(1500*(minutosTrabajados/60)) * 1.2
	case categoriaA:
		return float64(3000*(minutosTrabajados/60)) * 1.5
	default:
		return 0
	}
}
