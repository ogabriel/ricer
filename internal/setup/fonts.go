package setup

import "github.com/ogabriel/ricer/internal"

var Fonts = internal.Setup{
	Name: "fonts",
	Packages: []string{
		"noto-fonts-emoji",
		"nerd-fonts-git",
	},
}
