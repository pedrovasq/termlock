package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	
	"termlock/internal/models"
)

func ImportCSV(path string) ([]models.PasswordEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	if len(records) < 1 {
		return nil, fmt.Errorf("CSV is empty")
	}

	var entries []models.PasswordEntry
	for i, row := range records {
		if i == 0 {
			continue
		}
		if len(row) < 5 {
			continue
		}

		entries = append(entries, models.PasswordEntry {
			Title:		row[0],
			Sites:		[]string{row[1]},
			Username:	row[2],
			Password:	row[3],
			Note:		row[4],
		})
	}
	return entries, nil
}

