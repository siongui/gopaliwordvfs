package gopaliwordvfs

import (
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	
	content0, ok0 := vfs0[filename]
	if ok0 {
		return []byte(content0), nil
	}

	content1, ok1 := vfs1[filename]
	if ok1 {
		return []byte(content1), nil
	}

	content2, ok2 := vfs2[filename]
	if ok2 {
		return []byte(content2), nil
	}

	content3, ok3 := vfs3[filename]
	if ok3 {
		return []byte(content3), nil
	}

	return nil, os.ErrNotExist
}

func MapKeys() []string {
	length := 0
	
	length += len(vfs0)

	length += len(vfs1)

	length += len(vfs2)

	length += len(vfs3)

	keys := make([]string, length)
	i := 0
	
	for k := range vfs0 {
		keys[i] = k
		i++
	}

	for k := range vfs1 {
		keys[i] = k
		i++
	}

	for k := range vfs2 {
		keys[i] = k
		i++
	}

	for k := range vfs3 {
		keys[i] = k
		i++
	}

	return keys
}
