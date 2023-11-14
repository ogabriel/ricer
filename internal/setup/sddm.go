package setup

import "github.com/ogabriel/ricer/internal"

var Sddm = internal.Setup{
	Name: "sddm",
	Packages: []string{
		"sddm",
	},
	After: func() {
		internal.RunCommand("sudo systemctl enable sddm")
	},
}
