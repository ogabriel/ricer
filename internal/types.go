package internal

type Setup struct {
	Name string
	// PackageInstaller Installer
	Packages []string
	Before   func()
	After    func()
}
