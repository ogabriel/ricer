package profile

import (
	"github.com/ogabriel/ricer/internal"
	"github.com/ogabriel/ricer/internal/setup"
)

var extraHyprland = internal.Setup{
	Name: "extra hyprland",
	Packages: []string{
		"mako",
		"alacritty",
		"waybar",
		"swaylock",
		"swayidle",
		"rofi-lbonn-wayland",
	},
}

var Hyprland = internal.Profile{
	Name: "hyprland",
	Setups: []internal.Setup{
		setup.Pacman,
		setup.Yay,
		setup.Sddm,
		setup.Audio,
		setup.Bluetooth,
		setup.Network,
		setup.Screen,
		setup.Fonts,
		setup.Thunar,
		setup.Gnome,
		setup.Hyprland,
		extraHyprland,
	},
}
