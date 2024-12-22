package main

import (
	"fmt"
	"os"
)

type Config struct {
	Env      string `json:"env"`
	Test_env string `json:"test_env"`
}

func (c *Config) genJSON() error {
	file, err := os.Create("zed-platformio.json")
	if err != nil {
		return fmt.Errorf("error creating zed-platformio.json: %v", err)
	}
	defer file.Close()
	strConfig := fmt.Sprintf(`{\n\t"env": "%s",\n\t"test_env": "%s"\n}`, c.Env, c.Test_env)
	_, err = file.WriteString(strConfig)
	if err != nil {
		return fmt.Errorf("error writing to zed-platformio.json: %v", err)
	}
	return nil
}
