package src

import (
	"context"

	"github.com/graphql-services/graphql-configurator/gen"
	"github.com/novacloudcz/graphql-orm/events"
)

func New(db *gen.DB, ec *events.EventController) *Resolver {
	resolver := NewResolver(db, ec)

	// resolver.Handlers.CreateUser = func(ctx context.Context, r *gen.GeneratedMutationResolver, input map[string]interface{}) (item *gen.Company, err error) {
	// 	return gen.CreateUserHandler(ctx, r, input)
	// }

	return resolver
}

func (r *MutationResolver) CreateConfiguratorAssembly(ctx context.Context, input gen.ConfiguratorAssemblyInput) (item *gen.ConfiguratorAssembly, err error) {
	var itemID string
	err = r.BeginTransaction(ctx, func(ctx context.Context) error {
		inputHelper := AssemblyInputHelper{
			ConfiguratorAssemblyInput: &input,
		}
		itemID, err = inputHelper.Assemble(ctx, r.GeneratedResolver)
		return err
	})

	helper := AssemblyHelper{r.GeneratedResolver}
	item, err = helper.Load(ctx, itemID)

	return
}

func (r *QueryResolver) ConfiguratorAssembly(ctx context.Context, id string) (item *gen.ConfiguratorAssembly, err error) {
	helper := AssemblyHelper{r.GeneratedResolver}
	item, err = helper.Load(ctx, id)
	return
}
