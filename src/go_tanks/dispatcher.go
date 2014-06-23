package go_tanks

import (
	controllers "./controllers"
	i "./interfaces"
	log "./log"
)

type Dispatcher struct {
	Rooms  *Rooms
	Client *Client
}

func (d *Dispatcher) run() {
	d.dispatch()
}

func (d *Dispatcher) dispatch() error {
	defer d.Client.Close()
	var err error

	for {
		err = d.authStep()
		if err != nil {
			return d.sendError(err)
		}

		err = d.gameStep()
		if err != nil {
			return d.sendError(err)
		}

		d.Client.ReInit()
	}

	return nil
}

func (d *Dispatcher) gameStep() error {
	controller := controllers.NewGameController(d.Client, d.Rooms)
	return controller.JoinToGame()
}

func (d *Dispatcher) authStep() error {
	controller := controllers.AuthController{controllers.Controller{Client: d.Client, Rooms: d.Rooms}}
	return controller.Authorize()
}

func (d *Dispatcher) sendError(err error) error {
	log.Error(err)
	d.Client.SendMessage(i.ErrorMessage(err.Error()))

	return err
}
