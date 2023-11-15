package installer

import "github.com/ogabriel/ricer/internal"

var Yay = internal.Installer{
	Name:     "yay",
	Preprend: "yay -S --noconfirm",
}
