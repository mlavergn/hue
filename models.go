package hue

// http://www.developers.meethue.com/documentation/supported-lights

// return the model name for the given Light
func Model(light *Light) string {
	info := bulbs[light.ModelID]
	if info != nil {
		return info[0]
	} else {
		return ""
	}
}

// return the gamut type for the given Light
func Gamut(light *Light) string {
	info := bulbs[light.ModelID]
	if info != nil {
		return info[1]
	} else {
		return GamutUnknown
	}
}

// hashmap of model info as array {model, gamut}
// NOTE: Lux bulbs don't have a gamut so return blank
var bulbs = map[string][]string{
	"LCT001": {"Hue A19", GamutA},
	"LCT002": {"Hue BR30", GamutA},
	"LCT003": {"Hue GU10", GamutA},
	"LCT007": {"Hue A19", GamutB},
	"LLC005": {"Bloom", GamutA},
	"LLC006": {"Iris", GamutA},
	"LLC007": {"Bloom", GamutA},
	"LLC011": {"Bloom", GamutA},
	"LLC012": {"Bloom", GamutA},
	"LLC013": {"Storylight", GamutA},
	"LST001": {"Light Strip", GamutA},
	"LWB004": {"Lux A19", GamutUnknown},
}
