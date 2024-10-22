# Generate Test Scenario GitHub Action

This GitHub Action, named "Generate Test Scenario", is designed to generate a test scenario based on variants, lab number, and students' configurations.

## Inputs

- `variant_number`: Variant number (required)
- `lab_number`: Lab number (required)
- `wokwi_toml_path`: Path to the wokwi.toml file (required)

## Outputs

- `test_scenario`: Path to the generated test scenario file

## Steps

1. Run the parsing program with the provided flags:
   - `lab1-parse`
   - `-variant=${{ inputs.variant_number }}`
   - `-variants-path=${{ github.action_paths }}/variants/lab${{ inputs.lab_number }}.yml`
   - `-wokwi-config=${{ inputs.wokwi_toml_path }}`
   - `-output=test_scenario.yml`
   - Displays "Generated test scenario file: test_scenario.yml"

2. Set the output to the generated scenario file path
