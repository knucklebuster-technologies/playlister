package databases

import (
	"gopkg.in/mgo.v2"
)

//MongoDB type contains func to operate mongo
type MongoDB struct {
	Address string
	session *mgo.Session
}

//GetSession retun a mongo database session
func (m *MongoDB) GetSession(address string) *mgo.Session {
	if m.session == nil {
		s, err := mgo.Dial(address)

		// Check if connection error, is mongo running?
		if err != nil {
			panic(err)
		}
		m.session = s
		return m.session.Copy()
	}

	return m.session.Copy()
}
