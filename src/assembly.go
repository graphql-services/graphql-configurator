package src

import (
	"context"

	"github.com/graphql-services/graphql-configurator/gen"
)

type AssemblyHelper struct {
	r *gen.GeneratedResolver
}

func (as *AssemblyHelper) Load(ctx context.Context, ID string) (*gen.ConfiguratorAssembly, error) {
	item, err := LoadItem(ctx, as.r, ID)
	return &gen.ConfiguratorAssembly{
		ID:   ID,
		Item: item,
	}, err
}
func LoadItem(ctx context.Context, r *gen.GeneratedResolver, ID string) (*gen.ConfiguratorAssemblyItem, error) {
	item, err := r.Handlers.QueryConfiguratorItem(ctx, r, gen.QueryConfiguratorItemHandlerOptions{ID: &ID})
	if err != nil {
		return nil, err
	}

	// isNotTemplate := item.sto == nil
	// if !isNotTemplate {
	// 	rawData := *item.RawData
	// 	var item *gen.ConfiguratorAssemblyItem
	// 	err := json.Unmarshal([]byte(rawData), &item)
	// 	return item, err
	// }

	_attributes, err := r.Handlers.ConfiguratorItemAttributes(ctx, r, item)
	if err != nil {
		return nil, err
	}

	attributes := []*gen.ConfiguratorAssemblyAttribute{}
	for _, attr := range _attributes {
		attributes = append(attributes, &gen.ConfiguratorAssemblyAttribute{
			ID:           &attr.ID,
			DefinitionID: *attr.DefinitionID,
			StringValue:  attr.StringValue,
			IntValue:     attr.IntValue,
			FloatValue:   attr.FloatValue,
		})
	}

	_slots, err := r.Handlers.ConfiguratorItemSlots(ctx, r, item)
	if err != nil {
		return nil, err
	}
	slots := []*gen.ConfiguratorAssemblySlot{}
	for _, slot := range _slots {
		slotItem, err := LoadItem(ctx, r, *slot.ItemID)
		if err != nil {
			return nil, err
		}

		slots = append(slots, &gen.ConfiguratorAssemblySlot{
			ID:           &slot.ID,
			DefinitionID: *slot.DefinitionID,
			Item:         slotItem,
		})
	}

	return &gen.ConfiguratorAssemblyItem{
		ID:           &item.ID,
		DefinitionID: item.DefinitionID,
		StockItemID:  item.StockItemID,
		Code:         item.Code,
		Name:         item.Name,
		Attributes:   attributes,
		Slots:        slots,
	}, err
}
