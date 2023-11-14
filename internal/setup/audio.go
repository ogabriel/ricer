package setup

import "github.com/ogabriel/ricer/internal"

var Audio = internal.Setup{
	Name: "audio",
	Packages: []string{
		"pipewire",
		"pipewire-alsa",
		"pipewire-pulse",
		"libpulse",
		"pipewire-audio",
		"pipewire-jack",
		"gst-plugin-pipewire",
		"libpipewire",
		"wireplumber",
		"libwireplumber",
		"pavucontrol",
		"playerctl",
	},
}
