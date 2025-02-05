package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"path"
)

type Config struct {
	Server Server `toml:"Server"`
	JWT    JWT    `toml:"JWT"`
	Mysql  Mysql  `toml:"Mysql"`
	Redis  Redis  `toml:"Redis"`
	Log    Log    `toml:"log"`
	Minio  Minio  `toml:"Minio"`
}

type Server struct {
	Env    string `toml:"Env"`
	Host   string `toml:"Host"`
	Port   int    `toml:"Port"`
	Prefix string `toml:"Prefix"`
}

type JWT struct {
	Key string `toml:"Key"`
}

type Mysql struct {
	User     string `toml:"User"`
	Password string `toml:"Password"`
	Database string `toml:"Database"`
	Addr     string `toml:"Addr"`
	Other    string `toml:"Other"`
	Debug    bool   `toml:"Debug"`
}

type Redis struct {
	Addr         string `toml:"Addr"`
	Password     string `toml:"Password"`
	Database     int    `toml:"Database"`
	ReadTimeout  string `toml:"ReadTimeout"`
	WriteTimeout string `toml:"WriteTimeout"`
}

type Log struct {
	Level      string `toml:"Level"`
	FilePath   string `toml:"FilePath"`
	MaxBackups int    `toml:"MaxBackups"`
	MaxAge     int    `toml:"MaxAge"`
	MaxSize    int    `toml:"MaxSize"`
	Compress   bool   `toml:"Compress"`
}

type Minio struct {
	Endpoint        string `toml:"Endpoint"`
	AccessKeyId     string `toml:"AccessKeyId"`
	SecretAccessKey string `toml:"SecretAccessKey"`
	BucketName      string `toml:"BucketName"`
	UseSSL          bool   `toml:"UseSSL"`
}

var CFG *Config

func init() {
	CFG = getConfig()
}

func getConfig() *Config {
	root, _ := os.Getwd()

	var conf Config
	if _, err := toml.DecodeFile(path.Join(root, "configs", "config.dev.toml"), &conf); err != nil {
		panic(err)
	}
	return &conf
}
