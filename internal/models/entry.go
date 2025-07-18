package models

type PasswordEntry struct {
	Title		string
	Username	string
	Password	string
	Sites		[]string
	Note		string
}

type EntryMatch struct {
	Entry	PasswordEntry
}

func (e EntryMatch) String() string {
	return e.Entry.Title
}
