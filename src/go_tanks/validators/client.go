package go_tanks

import (
  i "../interfaces"
  "errors"
)

func ValidateClientMessage ( m *i.Message ) error {
  message := *m

  if message["WorldRecieveDisabled"] != nil {
    if _, ok := message["WorldRecieveDisabled"].(bool); !ok { return errors.New("WorldRecieveDisabled field should be boolean.") }
  }

  return nil
}
