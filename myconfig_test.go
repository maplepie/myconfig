package myconfig_test

import (
	. "myconfig"
	"testing"
)

func Test_Marshal(t *testing.T) {
	conf := &Config{"2016-02-13", "30s", "10s"}
	b, err := conf.Marshal()

	if err != nil {
		t.Errorf("Test_Marshal Error: %v\n", err)
	}
	t.Logf("Test_Marshal Log: %v\n", string(b))
}

func Test_UnMarshal(t *testing.T) {
	conf2 := new(Config)
	b := `{"BeginDate":"2016-01-13"}`
	err := conf2.UnMarshal([]byte(b))
	if err != nil {
		t.Errorf("Test_UnMarshal Error: %v\n", err)
	}
	t.Logf("Test_UnMarshal Log: %v\n", conf2)
}

func Test_SaveToFile(t *testing.T) {
	conf2 := new(Config)
	b := `{"BeginDate":"2016-01-13"}`
	err := conf2.SaveToFile([]byte(b))
	if err != nil {
		t.Errorf("Test_SaveToFile Error: %v\n", err)
	}
	t.Logf("Test_SaveToFile Sucess\n", "")
}
