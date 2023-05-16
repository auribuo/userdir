//go:build !windows

package userdir

import (
	"os"
	"filepath"
)

type userdirDefault struct{}

func (u *userdirDefault) Home() (string, error) {
	return os.UserHomeDir()
}

func (u *userdirDefault) Cache() (string, error) {
	return os.UserCacheDir()
}

func (u *userdirDefault) Config() (string, error) {
	return os.UserConfigDir()
}

func (u *userdirDefault) Data() (string, error) {
	home, err := u.Home()
	if err != nil {
		return "", err
	}
	return filepath..Join(home, ".local", "share"), nil
}