package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("APP_ENV", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.Port != 8080 {
		t.Errorf("Port = %d, want 8080", cfg.Port)
	}
	if cfg.Env != EnvDev {
		t.Errorf("Env = %q, want dev", cfg.Env)
	}
}

func TestLoadInvalidPort(t *testing.T) {
	for _, bad := range []string{"abc", "0", "70000", "-1"} {
		t.Setenv("PORT", bad)
		if _, err := Load(); err == nil {
			t.Errorf("Load() with PORT=%q: want error, got nil", bad)
		}
	}
}

func TestLoadInvalidEnv(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("APP_ENV", "staging")
	if _, err := Load(); err == nil {
		t.Error("Load() with APP_ENV=staging: want error, got nil")
	}
}
