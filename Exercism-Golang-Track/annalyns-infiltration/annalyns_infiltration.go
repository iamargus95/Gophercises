package annalyn

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {

	var isAwake bool
	if !knightIsAwake {
		isAwake = true
	}

	return isAwake
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {

	var spying bool
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		spying = true
	}

	return spying
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {

	var signal bool
	if !archerIsAwake && prisonerIsAwake {
		signal = true
	}

	return signal
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	var free bool
	if prisonerIsAwake && !knightIsAwake && !archerIsAwake || !archerIsAwake && petDogIsPresent {
		free = true
	}

	return free
}
