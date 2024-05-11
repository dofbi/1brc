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
