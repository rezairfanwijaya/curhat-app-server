package note

import (
	"net/http"
)

type IService interface {
	Save(input InputNewNote) (Note, int, error)
	GetAll() ([]Note, int, error)
	Delete() (int, error)
}

type service struct {
	repoNote IRespository
}

func NewService(repoNote IRespository) *service {
	return &service{repoNote}
}

func (s *service) Save(input InputNewNote) (Note, int, error) {
	author := ""
	if input.Author == "" {
		author = "anonymouse"
	} else {
		author = input.Author
	}

	newNote := Note{
		Author: author,
		Note:   input.Note,
	}

	noteSaved, err := s.repoNote.Save(newNote)
	if err != nil {
		return noteSaved, http.StatusInternalServerError, err
	}

	return noteSaved, http.StatusOK, nil
}

func (s *service) GetAll() ([]Note, int, error) {
	notes, err := s.repoNote.FindAll()
	if err != nil {
		return notes, http.StatusInternalServerError, err
	}

	return notes, http.StatusOK, nil
}

func (s *service) Delete() (int, error) {
	if err := s.repoNote.Delete(); err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, nil
}
