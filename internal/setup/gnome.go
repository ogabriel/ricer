package setup

import "github.com/ogabriel/ricer/internal"

var Gnome = internal.Setup{
	Name: "gnome",
	Packages: []string{
		"gnome-calculator",
		"gnome-keyring",
		"polkit-gnome",
	},
}
