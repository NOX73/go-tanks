package go_tanks

import "log"
import "fmt"

const (
  CLIENT = "CLIENT:\t%s"
  ERROR = "ERROR:\t%s"
  FATAL = "FATAL:\t%s"
  SERVER = "SERVER:\t%s"
  WORLD = "WORLD:\t%s"
  DEBUG = "DEBUG:\t%s"
)

func Client(v ...interface{}) {
  log.Printf(CLIENT, fmt.Sprint(v...))
}

func Error(v ...interface{}) {
  log.Printf(ERROR, fmt.Sprint(v...))
}

func Server(v ...interface{}) {
  log.Printf(SERVER, fmt.Sprint(v...))
}

func Fatal(v ...interface{}) {
  log.Fatalf(FATAL, fmt.Sprint(v...))
}

func World(v ...interface{}) {
  log.Printf(WORLD, fmt.Sprint(v...))
}

func Debug(v ...interface{}) {
  log.Printf(DEBUG, fmt.Sprint(v...))
}
