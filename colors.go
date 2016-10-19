package hue

import (
	"errors"
)

// http://www.developers.meethue.com/documentation/hue-xy-values
// XY offers the fastest color conversion for bulbs

// convenience for red
func ColorRed(light *Light) *State {
	result, _ := Color(light, Red)
	return result
}

// convenience for green
func ColorGreen(light *Light) *State {
	result, _ := Color(light, Green)
	return result
}

// convenience for blue
func ColorBlue(light *Light) *State {
	result, _ := Color(light, Blue)
	return result
}

// convenience for yellow
func ColorYellow(light *Light) *State {
	result, _ := Color(light, Yellow)
	return result
}

// convenience for orange
func ColorOrange(light *Light) *State {
	result, _ := Color(light, Orange)
	return result
}

// convenience for purple
func ColorPurple(light *Light) *State {
	result, _ := Color(light, Purple)
	return result
}

// convenience for white
func ColorWhite(light *Light) *State {
	result, _ := Color(light, White)
	return result
}

// generate a State for the given color name
func Color(light *Light, name string) (*State, error) {
	xy, err := XY(light, name)
	return &State{XY: &xy, Saturation: 254}, err
}

// lookup XY values for the given color name
func XY(light *Light, name string) ([2]float64, error) {
	gamutMap := colorMap[name]
	var xy []float64
	if gamutMap != nil {
		switch Gamut(light) {
		case GamutA:
			xy = gamutMap[1]
		case GamutB:
			xy = gamutMap[2]
		case GamutC:
			xy = gamutMap[3]
		default:
			xy = gamutMap[1]
		}
		return [2]float64{xy[0], xy[1]}, nil
	} else {
		err := errors.New("Unknown XY color: " + name)
		return [2]float64{-1.0, -1.0}, err
	}
}

const (
	GamutUnknown = "GamutUnknown"
	GamutA       = "GamutA"
	GamutB       = "GamutB"
	GamutC       = "GamutC"
)

const (
	AliceBlue         = "Alice Blue"
	AntiqueWhite      = "Antique White"
	Aqua              = "Aqua"
	Aquamarine        = "Aquamarine"
	Azure             = "Azure"
	Beige             = "Beige"
	Bisque            = "Bisque"
	Black             = "Black"
	BlanchedAlmond    = "Blanched Almond"
	Blue              = "Blue"
	BlueViolet        = "Blue Violet"
	Brown             = "Brown"
	Burlywood         = "Burlywood"
	CadetBlue         = "Cadet Blue"
	Chartreuse        = "Chartreuse"
	Chocolate         = "Chocolate"
	Coral             = "Coral"
	Cornflower        = "Cornflower"
	Cornsilk          = "Cornsilk"
	Crimson           = "Crimson"
	Cyan              = "Cyan"
	DarkBlue          = "Dark Blue"
	DarkCyan          = "Dark Cyan"
	DarkGoldenrod     = "Dark Goldenrod"
	DarkGray          = "Dark Gray"
	DarkGreen         = "Dark Green"
	DarkKhaki         = "Dark Khaki"
	DarkMagenta       = "Dark Magenta"
	DarkOliveGreen    = "Dark Olive Green"
	DarkOrange        = "Dark Orange"
	DarkOrchid        = "Dark Orchid"
	DarkRed           = "Dark Red"
	DarkSalmon        = "Dark Salmon"
	DarkSeaGreen      = "Dark Sea Green"
	DarkSlateBlue     = "Dark Slate Blue"
	DarkSlateGray     = "Dark Slate Gray"
	DarkTurquoise     = "Dark Turquoise"
	DarkViolet        = "Dark Violet"
	DeepPink          = "Deep Pink"
	DeepSkyBlue       = "Deep Sky Blue"
	DimGray           = "Dim Gray"
	DodgerBlue        = "Dodger Blue"
	Firebrick         = "Firebrick"
	FloralWhite       = "Floral White"
	ForestGreen       = "Forest Green"
	Fuchsia           = "Fuchsia"
	Gainsboro         = "Gainsboro"
	GhostWhite        = "Ghost White"
	Gold              = "Gold"
	Goldenrod         = "Goldenrod"
	Gray              = "Gray"
	WebGray           = "Web Gray"
	Green             = "Green"
	WebGreen          = "Web Green"
	GreenYellow       = "Green Yellow"
	Honeydew          = "Honeydew"
	HotPink           = "Hot Pink"
	IndianRed         = "Indian Red"
	Indigo            = "Indigo"
	Ivory             = "Ivory"
	Khaki             = "Khaki"
	Lavender          = "Lavender"
	LavenderBlush     = "Lavender Blush"
	LawnGreen         = "Lawn Green"
	LemonChiffon      = "Lemon Chiffon"
	LightBlue         = "Light Blue"
	LightCoral        = "Light Coral"
	LightCyan         = "Light Cyan"
	LightGoldenrod    = "Light Goldenrod"
	LightGray         = "Light Gray"
	LightGreen        = "Light Green"
	LightPink         = "Light Pink"
	LightSalmon       = "Light Salmon"
	LightSeaGreen     = "Light Sea Green"
	LightSkyBlue      = "Light Sky Blue"
	LightSlateGray    = "Light Slate Gray"
	LightSteelBlue    = "Light Steel Blue"
	LightYellow       = "Light Yellow"
	Lime              = "Lime"
	LimeGreen         = "Lime Green"
	Linen             = "Linen"
	Magenta           = "Magenta"
	Maroon            = "Maroon"
	WebMaroon         = "Web Maroon"
	MediumAquamarine  = "Medium Aquamarine"
	MediumBlue        = "Medium Blue"
	MediumOrchid      = "Medium Orchid"
	MediumPurple      = "Medium Purple"
	MediumSeaGreen    = "Medium Sea Green"
	MediumSlateBlue   = "Medium Slate Blue"
	MediumSpringGreen = "Medium Spring Green"
	MediumTurquoise   = "Medium Turquoise"
	MediumVioletRed   = "Medium Violet Red"
	MidnightBlue      = "Midnight Blue"
	MintCream         = "Mint Cream"
	MistyRose         = "Misty Rose"
	Moccasin          = "Moccasin"
	NavajoWhite       = "Navajo White"
	NavyBlue          = "Navy Blue"
	OldLace           = "Old Lace"
	Olive             = "Olive"
	OliveDrab         = "Olive Drab"
	Orange            = "Orange"
	OrangeRed         = "Orange Red"
	Orchid            = "Orchid"
	PaleGoldenrod     = "Pale Goldenrod"
	PaleGreen         = "Pale Green"
	PaleTurquoise     = "Pale Turquoise"
	PaleVioletRed     = "Pale Violet Red"
	PapayaWhip        = "Papaya Whip"
	PeachPuff         = "Peach Puff"
	Peru              = "Peru"
	Pink              = "Pink"
	Plum              = "Plum"
	PowderBlue        = "Powder Blue"
	Purple            = "Purple"
	WebPurple         = "Web Purple"
	RebeccaPurple     = "Rebecca Purple"
	Red               = "Red"
	RosyBrown         = "Rosy Brown"
	RoyalBlue         = "Royal Blue"
	SaddleBrown       = "Saddle Brown"
	Salmon            = "Salmon"
	SandyBrown        = "Sandy Brown"
	SeaGreen          = "Sea Green"
	Seashell          = "Seashell"
	Sienna            = "Sienna"
	Silver            = "Silver"
	SkyBlue           = "Sky Blue"
	SlateBlue         = "Slate Blue"
	SlateGray         = "Slate Gray"
	Snow              = "Snow"
	SpringGreen       = "Spring Green"
	SteelBlue         = "Steel Blue"
	Tan               = "Tan"
	Teal              = "Teal"
	Thistle           = "Thistle"
	Tomato            = "Tomato"
	Turquoise         = "Turquoise"
	Violet            = "Violet"
	Wheat             = "Wheat"
	White             = "White"
	WhiteSmoke        = "White Smoke"
	Yellow            = "Yellow"
	YellowGreen       = "Yellow Green"
)

// hashmap of XY color values as array {rgb, GamutA GamutB, GamutC}
var colorMap = map[string][][]float64{
	AliceBlue:         {{239, 247, 255}, {0.3088, 0.3212}, {0.3092, 0.321}, {0.3088, 0.3212}},
	AntiqueWhite:      {{249, 234, 214}, {0.3548, 0.3489}, {0.3548, 0.3489}, {0.3548, 0.3489}},
	Aqua:              {{0, 255, 255}, {0.17, 0.3403}, {0.2858, 0.2747}, {0.1607, 0.3423}},
	Aquamarine:        {{127, 255, 211}, {0.2138, 0.4051}, {0.3237, 0.3497}, {0.2138, 0.4051}},
	Azure:             {{239, 255, 255}, {0.3059, 0.3303}, {0.3123, 0.3271}, {0.3059, 0.3303}},
	Beige:             {{244, 244, 219}, {0.3402, 0.356}, {0.3402, 0.356}, {0.3402, 0.356}},
	Bisque:            {{255, 226, 196}, {0.3806, 0.3576}, {0.3806, 0.3576}, {0.3806, 0.3576}},
	Black:             {{0, 0, 0}, {0.139, 0.081}, {0.168, 0.041}, {0.153, 0.048}},
	BlanchedAlmond:    {{255, 234, 204}, {0.3695, 0.3584}, {0.3695, 0.3584}, {0.3695, 0.3584}},
	Blue:              {{0, 0, 255}, {0.139, 0.081}, {0.168, 0.041}, {0.153, 0.048}},
	BlueViolet:        {{137, 43, 226}, {0.245, 0.1214}, {0.251, 0.1056}, {0.251, 0.1056}},
	Brown:             {{165, 40, 40}, {0.6399, 0.3041}, {0.6399, 0.3041}, {0.6399, 0.3041}},
	Burlywood:         {{221, 183, 135}, {0.4236, 0.3811}, {0.4236, 0.3811}, {0.4236, 0.3811}},
	CadetBlue:         {{94, 158, 160}, {0.2211, 0.3328}, {0.2961, 0.295}, {0.2211, 0.3328}},
	Chartreuse:        {{127, 255, 0}, {0.2682, 0.6632}, {0.408, 0.517}, {0.2505, 0.6395}},
	Chocolate:         {{209, 104, 30}, {0.6009, 0.3684}, {0.6009, 0.3684}, {0.6009, 0.3684}},
	Coral:             {{255, 127, 79}, {0.5763, 0.3486}, {0.5763, 0.3486}, {0.5763, 0.3486}},
	Cornflower:        {{99, 147, 237}, {0.1905, 0.1945}, {0.2343, 0.1725}, {0.1905, 0.1945}},
	Cornsilk:          {{255, 247, 219}, {0.3511, 0.3574}, {0.3511, 0.3574}, {0.3511, 0.3574}},
	Crimson:           {{219, 20, 61}, {0.6531, 0.2834}, {0.6417, 0.304}, {0.6508, 0.2881}},
	Cyan:              {{0, 255, 255}, {0.17, 0.3403}, {0.2858, 0.2747}, {0.1607, 0.3423}},
	DarkBlue:          {{0, 0, 140}, {0.139, 0.081}, {0.168, 0.041}, {0.153, 0.048}},
	DarkCyan:          {{0, 140, 140}, {0.17, 0.3403}, {0.2858, 0.2747}, {0.1607, 0.3423}},
	DarkGoldenrod:     {{183, 135, 10}, {0.5265, 0.4428}, {0.5204, 0.4346}, {0.5214, 0.4361}},
	DarkGray:          {{168, 168, 168}, {0.3227, 0.329}, {0.3227, 0.329}, {0.3227, 0.329}},
	DarkGreen:         {{0, 99, 0}, {0.214, 0.709}, {0.408, 0.517}, {0.17, 0.7}},
	DarkKhaki:         {{188, 183, 107}, {0.4004, 0.4331}, {0.4004, 0.4331}, {0.4004, 0.4331}},
	DarkMagenta:       {{140, 0, 140}, {0.3787, 0.1724}, {0.3824, 0.1601}, {0.3833, 0.1591}},
	DarkOliveGreen:    {{84, 107, 45}, {0.3475, 0.5047}, {0.3908, 0.4829}, {0.3475, 0.5047}},
	DarkOrange:        {{255, 140, 0}, {0.5951, 0.3872}, {0.5916, 0.3824}, {0.5921, 0.3831}},
	DarkOrchid:        {{153, 51, 204}, {0.296, 0.1409}, {0.2986, 0.1341}, {0.2986, 0.1341}},
	DarkRed:           {{140, 0, 0}, {0.7, 0.2986}, {0.674, 0.322}, {0.692, 0.308}},
	DarkSalmon:        {{232, 150, 122}, {0.4837, 0.3479}, {0.4837, 0.3479}, {0.4837, 0.3479}},
	DarkSeaGreen:      {{142, 188, 142}, {0.2924, 0.4134}, {0.3429, 0.3879}, {0.2924, 0.4134}},
	DarkSlateBlue:     {{71, 61, 140}, {0.2206, 0.1484}, {0.2218, 0.1477}, {0.2206, 0.1484}},
	DarkSlateGray:     {{45, 79, 79}, {0.2239, 0.3368}, {0.2982, 0.2993}, {0.2239, 0.3368}},
	DarkTurquoise:     {{0, 206, 209}, {0.1693, 0.3347}, {0.2835, 0.2701}, {0.1605, 0.3366}},
	DarkViolet:        {{147, 0, 211}, {0.2742, 0.1326}, {0.2836, 0.1079}, {0.2824, 0.1104}},
	DeepPink:          {{255, 20, 147}, {0.5454, 0.2359}, {0.5386, 0.2468}, {0.5445, 0.2369}},
	DeepSkyBlue:       {{0, 191, 255}, {0.1576, 0.2368}, {0.2428, 0.1893}, {0.158, 0.2379}},
	DimGray:           {{104, 104, 104}, {0.3227, 0.329}, {0.3227, 0.329}, {0.3227, 0.329}},
	DodgerBlue:        {{30, 142, 255}, {0.1484, 0.1599}, {0.2115, 0.1273}, {0.1559, 0.1599}},
	Firebrick:         {{178, 33, 33}, {0.6621, 0.3023}, {0.6566, 0.3123}, {0.6621, 0.3023}},
	FloralWhite:       {{255, 249, 239}, {0.3361, 0.3388}, {0.3361, 0.3388}, {0.3361, 0.3388}},
	ForestGreen:       {{33, 140, 33}, {0.2097, 0.6732}, {0.408, 0.517}, {0.1984, 0.6746}},
	Fuchsia:           {{255, 0, 255}, {0.3787, 0.1724}, {0.3824, 0.1601}, {0.3833, 0.1591}},
	Gainsboro:         {{219, 219, 219}, {0.3227, 0.329}, {0.3227, 0.329}, {0.3227, 0.329}},
	GhostWhite:        {{247, 247, 255}, {0.3174, 0.3207}, {0.3174, 0.3207}, {0.3174, 0.3207}},
	Gold:              {{255, 214, 0}, {0.4947, 0.472}, {0.4859, 0.4599}, {0.4871, 0.4618}},
	Goldenrod:         {{216, 165, 33}, {0.5136, 0.4444}, {0.5113, 0.4413}, {0.5125, 0.4428}},
	Gray:              {{191, 191, 191}, {0.3227, 0.329}, {0.3227, 0.329}, {0.3227, 0.329}},
	WebGray:           {{127, 127, 127}, {0.3227, 0.329}, {0.3227, 0.329}, {0.3227, 0.329}},
	Green:             {{0, 255, 0}, {0.214, 0.709}, {0.408, 0.517}, {0.17, 0.7}},
	WebGreen:          {{0, 127, 0}, {0.214, 0.709}, {0.408, 0.517}, {0.17, 0.7}},
	GreenYellow:       {{173, 255, 45}, {0.3298, 0.5959}, {0.408, 0.517}, {0.3221, 0.5857}},
	Honeydew:          {{239, 255, 239}, {0.316, 0.3477}, {0.3213, 0.345}, {0.316, 0.3477}},
	HotPink:           {{255, 104, 181}, {0.4682, 0.2452}, {0.4682, 0.2452}, {0.4682, 0.2452}},
	IndianRed:         {{204, 91, 91}, {0.5488, 0.3112}, {0.5488, 0.3112}, {0.5488, 0.3112}},
	Indigo:            {{73, 0, 130}, {0.2332, 0.1169}, {0.2437, 0.0895}, {0.2428, 0.0913}},
	Ivory:             {{255, 255, 239}, {0.3334, 0.3455}, {0.3334, 0.3455}, {0.3334, 0.3455}},
	Khaki:             {{239, 229, 140}, {0.4019, 0.4261}, {0.4019, 0.4261}, {0.4019, 0.4261}},
	Lavender:          {{229, 229, 249}, {0.3085, 0.3071}, {0.3085, 0.3071}, {0.3085, 0.3071}},
	LavenderBlush:     {{255, 239, 244}, {0.3369, 0.3225}, {0.3369, 0.3225}, {0.3369, 0.3225}},
	LawnGreen:         {{124, 252, 0}, {0.2663, 0.6649}, {0.408, 0.517}, {0.2485, 0.641}},
	LemonChiffon:      {{255, 249, 204}, {0.3608, 0.3756}, {0.3608, 0.3756}, {0.3608, 0.3756}},
	LightBlue:         {{173, 216, 229}, {0.2621, 0.3157}, {0.2975, 0.2979}, {0.2621, 0.3157}},
	LightCoral:        {{239, 127, 127}, {0.5075, 0.3145}, {0.5075, 0.3145}, {0.5075, 0.3145}},
	LightCyan:         {{224, 255, 255}, {0.2901, 0.3316}, {0.3096, 0.3218}, {0.2901, 0.3316}},
	LightGoldenrod:    {{249, 249, 209}, {0.3504, 0.3717}, {0.3504, 0.3717}, {0.3504, 0.3717}},
	LightGray:         {{211, 211, 211}, {0.3227, 0.329}, {0.3227, 0.329}, {0.3227, 0.329}},
	LightGreen:        {{142, 237, 142}, {0.2648, 0.4901}, {0.3682, 0.438}, {0.2648, 0.4901}},
	LightPink:         {{255, 181, 193}, {0.4112, 0.3091}, {0.4112, 0.3091}, {0.4112, 0.3091}},
	LightSalmon:       {{255, 160, 122}, {0.5016, 0.3531}, {0.5016, 0.3531}, {0.5016, 0.3531}},
	LightSeaGreen:     {{33, 178, 170}, {0.1721, 0.358}, {0.2946, 0.292}, {0.1611, 0.3593}},
	LightSkyBlue:      {{135, 206, 249}, {0.214, 0.2749}, {0.2714, 0.246}, {0.214, 0.2749}},
	LightSlateGray:    {{119, 135, 153}, {0.2738, 0.297}, {0.2924, 0.2877}, {0.2738, 0.297}},
	LightSteelBlue:    {{175, 196, 221}, {0.276, 0.2975}, {0.293, 0.2889}, {0.276, 0.2975}},
	LightYellow:       {{255, 255, 224}, {0.3436, 0.3612}, {0.3436, 0.3612}, {0.3436, 0.3612}},
	Lime:              {{0, 255, 0}, {0.214, 0.709}, {0.408, 0.517}, {0.17, 0.7}},
	LimeGreen:         {{51, 204, 51}, {0.2101, 0.6765}, {0.408, 0.517}, {0.1972, 0.6781}},
	Linen:             {{249, 239, 229}, {0.3411, 0.3387}, {0.3411, 0.3387}, {0.3411, 0.3387}},
	Magenta:           {{255, 0, 255}, {0.3787, 0.1724}, {0.3824, 0.1601}, {0.3833, 0.1591}},
	Maroon:            {{175, 48, 96}, {0.5383, 0.2566}, {0.5383, 0.2566}, {0.5383, 0.2566}},
	WebMaroon:         {{127, 0, 0}, {0.7, 0.2986}, {0.674, 0.322}, {0.692, 0.308}},
	MediumAquamarine:  {{102, 204, 170}, {0.215, 0.4014}, {0.3224, 0.3473}, {0.215, 0.4014}},
	MediumBlue:        {{0, 0, 204}, {0.139, 0.081}, {0.168, 0.041}, {0.153, 0.048}},
	MediumOrchid:      {{186, 84, 211}, {0.3365, 0.1735}, {0.3365, 0.1735}, {0.3365, 0.1735}},
	MediumPurple:      {{147, 112, 219}, {0.263, 0.1773}, {0.263, 0.1773}, {0.263, 0.1773}},
	MediumSeaGreen:    {{61, 178, 112}, {0.1979, 0.5005}, {0.3588, 0.4194}, {0.1979, 0.5005}},
	MediumSlateBlue:   {{122, 104, 237}, {0.2179, 0.1424}, {0.2189, 0.1419}, {0.2179, 0.1424}},
	MediumSpringGreen: {{0, 249, 153}, {0.1919, 0.524}, {0.3622, 0.4262}, {0.1655, 0.5275}},
	MediumTurquoise:   {{71, 209, 204}, {0.176, 0.3496}, {0.2937, 0.2903}, {0.176, 0.3496}},
	MediumVioletRed:   {{198, 20, 132}, {0.504, 0.2201}, {0.5002, 0.2255}, {0.5047, 0.2177}},
	MidnightBlue:      {{25, 25, 112}, {0.1585, 0.0884}, {0.1825, 0.0697}, {0.1616, 0.0802}},
	MintCream:         {{244, 255, 249}, {0.315, 0.3363}, {0.3165, 0.3355}, {0.315, 0.3363}},
	MistyRose:         {{255, 226, 224}, {0.3581, 0.3284}, {0.3581, 0.3284}, {0.3581, 0.3284}},
	Moccasin:          {{255, 226, 181}, {0.3927, 0.3732}, {0.3927, 0.3732}, {0.3927, 0.3732}},
	NavajoWhite:       {{255, 221, 173}, {0.4027, 0.3757}, {0.4027, 0.3757}, {0.4027, 0.3757}},
	NavyBlue:          {{0, 0, 127}, {0.139, 0.081}, {0.168, 0.041}, {0.153, 0.048}},
	OldLace:           {{252, 244, 229}, {0.3421, 0.344}, {0.3421, 0.344}, {0.3421, 0.344}},
	Olive:             {{127, 127, 0}, {0.4432, 0.5154}, {0.4317, 0.4996}, {0.4334, 0.5022}},
	OliveDrab:         {{107, 142, 35}, {0.354, 0.5561}, {0.408, 0.517}, {0.354, 0.5561}},
	Orange:            {{255, 165, 0}, {0.5614, 0.4156}, {0.5562, 0.4084}, {0.5569, 0.4095}},
	OrangeRed:         {{255, 68, 0}, {0.6726, 0.3217}, {0.6733, 0.3224}, {0.6731, 0.3222}},
	Orchid:            {{216, 112, 214}, {0.3688, 0.2095}, {0.3688, 0.2095}, {0.3688, 0.2095}},
	PaleGoldenrod:     {{237, 232, 170}, {0.3751, 0.3983}, {0.3751, 0.3983}, {0.3751, 0.3983}},
	PaleGreen:         {{153, 249, 153}, {0.2675, 0.4826}, {0.3657, 0.4331}, {0.2675, 0.4826}},
	PaleTurquoise:     {{175, 237, 237}, {0.2539, 0.3344}, {0.3034, 0.3095}, {0.2539, 0.3344}},
	PaleVioletRed:     {{219, 112, 147}, {0.4658, 0.2773}, {0.4658, 0.2773}, {0.4658, 0.2773}},
	PapayaWhip:        {{255, 239, 214}, {0.3591, 0.3536}, {0.3591, 0.3536}, {0.3591, 0.3536}},
	PeachPuff:         {{255, 216, 186}, {0.3953, 0.3564}, {0.3953, 0.3564}, {0.3953, 0.3564}},
	Peru:              {{204, 132, 63}, {0.5305, 0.3911}, {0.5305, 0.3911}, {0.5305, 0.3911}},
	Pink:              {{255, 191, 204}, {0.3944, 0.3093}, {0.3944, 0.3093}, {0.3944, 0.3093}},
	Plum:              {{221, 160, 221}, {0.3495, 0.2545}, {0.3495, 0.2545}, {0.3495, 0.2545}},
	PowderBlue:        {{175, 224, 229}, {0.262, 0.3269}, {0.302, 0.3068}, {0.262, 0.3269}},
	Purple:            {{160, 33, 239}, {0.2651, 0.1291}, {0.2725, 0.1096}, {0.2725, 0.1096}},
	WebPurple:         {{127, 0, 127}, {0.3787, 0.1724}, {0.3824, 0.1601}, {0.3833, 0.1591}},
	RebeccaPurple:     {{102, 51, 153}, {0.2703, 0.1398}, {0.2703, 0.1398}, {0.2703, 0.1398}},
	Red:               {{255, 0, 0}, {0.7, 0.2986}, {0.674, 0.322}, {0.692, 0.308}},
	RosyBrown:         {{188, 142, 142}, {0.4026, 0.3227}, {0.4026, 0.3227}, {0.4026, 0.3227}},
	RoyalBlue:         {{63, 104, 224}, {0.1649, 0.1338}, {0.2047, 0.1138}, {0.1649, 0.1338}},
	SaddleBrown:       {{140, 68, 17}, {0.5993, 0.369}, {0.5993, 0.369}, {0.5993, 0.369}},
	Salmon:            {{249, 127, 114}, {0.5346, 0.3247}, {0.5346, 0.3247}, {0.5346, 0.3247}},
	SandyBrown:        {{244, 163, 96}, {0.5104, 0.3826}, {0.5104, 0.3826}, {0.5104, 0.3826}},
	SeaGreen:          {{45, 140, 86}, {0.1968, 0.5047}, {0.3602, 0.4223}, {0.1968, 0.5047}},
	Seashell:          {{255, 244, 237}, {0.3397, 0.3353}, {0.3397, 0.3353}, {0.3397, 0.3353}},
	Sienna:            {{160, 81, 45}, {0.5714, 0.3559}, {0.5714, 0.3559}, {0.5714, 0.3559}},
	Silver:            {{191, 191, 191}, {0.3227, 0.329}, {0.3227, 0.329}, {0.3227, 0.329}},
	SkyBlue:           {{135, 206, 234}, {0.2206, 0.2948}, {0.2807, 0.2645}, {0.2206, 0.2948}},
	SlateBlue:         {{107, 89, 204}, {0.2218, 0.1444}, {0.2218, 0.1444}, {0.2218, 0.1444}},
	SlateGray:         {{112, 127, 142}, {0.2762, 0.3009}, {0.2944, 0.2918}, {0.2762, 0.3009}},
	Snow:              {{255, 249, 249}, {0.3292, 0.3285}, {0.3292, 0.3285}, {0.3292, 0.3285}},
	SpringGreen:       {{0, 255, 127}, {0.1994, 0.5864}, {0.3882, 0.4777}, {0.1671, 0.5906}},
	SteelBlue:         {{68, 130, 181}, {0.183, 0.2325}, {0.248, 0.1997}, {0.183, 0.2325}},
	Tan:               {{209, 181, 140}, {0.4035, 0.3772}, {0.4035, 0.3772}, {0.4035, 0.3772}},
	Teal:              {{0, 127, 127}, {0.17, 0.3403}, {0.2858, 0.2747}, {0.1607, 0.3423}},
	Thistle:           {{216, 191, 216}, {0.3342, 0.2971}, {0.3342, 0.2971}, {0.3342, 0.2971}},
	Tomato:            {{255, 99, 71}, {0.6112, 0.3261}, {0.6112, 0.3261}, {0.6112, 0.3261}},
	Turquoise:         {{63, 224, 209}, {0.1732, 0.3672}, {0.2997, 0.3022}, {0.1702, 0.3675}},
	Violet:            {{237, 130, 237}, {0.3644, 0.2133}, {0.3644, 0.2133}, {0.3644, 0.2133}},
	Wheat:             {{244, 221, 178}, {0.3852, 0.3737}, {0.3852, 0.3737}, {0.3852, 0.3737}},
	White:             {{255, 255, 255}, {0.3227, 0.329}, {0.3227, 0.329}, {0.3227, 0.329}},
	WhiteSmoke:        {{244, 244, 244}, {0.3227, 0.329}, {0.3227, 0.329}, {0.3227, 0.329}},
	Yellow:            {{255, 255, 0}, {0.4432, 0.5154}, {0.4317, 0.4996}, {0.4334, 0.5022}},
	YellowGreen:       {{153, 204, 51}, {0.3517, 0.5618}, {0.408, 0.517}, {0.3517, 0.5618}},
}
