package behavior

import "github.com/Mindgamesnl/Vector/api"

type BehaviorModule struct {
	
}

func CreateBehaviorModule(api.Robot) BehaviorModule {
	return BehaviorModule{}
}