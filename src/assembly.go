package src

import (
	"context"

	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/graphql-services/graphql-configurator/gen"
)

type AssemblyHelper struct {
	r *gen.GeneratedResolver
}

func (as *AssemblyHelper) Load(ctx context.Context, ID string) (*gen.ConfiguratorAssembly, error) {
	cache := map[string]*gen.ConfiguratorAssemblyItem{}
	ctx, seg := xray.BeginSubsegment(ctx, "load-item")
	item, err := LoadItem(ctx, as.r, ID, cache)
	seg.Close(err)
	return &gen.ConfiguratorAssembly{
		ID:   ID,
		Item: item,
	}, err
}
func LoadItem(ctx context.Context, r *gen.GeneratedResolver, ID string, itemCache map[string]*gen.ConfiguratorAssemblyItem) (*gen.ConfiguratorAssemblyItem, error) {
	if item, exists := itemCache[ID]; exists {
		return item, nil
	}

	item, err := r.Handlers.QueryConfiguratorItem(ctx, r, gen.QueryConfiguratorItemHandlerOptions{ID: &ID})
	if err != nil {
		return nil, err
	}

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
	_item := &gen.ConfiguratorAssemblyItem{
		ID:           &item.ID,
		DefinitionID: item.DefinitionID,
		StockItemID:  item.StockItemID,
		Code:         item.Code,
		Name:         item.Name,
		Attributes:   attributes,
	}
	itemCache[ID] = _item

	_slots, err := r.Handlers.ConfiguratorItemSlots(ctx, r, item)
	if err != nil {
		return nil, err
	}
	slots := []*gen.ConfiguratorAssemblySlot{}
	for _, slot := range _slots {
		var slotItem *gen.ConfiguratorAssemblyItem
		if slot.ItemID != nil {
			slotItem, err = LoadItem(ctx, r, *slot.ItemID, itemCache)
			if err != nil {
				return nil, err
			}
		}

		if slot.DefinitionID != nil {
			slots = append(slots, &gen.ConfiguratorAssemblySlot{
				ID:           &slot.ID,
				DefinitionID: *slot.DefinitionID,
				Count:        slot.Count,
				Item:         slotItem,
			})
		}
	}

	_item.Slots = slots

	return _item, err
}
