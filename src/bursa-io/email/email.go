package email

type Mailer interface {
	subscribe(email string)
}
