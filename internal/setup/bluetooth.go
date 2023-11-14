package setup

import "github.com/ogabriel/ricer/internal"

var Bluetooth = internal.Setup{
	Name: "bluetooth",
	Packages: []string{
		"bluez",
		"bluez-utils",
		"bluez-libs",
		"blueman",
		"libappindicator-gtk3",
	},
	After: func() {
		commands := []string{
			"modprobe btusb",
			"sudo systemctl enable bluetooth",
			"sudo systemctl start bluetooth",
		}

		internal.RunCommands(commands)
	},
}
