package go_tanks

import (
  i "../interfaces"
  "errors"
)

func ValidateClientMessage ( m *i.Message ) error {
  message := *m

  if _, ok := message["WorldRecieveDisabled"].(bool); !ok { errors.New("WorldRecieveDisabled field should be boolean.") }

  return nil
}
