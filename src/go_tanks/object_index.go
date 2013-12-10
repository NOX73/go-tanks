package go_tanks

import (
  "container/list"
  log "./log"
)

type Coordsable interface {
  GetCoords() *Coords
}

type ObjectIndex struct {
  listX  *list.List
  listY  *list.List
}

func NewObjectIndex() *ObjectIndex {
  return &ObjectIndex{ list.New(), list.New() }
}

func ( o *ObjectIndex ) Add ( item Coordsable ) {
  o.addX( item )
  o.addY( item )
}

func ( o *ObjectIndex ) addX ( item Coordsable ) {
  after := o.findForAddX( item )
  if after == nil {
    o.listX.PushFront( item )
  } else {
    o.listX.InsertAfter( item, after )
  }
}

func ( o *ObjectIndex ) addY ( item Coordsable ) {
  after := o.findForAddY( item )
  if after == nil {
    o.listY.PushFront( item )
  } else {
    o.listY.InsertAfter( item, after )
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

func ( o *ObjectIndex ) Loging () {

  log.Debug("Log X")

  for e := o.listX.Front(); e != nil; e = e.Next() {
    log.Debug(e.Value.( Coordsable ).GetCoords().X)
  }

}
