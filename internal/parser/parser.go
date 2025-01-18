package parser

import (
	"encoding/json"
	"os"

	"github.com/mattizzz241/cyclonedx-human-parser/internal/models"
)

func ParseCycloneDXFile(filePath string) (*models.CycloneDX, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cycloneDX models.CycloneDX
	err = json.Unmarshal(data, &cycloneDX)
	if err != nil {
		return nil, err
	}

	return &cycloneDX, nil
}
