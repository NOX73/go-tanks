package go_tanks

import (
  "container/list"
  "math"
  log "./log"
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
  listY         *list.List // Not use yet
  itemsMap      ItemsMap
  maxDistance   float64
}

func NewObjectIndex( maxDistance float64 ) *ObjectIndex {
  return &ObjectIndex{ list.New(), list.New(), make(ItemsMap), maxDistance }
}

func ( o *ObjectIndex ) Add ( item Coordsable ) {
  x := o.addX( item )
  y := o.addY( item )
  o.itemsMap[item] = &IndexItem{x,y}
}

func ( o *ObjectIndex ) Remove ( item Coordsable ) {
  e, ok := o.itemsMap[item]
  if !ok { log.Warning("Index hasn't item for remove.", item); return }

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

func ( o *ObjectIndex ) checkApplyTankPosition ( tank *Tank, e *list.Element, radius float64, coords *Coords ) bool {
  value := e.Value.( Coordsable )
  c := value.GetCoords()

  if math.Abs(c.X - coords.X) >= o.maxDistance { return true }

  tank, ok := value.(*Tank)
  if !ok { return false }

  h := math.Hypot(coords.X - c.X, coords.Y - c.Y)

  diff := h - radius*2

  if diff < -1 { o.fixCoords( diff, coords, tank.Direction ) }

  return false
}

func ( o *ObjectIndex ) ApplyTankPosition ( coords *Coords, direction float64, tank *Tank, m *Map ) {
  radius := float64(tank.Radius)

  if( coords.X - radius < 0 ) { coords.X = radius }
  if( coords.Y - radius< 0 ) { coords.Y = radius }

  if( coords.X + radius > float64(m.Width) ) { coords.X = float64( m.Width ) - radius }
  if( coords.Y + radius > float64(m.Height) ) { coords.Y = float64( m.Height ) - radius }

  element := o.itemsMap[tank].X

  for e := element.Prev(); e != nil; e = e.Prev() {
    if o.checkApplyTankPosition(tank, e, radius, coords) {break}
  }

  for e := element.Next(); e != nil; e = e.Next() {
    if o.checkApplyTankPosition(tank, e, radius, coords) {break}
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
  diff = math.Abs(diff)
  radDirection := (math.Pi * direction) / 180

  diffX := diff * math.Cos( radDirection )
  diffY := diff * math.Sin( radDirection )

  coords.X -= diffX * 1.1
  coords.Y -= diffY * 1.1
}

func ( o *ObjectIndex ) ValidateBulletPosition ( coords *Coords, direction float64, bullet *Bullet, m *Map ) ( tankHit *Tank, inMap bool ) {
  element := o.itemsMap[bullet].X

  if(coords.X < 0 || coords.X > float64(m.Width) || coords.Y < 0 || coords.Y > float64(m.Height)) {return nil, false}

  for e := element.Prev(); e != nil; e = e.Prev() {
    hit, b := o.checkApplyBulletPosition(e, coords)
    if hit {return e.Value.(*Tank), true }
    if b { break }
  }

  for e := element.Next(); e != nil; e = e.Next() {
    hit, b := o.checkApplyBulletPosition(e, coords)
    if hit { return e.Value.(*Tank), true }
    if b { break }
  }

  o.ResortElement(element)

  return nil, true
}

func ( o *ObjectIndex ) checkApplyBulletPosition (e *list.Element, coords *Coords) (isHit, isBreak bool) {

  value := e.Value.( Coordsable )
  c := value.GetCoords()

  if c.X - coords.X >= o.maxDistance { return false, true }

  tank, ok := value.(*Tank)
  if !ok { return false, false }

  h := math.Hypot(float64(coords.X - c.X), float64(coords.Y - c.Y))

  if h < float64(tank.GetRadius()) { return true, true }

  return false, false
}
