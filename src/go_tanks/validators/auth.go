package go_tanks

import (
	"errors"

	i "../interfaces"
)

func ValidateAuthForm(m *i.Message) error {
	message := *m

	err := ValidateType(m)
	if err != nil {
		return err
	}

	if message["Type"] != "Auth" {
		return errors.New("Field 'Type' should be 'Auth' for authicate form.")
	}

	if _, ok := message["Login"].(string); !ok {
		return errors.New("Field 'Login' should be string.")
	}
	if _, ok := message["Password"].(string); !ok {
		return errors.New("Field 'Password' should be string.")
	}
	if _, ok := message["Room"].(string); !ok {
		message["Room"] = string("main")
	}

	return nil
}
