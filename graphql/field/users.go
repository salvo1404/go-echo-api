package field

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	user "github.com/salvo1404/go-echo-api/models"
)

var graphqlUser = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.ID},
			"name":      &graphql.Field{Type: graphql.String},
			"email":     &graphql.Field{Type: graphql.String},
			"createdAt": &graphql.Field{Type: graphql.String},
			"updatedAt": &graphql.Field{Type: graphql.String},
			"deletedAt": &graphql.Field{Type: graphql.String},
		},
		Description: "Users data",
	},
)

func GetUsersField(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(graphqlUser),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var u []*user.User
			if err := db.Find(&u).Error; err != nil {
				// do something
			}

			return u, nil
		},
		Description: "user",
	}
}
