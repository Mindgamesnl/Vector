package events

import "craftmend.com/Mindgamesnl/Vector/api"

type EventModule struct {

}

func CreateEventModule(api.Robot) EventModule {
	return EventModule{}
}