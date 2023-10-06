package cosign

import (
	"context"

	"github.com/kyverno/kyverno/pkg/images"
)

type Cosign interface {
	VerifyImageSignatures(ctx context.Context, signedImgRef interface{}, co interface{}) ([]interface{}, bool, error)
	VerifyImageAttestations(ctx context.Context, signedImgRef interface{}, co interface{}) (checkedAttestations []interface{}, bundleVerified bool, err error)
}

type driver struct{}

func (d *driver) VerifyImageSignatures(ctx context.Context, signedImgRef interface{}, co interface{}) ([]interface{}, bool, error) {
	panic("stub")
}

func (d *driver) VerifyImageAttestations(ctx context.Context, signedImgRef interface{}, co interface{}) (checkedAttestations []interface{}, bundleVerified bool, err error) {
	panic("stub")
}

var ImageSignatureRepository string

type cosignVerifier struct{}

func NewVerifier() images.ImageVerifier {
	panic("stub")
}

func (v *cosignVerifier) VerifySignature(ctx context.Context, opts images.Options) (*images.Response, error) {
	panic("stub")
}

func (v *cosignVerifier) FetchAttestations(ctx context.Context, opts images.Options) (*images.Response, error) {
	panic("stub")
}

type mock struct{ data map[string][]interface{} }

type sig struct {
	Embedme
	cosignPayload interface{}
}

func SetMock(image string, data [][]byte) error {
	panic("stub")
}

func ClearMock() {
	panic("stub")
}

func (m *mock) VerifyImageSignatures(_ context.Context, signedImgRef interface{}, _ interface{}) ([]interface{}, bool, error) {
	panic("stub")
}

func (m *mock) VerifyImageAttestations(ctx context.Context, signedImgRef interface{}, co interface{}) (checkedAttestations []interface{}, bundleVerified bool, err error) {
	panic("stub")
}

func (s *sig) Payload() ([]byte, error) {
	panic("stub")
}

type Embedme interface{}
