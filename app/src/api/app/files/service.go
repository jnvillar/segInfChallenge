package files

import (
	"api/app/models"
	"database/sql"
	"strconv"
	"api/app/clients"
)

var (
	gClient clients.GoogleClient
)

type FilesService struct {
	DB *sql.DB
}

// Get File From db
func (s *FilesService) SearchInDB(id string) (*models.File, error) {
	var f models.File
	row := s.DB.QueryRow(`SELECT id, title, description FROM files WHERE id = ?`, id)

	if err := row.Scan(&f.ID, &f.Title, &f.Description); err != nil {
		return nil, err
	}

	return &f, nil
}

func (s *FilesService) RetrieveAllFilesFromDrive() {
	gClient = clients.GoogleClient{}
	gClient.ShowFiles()
}

// Search in drive
func (s *FilesService) SearchInDrive(id string) (*models.File, error) {
	gClient = clients.GoogleClient{}
	file, err := gClient.FindFile(id)

	if err != nil {
		return &models.File{}, err
	}

	return file, nil
}

func (s *FilesService) CreateFileInDrive(file *models.File) (string, error) {
	gClient = clients.GoogleClient{}
	id, err := gClient.UploadFile(file.Title, file.Description)

	if err != nil {
		return "", err
	}

	return id, nil
}

// CreateItem ...
func (s *FilesService) CreateFileInDB(file *models.File) error {
	stmt, err := s.DB.Prepare(`INSERT INTO files(id, title,description) values(?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(file.ID, file.Title, file.Description)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	file.ID = strconv.Itoa(int(id))
	return nil
}
