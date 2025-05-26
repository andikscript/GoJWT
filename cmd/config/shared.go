package config

var SharedConfig *Config

type Config struct {
	Users []struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	} `json:"users"`
}

func GetConfig() *Config {
	return SharedConfig
}
