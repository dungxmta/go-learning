package mongo

import (
	"context"
	models "testProject/learning/example/api_jwt_mongo/model"
)

// result -> &interface{}
func (ins *connector) FindOne(colName string, ctx context.Context, result interface{}, filter interface{}, opts ...*interface{}) error {
	// TODO: options
	cur := ins.DB.Collection(colName).FindOne(ctx, filter)

	err := cur.Decode(result)
	return err
}

// results -> &([]interface)
func (ins *connector) Find(colName string, ctx context.Context, results interface{}, filter interface{}, opts ...*interface{}) error {
	cur, err := ins.DB.Collection(colName).Find(ctx, filter)
	if err != nil {
		return err
	}

	switch t := results.(type) {
	case *[]models.User:
		for cur.Next(ctx) {
			var obj models.User
			err := cur.Decode(&obj)
			if err != nil {
				return err
			}
			*t = append(*t, obj)
		}
	case *[]interface{}:
		for cur.Next(ctx) {
			var obj map[string]interface{}
			err := cur.Decode(&obj)
			if err != nil {
				return err
			}
			*t = append(*t, obj)
		}
	}

	if err := cur.Err(); err != nil {
		return err
	}

	return nil
}
