package go_tanks

import (
	i "../interfaces"
)

type Controller struct {
	Client i.Client
	Rooms  i.Rooms
}
