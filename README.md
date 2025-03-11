# Countries

A Go package that provides ISO country code utilities, supporting both ISO 3166-1 alpha-2 (two-letter) and ISO 3166-1 alpha-3 (three-letter) country codes.

## Features

- Complete list of ISO 3166-1 alpha-2 country codes (two-letter)
- Complete list of ISO 3166-1 alpha-3 country codes (three-letter)
- Complete list of US state and territory codes (two-letter)

## Installation

```bash
go get github.com/KGKallasmaa/countries
```

## Usage

```go
import (
    "fmt"
    "github.com/KGKallasmaa/countries"
)

func main() {
    // Using ISO-2 country codes
    countryCode := countries.ISO2CountryCodeUS
    fmt.Println(countryCode) // Outputs: US

    // Using ISO-3 country codes
    countryCode3 := countries.ISO3CountryCodeUSA
    fmt.Println(countryCode3) // Outputs: USA
    
    // Using US state codes
    stateCode := countries.USStateCodeCA
    fmt.Println(stateCode) // Outputs: CA (California)
}
```

## Available Constants

The package provides constants for all ISO 3166-1 country codes:

- ISO-2 (alpha-2) codes: e.g., `ISO2CountryCodeUS` for United States
- ISO-3 (alpha-3) codes: e.g., `ISO3CountryCodeUSA` for United States

## US States

The package includes constants for all US states and territories using the standard ANSI/FIPS two-letter codes:

- All 50 US states: e.g., `USStateCodeCA` for California, `USStateCodeNY` for New York
- US territories: e.g., `USStateCodePR` for Puerto Rico
- Federal district: `USStateCodeDC` for District of Columbia

You can also access all US state codes through the `AllUSStates` slice.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).

## Author

Karl-Gustav Kallasmaa
