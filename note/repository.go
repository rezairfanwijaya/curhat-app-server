package note

import "gorm.io/gorm"

type IRespository interface {
	Save(note Note) (Note, error)
	FindAll() ([]Note, error)
	Delete() error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(note Note) (Note, error) {
	if err := r.db.Create(&note).Error; err != nil {
		return note, err
	}

	return note, nil
}

func (r *repository) FindAll() ([]Note, error) {
	var notes []Note

	if err := r.db.Order("id desc").Find(&notes).Error; err != nil {
		return notes, err
	}

	return notes, nil
}

func (r *repository) Delete() error {
	if err := r.db.Exec("truncate notes").Error; err != nil {
		return err
	}

	return nil
}
