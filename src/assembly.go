package src

import (
	"context"

	"github.com/graphql-services/graphql-configurator/gen"
)

type AssemblyHelper struct {
	r *gen.GeneratedResolver
}

func (as *AssemblyHelper) Load(ctx context.Context, ID string) (*gen.ConfiguratorAssembly, error) {
	item, err := as.LoadItem(ctx, ID)
	return &gen.ConfiguratorAssembly{
		ID:   ID,
		Item: item,
	}, err
}
func (as *AssemblyHelper) LoadItem(ctx context.Context, ID string) (*gen.ConfiguratorAssemblyItem, error) {
	item, err := as.r.Handlers.QueryConfiguratorItem(ctx, as.r, gen.QueryConfiguratorItemHandlerOptions{ID: &ID})
	if err != nil {
		return nil, err
	}

	_attributes, err := as.r.Handlers.ConfiguratorItemAttributes(ctx, as.r, item)
	if err != nil {
		return nil, err
	}

	attributes := []*gen.ConfiguratorAssemblyAttribute{}
	for _, attr := range _attributes {
		attributes = append(attributes, &gen.ConfiguratorAssemblyAttribute{
			ID:           attr.ID,
			DefinitionID: *attr.DefinitionID,
			StringValue:  attr.StringValue,
			IntValue:     attr.IntValue,
			FloatValue:   attr.FloatValue,
		})
	}

	_slots, err := as.r.Handlers.ConfiguratorItemSlots(ctx, as.r, item)
	if err != nil {
		return nil, err
	}
	slots := []*gen.ConfiguratorAssemblySlot{}
	for _, slot := range _slots {
		slotItem, err := as.LoadItem(ctx, *slot.ItemID)
		if err != nil {
			return nil, err
		}

		slots = append(slots, &gen.ConfiguratorAssemblySlot{
			ID:           slot.ID,
			DefinitionID: *slot.DefinitionID,
			Item:         slotItem,
		})
	}

	return &gen.ConfiguratorAssemblyItem{
		ID:           item.ID,
		Code:         item.Code,
		Name:         item.Name,
		DefinitionID: *item.DefinitionID,
		Attributes:   attributes,
		Slots:        slots,
	}, err
}
