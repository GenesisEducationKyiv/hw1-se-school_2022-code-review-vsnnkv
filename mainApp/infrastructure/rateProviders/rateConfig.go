package rateProviders

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type rateConfig struct {
	BinanceUrl  string `yaml:"binanceUrl"`
	CoingekoUrl string `yaml:"coingekoUrl"`
	CoinbaseUrl string `yaml:"coinbaseUrl"`
}

func (c *rateConfig) getConf() *rateConfig {

	yamlFile, err := ioutil.ReadFile("infrastructure/rateProviders/rateConfig.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
