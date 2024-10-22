package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type Variant struct {
	ID       string `yaml:"id"`
	Template string `yaml:"template"`
	Steps    string `yaml:"steps"`
}

type VariantsData struct {
	Variants []Variant `yaml:"variants"`
}

func getVariantByID(filePath string, variantID string) (Variant, error) {
	// Read the YAML file
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return Variant{}, fmt.Errorf("could not read file: %v", err)
	}

	// Parse the YAML data
	var data VariantsData
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return Variant{}, fmt.Errorf("error parsing YAML: %v", err)
	}

	// Search for the variant by ID
	for _, variant := range data.Variants {
		if variant.ID == variantID {
			return variant, nil
		}
	}

	return Variant{}, fmt.Errorf("variant with ID %s not found", variantID)
}

func main() {
	// Define the flags
	variantID := flag.String("variant", "", "ID of the variant (required)")
	variantsPath := flag.String("variants-path", "", "Path to the variants file (required)")
	wokwiConfig := flag.String("wokwi-config", "", "Path to the Wokwi configuration file (required)")
	outputFile := flag.String("output", "", "Path to the output file (required)")

	// Display help message if no flags are provided
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	// Parse the flags
	flag.Parse()

	// Validate required flags
	if *variantID == "" || *variantsPath == "" || *wokwiConfig == "" || *outputFile == "" {
		fmt.Println("Error: All flags are required.")
		flag.Usage()
		os.Exit(1)
	}

	variant, err := getVariantByID(*variantsPath, *variantID)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Variant: %v", variant)

	order := strings.Split(variant.Steps, "->")

	log.Printf("Order: %v", order)

	var config map[string]interface{}

	_, err = toml.DecodeFile(*wokwiConfig, &config)
	if err != nil {
		log.Fatalf("Error loading TOML file: %v", err)
	}

	labSetup := config["lab_setup"].(map[string]interface{})

	log.Printf("Lab setup: %v", labSetup)

	labConfig := make(map[string]interface{})
	labConfig["btn1"] = labSetup["btn1"]
	labConfig["led_delay"] = labSetup["led_delay"]
	for i, step := range order {
		labConfig[fmt.Sprintf("led%d", i+1)] = labSetup[step]
	}
	log.Printf("Lab config: %v", labConfig)

	f, err := os.Create(*outputFile)
	if err != nil {
		log.Fatalf("Error creating test file: %v", err)
	}
	defer f.Close()

	tmpl, err := template.ParseFiles(variant.Template)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	err = tmpl.Execute(f, labConfig)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
