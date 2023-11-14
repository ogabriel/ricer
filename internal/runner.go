package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ogabriel/ricer/internal/setup"
)

func Run(commands []string) {
	command := strings.Join(commands, ";")
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func RunSetup(setup setup.Setup) {
	if setup.Before != nil {
		fmt.Println(setup.Name)
		setup.Before()
	}

	if setup.Packages != nil {
		fmt.Println(setup.Name)
		Run(setup.Packages)
	}

	if setup.After != nil {
		fmt.Println(setup.Name)
		setup.After()
	}
}

func RunSetups(setups []setup.Setup) {
	for _, setup := range setups {
		if setup.Before != nil {
			fmt.Println(setup.Name)
			setup.Before()
		}
	}

	var packages []string
	for _, setup := range setups {
		if setup.Packages != nil {
			packages = append(packages, setup.Packages...)
		}
	}

	if packages != nil {
		Run(packages)
	}

	for _, setup := range setups {
		if setup.After != nil {
			fmt.Println(setup.Name)
			setup.After()
		}
	}

}
