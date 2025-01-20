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

	// Showing the Main components detected by the SBOM and then all the dependencies
	fmt.Printf("Parsed SBOM File - Format: %s, SpecVersion: %s\n", cyclonedx.BomFormat, cyclonedx.SpecVersion)
	for _, component := range cyclonedx.Components {
		fmt.Printf("Main Component: %s - %s\n", component.Name, component.Version)
	}
	fmt.Println("---")
	for _, dependency := range cyclonedx.Dependencies {
		fmt.Printf("%s\n", dependency.Ref)
		for _, childDependency := range dependency.DependsOn {
			fmt.Printf("\t%s\n", childDependency)
		}
	}
}
