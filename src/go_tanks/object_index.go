package go_tanks

import (
  "container/list"
  "math"
  //log "./log"
)

type IndexItem struct {
  X     *list.Element
  Y     *list.Element
}

type ItemsMap map[Coordsable]*IndexItem

type Coordsable interface {
  GetCoords() *Coords
  GetRadius() int
}

type ObjectIndex struct {
  listX         *list.List
  listY         *list.List
  itemsMap      ItemsMap
}

func NewObjectIndex() *ObjectIndex {
  return &ObjectIndex{ list.New(), list.New(), make(ItemsMap) }
}

func ( o *ObjectIndex ) Add ( item Coordsable ) {
  x := o.addX( item )
  y := o.addY( item )
  o.itemsMap[item] = &IndexItem{x,y}
}

func ( o *ObjectIndex ) Remove ( item Coordsable ) {
  e := o.itemsMap[item]
  o.listX.Remove(e.X)
  o.listY.Remove(e.Y)
  delete(o.itemsMap, item)
}

func ( o *ObjectIndex ) addX ( item Coordsable ) *list.Element {
  after := o.findForAddX( item )
  if after == nil {
    return o.listX.PushFront( item )
  } else {
    return o.listX.InsertAfter( item, after )
  }
}

func ( o *ObjectIndex ) addY ( item Coordsable ) *list.Element {
  after := o.findForAddY( item )
  if after == nil {
    return o.listY.PushFront( item )
  } else {
    return o.listY.InsertAfter( item, after )
  }
}

func ( o *ObjectIndex ) findForAddX ( item Coordsable ) *list.Element {
  x := item.GetCoords().X

  for e := o.listX.Front(); e != nil; e = e.Next() {
    if e.Value.( Coordsable ).GetCoords().X > x {
      return e.Prev()
    }
    if e.Next() == nil {
      return e
    }
  }

  return nil
}

func ( o *ObjectIndex ) findForAddY ( item Coordsable ) *list.Element {
  y := item.GetCoords().Y

  var e *list.Element
  for e = o.listY.Front(); e != nil; e = e.Next() {
    if e.Value.( Coordsable ).GetCoords().Y > y {
      return e.Prev()
    }
    if e.Next() == nil {
      return e
    }
  }

  return nil
}

func ( o *ObjectIndex ) ApplyTankPosition ( coords *Coords, direction float64, tank *Tank, m *Map ) {

  if( coords.X - tank.Radius < 0 ) { coords.X = tank.Radius }
  if( coords.Y - tank.Radius < 0 ) { coords.Y = tank.Radius }

  if( coords.X + tank.Radius > m.Width ) { coords.X = m.Width - tank.Radius }
  if( coords.Y + tank.Radius > m.Height ) { coords.Y = m.Height - tank.Radius }

  element := o.itemsMap[tank].X
  maxDist := 2*tank.Radius //Max distance for check

  for e := element.Prev(); e != nil; e = e.Prev() {
    v, ok := e.Value.(*Tank) //Prev tank
    if !ok {continue;}
    c := v.GetCoords() // Prev tank coords

    if coords.X - c.X > maxDist {break;}

    h := math.Hypot(float64(coords.X - c.X), float64(coords.Y - c.Y))

    diff := h - float64(maxDist)

    if diff < -1 { o.fixCoords(diff, coords, tank.Direction) }
  }

  for e := element.Next(); e != nil; e = e.Next() {
    v, ok := e.Value.(*Tank) //Prev tank
    if !ok {continue;}
    c := v.GetCoords() // Prev tank coords

    if c.X - coords.X > maxDist {break;}

    h := math.Hypot(float64(coords.X - c.X), float64(coords.Y - c.Y))

    diff := h - float64(maxDist)

    if diff < -1 { o.fixCoords(diff, coords, tank.Direction) }
  }

  tank.ApplyMove( coords, direction )

  o.ResortElement(element)

}

func ( o *ObjectIndex ) ResortElement ( element *list.Element ) {
  coords := element.Value.(Coordsable).GetCoords()

  for e := element.Prev(); e != nil; e = element.Prev() {
    if e.Value.(Coordsable).GetCoords().X < coords.X {break}
    o.listX.MoveBefore(element, e)
  }

  for e := element.Next(); e != nil; e = element.Next() {
    if e.Value.(Coordsable).GetCoords().X > coords.X {break}
    o.listX.MoveAfter(element, e)
  }
}

func ( o *ObjectIndex ) fixCoords (diff float64, coords *Coords, direction float64) {
  //log.Debug("")
  //log.Debug("fix coords")
  diff = math.Abs(diff) 
  //log.Debug("diff=", diff)
  radDirection := (math.Pi * direction) / 180

  diffX := diff * math.Cos( radDirection )
  diffY := diff * math.Sin( radDirection )

  //log.Debug("direction=", direction)
  //log.Debug("diffX=", diffX)
  //log.Debug("diffY=", diffY)

  coords.X -= int(diffX)
  coords.Y -= int(diffY)
}

func ( o *ObjectIndex ) ValidateBulletPosition ( coords *Coords, direction float64, bullet *Bullet, m *Map ) ( bool ) {
    return !(coords.X < 0 || coords.X > m.Width || coords.Y < 0 || coords.Y > m.Height)
}
