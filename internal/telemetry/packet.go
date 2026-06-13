package telemetry

// Packet holds one 324-byte telemetry frame broadcast by Forza Horizon 6.
// Field order must match the binary layout exactly — binary.Read fills them sequentially.
// https://support.forza.net/hc/en-us/articles/51744149102611-Forza-Horizon-6-Data-Out-Documentation
type Packet struct {
	// Race state
	IsRaceOn int32 `json:"isRaceOn" doc:"1 when race is active, 0 when in menus or race stopped"`

	// Timestamps
	TimestampMS uint32 `json:"timestampMS" doc:"Millisecond timestamp; can overflow back to 0"`

	// Engine
	EngineMaxRpm     float32 `json:"engineMaxRpm" doc:"Maximum engine RPM"`
	EngineIdleRpm    float32 `json:"engineIdleRpm" doc:"Idle engine RPM"`
	CurrentEngineRpm float32 `json:"currentEngineRpm" doc:"Current engine RPM"`

	// Acceleration (car's local space; X = right, Y = up, Z = forward)
	AccelerationX float32 `json:"accelerationX" doc:"Lateral acceleration in car local space (m/s²), positive = right"`
	AccelerationY float32 `json:"accelerationY" doc:"Vertical acceleration in car local space (m/s²), positive = up"`
	AccelerationZ float32 `json:"accelerationZ" doc:"Longitudinal acceleration in car local space (m/s²), positive = forward"`

	// Velocity (car's local space; X = right, Y = up, Z = forward)
	VelocityX float32 `json:"velocityX" doc:"Lateral velocity in car local space (m/s), positive = right"`
	VelocityY float32 `json:"velocityY" doc:"Vertical velocity in car local space (m/s), positive = up"`
	VelocityZ float32 `json:"velocityZ" doc:"Longitudinal velocity in car local space (m/s), positive = forward"`

	// Angular velocity (car's local space, rad/s; X = pitch, Y = yaw, Z = roll)
	AngularVelocityX float32 `json:"angularVelocityX" doc:"Pitch angular velocity (rad/s)"`
	AngularVelocityY float32 `json:"angularVelocityY" doc:"Yaw angular velocity (rad/s)"`
	AngularVelocityZ float32 `json:"angularVelocityZ" doc:"Roll angular velocity (rad/s)"`

	// Orientation (radians)
	Yaw   float32 `json:"yaw" doc:"Yaw orientation (radians)"`
	Pitch float32 `json:"pitch" doc:"Pitch orientation (radians)"`
	Roll  float32 `json:"roll" doc:"Roll orientation (radians)"`

	// Suspension travel normalized (0.0 = max stretch, 1.0 = max compression)
	NormalizedSuspensionTravelFrontLeft  float32 `json:"normalizedSuspensionTravelFrontLeft" doc:"Front-left suspension travel normalized (0.0 = max stretch, 1.0 = max compression)"`
	NormalizedSuspensionTravelFrontRight float32 `json:"normalizedSuspensionTravelFrontRight" doc:"Front-right suspension travel normalized (0.0 = max stretch, 1.0 = max compression)"`
	NormalizedSuspensionTravelRearLeft   float32 `json:"normalizedSuspensionTravelRearLeft" doc:"Rear-left suspension travel normalized (0.0 = max stretch, 1.0 = max compression)"`
	NormalizedSuspensionTravelRearRight  float32 `json:"normalizedSuspensionTravelRearRight" doc:"Rear-right suspension travel normalized (0.0 = max stretch, 1.0 = max compression)"`

	// Tire normalized slip ratio (0 = 100% grip, |ratio| > 1.0 = loss of grip)
	TireSlipRatioFrontLeft  float32 `json:"tireSlipRatioFrontLeft" doc:"Front-left tire slip ratio (0 = full grip, |value| > 1.0 = loss of grip)"`
	TireSlipRatioFrontRight float32 `json:"tireSlipRatioFrontRight" doc:"Front-right tire slip ratio (0 = full grip, |value| > 1.0 = loss of grip)"`
	TireSlipRatioRearLeft   float32 `json:"tireSlipRatioRearLeft" doc:"Rear-left tire slip ratio (0 = full grip, |value| > 1.0 = loss of grip)"`
	TireSlipRatioRearRight  float32 `json:"tireSlipRatioRearRight" doc:"Rear-right tire slip ratio (0 = full grip, |value| > 1.0 = loss of grip)"`

	// Wheel rotation speed (radians/sec)
	WheelRotationSpeedFrontLeft  float32 `json:"wheelRotationSpeedFrontLeft" doc:"Front-left wheel rotation speed (rad/s)"`
	WheelRotationSpeedFrontRight float32 `json:"wheelRotationSpeedFrontRight" doc:"Front-right wheel rotation speed (rad/s)"`
	WheelRotationSpeedRearLeft   float32 `json:"wheelRotationSpeedRearLeft" doc:"Rear-left wheel rotation speed (rad/s)"`
	WheelRotationSpeedRearRight  float32 `json:"wheelRotationSpeedRearRight" doc:"Rear-right wheel rotation speed (rad/s)"`

	// Wheel on rumble strip (1 = on, 0 = off)
	WheelOnRumbleStripFrontLeft  int32 `json:"wheelOnRumbleStripFrontLeft" doc:"1 if front-left wheel is on a rumble strip, 0 otherwise"`
	WheelOnRumbleStripFrontRight int32 `json:"wheelOnRumbleStripFrontRight" doc:"1 if front-right wheel is on a rumble strip, 0 otherwise"`
	WheelOnRumbleStripRearLeft   int32 `json:"wheelOnRumbleStripRearLeft" doc:"1 if rear-left wheel is on a rumble strip, 0 otherwise"`
	WheelOnRumbleStripRearRight  int32 `json:"wheelOnRumbleStripRearRight" doc:"1 if rear-right wheel is on a rumble strip, 0 otherwise"`

	// Wheel in puddle (1 = in puddle, 0 = not)
	WheelInPuddleFrontLeft  int32 `json:"wheelInPuddleFrontLeft" doc:"1 if front-left wheel is in a puddle, 0 otherwise"`
	WheelInPuddleFrontRight int32 `json:"wheelInPuddleFrontRight" doc:"1 if front-right wheel is in a puddle, 0 otherwise"`
	WheelInPuddleRearLeft   int32 `json:"wheelInPuddleRearLeft" doc:"1 if rear-left wheel is in a puddle, 0 otherwise"`
	WheelInPuddleRearRight  int32 `json:"wheelInPuddleRearRight" doc:"1 if rear-right wheel is in a puddle, 0 otherwise"`

	// Surface rumble (non-dimensional, passed to controller force feedback)
	SurfaceRumbleFrontLeft  float32 `json:"surfaceRumbleFrontLeft" doc:"Front-left surface rumble intensity for force feedback"`
	SurfaceRumbleFrontRight float32 `json:"surfaceRumbleFrontRight" doc:"Front-right surface rumble intensity for force feedback"`
	SurfaceRumbleRearLeft   float32 `json:"surfaceRumbleRearLeft" doc:"Rear-left surface rumble intensity for force feedback"`
	SurfaceRumbleRearRight  float32 `json:"surfaceRumbleRearRight" doc:"Rear-right surface rumble intensity for force feedback"`

	// Tire normalized slip angle (0 = 100% grip, |angle| > 1.0 = loss of grip)
	TireSlipAngleFrontLeft  float32 `json:"tireSlipAngleFrontLeft" doc:"Front-left tire slip angle (0 = full grip, |value| > 1.0 = loss of grip)"`
	TireSlipAngleFrontRight float32 `json:"tireSlipAngleFrontRight" doc:"Front-right tire slip angle (0 = full grip, |value| > 1.0 = loss of grip)"`
	TireSlipAngleRearLeft   float32 `json:"tireSlipAngleRearLeft" doc:"Rear-left tire slip angle (0 = full grip, |value| > 1.0 = loss of grip)"`
	TireSlipAngleRearRight  float32 `json:"tireSlipAngleRearRight" doc:"Rear-right tire slip angle (0 = full grip, |value| > 1.0 = loss of grip)"`

	// Tire normalized combined slip (0 = 100% grip, |slip| > 1.0 = loss of grip)
	TireCombinedSlipFrontLeft  float32 `json:"tireCombinedSlipFrontLeft" doc:"Front-left tire combined slip (0 = full grip, |value| > 1.0 = loss of grip)"`
	TireCombinedSlipFrontRight float32 `json:"tireCombinedSlipFrontRight" doc:"Front-right tire combined slip (0 = full grip, |value| > 1.0 = loss of grip)"`
	TireCombinedSlipRearLeft   float32 `json:"tireCombinedSlipRearLeft" doc:"Rear-left tire combined slip (0 = full grip, |value| > 1.0 = loss of grip)"`
	TireCombinedSlipRearRight  float32 `json:"tireCombinedSlipRearRight" doc:"Rear-right tire combined slip (0 = full grip, |value| > 1.0 = loss of grip)"`

	// Suspension travel in meters
	SuspensionTravelMetersFrontLeft  float32 `json:"suspensionTravelMetersFrontLeft" doc:"Front-left suspension travel (meters)"`
	SuspensionTravelMetersFrontRight float32 `json:"suspensionTravelMetersFrontRight" doc:"Front-right suspension travel (meters)"`
	SuspensionTravelMetersRearLeft   float32 `json:"suspensionTravelMetersRearLeft" doc:"Rear-left suspension travel (meters)"`
	SuspensionTravelMetersRearRight  float32 `json:"suspensionTravelMetersRearRight" doc:"Rear-right suspension travel (meters)"`

	// Car info
	CarOrdinal          int32   `json:"carOrdinal" doc:"Unique ID of the car make/model"`
	CarClass            int32   `json:"carClass" doc:"Car class: 0=D, 1=C, 2=B, 3=A, 4=S1, 5=S2, 6=X"`
	CarPerformanceIndex int32   `json:"carPerformanceIndex" doc:"Car performance index: 100 (worst) to 999 (best)"`
	DrivetrainType      int32   `json:"drivetrainType" doc:"Drivetrain: 0=FWD, 1=RWD, 2=AWD"`
	NumCylinders        int32   `json:"numCylinders" doc:"Number of cylinders in the engine"`
	CarGroup            uint32  `json:"carGroup" doc:"Car group identifier"`
	SmashableVelDiff    float32 `json:"smashableVelDiff" doc:"Velocity loss from smashable object collision (m/s)"`
	SmashableMass       float32 `json:"smashableMass" doc:"Mass of the most recently hit smashable object (kg)"`

	// Position in world space (meters)
	PositionX float32 `json:"positionX" doc:"World-space X position (meters)"`
	PositionY float32 `json:"positionY" doc:"World-space Y position (meters)"`
	PositionZ float32 `json:"positionZ" doc:"World-space Z position (meters)"`

	// Performance
	Speed  float32 `json:"speed" doc:"Speed (m/s)"`
	Power  float32 `json:"power" doc:"Engine power output (watts)"`
	Torque float32 `json:"torque" doc:"Engine torque output (newton-meters)"`

	// Tire temperature
	TireTempFrontLeft  float32 `json:"tireTempFrontLeft" doc:"Front-left tire surface temperature (°C)"`
	TireTempFrontRight float32 `json:"tireTempFrontRight" doc:"Front-right tire surface temperature (°C)"`
	TireTempRearLeft   float32 `json:"tireTempRearLeft" doc:"Rear-left tire surface temperature (°C)"`
	TireTempRearRight  float32 `json:"tireTempRearRight" doc:"Rear-right tire surface temperature (°C)"`

	// Boost / fuel
	Boost float32 `json:"boost" doc:"Turbo/supercharger boost pressure (PSI above atmospheric)"`
	Fuel  float32 `json:"fuel" doc:"Fuel level: 0.0 = empty, 1.0 = full"`

	// Race progress
	DistanceTraveled float32 `json:"distanceTraveled" doc:"Total distance traveled in the session (meters)"`
	BestLap          float32 `json:"bestLap" doc:"Best lap time (seconds); 0.0 if not set"`
	LastLap          float32 `json:"lastLap" doc:"Last completed lap time (seconds); 0.0 if not set"`
	CurrentLap       float32 `json:"currentLap" doc:"Current lap time (seconds); 0.0 if not set"`
	CurrentRaceTime  float32 `json:"currentRaceTime" doc:"Time elapsed since driving started (seconds)"`
	LapNumber        uint16  `json:"lapNumber" doc:"Current lap number"`
	RacePosition     uint8   `json:"racePosition" doc:"Current race position"`

	// Inputs (0 to 255)
	Accel     uint8 `json:"accel" doc:"Throttle input: 0 = none, 255 = full"`
	Brake     uint8 `json:"brake" doc:"Brake input: 0 = none, 255 = full"`
	Clutch    uint8 `json:"clutch" doc:"Clutch input: 0 = none, 255 = full"`
	HandBrake uint8 `json:"handBrake" doc:"Handbrake input: 0 = none, 255 = full"`
	Gear      uint8 `json:"gear" doc:"Current gear: 0 = reverse, 1–10 = forward gears"`

	// Steering and driving line
	Steer                       int8 `json:"steer" doc:"Steering input: -127 = full left, 0 = center, 127 = full right"`
	NormalizedDrivingLine       int8 `json:"normalizedDrivingLine" doc:"Position relative to the racing line: -127 to 127"`
	NormalizedAIBrakeDifference int8 `json:"normalizedAIBrakeDifference" doc:"AI brake difference: -127 to 127"`
}
