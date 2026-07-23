package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ChannelType int

const (
	ChannelTypeText ChannelType = iota
	ChannelTypeForum
	ChannelTypeVoice
)

func (t *ChannelType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw string
	if err := unmarshal(&raw); err != nil {
		return err
	}

	switch raw {
	case "text":
		*t = ChannelTypeText
	case "forum":
		*t = ChannelTypeForum
	case "voice":
		*t = ChannelTypeVoice
	default:
		return fmt.Errorf("invalid channel type: %s", raw)
	}
	return nil
}

type ServerConfig struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type Role struct {
	Key   string `yaml:"key"`
	Name  string `yaml:"name"`
	Color string `yaml:"color"`
}

type Channel struct {
	Key   string      `yaml:"key"`
	Name  string      `yaml:"name"`
	Type  ChannelType `yaml:"type"`
	Emoji string      `yaml:"emoji"`
	Tags  []string    `yaml:"tags"`
}

type Category struct {
	Key      string    `yaml:"key"`
	Name     string    `yaml:"name"`
	Emoji    string    `yaml:"emoji"`
	Channels []Channel `yaml:"channels"`
}

type Config struct {
	Version    int          `yaml:"version"`
	Server     ServerConfig `yaml:"server"`
	Categories []Category   `yaml:"categories"`
	Roles      []Role       `yaml:"roles"`
}

func Load(path string) (*Config, error) {
	var c Config

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling yaml: %w", err)
	}
	return &c, nil
}
