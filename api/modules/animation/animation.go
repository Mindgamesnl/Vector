package animation

import "github.com/Mindgamesnl/Vector/api"

type AnimationModule struct {

}

func CreateAnimationModule(api.Robot) AnimationModule {
	return AnimationModule{}
}