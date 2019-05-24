package aircraft

// State defines acceleration direction
type State int

const (
	decreasing State = -1
	neutral    State = 0
	increasing State = 1
)

// Aircraft struct
type Aircraft struct {
	speedAcceleration    float64
	altitude             int
	altitudeAcceleration int
	altitudeState        State
	altitudeTarget       int
	heading              int
	headingAcceleration  float64
	headingState         State
	headingTarget        int
	speed                float64
	speedState           State
	speedTarget          float64
}

// UpdateAltitude update the `altitude` value
func (a *Aircraft) UpdateAltitude(deltaTime float64) {
	switch a.altitudeState {
	case decreasing:
		a.decreaseAltitude(deltaTime)
	case neutral:
		a.neutralAltitude(deltaTime)
	case increasing:
		a.increaseAltitude(deltaTime)
	}
}

func (a *Aircraft) decreaseAltitude(deltaTime float64) {
	nextAltitude := a.altitude - (a.altitudeAcceleration * int(deltaTime))

	if nextAltitude < a.altitudeTarget {
		nextAltitude = a.altitudeTarget
		a.altitudeState = neutral
	}

	a.altitude = nextAltitude
}

func (a *Aircraft) neutralAltitude(deltaTime float64) {
	if a.altitude == a.altitudeTarget {
		a.altitudeState = neutral

		return
	}

	a.altitudeTarget = a.altitude
}

func (a *Aircraft) increaseAltitude(deltaTime float64) {
	nextAltitude := a.altitude + (a.altitudeAcceleration * int(deltaTime))

	if nextAltitude > a.altitudeTarget {
		nextAltitude = a.altitudeTarget
		a.altitudeState = neutral
	}

	a.altitude = nextAltitude
}

// UpdateHeading updates
func (a *Aircraft) UpdateHeading(deltaTime float64) {
	switch a.headingState {
	case decreasing:
		a.decreaseHeading(deltaTime)
	case neutral:
		a.neutralHeading(deltaTime)
	case increasing:
		a.increaseHeading(deltaTime)
	}
}

func (a *Aircraft) decreaseHeading(deltaTime float64) {
	headingDecrementor := a.headingAcceleration * deltaTime
	nextHeading := int(float64(a.heading) - headingDecrementor)

	if nextHeading < 0 {
		nextHeading += 360
	}

	if nextHeading < a.headingTarget {
		nextHeading = a.headingTarget
		a.headingState = neutral
	}

	a.heading = nextHeading
}

func (a *Aircraft) neutralHeading(deltaTime float64) {
	if a.heading == a.headingTarget {
		a.headingState = neutral

		return
	}

	a.headingTarget = a.heading
}

func (a *Aircraft) increaseHeading(deltaTime float64) {
	headingDecrementor := a.headingAcceleration * deltaTime
	nextHeading := int(float64(a.heading) + headingDecrementor)

	if nextHeading >= 360 {
		nextHeading -= 360
	}

	if nextHeading > a.headingTarget {
		nextHeading = a.headingTarget
		a.headingState = neutral
	}

	a.heading = nextHeading
}

// UpdateSpeed updates the `speed` value
func (a *Aircraft) UpdateSpeed(deltaTime float64) {
	switch a.speedState {
	case decreasing:
		a.decreaseSpeed(deltaTime)
	case neutral:
		a.neutralSpeed(deltaTime)
	case increasing:
		a.increaseSpeed(deltaTime)
	}
}

func (a *Aircraft) decreaseSpeed(deltaTime float64) {
	nextSpeed := a.speed - (a.speedAcceleration * deltaTime)

	if nextSpeed < a.speedTarget {
		nextSpeed = a.speedTarget
		a.speedState = neutral
	}

	a.speed = nextSpeed
}

func (a *Aircraft) neutralSpeed(deltaTime float64) {
	if a.speed == a.speedTarget {
		a.speedState = neutral

		return
	}

	a.speedTarget = a.speed
}

func (a *Aircraft) increaseSpeed(deltaTime float64) {
	nextSpeed := a.speed + (a.speedAcceleration * deltaTime)

	if nextSpeed > a.speedTarget {
		nextSpeed = a.speedTarget
		a.speedState = neutral
	}

	a.speed = nextSpeed
}
