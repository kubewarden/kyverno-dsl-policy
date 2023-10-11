package github

import "github.com/google/go-containerregistry/pkg/authn"

const ghcrHostname = "ghcr.io"

type githubKeychain struct{}

type githubAuthenticator struct {
	username
	password string
}

func (g githubAuthenticator) Authorization() (*authn.AuthConfig, error) {
	panic("stub")
}

type Embedme interface{}
