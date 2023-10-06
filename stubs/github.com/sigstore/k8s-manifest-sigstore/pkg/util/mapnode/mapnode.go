package mapnode

import "reflect"

type DiffResult struct{ Items []Difference }

type Difference struct {
	Key    string
	Values map[string]interface{}
}

type DiffPattern Difference

func (d *Difference) Equal(d2 *Difference) bool {
	panic("stub")
}

func (dp *DiffPattern) Match(diff *Difference) bool {
	panic("stub")
}

func (d *DiffResult) Keys() []string {
	panic("stub")
}

func (d *DiffResult) Values() []map[string]interface{} {
	panic("stub")
}

func (dr *DiffResult) Size() int {
	panic("stub")
}

func (dr *DiffResult) Remove(patterns []*DiffPattern) *DiffResult {
	panic("stub")
}

func (dr *DiffResult) Filter(maskKeys []string) (*DiffResult, *DiffResult, []string) {
	panic("stub")
}

func (d *DiffResult) ToJson() string {
	panic("stub")
}

func (d *DiffResult) String() string {
	panic("stub")
}

func (d *DiffResult) KeyString() string {
	panic("stub")
}

const MapnodeVersion = "0.0.2"

const DefaultMatchMode = ""

const MatchDefaultWithNil = "match"

const DistinguishDefaultFromNil = "distiguish"

type NodeValue struct {
	Type  reflect.Type
	Value interface{}
}

type Node struct {
	Value    *NodeValue
	Children interface{}
}

type matchMode string

func NewNodeValue(val interface{}) *NodeValue {
	panic("stub")
}

func (nv *NodeValue) Interface() interface{} {
	panic("stub")
}

func (nv *NodeValue) String() string {
	panic("stub")
}

func NewFromBytes(rawObj []byte) (*Node, error) {
	panic("stub")
}

func NewFromYamlBytes(rawObj []byte) (*Node, error) {
	panic("stub")
}

func NewFromInterfaceBytes(rawObj []byte) (*Node, error) {
	panic("stub")
}

func NewFromMap(objMap map[string]interface{}) (*Node, error) {
	panic("stub")
}

func NewNode(val interface{}) *Node {
	panic("stub")
}

func (n *Node) KeyExists(concatKey string) bool {
	panic("stub")
}

func (n *Node) DeepCopyInto(n2 *Node) {
	panic("stub")
}

func (n *Node) Copy() *Node {
	panic("stub")
}

func (n *Node) Merge(n2 *Node) (*Node, error) {
	panic("stub")
}

func (n *Node) Extract(filterKeys []string) *Node {
	panic("stub")
}

func (t *Node) Mask(keys []string) *Node {
	panic("stub")
}

func (n *Node) IsValue() bool {
	panic("stub")
}

func (n *Node) IsMap() bool {
	panic("stub")
}

func (n *Node) IsSlice() bool {
	panic("stub")
}

func (n *Node) Size() int {
	panic("stub")
}

func (n *Node) GetChildrenMap() map[string]*Node {
	panic("stub")
}

func (n *Node) GetChildrenSlice() []*Node {
	panic("stub")
}

func (n *Node) GetNodeByJSONPath(jpathKey string) (*Node, error) {
	panic("stub")
}

func GetConcreteKeys(keys []string, n *Node) []string {
	panic("stub")
}

func (t *Node) MultipleSubNode(concatKey string) []*Node {
	panic("stub")
}

func (t *Node) GetNode(concatKey string) (*Node, bool) {
	panic("stub")
}

func (t *Node) SubNode(concatKey string) *Node {
	panic("stub")
}

func (t *Node) GetString(concatKey string) string {
	panic("stub")
}

func (t *Node) GetBool(concatKey string, defaultValue bool) bool {
	panic("stub")
}

func (t *Node) Get(concatKey string) (interface{}, bool) {
	panic("stub")
}

func (n *Node) GetChild(key string) (*Node, bool) {
	panic("stub")
}

func (t *Node) Ravel() map[string]interface{} {
	panic("stub")
}

func (n *Node) Interface() interface{} {
	panic("stub")
}

func (t *Node) ToMap() map[string]interface{} {
	panic("stub")
}

func (n *Node) ToJson() string {
	panic("stub")
}

func (n *Node) ToYaml() string {
	panic("stub")
}

func (n *Node) String() string {
	panic("stub")
}

func (t *Node) Diff(t2 *Node) *DiffResult {
	panic("stub")
}

func (t *Node) DiffStrict(t2 *Node) *DiffResult {
	panic("stub")
}

func (t *Node) DiffSpecificType(t2 *Node, findTypeList []string) *DiffResult {
	panic("stub")
}

func (t *Node) FindUpdatedAndDeleted(t2 *Node) *DiffResult {
	panic("stub")
}

func (t *Node) FindUpdatedAndCreated(t2 *Node) *DiffResult {
	panic("stub")
}

func FindDiffBetweenNodes(t1, t2 *Node, findType map[string]bool, matchMode matchMode) *DiffResult {
	panic("stub")
}

func GetValueByLongKey(m map[string]interface{}, longKey string) (interface{}, error) {
	panic("stub")
}

func SplitCommaSeparatedKeys(key string) []string {
	panic("stub")
}

type Embedme interface{}
