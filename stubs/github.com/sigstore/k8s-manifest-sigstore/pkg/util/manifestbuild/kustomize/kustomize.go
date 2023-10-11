package kustomize

import (
	"crypto"
	"crypto/ecdsa"
	"time"
)

const kustomizationFileName = "kustomization.yaml"

const gitCmd = "git"

const kustCmd = "kustomize"

const refQueryRegex = "\\?(version|ref)="

const gitSuffix = ".git"

const gitDelimiter = "_git/"

type FileInfo struct {
	Name string
	Hash string
}

type GitRepoResult struct {
	RootDir  string
	URL      string
	Revision string
	CommitID string
	Path     string
}

type KustomizationResource struct {
	GitRepo *GitRepoResult
	File    *FileInfo
}

func LoadKustomization(fpath, baseDir, gitURL, gitRevision string, inRemoteRepo bool) ([]*KustomizationResource, error) {
	panic("stub")
}

func Sha256Hash(fpath string) (string, error) {
	panic("stub")
}

func CmdExec(baseCmd, dir string, args ...string) (string, error) {
	panic("stub")
}

func KustomizeExec(dir string, args ...string) (string, error) {
	panic("stub")
}

func GitExec(dir string, args ...string) (string, error) {
	panic("stub")
}

func IsRepositoryResource(path string) bool {
	panic("stub")
}

func IsFileResource(path string) bool {
	panic("stub")
}

func IsFile(name string) (bool, error) {
	panic("stub")
}

func IsDir(name string) (bool, error) {
	panic("stub")
}

func FileExists(fpath string) bool {
	panic("stub")
}

const cosignPwdEnvKey = "COSIGN_PASSWORD"

type IntotoSigner struct {
	key   *ecdsa.PrivateKey
	keyID string
}

func GenerateProvenance(artifactName, digest, kustomizeBase string, startTime, finishTime time.Time, recipeCmd []string) (interface{}, error) {
	panic("stub")
}

func GenerateAttestation(provPath, privKeyPath string) (interface{}, error) {
	panic("stub")
}

func GetDigestOfArtifact(artifactPath string) (string, error) {
	panic("stub")
}

func OverwriteArtifactInProvenance(provPath, overwriteArtifact string) (string, error) {
	panic("stub")
}

func GetImageDigest(resBundleRef string) (string, error) {
	panic("stub")
}

func (it *IntotoSigner) Sign(data []byte) ([]byte, error) {
	panic("stub")
}

func (it *IntotoSigner) Verify(data, sig []byte) error {
	panic("stub")
}

func (es *IntotoSigner) KeyID() (string, error) {
	panic("stub")
}

func (es *IntotoSigner) Public() crypto.PublicKey {
	panic("stub")
}

type Embedme interface{}
