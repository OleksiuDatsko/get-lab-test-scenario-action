package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type AlgoVariantDescription struct {
	VariantNumber string `yaml:"variant"`
	FilePath      string `yaml:"file"`
}

type Variant struct {
	ID                     string `yaml:"id"`
	AlgoVariantDescription `yaml:"algo"`
	Delay                  string `yaml:"delay"`
}

type VariantsData struct {
	Variants []Variant `yaml:"variants"`
}

type AlgoVariant struct {
	ID       string `yaml:"id"`
	Template string `yaml:"template"`
	Steps    string `yaml:"steps"`
}

type AlgoVariantsData struct {
	Variants []AlgoVariant `yaml:"variants"`
}

func getAlgoVariantByID(filePath string, variantID string) (AlgoVariant, error) {
	// Read the YAML file
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return AlgoVariant{}, fmt.Errorf("could not read file: %v", err)
	}

	// Parse the YAML data
	var data AlgoVariantsData
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return AlgoVariant{}, fmt.Errorf("error parsing YAML: %v", err)
	}

	// Search for the variant by ID
	for _, variant := range data.Variants {
		if variant.ID == variantID {
			return variant, nil
		}
	}

	return AlgoVariant{}, fmt.Errorf("variant with ID %s not found", variantID)
}

func getVariantById(filePath string, variantID string) (Variant, error) {
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

func getLabSetupFromToml(tomlPath string) (map[string]interface{}, error) {
	var config map[string]interface{}

	_, err := toml.DecodeFile(tomlPath, &config)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}
	labSetup := config["lab_setup"].(map[string]interface{})

	return labSetup, nil
}
