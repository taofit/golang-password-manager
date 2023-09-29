package fileStorage

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

const FileSTGRoot = "storage"

func GetFileSTGRoot() (fyne.URI, error) {
	return storage.Child(fyne.CurrentApp().Storage().RootURI(), FileSTGRoot)
}

func GetFileRoot() (fyne.URI, error) {
	fileRoot, err := GetFileSTGRoot()

	if err != nil {
		return nil, err
	}
	err = createDirIfNotExist(fileRoot)
	if err != nil {
		return nil, err
	}

	return fileRoot, nil
}

func GetDirURI(parent fyne.URI, name string) (fyne.URI, error) {
	fileDir, err := storage.Child(parent, name)
	if err != nil {
		return nil, err
	}
	err = createDirIfNotExist(fileDir)
	if err != nil {
		return nil, err
	}
	return fileDir, nil
}

func createDirIfNotExist(rscURI fyne.URI) error {
	ok, err := storage.Exists(rscURI)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	err = storage.CreateListable(rscURI)
	if err != nil {
		return err
	}

	return nil
}
