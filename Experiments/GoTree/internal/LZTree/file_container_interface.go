package LZTree

type FileContainerInterface interface {
	AddFile(path string)
	ExtractFile(name string, path string)
}
