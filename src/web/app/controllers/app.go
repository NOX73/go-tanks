package controllers

import ( 
  "github.com/robfig/revel"
  "code.google.com/p/go.net/websocket"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Ws(ws *websocket.Conn) revel.Result {
  defer ws.Close()

  ws.Write([]byte("Hello"))
  msg := make([]byte, 255)

  for {
    _, err := ws.Read(msg)
    if err != nil { break }
    ws.Write(msg)
  }

  return nil
}
