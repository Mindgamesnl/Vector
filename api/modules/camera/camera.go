package camera

import "craftmend.com/Mindgamesnl/Vector/api"

type CameraModule struct {

}

func CreateCameraModule(api.Robot) CameraModule {
	return CameraModule{}
}