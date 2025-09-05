package config

import "github.com/spf13/viper"

type Config struct {
	ConfigFile string
	Format     string
	Verbose    bool
}

func Load(configFile string) (*Config, error) {
	v := viper.New()
	v.SetDefault("format", DefaultFormat)
	v.SetDefault("verbose", false)
	v.AutomaticEnv()

	if configFile != "" {
		v.SetConfigFile(configFile)
	}

	return &Config{
		ConfigFile: configFile,
		Format:     v.GetString("format"),
		Verbose:    v.GetBool("verbose"),
	}, nil
}
