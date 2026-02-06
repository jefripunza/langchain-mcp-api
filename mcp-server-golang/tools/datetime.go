package tools

import (
	"fmt"
	"mcp-server/types"
	"time"
)

func GetDatetimeTools() []types.Tool {
	return []types.Tool{
		{
			Name:        "get_current_time",
			Description: "Mendapatkan waktu saat ini",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"timezone": map[string]interface{}{
						"type":        "string",
						"description": "Timezone (default: Asia/Jakarta)",
					},
				},
				"required": []string{},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				timezone := "Asia/Jakarta"
				if tz, ok := args["timezone"].(string); ok && tz != "" {
					timezone = tz
				}

				loc, err := time.LoadLocation(timezone)
				if err != nil {
					loc = time.UTC
				}

				now := time.Now().In(loc)
				fmt.Printf("✅ MCP2 DateTime: %s\n", now.Format("2006-01-02 15:04:05"))

				return map[string]interface{}{
					"iso":       now.Format(time.RFC3339),
					"timestamp": now.Unix(),
					"timezone":  timezone,
					"formatted": now.Format("2006-01-02 15:04:05"),
				}, nil
			},
		},
		{
			Name:        "calculate_age",
			Description: "Menghitung umur berdasarkan tanggal lahir",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"birthdate": map[string]interface{}{
						"type":        "string",
						"description": "Tanggal lahir (format: YYYY-MM-DD)",
					},
				},
				"required": []string{"birthdate"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				birthdate := args["birthdate"].(string)
				birth, err := time.Parse("2006-01-02", birthdate)
				if err != nil {
					return nil, err
				}

				today := time.Now()
				age := today.Year() - birth.Year()
				if today.YearDay() < birth.YearDay() {
					age--
				}

				nextBirthday := time.Date(today.Year(), birth.Month(), birth.Day(), 0, 0, 0, 0, time.UTC)
				if nextBirthday.Before(today) {
					nextBirthday = nextBirthday.AddDate(1, 0, 0)
				}

				fmt.Printf("✅ MCP2 DateTime: %d years old\n", age)

				return map[string]interface{}{
					"age":           age,
					"birthdate":     birthdate,
					"next_birthday": nextBirthday.Format(time.RFC3339),
				}, nil
			},
		},
		{
			Name:        "add_days",
			Description: "Menambahkan atau mengurangi hari dari tanggal tertentu",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"date": map[string]interface{}{
						"type":        "string",
						"description": "Tanggal awal (format: YYYY-MM-DD atau ISO string)",
					},
					"days": map[string]interface{}{
						"type":        "number",
						"description": "Jumlah hari yang akan ditambahkan (negatif untuk mengurangi)",
					},
				},
				"required": []string{"date", "days"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				dateStr := args["date"].(string)
				days := int(args["days"].(float64))

				targetDate, err := time.Parse("2006-01-02", dateStr)
				if err != nil {
					targetDate, err = time.Parse(time.RFC3339, dateStr)
					if err != nil {
						return nil, err
					}
				}

				resultDate := targetDate.AddDate(0, 0, days)
				fmt.Printf("✅ MCP2 DateTime: %s\n", resultDate.Format("2006-01-02"))

				return map[string]interface{}{
					"original_date": dateStr,
					"days_added":    days,
					"result_date":   resultDate.Format(time.RFC3339),
					"formatted":     resultDate.Format("2006-01-02"),
				}, nil
			},
		},
		{
			Name:        "day_of_week",
			Description: "Mendapatkan hari dalam seminggu dari tanggal tertentu",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"date": map[string]interface{}{
						"type":        "string",
						"description": "Tanggal (format: YYYY-MM-DD atau ISO string)",
					},
				},
				"required": []string{"date"},
			},
			Handler: func(args map[string]interface{}) (interface{}, error) {
				dateStr := args["date"].(string)
				targetDate, err := time.Parse("2006-01-02", dateStr)
				if err != nil {
					targetDate, err = time.Parse(time.RFC3339, dateStr)
					if err != nil {
						return nil, err
					}
				}

				daysID := []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
				daysEN := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

				dayNum := int(targetDate.Weekday())
				fmt.Printf("✅ MCP2 DateTime: %s\n", daysID[dayNum])

				return map[string]interface{}{
					"date":        dateStr,
					"day_name_id": daysID[dayNum],
					"day_name_en": daysEN[dayNum],
					"day_number":  dayNum,
				}, nil
			},
		},
	}
}
