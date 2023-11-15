package setup

import "github.com/ogabriel/ricer/internal"

var Hyprland = internal.Setup{
	Name: "hyprland",
	Before: func() {
		internal.RunCommand("sudo pacman -R xdg-desktop-portal-gtk xdg-desktop-portal-wlr")
	},
	Packages: []string{
		"hyprland",
		"qt5-wayland",
		"qt6-wayland",
		"xdg-desktop-portal",
		"xdg-desktop-portal-hyprland",
	},
}
