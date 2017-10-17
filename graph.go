package facebook

// Graph represents a default configuration
// for facebook graph api
type Graph struct {
	Version string
	URL     string
}

// GraphInterface represents all methods for using
// Facebook Graph API
type GraphInterface interface {
	SetVersion(version string) GraphInterface
	GetVersion() string
	GetGraphURL() string

	GenerateURL(path string) string
	GenerateURLWithFields(path string, fields string) string
	GenerateRawURL(path string, querystring string) string
}

// SetVersion is use to set the Facebook Graph API version
// and also the default Graph API url
func (g *Graph) SetVersion(version string) GraphInterface {
	graphURL := `https://graph.facebook.com`
	defaultVersion := `v2.10`

	if version == "" {
		version = defaultVersion
	}

	return &Graph{
		Version: version,
		URL:     graphURL,
	}
}

// GetVersion is use to get current version
// used by graph functions
func (g *Graph) GetVersion() string {
	return g.Version
}

// GetGraphURL is use to get current graph API URL
func (g *Graph) GetGraphURL() string {
	return g.URL
}

// GenerateURL is use to generate a URL
// to be used when requesting to Facebook Graph API
func (g *Graph) GenerateURL(path string) string {
	return g.URL + `/` + g.Version + path
}

// GenerateURLWithFields is use to generate a URL
// to be used when accessing Facebook Graph API with `fields` as query string
func (g *Graph) GenerateURLWithFields(path string, fields string) string {
	return g.GenerateURL(path) + `?fields=` + fields
}

// GenerateRawURL is use to generate a URL
// with raw query string
func (g *Graph) GenerateRawURL(path string, querystring string) string {
	return g.GenerateURL(path) + querystring
}
