package account

import (
	"encoding/json"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
	"github.com/taofit/golang-password-manager/internal/fileStorage"
)

type account struct {
	Name     string
	Password string
}

func newAccount(name string, password string) account {
	return account{Name: name, Password: password}
}

func SaveAccount(name string, password string) (fyne.URI, error) {
	acc := newAccount(name, password)
	accFileURI, err := getAccountURI(name)
	if err != nil {
		return nil, err
	}
	ok, err := storage.CanWrite(accFileURI)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("can not write to file: %s", accFileURI)
	}
	w, err := storage.Writer(accFileURI)
	if err != nil {
		return nil, err
	}
	defer w.Close()
	err = json.NewEncoder(w).Encode(acc)

	return accFileURI, err
}

func getAccountURI(name string) (fyne.URI, error) {
	fileRoot, err := fileStorage.GetFileRoot()
	if err != nil {
		return nil, err
	}
	accDir, err := fileStorage.GetDirURI(fileRoot, name)
	if err != nil {
		return nil, err
	}
	accFileURI, err := storage.Child(accDir, name)
	if err != nil {
		return nil, err
	}
	return accFileURI, nil
}

func GetAccount(name string) (account, error) {
	acc := account{}
	accFileURI, err := getAccountURI(name)
	if err != nil {
		return acc, err
	}
	ok, err := storage.CanRead(accFileURI)
	if err != nil {
		return acc, err
	}
	if !ok {
		return acc, fmt.Errorf("can not read file: %s", accFileURI)
	}
	r, err := storage.Reader(accFileURI)
	if err != nil {
		return acc, err
	}
	defer r.Close()
	json.NewDecoder(r).Decode(&acc)

	return acc, nil
}
