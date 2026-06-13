package telemetry

// Packet holds one 324-byte telemetry frame broadcast by Forza Horizon 6.
// Field order must match the binary layout exactly — binary.Read fills them sequentially.
// https://support.forza.net/hc/en-us/articles/51744149102611-Forza-Horizon-6-Data-Out-Documentation
type Packet struct {
	// Race state
	IsRaceOn int32 // 1 when race is on, 0 when in menus/race stopped

	// Timestamps
	TimestampMS uint32 // can overflow to 0 eventually

	// Engine
	EngineMaxRpm     float32
	EngineIdleRpm    float32
	CurrentEngineRpm float32

	// Acceleration (car's local space; X = right, Y = up, Z = forward)
	AccelerationX float32
	AccelerationY float32
	AccelerationZ float32

	// Velocity (car's local space; X = right, Y = up, Z = forward)
	VelocityX float32
	VelocityY float32
	VelocityZ float32

	// Angular velocity (car's local space, rad/s; X = pitch, Y = yaw, Z = roll)
	AngularVelocityX float32
	AngularVelocityY float32
	AngularVelocityZ float32

	// Orientation (radians)
	Yaw   float32
	Pitch float32
	Roll  float32

	// Suspension travel normalized (0.0 = max stretch, 1.0 = max compression)
	NormalizedSuspensionTravelFrontLeft  float32
	NormalizedSuspensionTravelFrontRight float32
	NormalizedSuspensionTravelRearLeft   float32
	NormalizedSuspensionTravelRearRight  float32

	// Tire normalized slip ratio (0 = 100% grip, |ratio| > 1.0 = loss of grip)
	TireSlipRatioFrontLeft  float32
	TireSlipRatioFrontRight float32
	TireSlipRatioRearLeft   float32
	TireSlipRatioRearRight  float32

	// Wheel rotation speed (radians/sec)
	WheelRotationSpeedFrontLeft  float32
	WheelRotationSpeedFrontRight float32
	WheelRotationSpeedRearLeft   float32
	WheelRotationSpeedRearRight  float32

	// Wheel on rumble strip (1 = on, 0 = off)
	WheelOnRumbleStripFrontLeft  int32
	WheelOnRumbleStripFrontRight int32
	WheelOnRumbleStripRearLeft   int32
	WheelOnRumbleStripRearRight  int32

	// Wheel in puddle (1 = in puddle, 0 = not)
	WheelInPuddleFrontLeft  int32
	WheelInPuddleFrontRight int32
	WheelInPuddleRearLeft   int32
	WheelInPuddleRearRight  int32

	// Surface rumble (non-dimensional, passed to controller force feedback)
	SurfaceRumbleFrontLeft  float32
	SurfaceRumbleFrontRight float32
	SurfaceRumbleRearLeft   float32
	SurfaceRumbleRearRight  float32

	// Tire normalized slip angle (0 = 100% grip, |angle| > 1.0 = loss of grip)
	TireSlipAngleFrontLeft  float32
	TireSlipAngleFrontRight float32
	TireSlipAngleRearLeft   float32
	TireSlipAngleRearRight  float32

	// Tire normalized combined slip (0 = 100% grip, |slip| > 1.0 = loss of grip)
	TireCombinedSlipFrontLeft  float32
	TireCombinedSlipFrontRight float32
	TireCombinedSlipRearLeft   float32
	TireCombinedSlipRearRight  float32

	// Suspension travel in meters
	SuspensionTravelMetersFrontLeft  float32
	SuspensionTravelMetersFrontRight float32
	SuspensionTravelMetersRearLeft   float32
	SuspensionTravelMetersRearRight  float32

	// Car info
	CarOrdinal          int32   // unique ID of the car make/model
	CarClass            int32   // 0 (D) to 7 (X class)
	CarPerformanceIndex int32   // 100 (worst) to 999 (best)
	DrivetrainType      int32   // 0 = FWD, 1 = RWD, 2 = AWD
	NumCylinders        int32   // number of cylinders in the engine
	CarGroup            uint32  // car group identifier
	SmashableVelDiff    float32 // velocity loss from smashable object collision (m/s)
	SmashableMass       float32 // mass of recently hit smashable object (kg)

	// Position in world space (meters)
	PositionX float32
	PositionY float32
	PositionZ float32

	// Performance
	Speed  float32 // meters per second
	Power  float32 // watts
	Torque float32 // newton-meters

	// Tire temperature
	TireTempFrontLeft  float32
	TireTempFrontRight float32
	TireTempRearLeft   float32
	TireTempRearRight  float32

	// Boost / fuel
	Boost float32 // turbo/supercharger boost (PSI above atmospheric)
	Fuel  float32 // 0.0 = empty, 1.0 = full

	// Race progress
	DistanceTraveled float32 // total distance traveled (meters)
	BestLap          float32 // seconds; 0.0 if not applicable
	LastLap          float32 // seconds; 0.0 if not applicable
	CurrentLap       float32 // seconds; 0.0 if not applicable
	CurrentRaceTime  float32 // seconds since driving started
	LapNumber        uint16
	RacePosition     uint8

	// Inputs (0 to 255)
	Accel     uint8
	Brake     uint8
	Clutch    uint8
	HandBrake uint8
	Gear      uint8

	// Steering and driving line
	Steer                       int8 // -127 = full left, 0 = center, 127 = full right
	NormalizedDrivingLine       int8 // -127 to 127
	NormalizedAIBrakeDifference int8 // -127 to 127
}
