package files

import (
	"api/app/models"
	"database/sql"
	"strconv"
)

type FilesService struct {
	DB *sql.DB
}

// Get File From db
func (s *FilesService) GetFile(id string) (*models.File,error) {
	var f models.File
	row := s.DB.QueryRow(`SELECT id, title, description FROM files WHERE id = ?`, id)

	if err := row.Scan(&f.ID, &f.Title, &f.Description); err != nil {
		return nil, err
	}

	return &f, nil
}

// CreateItem ...
func (s *FilesService) CreateFile(file *models.File) error {
	stmt, err := s.DB.Prepare(`INSERT INTO files(title,description) values(?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(file.Title, file.Description)
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
