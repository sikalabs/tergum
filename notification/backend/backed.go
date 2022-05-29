package backend

import "github.com/sikalabs/tergum/notification/backend/email"

type Backend struct {
	Email *email.EmailBackend `yaml:"Email"`
}

func (b Backend) Validate() error {
	if b.Email != nil {
		err := b.Email.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
