// Our base email management package. Handles adding users to mailing lists, sending
// emails, etc.

package emailer

import (
	"github.com/derekdowling/bursa/config"
	"github.com/derekdowling/bursa/emailer/mailchimp"
)

// Returns the user's email if successful, nil if not
func Subscribe(email string) error {
	return mailchimp.Subscribe(email, Enabled(), getMailListId())
}

// Determines which mailing list to add user to based on context
func getMailListId() string {
	return config.App.GetString("email.list_id")
}

// Checks whether or not we are in production to avoid spamming ourselves
// with email
func Enabled() bool {
	return config.App.GetBool("email.enabled")
}
