package repository

type CommonRepo interface {
	FindAll(colName string, queryData map[string]interface{}) ([]interface{}, error)
	FindOne(colName string, queryData map[string]interface{}) (interface{}, error)
	// Insert(colName string, obj *interface{}) (string, error)
}
