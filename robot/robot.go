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

type Robot struct {
  Position Position
}

func New() *Robot {
  robot := Robot{}
  return &robot
}

func (this *Robot) Place(pos Position) error {
  if !ValidDirection(pos.Direction) {
    return errors.New("Not a valid direction")
  }
  return nil
}

func ValidDirection(dir string) bool {
  return dir == NORTH || dir == SOUTH || dir == EAST || dir == WEST
}

func (this *Robot) Move() error {
  return nil
}

func (this *Robot) Left() error {
  return nil
}

func (this *Robot) Right() error {
  return nil
}
