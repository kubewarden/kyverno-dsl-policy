package cosign

import (
	"context"
	"crypto/x509"
)

const rekorServerEnvKey = "REKOR_SERVER"

const defaultRekorServerURL = "https://rekor.sigstore.dev"

const defaultOIDCIssuer = "https://oauth2.sigstore.dev/auth"

const defaultOIDCClientID = "sigstore"

const cosignPasswordEnvKey = "COSIGN_PASSWORD"

const defaultTlogUploadTimeout = 90

const treeIDHexStringLen = 16

const uuidHexStringLen = 64

func SignImage(resBundleRef string, keyPath, certPath *string, rekorURL string, noTlogUpload, force bool, pf interface{}, imageAnnotations map[string]interface{}, allowInsecure bool) error {
	panic("stub")
}

func SignBlob(blobPath string, keyPath, certPath *string, rekorURL string, noTlogUpload, force bool, pf interface{}) (map[string][]byte, error) {
	panic("stub")
}

func GetRekorServerURL() string {
	panic("stub")
}

func GetTlogEntry(ctx context.Context, rekorClient interface{}, uuid string) (interface{}, error) {
	panic("stub")
}

func ComputeLeafHash(e interface{}) ([]byte, error) {
	panic("stub")
}

type cosignBundleSignature struct {
	Embedme
	message         []byte
	base64Signature []byte
	cert            []byte
	base64Bundle    []byte
}

func VerifyImage(resBundleRef, pubkeyPath, certRef, certChain, rekorURL, oidcIssuer string, rootCerts *x509.CertPool, allowInsecure bool) (bool, string, *int64, error) {
	panic("stub")
}

func VerifyBlob(msgBytes, sigBytes, certBytes, bundleBytes []byte, pubkeyPath *string, certRef, certChain, rekorURL, oidcIssuer string) (bool, string, *int64, error) {
	panic("stub")
}

func (s *cosignBundleSignature) Annotations() (map[string]string, error) {
	panic("stub")
}

func (s *cosignBundleSignature) Payload() ([]byte, error) {
	panic("stub")
}

func (s *cosignBundleSignature) Base64Signature() (string, error) {
	panic("stub")
}

func (s *cosignBundleSignature) Cert() (*x509.Certificate, error) {
	panic("stub")
}

func (s *cosignBundleSignature) Chain() ([]*x509.Certificate, error) {
	panic("stub")
}

func (s *cosignBundleSignature) Bundle() (interface{}, error) {
	panic("stub")
}

type Embedme interface{}
