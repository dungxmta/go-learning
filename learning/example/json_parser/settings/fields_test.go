package settings

import (
	"fmt"
	"testing"
)

func TestProperTitle(t *testing.T) {
	for _, v := range []string{"name", "ip", "full_name", "status_os", "os_ver", "static_id"} {
		s := properTitle(v)
		t.Log(fmt.Sprintf("%10v -> %v", v, s))
	}
}

func TestProcess_01_QueryEmpty(t *testing.T) {
	settingType := "test"
	customerId := "test"
	fieldTypeMap := map[string]string{
		"os":            "string",
		"full_name":     "array",
		"os_version":    "array_object",
		"inserted_time": "datetime",
	}

	err := Process(settingType, customerId, fieldTypeMap)
	t.Log(err)
}

func TestProcess_02_QueryEmpty(t *testing.T) {
	settingType := "test"
	customerId := "test"
	fieldTypeMap := map[string]string{}

	err := Process(settingType, customerId, fieldTypeMap)
	t.Log(err)
}

func TestProcess_03_Update(t *testing.T) {
	settingType := "test"
	customerId := "test"
	fieldTypeMap := map[string]string{
		"os":            "string",
		"full_name":     "array",
		"os_version":    "array_object",
		"inserted_time": "datetime",
	}

	orgFnQuery := connector.Query
	fnQueryMock := func(settingType, cId string) (doc Settings, status bool, err error) {
		doc = Settings{
			Id:         "test",
			CustomerId: "test",
			Type:       "test",
			Fields: []Field{
				{
					FieldId: "1",
					Name:    "1",
					Type:    "string",
					Deleted: false,
				},
				{
					FieldId: "inserted_time",
					Name:    "Inserted Time",
					Type:    "string",
					Deleted: false,
				},
				{
					FieldId: "os_version",
					Name:    "os_version",
					Type:    "abc",
					Deleted: true,
				},
			},
		}
		return doc, true, nil
	}
	connector.Query = fnQueryMock
	defer func() { connector.Query = orgFnQuery }()

	err := Process(settingType, customerId, fieldTypeMap)
	t.Log(err)
}
