package userdir

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const app = "testapp"

func TestUserdirApp(t *testing.T) {
	// Save the original HOME environment variable value
	originalHome, err := os.UserHomeDir()

	if err != nil {
		t.Errorf("os.UserHomeDir returned an error: %v", err)
	}

	// Create the expected paths for the config, data, and cache directories
	var expectedConfig, expectedData, expectedCache string
	if runtime.GOOS == "windows" {
		expectedConfig = filepath.Join(originalHome, "AppData", "Roaming", app)
		expectedData = filepath.Join(originalHome, "AppData", "Roaming", app, "data")
		expectedCache = filepath.Join(originalHome, "AppData", "Local", app)
	} else {
		expectedConfig = filepath.Join(originalHome, ".config", app)
		expectedData = filepath.Join(originalHome, ".local", "share", app)
		expectedCache = filepath.Join(originalHome, ".cache", app)
	}

	// Create a new userdirDefault instance
	u := App(app)

	// Test the Home method
	home, err := u.Home()
	if err != nil {
		t.Errorf("Home returned an error: %v", err)
	}
	if home != originalHome {
		t.Errorf("Home returned %q, expected %q", home, originalHome)
	}

	// Test the Config method
	config, err := u.Config()
	if err != nil {
		t.Errorf("Config returned an error: %v", err)
	}
	if config != expectedConfig {
		t.Errorf("Config returned %q, expected %q", config, expectedConfig)
	}

	// Test the Data method
	data, err := u.Data()
	if err != nil {
		t.Errorf("Data returned an error: %v", err)
	}
	if data != expectedData {
		t.Errorf("Data returned %q, expected %q", data, expectedData)
	}

	// Test the Cache method
	cache, err := u.Cache()
	if err != nil {
		t.Errorf("Cache returned an error: %v", err)
	}
	if cache != expectedCache {
		t.Errorf("Cache returned %q, expected %q", cache, expectedCache)
	}
}
