package dto

type Authenticate struct {
	Address     string `json:"email"`
	Signature   string `json:"signature"`
	Fingerprint string `json:"print"`
}
