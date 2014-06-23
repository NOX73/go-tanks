package go_tanks

type Rooms interface {
	Get(name string) World
}
