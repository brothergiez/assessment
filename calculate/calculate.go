package calculate

import (
	"fmt"
	"strconv"
)

func getIntFromMap(item map[string]interface{}, key string) (int, error) {
	value, ok := item[key]
	if !ok {
		return 0, fmt.Errorf("%s not found in item", key)
	}

	switch v := value.(type) {
	case string:
		intVal, err := strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("error converting %s to integer: %v", key, err)
		}
		return intVal, nil
	case int:
		return v, nil
	default:
		return 0, fmt.Errorf("invalid type for %s", key)
	}
}

func CalculateTotal(items []map[string]interface{}) (int, error) {
	total := 0
	for _, item := range items {
		price, err := getIntFromMap(item, "price")
		if err != nil {
			return 0, err
		}

		quantity, err := getIntFromMap(item, "quantity")
		if err != nil {
			return 0, err
		}

		total += price * quantity
	}
	return total, nil
}
