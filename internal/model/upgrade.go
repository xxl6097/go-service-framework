package model

type Upgrade struct {
	Windows struct {
		Arm64 string `json:"arm64"`
		Amd64 string `json:"amd64"`
	} `json:"windows"`
	Linux struct {
		Arm64 string `json:"arm64"`
		Amd64 string `json:"amd64"`
	} `json:"linux"`
	Darwin struct {
		Arm64 string `json:"arm64"`
		Amd64 string `json:"amd64"`
	} `json:"darwin"`
}
