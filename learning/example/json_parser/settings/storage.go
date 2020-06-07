package settings

type DB struct {
	Query  func(settingType, cId string) (doc Settings, status bool, err error)
	Insert func(doc interface{}) (string, error)
	Update func(id string, doc interface{}) error
}

var connector = DB{
	Query:  query,
	Insert: insert,
	Update: update,
}

func insert(doc interface{}) (string, error) {
	return "", nil
}

func update(id string, doc interface{}) error {
	return nil
}

func query(settingType, cId string) (doc Settings, status bool, err error) {
	return Settings{}, false, nil
}
