package usecase

import (
	"architecture_go/services/contact/internal/domain"
)

type ContactUseCase interface {

	CreateContact(contact domain.Contact) (int, error)
	GetContact(id int) (*domain.Contact, error)
	UpdateContact(contact domain.Contact) error
	DeleteContact(id int) error

	AddContactToGroup(contactID, groupID int) error

	CreateGroup(group domain.Group) (int, error)
	GetGroup(id int) (*domain.Group, error)

}
