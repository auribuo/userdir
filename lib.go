package userdir

type Userdir interface {
	Home() (string, error)
	Config() (string, error)
	Data() (string, error)
	Cache() (string, error)
}

func Base() Userdir {
	return &userdirDefault{}
}

func App(appName string) Userdir {
	return &userdirApp{
		app: appName,
	}
}

func Org(orgName string, appName string) Userdir {
	return &userdirOrg{
		org: orgName,
		app: appName,
	}
}
