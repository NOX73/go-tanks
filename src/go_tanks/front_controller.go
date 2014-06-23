package go_tanks

import (
	log "./log"
)

type FrontController struct {
	ClientsChannel <-chan *Client
	rooms          *Rooms
}

func NewFrontController(channel <-chan *Client, rooms *Rooms) *FrontController {
	return &FrontController{ClientsChannel: channel, rooms: rooms}
}

func (fc *FrontController) Accept() {
	for {
		client := <-fc.ClientsChannel
		log.Client("New client connected ( ", client.RemoteAddr(), " )")

		fc.processClient(client)
	}
}

func (fc *FrontController) processClient(client *Client) {
	dispatcher := Dispatcher{Client: client, Rooms: fc.rooms}
	go dispatcher.run()
}
