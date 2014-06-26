package mail

import (
	"net/smtp"
	"strconv"
	)

type MailService struct {
	Username 	string
	Password	string
	Server 		string
	Port		int
}

func DefaultMailService() *MailService {
	return &MailService {
		"redrepo.mail@gmail.com",
		"iamredrepo",
		"smtp.gmail.com",
		587,
	}
}

func(service MailService) Send(body []byte, recipients []string) error {
	
	var err error

	err = smtp.SendMail(
		service.Server+":"+strconv.Itoa(service.Port),
		service.createPlainAuthentication(),
		service.Username,
		recipients,
		body,
		)

	if err != nil {
		return err
	}

	return nil
}

func (service MailService) createPlainAuthentication() smtp.Auth {
	auth := smtp.PlainAuth(
		"",
		service.Username,
		service.Password,
		service.Server,
		)
	return auth
}
