package config

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func GetBanditThresholds() *BanditThreshoulds {
	c := &Configs{}

	var data []byte
	var err error
	data, err = os.ReadFile("./internal/config/bandit.yaml")
	if err != nil {
		panic(err.Error())
	}

	decoderOverride := yaml.NewDecoder(bytes.NewReader(data))
	decoderOverride.KnownFields(true)
	err = decoderOverride.Decode(&c)
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}

	return &c.BanditThreshoulds
}
