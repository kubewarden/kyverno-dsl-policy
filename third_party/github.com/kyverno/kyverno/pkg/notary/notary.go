package notary

import (
	"context"
	"crypto/x509"

	"github.com/kyverno/kyverno/pkg/images"
)

type notaryVerifier struct{ log interface{} }

func NewVerifier() images.ImageVerifier {
	panic("stub")
}

func (v *notaryVerifier) VerifySignature(ctx context.Context, opts images.Options) (*images.Response, error) {
	panic("stub")
}

func (v *notaryVerifier) FetchAttestations(ctx context.Context, opts images.Options) (*images.Response, error) {
	panic("stub")
}

type imageResource struct{ ref interface{} }

func (ir *imageResource) String() string {
	panic("stub")
}

func (ir *imageResource) RegistryStr() string {
	panic("stub")
}

type simpleTrustStore struct {
	name      string
	storeType interface{}
	certs     []*x509.Certificate
}

func NewTrustStore(name string, certs []*x509.Certificate) interface{} {
	panic("stub")
}

func (ts *simpleTrustStore) GetCertificates(ctx context.Context, storeType interface{}, name string) ([]*x509.Certificate, error) {
	panic("stub")
}

type Embedme interface{}
