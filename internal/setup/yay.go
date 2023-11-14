package setup

import "github.com/ogabriel/ricer/internal"

var Yay = internal.Setup{
	Name: "yay",
	Before: func() {
		commands := []string{
			"sudo pacman -S --needed --no-confirm git base-devel",
			"cd /tmp",
			"git clone https://aur.archlinux.org/yay-bin.git",
			"cd yay-bin",
			"makepkg -si",
		}

		internal.Run(commands)
	},
}
