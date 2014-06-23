package go_tanks

import (
	i "./interfaces"
)

type Rooms struct {
	rooms  map[string]*World
	config WorldConfig
}

func NewRooms(conf WorldConfig) *Rooms {
	return &Rooms{config: conf, rooms: make(map[string]*World)}
}

func (r *Rooms) Get(name string) i.World {
	var room *World
	var ok bool

	room, ok = r.rooms[name]

	if !ok {
		room = NewWorld(r.config, name)
		room.run()
		r.rooms[name] = room
	}

	return room
}
