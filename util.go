package hue

import (
	"log"
	"os"
	"os/user"
	"path"	
)

// GetBridge returns a reference to a Bridge
func GetBridge() (*Bridge, error) {
	b, err := Discover()
	if err != nil {
		log.Fatal(err)
	}
	if !b.IsPaired() {
		// link button must be pressed before calling
		if err := b.Pair(); err != nil {
			log.Fatal(err)
		}
	}

	return b, err
}

// ResetBridge removes the hue configuration for re-pairing with the bridge
func ResetBridge() {
	usr, err := user.Current()
	if err == nil {
		hueCfg := path.Join(usr.HomeDir, ".hue")
		_, err = os.Stat(hueCfg)
		if !os.IsNotExist(err) {
			os.Remove(hueCfg)
		}
	}
}

// Describe prints a status reports on known lights to stdout
func Describe(b *Bridge) {
	lights, _ := b.Lights().List()
	println("Count:", len(lights))
	for _, light := range lights {
		println(light.Name)
		println("\t", light.ID)
		println("\t", light.ModelID, light.Type)
		println("\t", Model(light))
		println("\t Reachable:", light.State.Reachable)
		println("\t State:", light.State.On)
	}
}

// LightMap returns a map of lights to names
func LightMap(b *Bridge) (map[string]*Light, error) {
	lights := make(map[string]*Light)
	lightList, err := b.Lights().List()
	for _, light := range lightList {
		lights[light.Name] = light
	}

	return lights, err
}
