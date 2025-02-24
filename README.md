# Countries

A Go package that provides ISO country code utilities, supporting both ISO 3166-1 alpha-2 (two-letter) and ISO 3166-1 alpha-3 (three-letter) country codes.

## Features

- Complete list of ISO 3166-1 alpha-2 country codes (two-letter)
- Complete list of ISO 3166-1 alpha-3 country codes (three-letter)

## Installation

```bash
go get github.com/KGKallasmaa/countries
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/KGKallasmaa/countries"
)

func main() {
    // Using ISO-2 country codes
    countryCode := ISO2CountryCodeUS
    fmt.Println(countryCode) // Outputs: US

    // Using ISO-3 country codes
    countryCode3 := ISO3CountryCodeUSA
    fmt.Println(countryCode3) // Outputs: USA
}
```

## Available Constants

The package provides constants for all ISO 3166-1 country codes:

- ISO-2 (alpha-2) codes: e.g., `ISO2CountryCodeUS` for United States
- ISO-3 (alpha-3) codes: e.g., `ISO3CountryCodeUSA` for United States

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).

## Author

Karl-Gustav Kallasmaa
