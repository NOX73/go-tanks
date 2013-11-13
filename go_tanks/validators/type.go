package go_tanks

import (
  i "../interfaces"
  "errors"
  //log "../log"
)

func ValidateType ( m *i.Message ) error {
  message := *m

  if _, ok := message["Type"].(string); !ok { return errors.New("Field 'Type' should be string.") }

  return nil
}

