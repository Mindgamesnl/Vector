package animation

import "craftmend.com/Mindgamesnl/Vector/api"

type AnimationModule struct {

}

func CreateAnimationModule(api.Robot) AnimationModule {
	return AnimationModule{}
}