package token

type JWTMaker struct {
	secretKey string
}

const minSecretKeySize = 32

//func NewJWTMaker(secretKey string) (Maker, error) {
//	if len(secretKey) < minSecretKeySize {
//		return nil, fmt.Errorf("invalid key size: mist be at least %d characters. ", minSecretKeySize)
//	}
//	return &JWTMaker{secretKey}, nil
//}
