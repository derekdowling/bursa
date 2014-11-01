package mailchimp

import (
	"bursa.io/config"
	"github.com/mattbaird/gochimp"
	"log"
	"strconv"
)

// Adds a user, via their email, to one of our MailChimp mailing lists
func SubscribeToChimp(userEmail string) string {
	chimp := getMailChimp()
	request := gochimp.ListsSubscribe{
		ListId:         getMailListId(),
		Email:          gochimp.Email{Email: userEmail},
		DoubleOptIn:    false,
		UpdateExisting: true,
		SendWelcome:    sendWelcomeEmail(),
	}

	resp, err := chimp.ListsSubscribe(request)
	if err != nil {
		log.Println(err.Error())
	}
	return resp.Email
}

// Checks whether or not we are in production to avoid spamming ourselves
// with email
func sendWelcomeEmail() bool {
	val, _ := strconv.ParseBool(config.GetStringMapString("email")["enabled"])
	return val
}

// Sets up the MailChimp API
func getMailChimp() *gochimp.ChimpAPI {
	api_key := config.GetStringMapString("email")["mailchimp_key"]
	return gochimp.NewChimp(api_key, true)
}

// Determines which mailing list to add user to based on context
func getMailListId() string {
	return config.GetStringMapString("email")["list_id"]
}
