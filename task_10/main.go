package main

import "fmt"

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
//15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
//градусов. Последовательность в подмножноствах не важна.
//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc

func main() {
	temperature := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groupedTemps := make(map[int][]float32)

	for _, temp := range temperature {
		key := roundToFormat(temp)
		groupedTemps[key] = append(groupedTemps[key], temp)
	}

	fmt.Println(groupedTemps)
}

func roundToFormat(val float32) int {
	// -25.4 -> -20 etc
	return int(val) / 10 * 10
}
