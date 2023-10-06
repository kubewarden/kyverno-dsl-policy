package pgp

func VerifyBlob(msgBytes, sigBytes []byte, pubkeyPathString *string) (bool, string, *int64, error) {
	panic("stub")
}

func LoadPublicKey(keyPath string) (interface{}, error) {
	panic("stub")
}

func GetFirstIdentity(signer interface{}) interface{} {
	panic("stub")
}

type Embedme interface{}
