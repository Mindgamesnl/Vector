package world

import "github.com/Mindgamesnl/Vector/api"

type WorldModule struct {

}

func CreateWorldModule(api.Robot) WorldModule {
	return WorldModule{}
}