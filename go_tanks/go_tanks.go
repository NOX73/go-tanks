package go_tanks

type Config struct {
        Address         string
        Port            int
}

var DefaultConfig = Config{
        Address:    "0.0.0.0",
        Port:       9292,
}
