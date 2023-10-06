package x509

import "crypto/x509"

const PEMTypePrivateKey = "RSA PRIVATE KEY"

const PEMTypePublicKey = "PUBLIC KEY"

const PEMTypeCertificate = "CERTIFICATE"

func VerifyBlob(msgBytes, sigBytes, certBytes []byte, caCertPathString *string) (bool, string, *int64, error) {
	panic("stub")
}

func LoadCertificate(certPath string) (*x509.Certificate, error) {
	panic("stub")
}

func LoadCertificateChain(certChainPath string) ([]*x509.Certificate, error) {
	panic("stub")
}

func PEMDecode(pemBytes []byte, mode string) []byte {
	panic("stub")
}

func GetPublicKeyFromCertificate(certPemBytes []byte) ([]byte, error) {
	panic("stub")
}

func GetNameInfoFromX509Cert(cert *x509.Certificate) string {
	panic("stub")
}

type Embedme interface{}
