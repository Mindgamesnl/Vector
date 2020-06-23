package vision

import "craftmend.com/Mindgamesnl/Vector/api"

type VisionModule struct {

}

func CreateVisionModule(api.Robot) VisionModule {
	return VisionModule{}
}