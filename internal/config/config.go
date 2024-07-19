package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path"
)

type Config struct {
	Env    string `yaml:"env"`
	Server struct {
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		Prefix string `yaml:"prefix"`
	} `yaml:"server"`
	Mysql struct {
		DSN   string `yaml:"dsn"`
		Debug bool   `yaml:"debug"`
	} `yaml:"mysql"`
	Redis struct {
		Addr         string `yaml:"addr"`
		Password     string `yaml:"password"`
		DB           int    `yaml:"db"`
		ReadTimeout  string `yaml:"read_timeout"`
		WriteTimeout string `yaml:"write_timeout"`
	}
	Log struct {
		LogLevel    string `yaml:"log_level"`
		LogFileName string `yaml:"log_file_name"`
		MaxBackups  int    `yaml:"max_backups"`
		MaxAge      int    `yaml:"max_age"`
		MaxSize     int    `yaml:"max_size"`
		Compress    bool   `yaml:"compress"`
	} `yaml:"log"`
	JWT struct {
		Key string `yaml:"key"`
	} `yaml:"jwt"`
	Wechat struct {
		AppId     string `yaml:"app_id"`
		AppSecret string `yaml:"app_secret"`
	} `yaml:"wechat"`
	Minio struct {
		Endpoint        string `yaml:"endpoint"`
		AccessKeyId     string `yaml:"access_key_id"`
		SecretAccessKey string `yaml:"secret_access_key"`
		BucketName      string `yaml:"bucket_name"`
		UseSSL          bool   `yaml:"use_ssl"`
	} `yaml:"minio"`
	CAS struct {
		TicketCheckUrl string `yaml:"ticket_check_url"`
		RedirectUrl    string `yaml:"redirect_url"`
	} `yaml:"cas"`
}

var CFG *Config

func init() {
	CFG = getConfig()
}

func getConfig() *Config {
	root, _ := os.Getwd()

	file, _ := os.Open(path.Join(root, "config.yaml"))

	d := yaml.NewDecoder(file)

	config := &Config{}
	if err := d.Decode(&config); err != nil {
		panic(err)
	}
	return config
}
