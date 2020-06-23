package behavior

import "craftmend.com/Mindgamesnl/Vector/api"

type BehaviorModule struct {
	
}

func CreateBehaviorModule(api.Robot) BehaviorModule {
	return BehaviorModule{}
}