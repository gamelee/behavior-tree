package behavoir_tree

type IDecorator interface {
	IAction
	Child() IAction
}

type Decorator struct {
	Action
	child IAction
}

func (d *Decorator) Child() IAction {
	return d.child
}

func (d *Decorator) Run(note *Note) (interface{}, error) {
	return d.Run(note)
}