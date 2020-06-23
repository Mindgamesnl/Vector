package main

import (
	"github.com/Mindgamesnl/Vector/api"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	robot := api.InitializeRobot("")
	spew.Dump(robot)
}
