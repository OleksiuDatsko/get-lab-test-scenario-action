# Generate Test Scenario GitHub Action

 This GitHub Action, named "Generate Test Scenario", is designed to generate a test scenario based on variants, lab number, and students' configurations.

## Inputs

- `variant_number`: Variant number (required)
- `lab_number`: Lab number (required)
- `wokwi_toml_path`: Path to the wokwi.toml file (required)
- `output`: Output file (optional, defaults to "test_scenario.yml")

## Outputs

- `test_scenario`: Path to the generated test scenario file

## Steps

1. Run the parsing program with the provided flags:
   Lab 1 parsing program:
   ```sh
   Usage: lab1-parser [options]
   Options:
   -output string
         Path to the output file (required)
   -root-path string
         Path to the root directory
   -variant string
         ID of the variant (required)
   -variants-path string
         Path to the variants file (required)
   -wokwi-config string
         Path to the Wokwi configuration file (required)
   ```

2. Set the output to the generated scenario file path
