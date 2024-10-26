package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	// Define the flags
	variantID := flag.String("variant", "", "ID of the variant (required)")
	variantsPath := flag.String("variants-path", "", "Path to the variants file (required)")
	wokwiConfig := flag.String("wokwi-config", "", "Path to the Wokwi configuration file (required)")
	rootPath := flag.String("root-path", "", "Path to the root directory")
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

	variant, err := getVariantById(*variantsPath, *variantID)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Variant: №%s, algo: %s, delay: %s", variant.ID, variant.AlgoVariantDescription.VariantNumber, variant.Delay)

	ap := variant.AlgoVariantDescription.FilePath
	if *rootPath != "" {
		ap = *rootPath + "/" + ap
	}

	algoVariant, err := getAlgoVariantByID(ap, variant.AlgoVariantDescription.VariantNumber)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Algo variant: %s", algoVariant.Steps)

	order := strings.Split(algoVariant.Steps, "->")

	labSetup, err := getLabSetupFromToml(*wokwiConfig)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}


	labConfig := make(map[string]interface{})
	labConfig["btn1"] = labSetup["btn1"]
	labConfig["led_delay"] = variant.Delay
	for i, step := range order {
		labConfig[fmt.Sprintf("led%d", i+1)] = labSetup[step]
	}
	log.Printf("Lab config: %v", labConfig)

	f, err := os.Create(*outputFile)
	if err != nil {
		log.Fatalf("Error creating test file: %v", err)
	}
	defer f.Close()

	tp := algoVariant.Template
	if *rootPath != "" {
		tp = *rootPath + "/" + algoVariant.Template
	}

	tmpl, err := template.ParseFiles(tp)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	err = tmpl.Execute(f, labConfig)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
