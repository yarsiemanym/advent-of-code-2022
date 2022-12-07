package day07

type file struct {
	Name string
	Size int
}

func NewFile(name string, size int) *file {
	return &file{
		Name: name,
		Size: size,
	}
}
