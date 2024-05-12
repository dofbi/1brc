package calc

import (
	"bytes"
	"fmt"
	"os"
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

func TestCalcTemperature(t *testing.T){
	tmpfile, err := os.CreateTemp("","example")
	if err != nil {
		t.Fatal(err)
	}

	text := []byte("Dakar;12\nTambacounda;8.9\nDakar;15\nTambacounda;10.0")

	if _, err := tmpfile.Write(text); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	statsParVille, err := CalcTemperature(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	stats := statsParVille["Dakar"]
	if stats.Min != 12.0 || stats.Max != 15.0 || stats.Total != 27.0 || stats.Count != 2 {
		t.Fatalf("Statistiques incorrectes pour Dakar: %+v", stats)
	}

	stats = statsParVille["Tambacounda"]
	if stats.Min != 8.9 || stats.Max != 10.0 || stats.Total != 18.9 || stats.Count != 2 {
		t.Fatalf("Statistiques incorrectes pour Tambacounda: %+v", stats)
	}

	os.Remove(tmpfile.Name())
}

func TestPrintStats(t *testing.T){

	statsParVille := map[string]Stats{
		"Dakar": {Min: 12.0, Max: 15.0, Total: 27.0, Count: 2},
		"Tambacounda": {Min: 8.9, Max: 10.0, Total: 18.9, Count: 2},
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintStats(statsParVille)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expected := "Dakar;12.0;13.5;15.0\nTambacounda;8.9;9.4;10.0\n"

	if output != expected {
		t.Fatalf("Expected %q, got %q", expected, output)
	}
}