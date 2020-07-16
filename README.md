### Basic Golang ICAN validator

Validations:
- Validate if country code and check digits are present
- Validate if country code is in accepted country code list
- Validate national BCAN code format
- Validate mod97 check digits

#### Testing
Test suite includes example of ICAN's for most countries and fake ICAN's. Errors are as specific as possible.

#### Example

	package main

	import (
		"fmt"
		"github.com/cryptohub-digital/go-ican"
	)

	func main() {
		ican, err := ican.NewICAN("NL40ABNA0517552264")

		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("%v\n", ican.PrintCode)
			fmt.Printf("%v\n", ican.Code)
		}
	}
