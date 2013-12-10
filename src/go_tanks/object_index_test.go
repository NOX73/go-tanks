package go_tanks

import (
  "testing"
  . "launchpad.net/gocheck"
  //"log"
)

type FakeTank struct {
  Coords *Coords
}

func ( t *FakeTank ) GetCoords () *Coords {
  return t.Coords
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }
type OAuthSuite struct{}
var _ = Suite(&OAuthSuite{})

func (s *OAuthSuite) TestAddToList(c *C) {

  o := NewObjectIndex()

  o.Add(&FakeTank{&Coords{100,200}})
  o.Add(&FakeTank{&Coords{50,250}})
  o.Add(&FakeTank{&Coords{200,250}})
  o.Add(&FakeTank{&Coords{50,40}})
  o.Add(&FakeTank{&Coords{150,0}})

  c.Assert(o.listX.Len(), Equals, 5)
  c.Assert(o.listX.Len(), Equals, 5)

  elementX := o.listX.Front()
  elementY := o.listY.Front()
  c.Assert(elementX.Value.(Coordsable).GetCoords().X, Equals, 50)
  c.Assert(elementY.Value.(Coordsable).GetCoords().Y, Equals, 0)

  elementX = elementX.Next()
  elementY = elementY.Next()
  c.Assert(elementX.Value.(Coordsable).GetCoords().X, Equals, 50)
  c.Assert(elementY.Value.(Coordsable).GetCoords().Y, Equals, 40)

  elementX = elementX.Next()
  elementY = elementY.Next()
  c.Assert(elementX.Value.(Coordsable).GetCoords().X, Equals, 100)
  c.Assert(elementY.Value.(Coordsable).GetCoords().Y, Equals, 200)

  elementX = elementX.Next()
  elementY = elementY.Next()
  c.Assert(elementX.Value.(Coordsable).GetCoords().X, Equals, 150)
  c.Assert(elementY.Value.(Coordsable).GetCoords().Y, Equals, 250)

  elementX = elementX.Next()
  elementY = elementY.Next()
  c.Assert(elementX.Value.(Coordsable).GetCoords().X, Equals, 200)
  c.Assert(elementY.Value.(Coordsable).GetCoords().Y, Equals, 250)

}
