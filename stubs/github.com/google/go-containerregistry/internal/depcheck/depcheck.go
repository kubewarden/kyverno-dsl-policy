package depcheck

import "testing"

type node struct {
	importpath string
	consumers  map[string]struct{}
}

type graph map[string]node

func StdlibPackages() []string {
	panic("stub")
}

func CheckNoDependency(ip string, banned []string, buildFlags ...string) error {
	panic("stub")
}

func AssertNoDependency(t *testing.T, banned map[string][]string, buildFlags ...string) {
	panic("stub")
}

func AssertOnlyDependencies(t *testing.T, allowed map[string][]string, buildFlags ...string) {
	panic("stub")
}

func CheckOnlyDependencies(ip string, allowed map[string]struct{}, buildFlags ...string) error {
	panic("stub")
}

type Embedme interface{}
