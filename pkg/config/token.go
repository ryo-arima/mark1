package config

import "io/ioutil"

type HomeData struct {
	Token string
}

func LoadToken(homeDir string) string {
	buf1, err := ioutil.ReadFile(homeDir + "/token")
	if err != nil {
		panic(err)
	}
	return string(buf1)
}
