//go:build !windows

package userdir

import (
	"os"
	"path/filepath"
)

type userdirOrg struct {
	org string
	app string
}

func (u *userdirOrg) Home() (string, error) {
	return os.UserHomeDir()
}

func (u *userdirOrg) Config() (string, error) {
	config, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(config, u.org, u.app), nil
}

func (u *userdirOrg) Cache() (string, error) {
	cache, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(cache, u.org, u.app), nil
}

func (u *userdirOrg) Data() (string, error) {
	home, err := u.Home()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "share", u.org, u.app), nil
}
