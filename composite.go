package behavoir_tree

type IComposite interface {
	INode
	IAction
	Children() []IAction
	Child(idx int) IAction
}

type Composite struct {
	Action
	children []IAction
}

func (c *Composite) Children() []IAction {
	return c.children
}

func (c *Composite) Child(idx int) IAction {
	return c.children[idx]
}

func (c *Composite) Run(note *Note) (interface{}, error) {
	rst := make([]interface{}, len(c.Children()))
	for i, child := range c.Children()	 {
		r, err := child.Run(note)
		rst[i] = r
		if err != nil {
			return rst, err
		}
	}
	return rst, nil
}
