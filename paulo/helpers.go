package paulo

import (
	"crypto/rand"
	"errors"
	"os"
)

const (
	randomString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321_+"
)

// RandomString generates a random string length n from values in the const randomString
func (p *Paulo) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomString)

	for i := range s {
		c, _ := rand.Prime(rand.Reader, len(r))
		x, y := c.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}
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
