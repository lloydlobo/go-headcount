package services

import (
	"strings"
	"sync"

	"github.com/lloydlobo/go-headcount/models"
)

type (
	// Action implements enumeration of actions
	Action int

	// Usage
	//
	// 	var filters = []services.Filter{
	// 			{Url: "#/", Name: "All", Selected: true},
	// 			{Url: "#/active", Name: "Active", Selected: false},
	// 			{Url: "#/completed", Name: "Completed", Selected: false},
	// 	}
	Filter struct {
		Url      string
		Name     string
		Selected bool
	}
)

// Enumerate Action related constants in one type
const ( // Hack: using `-1` as `default` case value to act as ActionGet operation.
	ActionCreate Action = iota
	ActionToggle
	ActionEdit
	ActionUpdate
	ActionDelete
)

type ContactService struct {
	Contacts  models.Contacts // FUTURE: map[int]*Contact
	lock      sync.Mutex      // Lock and defer Unlock during mutation of contacts.
	seq       int             // Tracks times contact is created while server is running. Start from 1.
	idCounter int             // Tracks current count of Contact till when session resets. Start from 0.
}

func NewContacts() *ContactService {
	return &ContactService{
		Contacts: models.Contacts{},
		seq:      1,
	}
}

func (c *ContactService) Seq() int {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.seq
}

func (c *ContactService) ResetContacts() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.Contacts = make([]models.Contact, 0)
	c.idCounter = 0
}

func (c *ContactService) CrudOps(action Action, contact models.Contact) models.Contact {
	c.lock.Lock()
	defer c.lock.Unlock()

	index := -1

	// Find index by ID
	if action != ActionCreate {
		for i, r := range c.Contacts {
			if r.ID == contact.ID {
				index = i
				break
			}
		}
	}

	switch action {
	case ActionCreate:
		c.Contacts = append(c.Contacts, contact)
		c.idCounter += 1
		c.seq += 1
		return contact

	case ActionToggle:
		c.Contacts[index].Status = contact.Status
		return contact

	case ActionUpdate:
		name := strings.Trim(contact.Name, " ")
		phone := strings.Trim(contact.Phone, " ")
		email := strings.Trim(contact.Email, " ") // TODO: add email regexp validation

		if len(name) != 0 && len(phone) != 0 && len(email) != 0 {
			c.Contacts[index].Name = contact.Name
			c.Contacts[index].Email = contact.Email
		} else {
			// remove if name is empty
			c.Contacts = append(c.Contacts[:index], c.Contacts[index+1:]...)
			return models.Contact{}
		}
		return contact

	case ActionDelete:
		c.Contacts = append(c.Contacts[:index], c.Contacts[index+1:]...)

	default:
		// ActionEdit should do nothing but return contact from store
	}

	if index != -1 && action != ActionDelete {
		return c.Contacts[index]
	}
	return models.Contact{}
}
