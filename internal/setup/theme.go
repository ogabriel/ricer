package setup

import "github.com/ogabriel/ricer/internal"

var Theme = internal.Setup{
	Name: "theme",
	Packages: []string{
		"wpgtk",
		"python-pywal",
	},
}
