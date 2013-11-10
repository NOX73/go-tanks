package go_tanks

import (
  i "../interfaces"
  "errors"
  //log "../log"
)

func ValidateTankCommandMessage ( m *i.Message  ) error {
  message := *m

  if LeftMotor, ok := message["LeftMotor"].(float64); !ok {
    errors.New("LeftMotor paramentr should be float.")
  } else {
    if LeftMotor < 0 || LeftMotor > 1 {
      return errors.New("LeftMotor paramentr should be in [0..1]")
    }
  }

  return nil
}

