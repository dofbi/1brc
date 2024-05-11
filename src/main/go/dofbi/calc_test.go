package calc

import (
	"fmt"
	"testing"
)

func TestFileReading(t *testing.T) {
	filePath := "../../../../data/weather_stations.csv"

	data, err := ReadFile(filePath)

	if err != nil {
		t.Fatal(err)
	}

	if len(data) == 0 {
		t.Fatalf("Le fichier %s est vide", filePath)
	}

	fmt.Println("Ligne: ",len(data))
}

func TestMinTemperature(t *testing.T){
	temps := []float64{12.0, 8.9, 38.8}
	min := MinTemperature(temps)

	if min != 8.9 {
		t.Fatalf("Expected 8.9, got %f", min)
	}
}

func TestMaxTemperature(t *testing.T){
	temps := []float64{12.0, 8.9, 38.8}
	min := MaxTemperature(temps)

	if min != 38.8 {
		t.Fatalf("Expected 38.8, got %f", min)
	}
	
}

func TestAverageTemperature(t *testing.T){
	temps := []float64{12.0, 8.9, 38.8}
	avg := AverageTemperature(temps)
	expected :=(12.0 + 8.9 + 38.8)/3

	if avg != expected {
		t.Fatalf("Expected %f, got %f", expected, avg)
	}
}