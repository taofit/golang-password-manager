package category

import (
	"github.com/taofit/golang-password-manager/internal/fileStorage"
)

var DefaultCateList = []string{"password", "credit card", "note", "identity", "software license", "connections"}

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

func SaveCategories(accName string) error {
	for _, cate := range DefaultCateList {
		err := AddCategory(accName, cate)
		if err != nil {
			return err
		}
	}

	return nil
}
