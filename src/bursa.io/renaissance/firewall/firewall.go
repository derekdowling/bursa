package firewall

// The purpose of this package is to provide a firewall layer that prevents users
// from accessing parts of the website without appropriate credentials

import (
	"bursa.io/models/users"
	"bursa.io/renaissance/satchel"
)

// Our type for specifying role flags
// These should be defined as such:
//
// const(
//		ROLE_1 Role = 1 << iota
//		ROLE_2
//		...
//		ROLE_N
// )
//
// Each route can have multiple rows assigned to it. By assigning roles to bit flags,
// it allows us to check all rows much quicker
type Role int

type Firewall struct {
	route_rules map[String]Role
}

// Implements our Mechanism Inteface for Davinci
func (self *Firewall) Trigger(s *satchel.Satchel) {
	authorized = session.LoggedIn()
}

// Checks our route rules vs what our session currently tells us about the user
func (self *Firewall) Check(r *http.Request) {

}
