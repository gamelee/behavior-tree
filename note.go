package behavoir_tree

import "sync"

type Note struct {
	*sync.Map
	*injector
}

func NewNote() *Note {
	return &Note{
		Map:      &sync.Map{},
		injector: NewInjector(),
	}
}

