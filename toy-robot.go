package main

import (
  "regexp"
  "os"
  "bufio"
  "fmt"
  "strconv"
  "errors"

  robot "github.com/grim-fendango/toy-robot/robot"
)

const (
  end = "END"
  place = "PLACE"
  move = "MOVE"
  left = "LEFT"
  right = "RIGHT"
  report = "REPORT"
)

// Error messages
const (
  not_enough_args = "Not enough arguments provided for this command"
  badly_formatted_args = "One or more arguments was poorly formatted"
)

func placeCmd(rbt *robot.Robot, args []string) error {
  if len(args) < 3 {
    return errors.New(not_enough_args)
  }
  x, err := strconv.Atoi(args[0])
  if err != nil {
    return errors.New(badly_formatted_args)
  }
  y, err := strconv.Atoi(args[1])
  if err != nil {
    return errors.New(badly_formatted_args)
  }
  dir := args[2]
  if !robot.ValidDirection(dir) {
    return fmt.Errorf("%v\n %v", badly_formatted_args, err)
  }
  pos := robot.Position{
    X: x,
    Y: y,
    Direction: dir,
  }
  err = rbt.Place(pos)
  if err == nil {
    fmt.Printf("Placing at X: %v, Y: %v, Direction: %v\n", pos.X, pos.Y, pos.Direction)
  }
  return err
}

func moveCmd(rbt *robot.Robot) error {
  err := rbt.Move()
  if err == nil {
    fmt.Println("Moving forward bleep bloop")
  }
  return err
}

func leftCmd(rbt *robot.Robot) error {
  err := rbt.Left()
  if err == nil {
    fmt.Println("Moving left bleep bloop")
  }
  return err
}

func rightCmd(rbt *robot.Robot) error {
  err := rbt.Right()
  if err == nil {
    fmt.Println("Moving right bloop blooorp")
  }
  return err
}

func reportCmd(rbt *robot.Robot) {
  position := rbt.Position
  fmt.Println("Reporting position...")
  fmt.Printf("X: %v, Y: %v, Direction: %v\n", position.X, position.Y, position.Direction)
}

func findCommand(raw string) string {
  rCommand, _ := regexp.Compile(fmt.Sprintf("^%s|%s|%s|%s|%s", end, move, left, right, place))
  command := rCommand.FindString(raw)
  return command
}

func findArgs(raw string) []string {
  rArgs, _ := regexp.Compile(fmt.Sprintf(`%s ([0-9]*) ([0-9]*) ([A-Z]*)`, place))
  args := rArgs.FindAllStringSubmatch(raw, -1)
  if len(args) == 0 {
    return []string{}
  }
  return args[0][1:]
}

func processInput(raw string) (command string, args []string) {
  command = findCommand(raw)
  args = findArgs(raw)
  return command, args
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  rbt := robot.New()
  fmt.Println("Robot is ready")
  for {
    fmt.Print("Enter Command: ")
    raw, _ := reader.ReadString('\n')
    command, args := processInput(raw)
    if len(args) > 0 {
      fmt.Printf("Command: %v Args: %v\n", command, args)
    } else {
      fmt.Printf("Command: %v\n", command)
    }
    var err error
    if command == end {
      break
    } else if command == place {
      err = placeCmd(rbt, args)
    } else if command == move {
      err = moveCmd(rbt)
    } else if command == left {
      err = leftCmd(rbt)
    } else if command == right {
      err = rightCmd(rbt)
    } else if command == report {
      reportCmd(rbt)
    } else {
      fmt.Println("Invalid command entered")
    }
    if err != nil {
      fmt.Printf("Command could not be performed: %v\n", err)
    }
  }
  fmt.Println("Good bye! :D")
}
