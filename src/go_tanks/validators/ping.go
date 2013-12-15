package go_tanks

import (
  i "../interfaces"
  "errors"
  //log "../log"
)

func ValidatePingMessage ( m *i.Message  ) error {
  message := *m

  if message["PingId"] == nil { return errors.New("PingId must be.") }
  if _, ok := message["PingId"].(float64); !ok { return errors.New("PingId must be integer.") }

  return nil
}
