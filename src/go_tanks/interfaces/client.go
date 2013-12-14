package go_tanks

const (
  NEW_TANK = iota
  NEW_CLIENT
  REMOVE_CLIENT
  TANK_COMMAND
  HIT_TANK
)

type Message            map[string]interface{}
type MessageChan        chan *Message

type Client interface {
  SendMessage ( *Message ) error
  ReadMessage () ( *Message, error )

  SetAuthCredentials ( login, password string)
  SetTank ( Tank )
  GetTank () Tank
  HasTank () bool

  Login () *string
  Password () *string

  OutBox () MessageChan
  InBox () MessageChan

  OutClientBox () MessageChan
  InClientBox () MessageChan

  OutBoxHasPlace () bool

  SendWorld ( *Message )
  SetWorldDisabled( bool )
  SetWorldFrequency ( val int )

  ReInit()
}

var typeToIdMaping = map[string]int {
  "TankCommand": TANK_COMMAND,
}

func ( m *Message ) GetTypeId () interface{} {
  message := *m
  if message["TypeId"] == nil {
    message["TypeId"] = typeToIdMaping[ message["Type"].(string) ]
  }

  return message["TypeId"].(int)
}

func ErrorMessage ( message string ) *Message {
  return &Message{"Type":"Error", "Message": message}
}
