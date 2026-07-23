package config

import "testing"

func TestLoad(t *testing.T) {
	cfg, err := Load("config.example.yaml")
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if cfg.Version != 1 {
		t.Errorf("expected version 1, got %d", cfg.Version)
	}

	if len(cfg.Categories) != 7 {
		t.Errorf("expected 7 categories, got %d", len(cfg.Categories))
	}

	if cfg.Categories[4].Channels[0].Type != ChannelTypeForum {
		t.Errorf("expected channel type forum, got %d", cfg.Categories[4].Channels[0].Type)
	}

}
