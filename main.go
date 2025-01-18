package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mattizzz241/cyclonedx-human-parser/internal/parser"
)

func main() {
	filePath := os.Args[1]

	cyclonedx, err := parser.ParseCycloneDXFile(filePath)
	if err != nil {
		log.Fatalf("Error parsing CycloneDX file: %v", err)
	}

	fmt.Printf("Parsed CycloneDX BOM Format: %s\n", cyclonedx.BomFormat)
	for _, component := range cyclonedx.Components {
		fmt.Printf("Component: %s, Version: %s, Type: %s\n", component.Name, component.Version, component.Type)
	}
}
