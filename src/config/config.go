package config

import "gopkg.in/ini.v1"

var (
	cfg *ini.File
)

func ConfigLoad(filename string) error {
	var err error
	cfg, err = ini.Load(filename)
	if err != nil {
		return err
	}

	return nil
}

func ConfigGet(section, key string) string {
	return cfg.Section(section).Key(key).Value()
}

func ConfigSet(section, key, value string) {
	cfg.Section(section).Key(key).SetValue(value)
}

func ConfigSave(filename string) {
	cfg.SaveTo(filename)
}
