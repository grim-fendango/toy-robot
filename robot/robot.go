package robot

import (
  "errors"
)

const (
  NORTH = "NORTH"
  SOUTH = "SOUTH"
  EAST = "EAST"
  WEST = "WEST"
)

type Position struct {
  X int
  Y int
  Direction string
}

type Table struct {
  Width int
  Height int
}

type Robot struct {
  Position Position
  Table Table
}

func New(pos Position, t Table) *Robot {
  robot := Robot{
    Position: pos,
    Table: t,
  }
  return &robot
}

func (r *Robot) Place(pos Position) error {
  err := ValidPosition(pos, r.Table)
  if err != nil {
    return err
  }
  r.Position = pos
  return nil
}

func ValidPosition(pos Position, t Table) error {
  if !ValidDirection(pos.Direction) {
    return errors.New("Not a valid direction")
  }
  if !ValidCoords(pos.X, pos.Y, t) {
    return errors.New("Position not on table")
  }
  return nil
}

func ValidDirection(dir string) bool {
  return dir == NORTH || dir == SOUTH || dir == EAST || dir == WEST
}

func ValidCoords(x, y int, t Table) bool {
  return x <= t.Width && x >= 0 && y <= t.Height && y >= 0
}

func (r *Robot) Move() error {
  dir := r.Position.Direction
  newPos := r.Position
  if dir == NORTH {
    newPos.Y++
  } else if dir == SOUTH {
    newPos.Y--
  } else if dir == EAST {
    newPos.X++
  } else if dir == WEST {
    newPos.X--
  } else {
    return errors.New("Unknown direction")
  }
  err := ValidPosition(newPos, r.Table)
  if err != nil {
    return err
  }
  r.Position = newPos
  return nil
}

func (r *Robot) Left() error {
  dir := r.Position.Direction
  if dir == NORTH {
    r.Position.Direction = WEST
  } else if dir == SOUTH {
    r.Position.Direction = EAST
  } else if dir == EAST {
    r.Position.Direction = NORTH
  } else if dir == WEST {
    r.Position.Direction = SOUTH
  } else {
    return errors.New("Unknown direction")
  }
  return nil
}

func (r *Robot) Right() error {
  dir := r.Position.Direction
  if dir == NORTH {
    r.Position.Direction = EAST
  } else if dir == SOUTH {
    r.Position.Direction = WEST
  } else if dir == EAST {
    r.Position.Direction = SOUTH
  } else if dir == WEST {
    r.Position.Direction = NORTH
  } else {
    return errors.New("Unknown direction")
  }
  return nil
}
