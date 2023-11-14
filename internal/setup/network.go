package setup

import "github.com/ogabriel/ricer/internal"

var Network = internal.Setup{
	Name: "network",
	Packages: []string{
		"networkmanager",
		"network-manager-applet",
		"libappindicator-gtk3",
	},
	After: func() {
		commands := []string{
			"sudo systemctl enable NetworkManager",
			"sudo systemctl start NetworkManager",
		}

		internal.RunCommands(commands)
	},
}
