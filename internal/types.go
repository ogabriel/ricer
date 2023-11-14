package internal

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
