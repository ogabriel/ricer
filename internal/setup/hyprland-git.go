package setup

import "github.com/ogabriel/ricer/internal"

var HyprlandGit = internal.Setup{
	Name: "hyprland-git",
	Before: func() {
		internal.RunCommand("sudo pacman -R --noconfirm xdg-desktop-portal-gtk xdg-desktop-portal-wlr")
	},
	Packages: []string{
		"hyprland-git",
		"qt5-wayland",
		"qt6-wayland",
		"xdg-desktop-portal",
		"xdg-desktop-portal-hyprland-git",
	},
}
