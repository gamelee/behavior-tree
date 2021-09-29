package behavoir_tree

type NodeType string

const (
	NodeTypeNode NodeType= "node"
)

type INode interface {
	ID() string
	Init(conf *NodeConfig) error
	Type() NodeType
}

type Node struct {
	conf *NodeConfig
}

func (node *Node) Init(conf *NodeConfig) error{
	node.conf = conf
	return nil
}

func (node *Node) ID() string {
	return node.conf.ID
}

func (node *Node) Type() NodeType {
	return node.conf.Type
}
