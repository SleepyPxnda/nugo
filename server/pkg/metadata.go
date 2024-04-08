package pkg

type Metadata struct {
	Id                       int
	Version                  string
	Title                    string
	Authors                  string
	Owners                   string
	LicenseUrl               string
	ProjectUrl               string
	IconUrl                  string
	RequireLicenseAcceptance bool
	Description              string
	ReleaseNotes             string
	Tags                     []string
	Dependencies             []Dependency
}
