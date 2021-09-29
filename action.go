package behavoir_tree

type IAction interface {
	INode
	Run(note *Note) (interface{}, error)
}

type Action struct {
	Node
}

func (action *Node) Run(note *Note) (interface{}, error) {
	return "null action node", nil
}
