package category

import (
	"github.com/graphql-go/graphql"
	"github.com/ncarlier/readflow/pkg/helper"
	"github.com/ncarlier/readflow/pkg/model"
	"github.com/ncarlier/readflow/pkg/schema"
	"github.com/ncarlier/readflow/pkg/service"
)

var createOrUpdateCategoryMutationField = &graphql.Field{
	Type:        Type,
	Description: "create or update a category",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type:        graphql.ID,
			Description: "category to update if provided",
		},
		"title": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "title of the category",
		},
	},
	Resolve: createOrUpdateCategoryResolver,
}

func createOrUpdateCategoryResolver(p graphql.ResolveParams) (interface{}, error) {
	title := helper.ParseGraphQLArgument[string](p.Args, "title")
	if id := helper.ParseGraphQLID(p.Args, "id"); id != nil {
		form := model.CategoryUpdateForm{
			ID:    *id,
			Title: title,
		}
		return service.Lookup().UpdateCategory(p.Context, form)
	}
	if title == nil {
		return nil, helper.RequireParameterError("title")
	}
	form := model.CategoryCreateForm{
		Title: *title,
	}
	return service.Lookup().CreateCategory(p.Context, form)
}

var deleteCategoriesMutationField = &graphql.Field{
	Type:        graphql.Int,
	Description: "delete categories",
	Args: graphql.FieldConfigArgument{
		"ids": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.NewList(graphql.ID)),
		},
	},
	Resolve: deleteCategoriesResolver,
}

func deleteCategoriesResolver(p graphql.ResolveParams) (interface{}, error) {
	idsArg, ok := p.Args["ids"].([]interface{})
	if !ok {
		return nil, helper.InvalidParameterError("ids")
	}
	var ids []uint
	for _, v := range idsArg {
		if id := helper.ConvGraphQLID(v); id != nil {
			ids = append(ids, *id)
		}
	}

	return service.Lookup().DeleteCategories(p.Context, ids)
}

func init() {
	schema.AddMutationField("createOrUpdateCategory", createOrUpdateCategoryMutationField)
	schema.AddMutationField("deleteCategories", deleteCategoriesMutationField)
}
