// Copyright 2015 JetHe. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func LoadConfig() *Config {
	var conf Config
	file, err := os.Open("./config.json")
	if err != nil {
		log.Println("Open file failed...", err)
		return nil
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("read file failed...", err)
		return nil
	}
	ConfigData = string(data)
	
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &conf
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//get configuration
	conf := LoadConfig()

	if conf == nil {
		return
	}

	StartCollect(*conf)
}
