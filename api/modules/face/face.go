package face

import "github.com/Mindgamesnl/Vector/api"

type FaceModule struct {
	
}

func CreateFaceModule(api.Robot) FaceModule {
	return FaceModule{}
}