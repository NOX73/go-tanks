package go_tanks

import (
  i "../interfaces"
  "errors"
)

func ValidateClientMessage ( m *i.Message ) error {
  message := *m

  if message["WorldDisabled"] != nil {
    if _, ok := message["WorldDisabled"].(bool); !ok { return errors.New("WorldDisabled field should be boolean.") }
  }


  if message["WorldFrequency"] != nil {
    val, ok := message["WorldFrequency"].(float64)
    if !ok { return errors.New("WorldFrequency field should be int.") }

    message["WorldFrequency"] = int(val)

    if val < 0 || val > 100 { return errors.New("WorldFrequency paramentr should be in [0..100]") }
  }

  return nil
}
