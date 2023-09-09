package device

import (
	"github.com/graphql-go/graphql"
	"github.com/ncarlier/readflow/pkg/helper"
	"github.com/ncarlier/readflow/pkg/schema"
	"github.com/ncarlier/readflow/pkg/service"
)

var createPushSubscriptionMutationField = &graphql.Field{
	Type:        deviceType,
	Description: "create push subscription for a device",
	Args: graphql.FieldConfigArgument{
		"sub": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: createPushSubscriptionResolver,
}

func createPushSubscriptionResolver(p graphql.ResolveParams) (interface{}, error) {
	sub, _ := p.Args["sub"].(string)
	return service.Lookup().CreateDevice(p.Context, sub)
}

var deletePushSubscriptionMutationField = &graphql.Field{
	Type:        deviceType,
	Description: "remove device push subscription",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
	Resolve: deletePushSubscriptionResolver,
}

func deletePushSubscriptionResolver(p graphql.ResolveParams) (interface{}, error) {
	id := helper.ParseGraphQLID(p.Args, "id")
	if id == nil {
		return nil, helper.InvalidParameterError("id")
	}
	return service.Lookup().DeleteDevice(p.Context, *id)
}

func init() {
	schema.AddMutationField("createPushSubscription", createPushSubscriptionMutationField)
	schema.AddMutationField("deletePushSubscription", deletePushSubscriptionMutationField)
}
