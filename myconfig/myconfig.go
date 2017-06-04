package myconfig

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	BeginDate string
	MinSleep  string
	DeltaTime string
}

func (conf *Config) Marshal() ([]byte, error) {
	return json.MarshalIndent(conf, "", "\t")
}

func (conf *Config) UnMarshal(b []byte) error {
	return json.Unmarshal(b, conf)
}

func (conf *Config) SaveToFile(b []byte) error {
	return ioutil.WriteFile("config.conf", b, 0755)
}

// func main() {
// 	s := "2016-01-01"
// 	conf := &Config{s, "30s", "10s"}
// 	b, err := conf.Marshal()
// 	checkError(err)
// 	fmt.Println(string(b))

// 	conf2 := new(Config)
// 	err = conf2.UnMarshal(b)
// 	checkError(err)
// 	fmt.Println(conf2)

// 	fmt.Println("ok")
// }

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
