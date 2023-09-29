package category

import (
	"github.com/taofit/golang-password-manager/internal/fileStorage"
)

var CategoryList = []string{"password", "note"}

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
	for _, cate := range CategoryList {
		err := AddCategory(accName, cate)
		if err != nil {
			return err
		}
	}

	return nil
}
