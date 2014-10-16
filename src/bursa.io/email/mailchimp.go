package email

import (
	"encoding/json"
	"github.com/spf13/viper"
	"net/http"
)

type MailChimp struct {
	apiKey string
}

func NewMailChimp() *MailChimp {
	m := new(MailChimp)
	m.apiKey = viper.GetString("mailchimp_key")
	return m
}

func (self *MailChimp) subscribe(email string) bool {

}
