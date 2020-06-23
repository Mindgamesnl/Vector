package api

import (
	"craftmend.com/Mindgamesnl/Vector/api/modules/animation"
	"craftmend.com/Mindgamesnl/Vector/api/modules/audio"
	"craftmend.com/Mindgamesnl/Vector/api/modules/behavior"
	"craftmend.com/Mindgamesnl/Vector/api/modules/camera"
	"craftmend.com/Mindgamesnl/Vector/api/modules/control"
	"craftmend.com/Mindgamesnl/Vector/api/modules/events"
	"craftmend.com/Mindgamesnl/Vector/api/modules/face"
	"craftmend.com/Mindgamesnl/Vector/api/modules/motor"
	"craftmend.com/Mindgamesnl/Vector/api/modules/photo"
	"craftmend.com/Mindgamesnl/Vector/api/modules/screen"
	"craftmend.com/Mindgamesnl/Vector/api/modules/vision"
	"craftmend.com/Mindgamesnl/Vector/api/modules/world"
	"craftmend.com/Mindgamesnl/Vector/api/helpers"
	"craftmend.com/Mindgamesnl/Vector/api/math"
	"craftmend.com/Mindgamesnl/Vector/api/objects"
	"craftmend.com/Mindgamesnl/Vector/api/state"
	"time"
)

type Robot struct {
	// events
	ErrorHandlers        helpers.EventHandler
	ConnectedHandlers    helpers.EventHandler
	DisconnectedHandlers helpers.EventHandler
	DelocalizedHandlers  helpers.EventHandler

	// modules
	Control       control.ControlModule
	Camera        camera.CameraModule
	Behavior      behavior.BehaviorModule
	Motors        motor.MotorModule
	Photos        photo.PhotoModule
	Events        events.EventModule
	Animation     animation.AnimationModule
	Face          face.FaceModule
	Audio         audio.AudioModule
	Screen        screen.ScreenModule
	Vision        vision.VisionModule
	World         world.WorldModule
	MapNavigation world.MapNavModule

	// state
	IsConnected          bool                        // if the robot is connected or nah
	IsDisconnecting      bool                        // wetter the connection is being shut down
	RobotAcceleration    math.Acceleration           // the current accelerometer reading
	CarryingObjectId     int8                       // the ID of the object currently being carried (-1 if none)
	HeadTrackingObjectId int8                       // the ID of the object the head is tracking to (-1 if none)
	AngularVelocity      math.Acceleration           // the current gyroscope reading (x, y, z)
	HeadAngleRad         float32                     // Vector's head angle (up/down) in radians
	LastImageTimestamp   time.Time                   // the robot's timestamp for the last image seen.
	LeftWheelSpeedMmPS   uint16                      // speed of the left wheel in milli meters per second
	RightWheelSpeedMmPS  uint16                      // speed of the right wheel in milli meters per second
	Location             math.Pose                   // Local position of the robot (location and orientation)
	ProximitySensor      objects.ProximityDataSensor // live proximity sensor data
	Status               state.RobotState            // current state and activity of the robot
	TouchSensor          objects.TouchSensorData     // current state of the touch sensor on the back
	IpV4                 string                      // local ip address of the robot
}

func InitializeRobot(ip string) Robot {
	r := Robot{}

	// initialize handlers
	r.ErrorHandlers = helpers.CreateEventHandler()
	r.ConnectedHandlers = helpers.CreateEventHandler()
	r.DisconnectedHandlers = helpers.CreateEventHandler()
	r.DelocalizedHandlers = helpers.CreateEventHandler()

	// initialize modules
	r.Control = control.CreateControlModule(r)
	r.Camera = camera.CreateCameraModule(r)
	r.Behavior = behavior.CreateBehaviorModule(r)
	r.Motors = motor.CreateMotorModule(r)
	r.Photos = photo.CreatePhotoModule(r)
	r.Events = events.CreateEventModule(r)
	r.Animation = animation.CreateAnimationModule(r)
	r.Face = face.CreateFaceModule(r)
	r.Audio = audio.CreateAudioModule(r)
	r.Screen = screen.CreateScreenModule(r)
	r.Vision = vision.CreateVisionModule(r)
	r.World = world.CreateWorldModule(r)
	r.MapNavigation = world.CreateMapNavModule(r)

	// default states
	r.IsConnected = false
	r.IsDisconnecting = false
	r.RobotAcceleration = math.Acceleration{
		X:0,
		Y:0,
		Z:0,
	}
	r.CarryingObjectId = -1
	r.HeadTrackingObjectId = -1
	r.AngularVelocity = math.Acceleration{
		X:0,
		Y:0,
		Z:0,
	}
	r.HeadAngleRad = 0
	r.LastImageTimestamp = time.Now()
	r.LeftWheelSpeedMmPS = 0
	r.RightWheelSpeedMmPS = 0
	r.Location = math.CreateZeroPose()
	r.ProximitySensor = objects.CreateZeroProximitySensorData()
	r.Status = state.CreateIdleRobotStatus()
	r.TouchSensor = objects.TouchSensorData{
		IsBeingTouched: false,
		RawTouchValue:  0,
	}
	r.IpV4 = ip

	return r
}
