package db

import arango "github.com/diegogub/aranGO"

func GetSessionDB() *arango.Session {

	//Connection ArangoDB
	s, err := arango.Connect("http://192.168.0.100:8529","boanergepro", "123456", false)

	if err != nil {
		panic(err)
	}
	return s
}