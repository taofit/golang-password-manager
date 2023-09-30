package category

import (
	"encoding/json"

	"fyne.io/fyne/v2/storage"
	"github.com/taofit/golang-password-manager/internal/fileStorage"
)

const (
	Website         = "website"
	CreditCard      = "credit card"
	Identity        = "identity"
	SoftwareLicense = "software license"
	Connections     = "connections"
	Note            = "note"
)

var DefaultCateList = []string{Website, CreditCard, Identity, SoftwareLicense, Connections, Note}

type category struct {
	name string
}

func CreateCategory(account string, name string) category {
	return category{name: name}
}

func AddCategory(accName string, cateName string) error {
	fileRoot, err := fileStorage.GetFileRoot()
	if err != nil {
		return err
	}
	accDir, err := fileStorage.GetDirURI(fileRoot, accName)
	if err != nil {
		return err
	}
	_, err = fileStorage.GetDirURI(accDir, cateName)
	if err != nil {
		return err
	}

	return nil
}

func AddCategories(accName string) error {
	for _, cate := range DefaultCateList {
		err := AddCategory(accName, cate)
		if err != nil {
			return err
		}
	}

	return nil
}

func SaveCateEntry(accName, cateName, entryName string) error {
	fileRoot, err := fileStorage.GetFileRoot()
	if err != nil {
		return err
	}
	accDir, err := fileStorage.GetDirURI(fileRoot, accName)
	if err != nil {
		return err
	}
	cateURI, err := fileStorage.GetDirURI(accDir, cateName)
	if err != nil {
		return err
	}
	fileURI, err := storage.Child(cateURI, entryName)
	if err != nil {
		return err
	}
	w, err := storage.Writer(fileURI)
	if err != nil {
		return err
	}
	defer w.Close()

	err = json.NewEncoder(w).Encode("entryName by dit")

	return err
}

func FetchCateList(accName string) ([]string, error) {
	cateNameList := []string{}
	fileRoot, err := fileStorage.GetFileRoot()
	if err != nil {
		return []string{}, err
	}
	accDir, err := fileStorage.GetDirURI(fileRoot, accName)
	if err != nil {
		return []string{}, err
	}
	uris, _ := storage.List(accDir)
	for _, uri := range uris {
		ok, _ := storage.CanList(uri)
		if ok {
			cateNameList = append(cateNameList, uri.Name())
		}
	}

	return cateNameList, nil
}

func FetchEntryList(accName, cateName string) ([]string, error) {
	entryList := []string{}
	fileRoot, err := fileStorage.GetFileRoot()
	if err != nil {
		return []string{}, err
	}
	accDir, err := fileStorage.GetDirURI(fileRoot, accName)
	if err != nil {
		return []string{}, err
	}
	cateURI, err := fileStorage.GetDirURI(accDir, cateName)
	if err != nil {
		return []string{}, err
	}
	uris, err := storage.List(cateURI)
	if err != nil {
		return []string{}, err
	}
	for _, uri := range uris {
		entryList = append(entryList, uri.Name())
	}

	return entryList, nil
}
