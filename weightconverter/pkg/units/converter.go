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

func (c *converter) Convert(input string) string {
	var sb strings.Builder

	grams, err := parseToGrams(input)
	if err != nil {
		c.logger.Printf("Failed to parse bad input: %v", input)
		return ""
	}

	sb.WriteString(fmt.Sprintf("%v is %v grams\n", input, grams))

	sb.WriteString(fmt.Sprintf("%v is %v lbs\n", input, ConvertGtoLb(grams)))

	return sb.String()
}

func parseToGrams(s string) (r float64, e error) {
	return 2000.0, nil
}
