package units

import (
	"fmt"
	"log"
	"strings"
)

type Converter interface {
	GetAllConversions(in string) string
}

type converter struct {
	logger *log.Logger
}

func GetConverter(l *log.Logger) converter {
	c := converter{
		logger: l,
	}
	return c
}

func weightToGrams(weight float64, units weightUnit) float64 {
	if units == Grams {
		return (weight)
	}
	return 0.0
}

func (c *converter) Convert(input string) (string, error) {

	grams, err := parseToGrams(input)
	if err != nil {
		return "", err
	} else {

		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("%v is %v grams\n", input, grams))
		sb.WriteString(fmt.Sprintf("%v is %v lbs\n", input, ConvertGtoLb(grams)))

		return sb.String(), nil
	}
}

func parseToGrams(s string) (float64, error) {
	value, units, err := parseMeasure(s)
	if err != nil {
		return 0.0, nil
	}
	switch units {
	case Grams:
		return value, nil

	}
	value = weightToGrams(value, units)
	return value, nil
}
