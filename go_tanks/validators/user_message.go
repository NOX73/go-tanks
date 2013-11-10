package go_tanks

import (
  i "../interfaces"
  "errors"
)

var UserMessageValidators = map[string]func( m *i.Message )error{
  "TankCommand": ValidateTankCommandMessage,
}

func ValidateUserMessage ( m *i.Message ) error {
  err := ValidateType(m);
  if err != nil { return err }

  message := *m

  validator := UserMessageValidators[message["Type"].(string)]
  if validator == nil { return errors.New("Wrond message type.") }

  return UserMessageValidators[message["Type"].(string)]( m )
}



