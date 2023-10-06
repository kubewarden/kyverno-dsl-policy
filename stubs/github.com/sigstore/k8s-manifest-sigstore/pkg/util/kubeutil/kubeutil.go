package kubeutil

func DryRunCreate(objBytes []byte, namespace string) ([]byte, error) {
	panic("stub")
}

func StrategicMergePatch(objBytes, patchBytes []byte, namespace string) ([]byte, error) {
	panic("stub")
}

func GetApplyPatchBytes(manifestBytes []byte, objNamespace string) ([]byte, []byte, error) {
	panic("stub")
}

const InClusterObjectPrefix = "k8s://"

var cfg interface{}

type podGetterFunc func(obj interface{}) ([]interface{}, error)

type ImageObject struct {
	PodName           string
	ContainerName     string
	ImageID           string
	ResourceBundleRef string
	Digest            string
}

func GetInClusterConfig() (interface{}, error) {
	panic("stub")
}

func IsInCluster() bool {
	panic("stub")
}

func GetOutOfClusterConfig() (interface{}, error) {
	panic("stub")
}

func GetKubeConfig() (interface{}, error) {
	panic("stub")
}

func SetKubeConfig(conf interface{}) {
	panic("stub")
}

func MatchLabels(obj interface{}, labelSelector interface{}) (bool, error) {
	panic("stub")
}

func GetAPIResources() ([]interface{}, error) {
	panic("stub")
}

func GetNamespaces() ([]interface{}, error) {
	panic("stub")
}

func ParseObjectRefInCluster(objRef string) (string, string, error) {
	panic("stub")
}

func ParseObjectRefInClusterWithKind(objRef string) (string, string, string, error) {
	panic("stub")
}

func GetSecret(namespace, name string) (interface{}, error) {
	panic("stub")
}

func GetResource(apiVersion, kind, namespace, name string) (interface{}, error) {
	panic("stub")
}

func ListResources(apiVersion, kind, namespace string) ([]interface{}, error) {
	panic("stub")
}

func GetAllImagesFromObject(obj interface{}) ([]ImageObject, error) {
	panic("stub")
}

func GetAllPodsFromObject(obj interface{}) ([]interface{}, error) {
	panic("stub")
}

type Embedme interface{}
