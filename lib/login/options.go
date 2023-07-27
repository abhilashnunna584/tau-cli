package loginLib

import (
	"github.com/taubyte/tau-cli/singletons/config"
)

func GetProfiles() (_default string, possible []string, err error) {
	profiles := config.Profiles()

	_profiles := profiles.List(true)

	var x int
	possible = make([]string, len(_profiles))
	for name, profile := range _profiles {
		possible[x] = name
		x++

		if profile.Default == true {
			_default = name
		}
	}

	return
}
