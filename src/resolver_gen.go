package src

import (
	"context"

	"github.com/graphql-services/graphql-configurator/gen"
	"github.com/novacloudcz/graphql-orm/events"
)

func NewResolver(db *gen.DB, ec *events.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{&gen.GeneratedResolver{Handlers: handlers, DB: db, EventController: ec}}
}

type Resolver struct {
	*gen.GeneratedResolver
}

type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

func (r *MutationResolver) BeginTransaction(ctx context.Context, fn func(context.Context) error) error {
	ctx = gen.EnrichContextWithMutations(ctx, r.GeneratedResolver)
	err := fn(ctx)
	if err != nil {
		tx := r.GeneratedResolver.GetDB(ctx)
		tx.Rollback()
		return err
	}
	return gen.FinishMutationContext(ctx, r.GeneratedResolver)
}

type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{&gen.GeneratedMutationResolver{r.GeneratedResolver}}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{&gen.GeneratedQueryResolver{r.GeneratedResolver}}
}

type ConfiguratorItemDefinitionResultTypeResolver struct {
	*gen.GeneratedConfiguratorItemDefinitionResultTypeResolver
}

func (r *Resolver) ConfiguratorItemDefinitionResultType() gen.ConfiguratorItemDefinitionResultTypeResolver {
	return &ConfiguratorItemDefinitionResultTypeResolver{&gen.GeneratedConfiguratorItemDefinitionResultTypeResolver{r.GeneratedResolver}}
}

type ConfiguratorItemDefinitionResolver struct {
	*gen.GeneratedConfiguratorItemDefinitionResolver
}

func (r *Resolver) ConfiguratorItemDefinition() gen.ConfiguratorItemDefinitionResolver {
	return &ConfiguratorItemDefinitionResolver{&gen.GeneratedConfiguratorItemDefinitionResolver{r.GeneratedResolver}}
}

type ConfiguratorAttributeDefinitionResultTypeResolver struct {
	*gen.GeneratedConfiguratorAttributeDefinitionResultTypeResolver
}

func (r *Resolver) ConfiguratorAttributeDefinitionResultType() gen.ConfiguratorAttributeDefinitionResultTypeResolver {
	return &ConfiguratorAttributeDefinitionResultTypeResolver{&gen.GeneratedConfiguratorAttributeDefinitionResultTypeResolver{r.GeneratedResolver}}
}

type ConfiguratorAttributeDefinitionResolver struct {
	*gen.GeneratedConfiguratorAttributeDefinitionResolver
}

func (r *Resolver) ConfiguratorAttributeDefinition() gen.ConfiguratorAttributeDefinitionResolver {
	return &ConfiguratorAttributeDefinitionResolver{&gen.GeneratedConfiguratorAttributeDefinitionResolver{r.GeneratedResolver}}
}

type ConfiguratorSlotDefinitionResultTypeResolver struct {
	*gen.GeneratedConfiguratorSlotDefinitionResultTypeResolver
}

func (r *Resolver) ConfiguratorSlotDefinitionResultType() gen.ConfiguratorSlotDefinitionResultTypeResolver {
	return &ConfiguratorSlotDefinitionResultTypeResolver{&gen.GeneratedConfiguratorSlotDefinitionResultTypeResolver{r.GeneratedResolver}}
}

type ConfiguratorSlotDefinitionResolver struct {
	*gen.GeneratedConfiguratorSlotDefinitionResolver
}

func (r *Resolver) ConfiguratorSlotDefinition() gen.ConfiguratorSlotDefinitionResolver {
	return &ConfiguratorSlotDefinitionResolver{&gen.GeneratedConfiguratorSlotDefinitionResolver{r.GeneratedResolver}}
}

type ConfiguratorItemResultTypeResolver struct {
	*gen.GeneratedConfiguratorItemResultTypeResolver
}

func (r *Resolver) ConfiguratorItemResultType() gen.ConfiguratorItemResultTypeResolver {
	return &ConfiguratorItemResultTypeResolver{&gen.GeneratedConfiguratorItemResultTypeResolver{r.GeneratedResolver}}
}

type ConfiguratorItemResolver struct {
	*gen.GeneratedConfiguratorItemResolver
}

func (r *Resolver) ConfiguratorItem() gen.ConfiguratorItemResolver {
	return &ConfiguratorItemResolver{&gen.GeneratedConfiguratorItemResolver{r.GeneratedResolver}}
}

type ConfiguratorAttributeResultTypeResolver struct {
	*gen.GeneratedConfiguratorAttributeResultTypeResolver
}

func (r *Resolver) ConfiguratorAttributeResultType() gen.ConfiguratorAttributeResultTypeResolver {
	return &ConfiguratorAttributeResultTypeResolver{&gen.GeneratedConfiguratorAttributeResultTypeResolver{r.GeneratedResolver}}
}

type ConfiguratorAttributeResolver struct {
	*gen.GeneratedConfiguratorAttributeResolver
}

func (r *Resolver) ConfiguratorAttribute() gen.ConfiguratorAttributeResolver {
	return &ConfiguratorAttributeResolver{&gen.GeneratedConfiguratorAttributeResolver{r.GeneratedResolver}}
}

type ConfiguratorSlotResultTypeResolver struct {
	*gen.GeneratedConfiguratorSlotResultTypeResolver
}

func (r *Resolver) ConfiguratorSlotResultType() gen.ConfiguratorSlotResultTypeResolver {
	return &ConfiguratorSlotResultTypeResolver{&gen.GeneratedConfiguratorSlotResultTypeResolver{r.GeneratedResolver}}
}

type ConfiguratorSlotResolver struct {
	*gen.GeneratedConfiguratorSlotResolver
}

func (r *Resolver) ConfiguratorSlot() gen.ConfiguratorSlotResolver {
	return &ConfiguratorSlotResolver{&gen.GeneratedConfiguratorSlotResolver{r.GeneratedResolver}}
}
