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

func ( t *FakeTank ) GetRadius () int { return 20 }

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type SuiteT struct {
  o   *ObjectIndex
}

func (s *SuiteT) SetUpTest (c *C) {
  s.o = NewObjectIndex()

  s.o.Add(&FakeTank{&Coords{100,200}})
  s.o.Add(&FakeTank{&Coords{50,250}})
  s.o.Add(&FakeTank{&Coords{200,250}})
  s.o.Add(&FakeTank{&Coords{50,40}})
  s.o.Add(&FakeTank{&Coords{150,0}})
}

var _ = Suite( &SuiteT{} )

func (s *SuiteT) TestResortElement (c *C) {
  o := s.o

  element := o.listX.Front()
  tank := element.Value.(*FakeTank)

  c.Assert(tank.Coords.X, Equals, 50)
  tank.Coords.X = 0

  o.ResortElement(element)

  c.Assert(o.listX.Front().Value.(*FakeTank).Coords.X, Equals, 0)

  tank.Coords.X = 120
  o.ResortElement(element)
  c.Assert(o.listX.Front().Value.(*FakeTank).Coords.X, Equals, 50)

  c.Assert(o.listX.Front().Next().Next().Value.(*FakeTank).Coords.X, Equals, 120)

}

func (s *SuiteT) TestAddToList(c *C) {

  o := s.o

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


func (s *SuiteT) TestRemoveToList(c *C) {

  o := s.o

  c.Assert(o.listX.Len(), Equals, 5)
  c.Assert(o.listX.Len(), Equals, 5)

  o.Remove(o.listX.Front().Value.(Coordsable))

  c.Assert(o.listX.Len(), Equals, 4)
  c.Assert(o.listX.Len(), Equals, 4)

  o.Remove(o.listX.Front().Value.(Coordsable))

  c.Assert(o.listX.Len(), Equals, 3)
  c.Assert(o.listX.Len(), Equals, 3)

}
