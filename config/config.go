package config

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
)

// Config represents the configuration information.
type Config struct {
  Url       string     `json:"url"`
  ApiKey    string     `json:"api_key"`
  MinorRev  string     `json:"minor_rev"`
}

var Conf Config

func init() {
  // Get the config file
  config_file, err := ioutil.ReadFile("./src/github.com/jacki3lene/goean/config/config.json")
  if err != nil {
    fmt.Printf("File error: %v\n", err)
  }
  json.Unmarshal(config_file, &Conf)
}