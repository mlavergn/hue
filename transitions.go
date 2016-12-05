package hue

import (
	"math/rand"
	"time"
)

// NOTE: TransitionTime is in units of 100ms

// ShortCircuit creates a short circuit transition
func ShortCircuit(light *Light, duration time.Duration, completionState *State) {
	complete := time.Now().Add(duration)

	hueTT := uint16(2)
	for time.Now().Before(complete) {
		// off
		light.Set(&State{TransitionTime: hueTT, Brightness: 150})
		time.Sleep(time.Duration(hueTT*100) * time.Millisecond)
		// on
		light.Set(&State{TransitionTime: hueTT, Brightness: 254})
		time.Sleep(time.Duration(rand.Intn(1000)+100) * time.Millisecond)
	}

	light.Set(completionState)
}

// Fade creates a slow fade
func Fade(light *Light, duration time.Duration, completionState *State) {
	ms := uint16(duration.Nanoseconds() / 1000000)
	hueTT := ms / 100 / 2

	// off
	light.Set(&State{TransitionTime: hueTT, Brightness: 1})
	time.Sleep(time.Duration(hueTT*100) * time.Millisecond)
	// on
	state := *completionState
	state.TransitionTime = hueTT
	state.Brightness = 254
	light.Set(&state)
	time.Sleep(time.Duration(hueTT*100) * time.Millisecond)

	light.Set(completionState)
}

// Flash creates a slow consistent flashing transition
func Flash(light *Light, duration time.Duration, completionState *State) {
	complete := time.Now().Add(duration)

	hueTT := uint16(5)
	for time.Now().Before(complete) {
		// off
		light.Set(&State{TransitionTime: hueTT, Brightness: 1})
		time.Sleep(time.Duration(hueTT*100) * time.Millisecond)
		// on
		light.Set(&State{TransitionTime: hueTT, Brightness: 254})
		time.Sleep(time.Duration(hueTT*100) * time.Millisecond)
	}

	light.Set(completionState)
}

// Strobe creates a strobe transition
// NOTE: strobe requires low latency which cannot be guaranteed by XBee's variable latency model
func Strobe(light *Light, duration time.Duration, completionState *State) {
	complete := time.Now().Add(duration)

	hueTT := uint16(1)
	for time.Now().Before(complete) {
		// off
		light.Set(&State{TransitionTime: hueTT, Brightness: 1})
		time.Sleep(time.Duration(hueTT*30) * time.Millisecond)
		// on
		light.Set(&State{TransitionTime: hueTT, Brightness: 254})
		time.Sleep(time.Duration(hueTT*30) * time.Millisecond)
	}

	light.Set(completionState)
}

// Lightning creates a white lighting transition
func Lightning(light *Light, duration time.Duration, completionState *State) {
	complete := time.Now().Add(duration)

	light.Set(ColorWhite(light))
	for time.Now().Before(complete) {
		// off
		light.Set(&State{TransitionTime: 1, Brightness: 1})
		time.Sleep(time.Duration(rand.Intn(500)+100) * time.Millisecond)
		// surge
		light.Set(&State{TransitionTime: 1, Brightness: uint8(rand.Intn(54) + 200)})
		time.Sleep(time.Duration(rand.Intn(50)+10) * time.Millisecond)
		// on
		light.Set(&State{TransitionTime: 1, Brightness: uint8(rand.Intn(100) + 54)})
		time.Sleep(time.Duration(rand.Intn(200)+10) * time.Millisecond)
	}

	light.Set(completionState)
}

// CurrentState convenience that generates an equivalent State struct from a LightState struct
func CurrentState(light *Light) *State {
	lightState := light.State
	result := &State{On: lightState.On, Brightness: lightState.Brightness, Hue: lightState.Hue, Saturation: lightState.Saturation, XY: &lightState.XY, Ct: lightState.ColorTemp, Alert: lightState.Alert, Effect: lightState.Effect}

	return result
}
