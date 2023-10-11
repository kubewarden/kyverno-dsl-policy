package authn

import "sync"

type anonymous struct{}

func (a *anonymous) Authorization() (*AuthConfig, error) {
	panic("stub")
}

type auth struct{ config AuthConfig }

func FromConfig(cfg AuthConfig) Authenticator {
	panic("stub")
}

func (a *auth) Authorization() (*AuthConfig, error) {
	panic("stub")
}

type Authenticator interface {
	Authorization() (*AuthConfig, error)
}

type AuthConfig struct {
	Username      string
	Password      string
	Auth          string
	IdentityToken string
	RegistryToken string
}

type authConfig AuthConfig

func (a *AuthConfig) UnmarshalJSON(data []byte) error {
	panic("stub")
}

func (a AuthConfig) MarshalJSON() ([]byte, error) {
	panic("stub")
}

type Basic struct {
	Username string
	Password string
}

func (b *Basic) Authorization() (*AuthConfig, error) {
	panic("stub")
}

type Bearer struct{ Token string }

func (b *Bearer) Authorization() (*AuthConfig, error) {
	panic("stub")
}

type wrapper struct{ h Helper }

type Keychain interface {
	Resolve(Resource) (Authenticator, error)
}

type defaultKeychain struct{ mu sync.Mutex }

type Resource interface {
	String() string
	RegistryStr() string
}

type Helper interface {
	Get(serverURL string) (string, string, error)
}

func (dk *defaultKeychain) Resolve(target Resource) (Authenticator, error) {
	panic("stub")
}

func NewKeychainFromHelper(h Helper) Keychain {
	panic("stub")
}

func (w wrapper) Resolve(r Resource) (Authenticator, error) {
	panic("stub")
}

type multiKeychain struct{ keychains []Keychain }

func NewMultiKeychain(kcs ...Keychain) Keychain {
	panic("stub")
}

func (mk *multiKeychain) Resolve(target Resource) (Authenticator, error) {
	panic("stub")
}

type Embedme interface{}
