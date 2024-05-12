package calc

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Stats struct{
	Min float64
	Max float64
	Total float64
	Count int
}

func (s Stats) String() string{
	moyenne := s.Total / float64(s.Count)
	return fmt.Sprintf("%.1f/%.1f/%.1f",s.Min,moyenne,s.Max)
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

	statsParVille := make(map[string]Stats)

	for _, ligne := range lignes {
		
		if strings.HasPrefix(ligne, "#") || strings.TrimSpace(ligne) == "" {
			continue
		}

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

func PrintStats(statsParVille map[string]Stats) error{
	file, err := os.Create("resultats.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	
	villes := make([]string, 0, len(statsParVille))

	for ville := range statsParVille{
		villes = append(villes, ville)
	}

	sort.Strings(villes)

	results := make([]string, len(villes))

	for i, ville := range villes {
		stats := statsParVille[ville]

		results[i] = fmt.Sprintf("%s;%s",ville,stats)
	}

	_, err = fmt.Fprintf(file,"{%s}\n", strings.Join(results, ", "))

	if err != nil {
		return err
	}

	return nil

}