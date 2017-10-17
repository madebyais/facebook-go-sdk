package facebook

import "testing"

func TestNewGraph(t *testing.T) {
	graph := new(Graph).SetVersion(``)

	url := graph.GenerateURL(`/me`)
	expectedURL := `https://graph.facebook.com/v2.10/me`
	if url != expectedURL {
		t.Fatalf(`Expected %s, BUT got %s`, expectedURL, url)
	}
}

func TestGetVersion(t *testing.T) {
	graph := new(Graph).SetVersion(``)
	if graph.GetVersion() != `v2.10` {
		t.Fatalf(`Expected "v2.10", BUT got "%s"`, graph.GetVersion())
	}
}

func TestGetGraphURL(t *testing.T) {
	graph := new(Graph).SetVersion(``)
	if graph.GetGraphURL() != `https://graph.facebook.com` {
		t.Fatalf(`Expected "https://graph.facebook.com", BUT got "%s"`, graph.GetGraphURL())
	}
}

func TestGenerateURL(t *testing.T) {
	version := `v2.10`
	graph := new(Graph).SetVersion(version)

	url := graph.GenerateURL(`/me`)
	expectedURL := `https://graph.facebook.com/` + version + `/me`
	if url != expectedURL {
		t.Fatalf(`Expected "%s", BUT got "%s"`, expectedURL, url)
	}
}

func TestGenerateURLWithFields(t *testing.T) {
	version := `v2.10`
	graph := new(Graph).SetVersion(version)

	url := graph.GenerateURLWithFields(`/me`, `id,name`)
	expectedURL := `https://graph.facebook.com/` + version + `/me?fields=id,name`
	if url != expectedURL {
		t.Fatalf(`Expected "%s", BUT got "%s"`, expectedURL, url)
	}
}

func TestGenerateRawURL(t *testing.T) {
	version := `v2.10`
	graph := new(Graph).SetVersion(version)

	url := graph.GenerateRawURL(`/me`, `?fields=id,name,birthday`)
	expectedURL := `https://graph.facebook.com/` + version + `/me?fields=id,name,birthday`
	if url != expectedURL {
		t.Fatalf(`Expected "%s", BUT got "%s"`, expectedURL, url)
	}
}
