package LZTree

type FileContainer struct {
	tree *Tree
}

func NewFileContainer(containerPath string) *FileContainer {

	c := &FileContainer{
		tree: NewTree(),
	}

	return c
}

func (c *FileContainer) AddFile(path string) {}

func (c *FileContainer) ExtractFile(name string, path string) {}
