package calc

import (
	"os"
)

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