package model

import (
	"encoding/json"
	"os"
)

type ProcModel struct {
	Name    string      `json:"name"`
	BinUrl  string      `json:"binUrl"`
	ConfUrl string      `json:"confUrl"`
	Upgrade bool        `json:"upgrade"`
	Args    []string    `json:"args"`
	Status  string      `json:"status"`
	Exit    bool        `json:"exit"`
	Proc    *os.Process `json:"-"`
}

func (u *ProcModel) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
func (u *ProcModel) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
