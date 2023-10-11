package util

import (
	"crypto/x509"
	"io"
	"math/big"
	"sync"
	"time"
)

var localCacheEnvKey = "K8S_MANIFEST_LOCAL_FILE_CACHE"

var globalCache Cache

var defaultLocalFileCacheTTL = 180

var defaultOnMemoryCacheTTL = 30

const CacheTypeUnknown = ""

const CacheTypeMemory = "memory"

const CacheTypeFile = "file"

type cachedObject struct {
	timestamp time.Time
	object    []interface{}
}

type CacheType string

type LocalFileCache struct {
	TTL     time.Duration
	baseDir string
	mem     *OnMemoryCache
}

type Cache interface {
	Set(key string, value ...interface{}) error
	Get(key string) ([]interface{}, error)
}

type OnMemoryCache struct {
	TTL  time.Duration
	data map[string]cachedObject
	mu   sync.RWMutex
}

func (c *OnMemoryCache) Set(key string, value ...interface{}) error {
	panic("stub")
}

func (c *OnMemoryCache) Get(key string) ([]interface{}, error) {
	panic("stub")
}

func (c *LocalFileCache) Set(key string, value ...interface{}) error {
	panic("stub")
}

func (c *LocalFileCache) Get(key string) ([]interface{}, error) {
	panic("stub")
}

func GetCacheBaseDir() string {
	panic("stub")
}

func IsLocalCacheEnabeld() bool {
	panic("stub")
}

func SetCache(key string, value ...interface{}) error {
	panic("stub")
}

func GetCache(key string) ([]interface{}, error) {
	panic("stub")
}

func GetNameInfoFromCert(cert *x509.Certificate) string {
	panic("stub")
}

func CmdExec(baseCmd string, args ...string) (string, error) {
	panic("stub")
}

const EnvVarFileRefPrefix = "env://"

type AnnotationWriter func([]byte, map[string]interface{}) ([]byte, error)

type MutateOptions struct {
	AW          AnnotationWriter
	Annotations map[string]interface{}
}

func TarGzCompress(src string, buf io.Writer, moList ...*MutateOptions) error {
	panic("stub")
}

func TarGzDecompress(src io.Reader, dst string) error {
	panic("stub")
}

func IsDir(path string) (bool, error) {
	panic("stub")
}

func GetHomeDir() string {
	panic("stub")
}

func LoadFileDataInEnvVar(envVarRef string) ([]byte, error) {
	panic("stub")
}

func GzipCompress(in []byte) []byte {
	panic("stub")
}

func GzipDecompress(in []byte) []byte {
	panic("stub")
}

func PullImage(resBundleRef string, allowInsecure bool) (interface{}, error) {
	panic("stub")
}

func GetImageDigest(resBundleRef string, allowInsecure bool) (string, error) {
	panic("stub")
}

func GetBlob(layer interface{}) ([]byte, error) {
	panic("stub")
}

func GenerateConcatYAMLsFromImage(img interface{}) ([]byte, error) {
	panic("stub")
}

func GetYAMLsInArtifact(blob []byte) ([][]byte, error) {
	panic("stub")
}

func MatchPattern(pattern, value string) bool {
	panic("stub")
}

func MatchSinglePattern(pattern, value string) bool {
	panic("stub")
}

func ExactMatch(pattern, value string) bool {
	panic("stub")
}

func ExactMatchWithPatternArray(value string, patternArray []string) bool {
	panic("stub")
}

func GetUnionOfArrays(array1, array2 []string) []string {
	panic("stub")
}

func MatchPatternWithArray(pattern string, valueArray []string) bool {
	panic("stub")
}

func MatchWithPatternArray(value string, patternArray []string) bool {
	panic("stub")
}

func MatchBigInt(pattern string, value *big.Int) bool {
	panic("stub")
}

func SplitRule(rules string) []string {
	panic("stub")
}

func SplitCommaSeparatedString(in string) []string {
	panic("stub")
}

var GitVersion = "develop"

var gitCommit = "unknown"

var gitTreeState = "unknown"

var buildDate = "unknown"

type VersionInfo struct {
	Major        string
	Minor        string
	GitVersion   string
	GitCommit    string
	GitTreeState string
	BuildDate    string
	GoVersion    string
	Compiler     string
	Platform     string
}

func GetVersionInfo() *VersionInfo {
	panic("stub")
}

const defaultmaxResourceManifestsNum = 3

type candidateManifest struct {
	yaml  []byte
	table map[string]interface{}
	count int
	name  string
}

type ResourceInfo struct {
	group     string
	version   string
	kind      string
	name      string
	namespace string
	raw       []byte
}

func (ri ResourceInfo) Map() map[string]string {
	panic("stub")
}

func FindYAMLsInDir(dirPath string) ([][]byte, error) {
	panic("stub")
}

func LoadYAMLsInDirWithMutationOptions(dirPath string, moList ...*MutateOptions) ([][]byte, error) {
	panic("stub")
}

func FindManifestYAML(concatYamlBytes, objBytes []byte, maxResourceManifestNum *int, ignoreFields []string) (bool, [][]byte) {
	panic("stub")
}

func ManifestSearchByGVKNameNamespace(concatYamlBytes []byte, apiVersion, kind, name, namespace string) (bool, []byte) {
	panic("stub")
}

func ManifestSearchByValue(concatYamlBytes, objBytes []byte, maxResourceManifests *int, ignoreFields []string) (bool, [][]byte) {
	panic("stub")
}

func ConcatenateYAMLs(yamls [][]byte) []byte {
	panic("stub")
}

func IsConcatYAMLs(yaml []byte) bool {
	panic("stub")
}

func SplitConcatYAMLs(yaml []byte) [][]byte {
	panic("stub")
}

func GetAnnotationsInYAML(yamlBytes []byte) map[string]string {
	panic("stub")
}

type Embedme interface{}
