package domain_helpers

import "fmt"

func ParseFloat(str, fieldName string) (float64, error) {
	var f float64
	_, err := fmt.Sscanf(str, "%f", &f)
	if err != nil {
		return 0, fmt.Errorf("invalid %s value: %v", fieldName, err)
	}
	return f, nil
}
