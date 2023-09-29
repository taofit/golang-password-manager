package account

import (
	"encoding/json"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
	"github.com/taofit/golang-password-manager/internal/category"
	"github.com/taofit/golang-password-manager/internal/fileStorage"
)

type account struct {
	Name         string   `json:"name"`
	Password     string   `json:"password"`
	CategoryList []string `json:"categoryList"`
}

func newAccount(name string, password string) account {
	return newAccountWithCateList(name, password, []string{})
}

func newAccountWithCateList(name string, password string, categoryList []string) account {
	if len(categoryList) == 0 {
		return account{Name: name, Password: password, CategoryList: category.DefaultCateList}
	}
	return account{Name: name, Password: password, CategoryList: category.DefaultCateList} //TODO fetch categorylist from the storage
}

func SaveAccount(name string, password string) (account, error) {
	acc := account{}
	accFileURI, err := getAccountURI(name)
	if err != nil {
		return acc, err
	}
	ok, err := storage.CanWrite(accFileURI)
	if err != nil {
		return acc, err
	}
	if !ok {
		return acc, fmt.Errorf("can not write to file: %s", accFileURI)
	}
	w, err := storage.Writer(accFileURI)
	if err != nil {
		return acc, err
	}
	defer w.Close()
	acc = newAccount(name, password)
	err = json.NewEncoder(w).Encode(acc)

	return acc, err
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
