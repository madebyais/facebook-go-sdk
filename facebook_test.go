package facebook

import (
	"testing"
)

func TestNewFacebook(t *testing.T) {
	fb := New()

	if fb.GetVersion() != `v2.10` {
		t.Fatalf(`Expected "v2.10", BUT got %s`, fb.GetVersion())
	}
}

func TestSetAppID(t *testing.T) {
	appID := `11111`
	fb := New().SetAppID(appID)

	if fb.GetAppID() != appID {
		t.Fatalf(`Expected "%s", BUT got %s`, appID, fb.GetAppID())
	}
}

func TestSetAppSecret(t *testing.T) {
	appSecret := `11111`
	fb := New().SetAppSecret(appSecret)

	if fb.GetAppSecret() != appSecret {
		t.Fatalf(`Expected "%s", BUT got %s`, appSecret, fb.GetAppSecret())
	}
}

func TestSetVersion(t *testing.T) {
	version := `v2.8`
	fb := New().SetVersion(version)

	if fb.GetVersion() != version {
		t.Fatalf(`Expected "%s", BUT got %s`, version, fb.GetVersion())
	}
}
