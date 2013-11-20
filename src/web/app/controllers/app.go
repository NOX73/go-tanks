package controllers

import ( 
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
  buff := make([]byte, 255)

  for {
    n, err := ws.Read(buff)
    if err != nil { break }
    ws.Write(buff[0:n])
  }

  return nil
}
