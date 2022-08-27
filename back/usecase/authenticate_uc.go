package usecase

import (
	"ideal-journey/clients/errors"
	"ideal-journey/dto"
	"ideal-journey/entity"
	"ideal-journey/usecase/repo"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type AuthenticateUC struct {
	r     repo.AuthenticateRepo
	jwtUC *JwtUC
}

func NewAuthenticateUC(deps *AuthenticateDependecies) *AuthenticateUC {
	return &AuthenticateUC{
		r:     deps.Repos,
		jwtUC: deps.JwtUC,
	}
}

func (uc *AuthenticateUC) Create(address string) (*entity.UserAuth, errors.RestErr) {
	user, err := uc.r.Create(address)
	if err != nil {
		return nil, errors.UnautorizedError()
	}
	return user, nil
}

func (uc *AuthenticateUC) FindById(address string) (*entity.UserAuth, errors.RestErr) {
	user, err := uc.r.FindById(address)
	if err != nil {
		return nil, errors.UnautorizedError()
	}
	if user == nil {
		return nil, errors.NotFoundError("no user")
	}
	return user, nil
}

func (uc *AuthenticateUC) Auth(dto *dto.Authenticate) (string, errors.RestErr) {
	user, err := uc.r.FindById(dto.Address)
	if err != nil || user == nil {
		return "", errors.UnautorizedError()
	}
	if !verifySig(dto.Address, dto.Signature, []byte(user.Nonce)) {
		return "", errors.UnautorizedError()
	}

	user.Fingerprint = dto.Fingerprint
	err = uc.r.Update(user)
	if err != nil {
		return "", errors.UnautorizedError()
	}

	return uc.jwtUC.GenerateJWT(user)
}

func verifySig(from, sigHex string, msg []byte) bool {
	sig := hexutil.MustDecode(sigHex)

	msg = accounts.TextHash(msg)
	sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	return from == recoveredAddr.Hex()
}
