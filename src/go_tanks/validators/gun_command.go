package go_tanks

import (
  i "../interfaces"
  "errors"
  //log "../log"
)

func ValidateGunCommand ( m *i.Message  ) error {
  message := *m

  gun, ok := message["Gun"].(map[string]interface{})
  if !ok { return errors.New("Gun paramentr should be map[string].") }

  if gun["Direction"] != nil {
    val, ok := gun["Direction"].(float64)
    if !ok { return errors.New("Gun/Direction paramentr should be float.") }

    if val < 0 || val > 360 { return errors.New("Gun/Direction paramentr should be in [0..360]") }
  }

  return nil
}
