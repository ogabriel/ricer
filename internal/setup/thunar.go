package setup

import "github.com/ogabriel/ricer/internal"

var Thunar = internal.Setup{
	Name: "thunar",
	Packages: []string{
		"thunar",
		"tumbler",
	},
}
