package entity

type UserAuth struct {
	Id          string
	Nonce       string
	Fingerprint string
	Created     int64
}
