package validator

import (
	"errors"
	"regexp"
)

var uuidRegex = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)

func UUID() FieldRule {
	return FieldRule{
		rule: func(value any) error {
			s, ok := value.(string)
			if !ok {
				return errors.New("must be a string UUID")
			}
			if !uuidRegex.MatchString(s) {
				return errors.New("invalid UUID format")
			}
			return nil
		},
	}
}
