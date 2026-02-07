package tools

import (
	"fmt"
	"mcp-server/types"
)

func GetConverterTools() []types.Tool {
	return []types.Tool{
		{
			Name:        "celsius_to_fahrenheit",
			Description: "Mengkonversi suhu dari Celsius ke Fahrenheit",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"celsius": map[string]interface{}{
						"type":        "number",
						"description": "Suhu dalam Celsius",
					},
				},
				"required": []string{"celsius"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				celsius := args["celsius"].(float64)
				fahrenheit := (celsius * 9 / 5) + 32
				kelvin := celsius + 273.15

				fmt.Printf("✅ MCP2 Converter c2f: %.2fC -> %.2fF\n", celsius, fahrenheit)

				return map[string]interface{}{
					"celsius":    celsius,
					"fahrenheit": fahrenheit,
					"kelvin":     kelvin,
				}, nil
			},
		},
		{
			Name:        "fahrenheit_to_celsius",
			Description: "Mengkonversi suhu dari Fahrenheit ke Celsius",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"fahrenheit": map[string]interface{}{
						"type":        "number",
						"description": "Suhu dalam Fahrenheit",
					},
				},
				"required": []string{"fahrenheit"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				fahrenheit := args["fahrenheit"].(float64)
				celsius := (fahrenheit - 32) * 5 / 9
				kelvin := celsius + 273.15

				fmt.Printf("✅ MCP2 Converter f2c: %.2fF -> %.2fC\n", fahrenheit, celsius)

				return map[string]interface{}{
					"fahrenheit": fahrenheit,
					"celsius":    celsius,
					"kelvin":     kelvin,
				}, nil
			},
		},
		{
			Name:        "km_to_miles",
			Description: "Mengkonversi jarak dari kilometer ke mil",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"km": map[string]interface{}{
						"type":        "number",
						"description": "Jarak dalam kilometer",
					},
				},
				"required": []string{"km"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				km := args["km"].(float64)
				miles := km * 0.621371
				meters := km * 1000
				feet := km * 3280.84

				fmt.Printf("✅ MCP2 Converter km2mil: %.2fkm -> %.2fmil\n", km, miles)

				return map[string]interface{}{
					"km":     km,
					"miles":  miles,
					"meters": meters,
					"feet":   feet,
				}, nil
			},
		},
		{
			Name:        "miles_to_km",
			Description: "Mengkonversi jarak dari mil ke kilometer",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"miles": map[string]interface{}{
						"type":        "number",
						"description": "Jarak dalam mil",
					},
				},
				"required": []string{"miles"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				miles := args["miles"].(float64)
				km := miles * 1.60934
				meters := miles * 1609.34
				feet := miles * 5280

				fmt.Printf("✅ MCP2 Converter mil2km: %.2fmil -> %.2fkm\n", miles, km)

				return map[string]interface{}{
					"miles":  miles,
					"km":     km,
					"meters": meters,
					"feet":   feet,
				}, nil
			},
		},
		{
			Name:        "kg_to_pounds",
			Description: "Mengkonversi berat dari kilogram ke pound",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"kg": map[string]interface{}{
						"type":        "number",
						"description": "Berat dalam kilogram",
					},
				},
				"required": []string{"kg"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				kg := args["kg"].(float64)
				pounds := kg * 2.20462
				grams := kg * 1000
				ounces := kg * 35.274

				fmt.Printf("✅ MCP2 Converter kg2lb: %.2fkg -> %.2flb\n", kg, pounds)

				return map[string]interface{}{
					"kg":     kg,
					"pounds": pounds,
					"grams":  grams,
					"ounces": ounces,
				}, nil
			},
		},
		{
			Name:        "pounds_to_kg",
			Description: "Mengkonversi berat dari pound ke kilogram",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"pounds": map[string]interface{}{
						"type":        "number",
						"description": "Berat dalam pound",
					},
				},
				"required": []string{"pounds"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				pounds := args["pounds"].(float64)
				kg := pounds * 0.453592
				grams := pounds * 453.592
				ounces := pounds * 16

				fmt.Printf("✅ MCP2 Converter lb2kg: %.2flb -> %.2fkg\n", pounds, kg)

				return map[string]interface{}{
					"pounds": pounds,
					"kg":     kg,
					"grams":  grams,
					"ounces": ounces,
				}, nil
			},
		},
	}
}
