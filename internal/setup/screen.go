package setup

import "github.com/ogabriel/ricer/internal"

var Screen = internal.Setup{
	Name: "screen",
	Packages: []string{
		"gammastep",
		"brightnessctl",
		"grim",
		"slurp",
		"flameshot-git",
	},
}
