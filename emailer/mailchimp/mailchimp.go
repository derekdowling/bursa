package mailchimp

import (
	"github.com/derekdowling/bursa/config"
	"github.com/mattbaird/gochimp"
)

// Adds a user, via their email, to one of our MailChimp mailing lists
func Subscribe(userEmail string, welcome bool, list string) error {
	chimp := getMailChimp()
	request := gochimp.ListsSubscribe{
		ListId:         list,
		Email:          gochimp.Email{Email: userEmail},
		DoubleOptIn:    false,
		UpdateExisting: true,
		SendWelcome:    welcome,
	}

	_, err := chimp.ListsSubscribe(request)
	return err
}

// Sets up the MailChimp API
func getMailChimp() *gochimp.ChimpAPI {
	api_key := config.App.GetString("email.mailchimp_key")
	return gochimp.NewChimp(api_key, true)
}
