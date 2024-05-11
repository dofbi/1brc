package calc

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stats struct{
	Min float64
	Max float64
	Total float64
	Count int
}

func ReadFile(filepath string) ([]byte, error) {
	body, err := os.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func MinTemperature(temps[]float64) float64{
	min := temps[0]
	for _, temp := range temps {
		if (temp < min){
			min = temp
		}
	}
	return min
}

func MaxTemperature(temps[]float64) float64{
	max := temps[0]
	for _, temp := range temps {
		if (temp > max){
			max = temp
		}
	}
	return max
}

func AverageTemperature(temps[]float64) float64{
	total := 0.0
	for _, temp := range temps {
		total += temp
	}
	return total/float64(len(temps))

}

func CalcTemperature(filePath string)(map[string]Stats, error){

	data, err := ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	lignes := strings.Split(string(data), "\n")

	fmt.Println("lignes:", lignes)

	statsParVille := make(map[string]Stats)

	for _, ligne := range lignes {
		fmt.Println("ligne:", ligne)
		parts := strings.Split(ligne, ";")
		ville := parts[0]
		temperature, err := strconv.ParseFloat(parts[1], 64)

		if err != nil {
			return nil, err
		}

		stats := statsParVille[ville]

		if stats.Count == 0 || temperature < stats.Min {
			stats.Min = temperature
		}

		if stats.Count == 0 || temperature > stats.Max {
			stats.Max = temperature
		}

		stats.Total += temperature
		stats.Count++

		statsParVille[ville] = stats
	}

	return statsParVille, nil

}