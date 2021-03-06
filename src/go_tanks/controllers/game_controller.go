package go_tanks

import (
	//log "../log"
	"errors"
	"fmt"

	i "../interfaces"
	v "../validators"
	h "./handlers"
)

type GameController struct {
	Controller
	World i.World
}

func NewGameController(client i.Client, rooms i.Rooms) *GameController {
	roomName := *client.Room()
	room := rooms.Get(roomName)
	controller := GameController{Controller: Controller{Client: client, Rooms: rooms}, World: room}
	return &controller
}

func (c *GameController) JoinToGame() error {
	c.addToWorld()

	c.World.NewTank(c.Client)
	message := *(<-c.Client.InBox())

	tank := message["Tank"].(i.Tank)
	c.Client.SetTank(tank)

	message["Message"] = fmt.Sprint("Your tank id = ", tank.GetId())
	c.Client.SendMessage(&message)

	inClientBox := c.Client.InClientBox()
	inBox := c.Client.InBox()

	for {
		select {
		case message, ok := <-inClientBox:

			if !ok {
				c.removeFromWorld()
				return errors.New("Receive channel closed.")
			}

			err := v.ValidateUserMessage(message)

			if err != nil {
				c.Client.SendMessage(i.ErrorMessage(err.Error()))
				continue
			}

			c.handleMessage(message)

		case message, _ := <-inBox:

			switch message.GetTypeId() {
			case i.DESTROY_TANK:
				c.Client.SendMessage(&i.Message{"Type": "Hit", "Message": "Your are die :("})
				return nil
			case i.HIT_TANK:
				c.Client.SendMessage(&i.Message{"Type": "Hit", "Message": "Tank hit."})
			}

		}
	}

	return nil
}

func (c *GameController) addToWorld() {
	c.World.AttachClient(c.Client)
}

func (c *GameController) removeFromWorld() {
	c.World.DetachClient(c.Client)
}

func (c *GameController) handleMessage(m *i.Message) error {
	return h.HandleMessage(c.World, c.Client, m)
}
