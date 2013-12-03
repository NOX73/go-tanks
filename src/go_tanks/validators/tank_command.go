package go_tanks

import (
  i "../interfaces"
  "errors"
  //log "../log"
)

func ValidateTankCommandMessage ( m *i.Message  ) error {
  message := *m

  if message["LeftMotor"] != nil {
    val, ok := message["LeftMotor"].(float64)
    if !ok { return errors.New("LeftMotor paramentr should be float.") }

    if val < -1 || val > 1 {
      return errors.New("LeftMotor paramentr should be in [-1..1]")
    }
  }

  if message["RightMotor"] != nil {
    val, ok := message["RightMotor"].(float64)
    if !ok { return errors.New("RightMotor paramentr should be float.") }

    if val < -1 || val > 1 {
      return errors.New("RightMotor paramentr should be in [-1..1]")
    }
  }

  if message["Gun"] != nil {
    err := ValidateGunCommand ( m )
    if err != nil { return err }
  }

  return nil
}

