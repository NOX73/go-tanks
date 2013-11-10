package go_tanks

import (
  i "../interfaces"
  "errors"
)

func ValidateAuthForm ( m *i.Message ) error {
  message := *m

  err := ValidateType(m);
  if err != nil { return err }

  if _, ok := message["Login"].(string); !ok { return errors.New("Field 'Login' should be string.") }
  if _, ok := message["Password"].(string); !ok { return errors.New("Field 'Password' should be string.") }

  if message["Type"] != "Auth" { return errors.New("Field 'Type' should be 'Auth' for authicate form.") }

  return nil
}

