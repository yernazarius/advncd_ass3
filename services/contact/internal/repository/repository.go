package repository

import (
	"architecture_go/services/contact/internal/domain"
)

type ContactRepositoryImpl struct{}

func NewContactRepository() ContactRepository {
	return &ContactRepositoryImpl{}
}

type ContactRepository interface {
	
	CreateContact(contact domain.Contact) (int, error)
	GetContact(id int) (*domain.Contact, error)
	UpdateContact(contact domain.Contact) error
	DeleteContact(id int) error

	CreateGroup(group domain.Group) (int, error)
	GetGroup(id int) (*domain.Group, error)
	UpdateGroup(group domain.Group) error
	DeleteGroup(id int) error

	AddContactToGroup(contactID, groupID int) error
	RemoveContactFromGroup(contactID, groupID int) error
	GetContactsByGroup(groupID int) ([]*domain.Contact, error)
}
