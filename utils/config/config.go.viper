package config

import (
//    "os"
    "log"
    "bytes"
    "io/ioutil"
    "github.com/spf13/viper"
)

func init() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err != nil {
        log.Panicln("Error loading config file")
    }
}

func LoadPath(path string, conftypes ...string) (conf *viper.Viper) {
    conf = viper.New()
    conf.SetConfigName("config")
    for _, conftype := range conftypes {
        conf.SetConfigType(conftype)
    }
    conf.AddConfigPath(path)
    err := conf.ReadInConfig()
    if err != nil {
        log.Panicln("Error loading config file")
    }
}

func LoadFile(file string, conftype ...string) (conf *viper.Viper) {
    content, err := ioutil.ReadFile(file)
    if err != nil {
        log.Panicln("Error loading config file")
    }
    conf = viper.New()
    for _, conftype := range conftypes {
        conf.SetConfigType(conftype)
    }
    err = conf.ReadConfig(bytes.NewBuffer(content))
    if err != nil {
        log.Panicln("Error loading config file")
    }
}
