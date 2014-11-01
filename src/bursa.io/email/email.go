// Our base email management package. Handles adding users to mailing lists, sending
// emails, etc.

package email

import (
	"bursa.io/email/mailchimp"
)

// Returns the user's email if successful, nil if not
func Subscribe(email string) string {
	return mailchimp.SubscribeToChimp(email)
}
