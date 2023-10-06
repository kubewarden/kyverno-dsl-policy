package sigtypes

const SigTypeUnknown = ""

const SigTypeCosign = "cosign"

const SigTypePGP = "pgp"

const SigTypeX509 = "x509"

type SigType string

func GetSignatureTypeFromPublicKey(keyPathPtr *string) SigType {
	panic("stub")
}

type Embedme interface{}
