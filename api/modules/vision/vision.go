package vision

import "github.com/Mindgamesnl/Vector/api"

type VisionModule struct {

}

func CreateVisionModule(api.Robot) VisionModule {
	return VisionModule{}
}