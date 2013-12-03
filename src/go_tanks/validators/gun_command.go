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

  if gun["TurnAngle"] != nil {
    val, ok := gun["TurnAngle"].(float64)
    if !ok { return errors.New("Gun/TurnAngle paramentr should be float.") }

    if val < -36000 || val > 36000 { return errors.New("Gun/TurnAngle paramentr should be in [-36000..36000]") }
  }

  return nil
}
