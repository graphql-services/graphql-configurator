package src

import (
	"context"

	"github.com/graphql-services/graphql-configurator/gen"
)

type AssemblyInputHelper struct {
}

func (h *AssemblyInputHelper) Assemble(ctx context.Context, r *gen.GeneratedResolver, rootItem *gen.ConfiguratorAssemblyItemInput) (id string, err error) {
	return h.CreateItem(ctx, r, rootItem)
}
func (h *AssemblyInputHelper) CreateItem(ctx context.Context, r *gen.GeneratedResolver, inputItem *gen.ConfiguratorAssemblyItemInput) (string, error) {
	return createOrUpdateItem(ctx, r, inputItem)
}

func createOrUpdateItem(ctx context.Context, r *gen.GeneratedResolver, inputItem *gen.ConfiguratorAssemblyItemInput) (id string, err error) {
	_, err = r.Handlers.QueryConfiguratorItemDefinition(ctx, r, gen.QueryConfiguratorItemDefinitionHandlerOptions{
		ID: &inputItem.DefinitionID,
	})
	if err != nil {
		return
	}

	values := map[string]interface{}{
		"code":         inputItem.Code,
		"name":         inputItem.Name,
		"definitionId": inputItem.DefinitionID,
	}
	var item *gen.ConfiguratorItem

	if inputItem.ID != nil {
		item, err = r.Handlers.UpdateConfiguratorItem(ctx, r, *inputItem.ID, values)
	} else {
		item, err = r.Handlers.CreateConfiguratorItem(ctx, r, values)
	}
	if err != nil {
		return
	}

	for _, attrInput := range inputItem.Attributes {
		err = createOrUpdateAttribute(ctx, r, item.ID, attrInput)
		if err != nil {
			return
		}
	}
	for _, slotInput := range inputItem.Slots {
		err = createOrUpdateSlot(ctx, r, item.ID, slotInput)
		if err != nil {
			return
		}
	}

	id = item.ID

	return
}

func createOrUpdateAttribute(ctx context.Context, r *gen.GeneratedResolver, itemID string, input *gen.ConfiguratorAssemblyAttributeInput) (err error) {
	_, _err := r.Handlers.QueryConfiguratorAttributeDefinition(ctx, r, gen.QueryConfiguratorAttributeDefinitionHandlerOptions{
		ID: &input.DefinitionID,
	})
	if _err != nil {
		err = _err
		return
	}

	attrValues := map[string]interface{}{
		"itemId":       itemID,
		"definitionId": input.DefinitionID,
	}
	if input.StringValue != nil {
		attrValues["stringValue"] = *input.StringValue
	}
	if input.IntValue != nil {
		attrValues["intValue"] = *input.IntValue
	}
	if input.FloatValue != nil {
		attrValues["floatValue"] = *input.FloatValue
	}
	if input.ID != nil {
		_, err = r.Handlers.UpdateConfiguratorAttribute(ctx, r, *input.ID, attrValues)
	} else {
		_, err = r.Handlers.CreateConfiguratorAttribute(ctx, r, attrValues)
	}

	return
}

func createOrUpdateSlot(ctx context.Context, r *gen.GeneratedResolver, itemID string, input *gen.ConfiguratorAssemblySlotInput) (err error) {
	_, _err := r.Handlers.QueryConfiguratorSlotDefinition(ctx, r, gen.QueryConfiguratorSlotDefinitionHandlerOptions{
		ID: &input.DefinitionID,
	})
	if _err != nil {
		err = _err
		return
	}

	slotValues := map[string]interface{}{
		"parentItemId": itemID,
		"definitionId": input.DefinitionID,
	}

	if input.Items != nil {
		itemIds := []string{}
		for _, item := range input.Items {
			subItem, _err := r.Handlers.QueryConfiguratorItem(ctx, r, gen.QueryConfiguratorItemHandlerOptions{ID: item.ID})
			if _err != nil {
				err = _err
				return
			}
			itemIds = append(itemIds, subItem.ID)
		}
		slotValues["itemsIds"] = itemIds
	}

	if input.ID != nil {
		_, err = r.Handlers.UpdateConfiguratorSlot(ctx, r, *input.ID, slotValues)
	} else {
		_, err = r.Handlers.CreateConfiguratorSlot(ctx, r, slotValues)
	}
	// fmt.Println("!!!", err)

	return
}
