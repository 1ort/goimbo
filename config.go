package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	API struct {
		BaseURL string `yaml:"base_url"`
		Enabled bool   `yaml:"enabled"`
	}
	Web struct {
		BaseURL       string `yaml:"base_url"`
		Enabled       bool   `yaml:"enabled"`
		CookieSecret  string `yaml:"cookie_storage_key"`
		XCSRFSecret   string `yaml:"x_csrf_secret"`
		EnableCaptcha bool   `yaml:"enable_captcha"`
	}
	Db struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	}
	Attachment struct {
		Folder            string   `yaml:"folder_path"`
		AllowedExtensions []string `yaml:"allowed_extensions"`
		MaxSize           float32  `yaml:"max_size_MB"`
		MaxAttachments    int      `yaml:"max_attachments"`
		MinAttachments    int      `yaml:"min_attachments"`
	}
}

func ReadConfig(configPath string) *Config {
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Sprintf("Can not open config file, error: %s", err))
	}
	var cfg Config
	err = yaml.Unmarshal(configFile, &cfg)
	if err != nil {
		panic(fmt.Sprintf("Can not parse config file, error: %s", err))
	}
	return &cfg
}

func (cfg *Config) GetDataBaseURL() string {
	templateURL := "postgresql://%s:%s/%s?user=%s&password=%s"
	dbURL := fmt.Sprintf(templateURL, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name, cfg.Db.User, cfg.Db.Pass)
	return dbURL
}

func (cfg *Config) GetAppAddr() string {
	templateAddr := "%s:%s"
	addr := fmt.Sprintf(templateAddr, cfg.App.Host, cfg.App.Port)
	return addr
}

func (cfg *Config) GetBaseAPIURL() string {
	return (cfg.API.BaseURL)
}

func (cfg *Config) GetBaseWebURL() string {
	return (cfg.Web.BaseURL)
}
