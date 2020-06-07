package settings

import (
	"encoding/json"
	"log"
	"strings"
)

type Settings struct {
	Id         string  `json:"_id" bson:"_id"`
	CustomerId string  `json:"customer_id" bson:"customer_id"`
	Type       string  `json:"type" bson:"type"`
	Fields     []Field `json:"fields" bson:"fields"`
}

func NewSetting(settingType, customerId string) Settings {
	return Settings{
		Id:         "rand()",
		CustomerId: customerId,
		Type:       settingType,
		Fields:     nil,
	}
}

type Field struct {
	FieldId string `json:"field_id" bson:"field_id"`
	Name    string `json:"name" bson:"name"`
	Type    string `json:"type" bson:"type"`
	Deleted bool   `json:"deleted" bson:"deleted"`
}

func NewField(id, fieldType string) Field {
	return Field{
		FieldId: id,
		Name:    properTitle(id),
		Type:    fieldType,
		Deleted: false,
	}
}

/*
func properTitle(input string) string {
	words := strings.Fields(input)
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
*/
func properTitle(s string) string {
	words := strings.Split(s, "_")
	w := strings.Join(words, " ")
	return strings.Title(w)
}

// input: asset_type, map_type {field_id: type}
func Process(settingType, customerId string, fieldTypeMap map[string]string) error {
	// query one by asset_type > doc
	doc, found, err := connector.Query(settingType, customerId)
	if err != nil {
		log.Fatal(err)
	}

	// no doc > build new fields from mapping > insert to db
	if !found {
		doc = NewSetting(settingType, customerId)
		fields := make([]Field, 0)

		for fieldId, fieldType := range fieldTypeMap {
			f := NewField(fieldId, fieldType)
			fields = append(fields, f)
		}

		doc.Fields = fields

		b, _ := json.MarshalIndent(doc, "", "   ")
		log.Println("Try to insert...")
		log.Println(string(b))

		_, err := connector.Insert(doc)
		if err != nil {
			log.Fatal(err)
		}
		return err

	} else { // found doc =/> update fields (no delete, just update or add more)
		fields := doc.Fields

		//  loop "fields" > if field in map_type > update "type" of field in "fields" & delete field in map_type
		for idx, field := range fields {
			if fieldType, ok := fieldTypeMap[field.FieldId]; ok {
				fields[idx].Type = fieldType
				fields[idx].Deleted = false

				log.Println("-> update key:", field.FieldId)
				delete(fieldTypeMap, field.FieldId)
			}
		}

		//  the left of map_type is all new fields > append to "fields"
		for fieldId, fieldType := range fieldTypeMap {
			f := NewField(fieldId, fieldType)
			fields = append(fields, f)

			log.Println("-> insert key:", fieldId)
		}

		//  update doc
		doc.Fields = fields

		b, _ := json.MarshalIndent(doc, "", "   ")
		log.Println("Try to update...")
		log.Println(string(b))

		err := connector.Update(doc.Id, doc)
		if err != nil {
			log.Fatal(err)
		}
		return err
	}
}
