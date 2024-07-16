package model

type ConfigModel struct {
	Password string      `json:"password" `
	Args     []string    `json:"args" `
	Procs    []ProcModel `json:"procs"`
}
