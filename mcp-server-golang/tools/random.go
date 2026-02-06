package tools

import (
	"fmt"
	"math/rand"
	"mcp-server/types"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomTools() []types.Tool {
	return []types.Tool{
		{
			Name:        "random_number",
			Description: "Generate angka random dalam range tertentu",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"min": map[string]interface{}{
						"type":        "number",
						"description": "Nilai minimum",
					},
					"max": map[string]interface{}{
						"type":        "number",
						"description": "Nilai maximum",
					},
				},
				"required": []string{"min", "max"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				min := int(args["min"].(float64))
				max := int(args["max"].(float64))
				result := rand.Intn(max-min+1) + min

				return map[string]interface{}{
					"result": result,
					"min":    min,
					"max":    max,
				}, nil
			},
		},
		{
			Name:        "random_string",
			Description: "Generate string random dengan panjang tertentu",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"length": map[string]interface{}{
						"type":        "number",
						"description": "Panjang string yang diinginkan",
					},
					"type": map[string]interface{}{
						"type":        "string",
						"description": "Tipe karakter: alphanumeric, alphabetic, numeric",
						"enum":        []string{"alphanumeric", "alphabetic", "numeric"},
					},
				},
				"required": []string{"length"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				length := int(args["length"].(float64))
				charType := "alphanumeric"
				if t, ok := args["type"].(string); ok && t != "" {
					charType = t
				}

				var chars string
				switch charType {
				case "numeric":
					chars = "0123456789"
				case "alphabetic":
					chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
				default:
					chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
				}

				result := make([]byte, length)
				for i := 0; i < length; i++ {
					result[i] = chars[rand.Intn(len(chars))]
				}

				return map[string]interface{}{
					"result": string(result),
					"length": length,
					"type":   charType,
				}, nil
			},
		},
		{
			Name:        "coin_flip",
			Description: "Lempar koin virtual (heads atau tails)",
			Parameters: map[string]interface{}{
				"type":       "object",
				"properties": map[string]interface{}{},
				"required":   []string{},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				result := "tails"
				resultID := "Ekor"
				if rand.Float64() < 0.5 {
					result = "heads"
					resultID = "Kepala"
				}

				return map[string]interface{}{
					"result":    result,
					"result_id": resultID,
				}, nil
			},
		},
		{
			Name:        "dice_roll",
			Description: "Lempar dadu virtual",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"sides": map[string]interface{}{
						"type":        "number",
						"description": "Jumlah sisi dadu (default: 6)",
					},
					"count": map[string]interface{}{
						"type":        "number",
						"description": "Jumlah dadu yang dilempar (default: 1)",
					},
				},
				"required": []string{},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				sides := 6
				count := 1

				if s, ok := args["sides"].(float64); ok {
					sides = int(s)
				}
				if c, ok := args["count"].(float64); ok {
					count = int(c)
				}

				rolls := make([]int, count)
				total := 0
				for i := 0; i < count; i++ {
					roll := rand.Intn(sides) + 1
					rolls[i] = roll
					total += roll
				}

				return map[string]interface{}{
					"rolls": rolls,
					"total": total,
					"sides": sides,
					"count": count,
				}, nil
			},
		},
		{
			Name:        "random_color",
			Description: "Generate warna random dalam format hex",
			Parameters: map[string]interface{}{
				"type":       "object",
				"properties": map[string]interface{}{},
				"required":   []string{},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				color := rand.Intn(16777215)
				hex := fmt.Sprintf("#%06x", color)

				r := (color >> 16) & 0xFF
				g := (color >> 8) & 0xFF
				b := color & 0xFF

				return map[string]interface{}{
					"hex": hex,
					"rgb": map[string]interface{}{
						"r": r,
						"g": g,
						"b": b,
					},
				}, nil
			},
		},
	}
}
