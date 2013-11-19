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
  ws.Write([]byte("Hello"))
  msg := make([]byte, 255)

  for {
    _, _ = ws.Read(msg)
    ws.Write(msg)
  }

  ws.Close()

  return nil
}
