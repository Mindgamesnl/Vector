package api

import (
	"craftmend.com/Mindgamesnl/Vector/api/components/animation"
	"craftmend.com/Mindgamesnl/Vector/api/components/audio"
	"craftmend.com/Mindgamesnl/Vector/api/components/behavior"
	"craftmend.com/Mindgamesnl/Vector/api/components/camera"
	"craftmend.com/Mindgamesnl/Vector/api/components/control"
	"craftmend.com/Mindgamesnl/Vector/api/components/events"
	"craftmend.com/Mindgamesnl/Vector/api/components/face"
	"craftmend.com/Mindgamesnl/Vector/api/components/motor"
	"craftmend.com/Mindgamesnl/Vector/api/components/photo"
	"craftmend.com/Mindgamesnl/Vector/api/components/screen"
	"craftmend.com/Mindgamesnl/Vector/api/components/vision"
	"craftmend.com/Mindgamesnl/Vector/api/components/world"
	"craftmend.com/Mindgamesnl/Vector/api/helpers"
	"craftmend.com/Mindgamesnl/Vector/api/math"
	"time"
)

type Robot struct {
	// events
	ErrorHandlers        helpers.EventHandler
	ConnectedHandlers    helpers.EventHandler
	DisconnectedHandlers helpers.EventHandler
	DelocalizedHandlers  helpers.EventHandler

	// components
	Control       control.ControlComponent
	Camera        camera.CameraComponent
	Behavior      behavior.BehaviorComponent
	Motors        motor.MotorComponent
	Photos        photo.PhotoComponent
	Events        events.EventComponent
	Animation     animation.AnimationComponent
	Face          face.FaceComponent
	Audio         audio.AudioComponent
	Screen        screen.ScreenComponent
	Vision        vision.VisionComponent
	World         world.WorldComponent
	MapNavigation world.MapNavComponent

	// attributes
	RobotAcceleration    math.Acceleration // the current accelerometer reading
	CarryingObjectId     uint8             // the ID of the object currently being carried (-1 if none)
	HeadTrackingObjectId uint8             // the ID of the object the head is tracking to (-1 if none)
	AngularVelocity      math.Acceleration // the current gyroscope reading (x, y, z)
	HeadAngleRad         float32           // Vector's head angle (up/down) in radians
	LastImageTimestamp   time.Time         // the robot's timestamp for the last image seen.
	LeftWheelSpeedMmPS   uint16            // speed of the left wheel in milli meters per second
	RightWheelSpeedMmPS  uint16            // speed of the right wheel in milli meters per second
}
