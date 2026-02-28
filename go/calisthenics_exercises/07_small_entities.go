package calisthenics_exercises

import (
	"encoding/json"
	"fmt"
)

// Regla 7: Mantener las unidades de código pequeñas
// Cuando más grande es una unidad de código, más probable es que esté haciéndose cargo
// de varias responsabilidades y tomando muchas decisiones, lo que aumenta la complejidad
// ciclomática y el riesgo de introducir errores.

// Ejercicio: Descompón el método GenerateCsvReportFromJson en unidades más pequeñas.

type ReportService struct {
	cache map[string]string
}

func NewReportService() *ReportService {
	return &ReportService{
		cache: make(map[string]string),
	}
}

func (r *ReportService) GenerateCsvReportFromJson(jsonInput string, delimiter string) (string, error) {
	if delimiter == "" {
		delimiter = ","
	}

	var data interface{}
	if err := json.Unmarshal([]byte(jsonInput), &data); err != nil {
		return "", fmt.Errorf("Invalid JSON")
	}

	arr, ok := data.([]interface{})
	if !ok {
		return "", fmt.Errorf("Expected array")
	}

	if len(arr) == 0 {
		return "", nil
	}

	firstRow := arr[0].(map[string]interface{})
	headers := []string{}
	for k := range firstRow {
		headers = append(headers, k)
	}

	lines := []string{joinStrings(headers, delimiter)}
	for _, item := range arr {
		row := item.(map[string]interface{})
		values := []string{}
		for _, h := range headers {
			val := row[h]
			if val == nil {
				values = append(values, "")
			} else {
				values = append(values, fmt.Sprintf("%v", val))
			}
		}
		lines.append(joinStrings(values, delimiter))
	}

	result := joinLines(lines)
	r.cache["last"] = result
	return result, nil
}

func joinStrings(strs []string, delimiter string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += delimiter + strs[i]
	}
	return result
}

func joinLines(lines []string) string {
	if len(lines) == 0 {
		return ""
	}
	result := lines[0]
	for i := 1; i < len(lines); i++ {
		result += "\n" + lines[i]
	}
	return result
}
