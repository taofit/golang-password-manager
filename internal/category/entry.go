package category

import "time"

type WebsiteEntry struct {
	Description    string `json:"description"`
	User           string
	Password       string
	Comment        string
	ExpireDate     time.Time
	CreatedDate    time.Time
	LastUpdateDate time.Time
}

type CreditCardEntry struct {
}

type entry interface {
	save()
	remove()
	fetch()
}

func NewWebsiteEntry(name, password string) WebsiteEntry {

	return WebsiteEntry{User: "username", Password: password}
}
