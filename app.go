//go:build !windows

package userdir

import (
	"os"
	"path/filepath"
)

type userdirApp struct {
	app string
}

func (u *userdirApp) Home() (string, error) {
	return os.UserHomeDir()
}

func (u *userdirApp) Config() (string, error) {
	config, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(config, u.app), nil
}

func (u *userdirApp) Cache() (string, error) {
	cache, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(cache, u.app), nil
}

func (u *userdirApp) Data() (string, error) {
	home, err := u.Home()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "share", u.app), nil
}
