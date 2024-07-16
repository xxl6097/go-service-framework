package model

type ConfigModel struct {
	AppStoreUrl string      `json:"appStoreUrl" `
	Password    string      `json:"password" `
	Args        []string    `json:"args" `
	Procs       []ProcModel `json:"procs"`
}
