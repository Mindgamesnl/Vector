package photo

import "github.com/Mindgamesnl/Vector/api"

type PhotoModule struct {

}

func CreatePhotoModule(api.Robot) PhotoModule {
	return PhotoModule{}
}