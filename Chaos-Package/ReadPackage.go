package READER

import (
	"errors"
	"io/ioutil"
)

type READER47 struct {
	Data  string
	State bool
}

func (RD *READER47) Starter() error {
	if !RD.State {
		RD.State = true
	}
	b, err := ioutil.ReadFile("Chaos-Package/Chaos.Destroy")
	if err != nil {
		return err
	}
	RD.Data = string(b)
	return nil
}
func (RD READER47) GETSTRing() (string, error) {
	if !RD.State {
		return "", errors.New("Failed: u dont used Starter()")
	}

	return RD.Data, nil
}
