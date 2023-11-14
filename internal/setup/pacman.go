package setup

import "github.com/ogabriel/ricer/internal"

var Pacman = internal.Setup{
	Name: "pacman",
	Before: func() {
		commands := []string{
			"sudo sed -i 's/#ILoveCandy/ILoveCandy/' /etc/pacman.conf",
			"sudo sed -i 's/#ParallelDownloads = 5/ParallelDownloads = 5/' /etc/pacman.conf",
		}

		internal.RunCommands(commands)
	},
}
