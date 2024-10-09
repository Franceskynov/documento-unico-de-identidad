package documentoUnicoDeIdentidad

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Validates if any DUI is valid or not
func Validate(document string) error {

	// Checks if the document has a valid format
	validDocument, _ := regexp.MatchString(`(^\d{8})-(\d$)`, document)

	if validDocument {

		Σ := 0
		segments := strings.Split(document, `-`)
		remaing := segments[0]
		validatorNumber, _ := strconv.Atoi(segments[1])
		
		for position := 0; position < len(remaing); position++ {
			digit, _ := strconv.Atoi(string(remaing[position]))
			Σ += digit * (9 - position)
		}

		if (10 - (Σ % 10)) % 10 == validatorNumber {
			return nil
		} else {
			return errors.New("invalid document")
		}
	} else {
		return errors.New("invalid format")
	}
}