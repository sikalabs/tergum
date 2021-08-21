package backend

import "github.com/sikalabs/tergum/tergum2/notification/backend/email"

type Backend struct {
	Email email.EmailBackend
}

func (b Backend) Validate() error {
	err := b.Email.Validate()
	if err != nil {
		return err
	}

	return nil
}
