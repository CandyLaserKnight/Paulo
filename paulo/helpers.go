package paulo

import (
	"errors"
	"os"
)

func (p *Paulo) CreateDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Paulo) CreateFileIfNotExists(path string) error {
	//var _, err = os.Stat(path)
	//errors.Is(err, os.ErrNotExist)
	//{
	//var file, err = os.Create(path)
	//if err != nil {
	//	return err
	//}

	//defer func(file *os.File) {
	//	_ = file.Close()
	//}(file)
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}
	return nil
}
