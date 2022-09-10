package dto

type Authenticate struct {
	Address     string `json:"publicAddress"`
	Signature   string `json:"signature"`
	Fingerprint string `json:"print"`
}
