package models

// Item ...
type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ItemServiceInterface ...
type ItemServiceInterface interface {
	Item(id string) (*Item, error)
	Items() ([]*Item, error)
	CreateItem(i *Item) error
	DeleteItem(id string) error
}

type File struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type FileServiceInterface interface {
	CreateFileInDB(f *File) error
	CreateFileInDrive(f *File) (string, error)
	SearchInDB(id string) (*File, error)
	SearchInDrive(id string) (*File, error)
	RetrieveAllFilesFromDrive()
}
