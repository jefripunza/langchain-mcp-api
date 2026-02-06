package tools

import (
	"mcp-server/types"
	"regexp"
	"strings"
)

func GetTextTools() []types.Tool {
	return []types.Tool{
		{
			Name:        "count_words",
			Description: "Menghitung jumlah kata dalam sebuah teks",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"text": map[string]interface{}{
						"type":        "string",
						"description": "Teks yang akan dihitung katanya",
					},
				},
				"required": []string{"text"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				text := args["text"].(string)
				trimmed := strings.TrimSpace(text)
				words := regexp.MustCompile(`\s+`).Split(trimmed, -1)

				wordCount := 0
				for _, word := range words {
					if len(word) > 0 {
						wordCount++
					}
				}

				noSpaces := regexp.MustCompile(`\s`).ReplaceAllString(text, "")

				return map[string]interface{}{
					"word_count":                wordCount,
					"character_count":           len(text),
					"character_count_no_spaces": len(noSpaces),
				}, nil
			},
		},
		{
			Name:        "reverse_text",
			Description: "Membalikkan urutan karakter dalam teks",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"text": map[string]interface{}{
						"type":        "string",
						"description": "Teks yang akan dibalik",
					},
				},
				"required": []string{"text"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				text := args["text"].(string)
				runes := []rune(text)
				for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
					runes[i], runes[j] = runes[j], runes[i]
				}
				return map[string]interface{}{
					"reversed": string(runes),
				}, nil
			},
		},
		{
			Name:        "to_uppercase",
			Description: "Mengubah teks menjadi huruf besar semua",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"text": map[string]interface{}{
						"type":        "string",
						"description": "Teks yang akan diubah",
					},
				},
				"required": []string{"text"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				text := args["text"].(string)
				return map[string]interface{}{
					"result": strings.ToUpper(text),
				}, nil
			},
		},
		{
			Name:        "to_lowercase",
			Description: "Mengubah teks menjadi huruf kecil semua",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"text": map[string]interface{}{
						"type":        "string",
						"description": "Teks yang akan diubah",
					},
				},
				"required": []string{"text"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				text := args["text"].(string)
				return map[string]interface{}{
					"result": strings.ToLower(text),
				}, nil
			},
		},
		{
			Name:        "to_title_case",
			Description: "Mengubah teks menjadi Title Case (huruf besar di awal kata)",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"text": map[string]interface{}{
						"type":        "string",
						"description": "Teks yang akan diubah",
					},
				},
				"required": []string{"text"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				text := args["text"].(string)
				return map[string]interface{}{
					"result": strings.Title(strings.ToLower(text)),
				}, nil
			},
		},
	}
}
