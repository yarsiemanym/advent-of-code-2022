package day07

import log "github.com/sirupsen/logrus"

type directory struct {
	Name           string
	Files          []*file
	Subdirectories []*directory
	Parent         *directory
}

func (thisDirectory *directory) Size() int {
	log.Debugf("Calculating size of directory \"%s\".", thisDirectory.Name)

	size := 0

	for _, file := range thisDirectory.Files {
		size += file.Size
	}

	for _, subdirectory := range thisDirectory.Subdirectories {
		size += subdirectory.Size()
	}

	log.Debugf("Size of directory \"%s\" is %d", thisDirectory.Name, size)

	return size
}

func (thisDirectory *directory) FindDirectoriesOfSizeLessThanOrEqualTo(maxSize int) []*directory {
	directories := []*directory{}

	if thisDirectory.Size() <= maxSize {
		directories = append(directories, thisDirectory)
	}

	for _, subdirectory := range thisDirectory.Subdirectories {
		directories = append(directories, subdirectory.FindDirectoriesOfSizeLessThanOrEqualTo(maxSize)...)
	}

	return directories
}

func (thisDirectory *directory) FindDirectoriesOfSizeGreaterThanOrEqualTo(minSize int) []*directory {
	directories := []*directory{}

	if thisDirectory.Size() >= minSize {
		directories = append(directories, thisDirectory)
	}

	for _, subdirectory := range thisDirectory.Subdirectories {
		directories = append(directories, subdirectory.FindDirectoriesOfSizeGreaterThanOrEqualTo(minSize)...)
	}

	return directories
}

func NewDirectory(name string, parent *directory) *directory {
	return &directory{
		Name:           name,
		Parent:         parent,
		Subdirectories: []*directory{},
		Files:          []*file{},
	}
}
