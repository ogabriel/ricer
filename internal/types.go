package internal

type Profile struct {
	Name   string
	Setups []Setup
}

type Setup struct {
	Name string
	// PackageInstaller Installer
	Packages  []string
	Services  []string
	Installer Installer
	Before    func()
	After     func()
}

type Installer struct {
	Name     string
	Preprend string
}
