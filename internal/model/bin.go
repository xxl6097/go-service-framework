package model

import (
	"encoding/json"
	"os"
)

type ProcModel struct {
	Name    string      `json:"name"`
	BinUrl  string      `json:"binUrl"`
	ConfUrl string      `json:"confUrl"`
	Args    []string    `json:"args"`
	Proc    *os.Process `json:"-"`
}

func (u *ProcModel) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
func (u *ProcModel) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
