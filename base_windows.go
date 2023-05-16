//go:build windows

package userdir

import (
	"os"
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
	return os.UserConfigDir()
}
