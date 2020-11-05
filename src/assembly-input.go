package src

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/graphql-services/graphql-configurator/gen"
	"github.com/patrickmn/go-cache"
)

type AssemblyInputHelper struct {
}

func (h *AssemblyInputHelper) CreateItem(ctx context.Context, r *gen.GeneratedResolver, inputItem *gen.ConfiguratorAssemblyItemInput) (id string, err error) {
	ctx, seg := xray.BeginSubsegment(ctx, "create-configurator-assembly")
	id, err = createItem(ctx, r, inputItem)
	seg.Close(err)
	return
}
func (h *AssemblyInputHelper) UpdateItem(ctx context.Context, r *gen.GeneratedResolver, inputItem *gen.ConfiguratorAssemblyItemInput) (id string, err error) {
	ctx, seg := xray.BeginSubsegment(ctx, "update-configurator-assembly")
	id, err = createOrUpdateItem(ctx, r, inputItem)
	seg.Close(err)
	return
}

func createItem(ctx context.Context, r *gen.GeneratedResolver, inputItem *gen.ConfiguratorAssemblyItemInput) (id string, err error) {
	values := map[string]interface{}{
		"code":         inputItem.Code,
		"name":         inputItem.Name,
		"definitionId": inputItem.DefinitionID,
		"stockItemId":  inputItem.StockItemID,
	}
	if inputItem.ID != nil {
		values["id"] = inputItem.ID
	}

	var item *gen.ConfiguratorItem
	item, err = r.Handlers.CreateConfiguratorItem(ctx, r, values)
	if err != nil {
		return
	}
	inputItem.ID = &item.ID
	return createOrUpdateItem(ctx, r, inputItem)
}
func createOrUpdateItem(ctx context.Context, r *gen.GeneratedResolver, inputItem *gen.ConfiguratorAssemblyItemInput) (id string, err error) {
	// _, err = r.Handlers.QueryConfiguratorItemDefinition(ctx, r, gen.QueryConfiguratorItemDefinitionHandlerOptions{
	// 	ID: inputItem.DefinitionID,
	// })
	// if err != nil {
	// 	return
	// }

	values := map[string]interface{}{
		"code":         inputItem.Code,
		"name":         inputItem.Name,
		"definitionId": inputItem.DefinitionID,
		"stockItemId":  inputItem.StockItemID,
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

	slotIDs := map[string]bool{}
	for _, slotInput := range inputItem.Slots {
		var slotID string
		slotID, err = createOrUpdateSlot(ctx, r, item.ID, slotInput)
		if err != nil {
			return
		}
		slotIDs[slotID] = true
	}

	slots, _err := r.Handlers.ConfiguratorItemSlots(ctx, r, item)
	if _err != nil {
		err = _err
		return
	}
	for _, slot := range slots {
		if _, ok := slotIDs[slot.ID]; !ok {
			fmt.Println("deleting slot", slot.ID)
			_, err = r.Handlers.DeleteConfiguratorSlot(ctx, r, slot.ID)
			if err != nil {
				return
			}
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

var slotDefinitionExistsCache *cache.Cache

func slotDefinitionExists(ctx context.Context, r *gen.GeneratedResolver, definitionId string) (res bool, err error) {
	if slotDefinitionExistsCache == nil {
		slotDefinitionExistsCache = cache.New(time.Minute, 10*time.Minute)
	}

	exists, found := slotDefinitionExistsCache.Get(definitionId)
	if found {
		res = exists.(bool)
		return
	}

	item, err := r.Handlers.QueryConfiguratorSlotDefinition(ctx, r, gen.QueryConfiguratorSlotDefinitionHandlerOptions{
		ID: &definitionId,
	})
	if err != nil {
		return
	}
	res = item != nil
	slotDefinitionExistsCache.Set(definitionId, res, cache.DefaultExpiration)
	return
}

func createOrUpdateSlot(ctx context.Context, r *gen.GeneratedResolver, itemID string, input *gen.ConfiguratorAssemblySlotInput) (slotID string, err error) {
	// _, _err := r.Handlers.QueryConfiguratorSlotDefinition(ctx, r, gen.QueryConfiguratorSlotDefinitionHandlerOptions{
	// 	ID: &input.DefinitionID,
	// })
	exists, err := slotDefinitionExists(ctx, r, input.DefinitionID)
	if err != nil {
		return
	}
	if !exists {
		err = fmt.Errorf("Slot definition '%s' not found", input.DefinitionID)
		return
	}

	slotValues := map[string]interface{}{
		"parentItemId": itemID,
		"definitionId": input.DefinitionID,
		"count":        input.Count,
	}

	if input.Item != nil {
		var subItemID string
		if input.Item.ID != nil {
			subItemID = *input.Item.ID
		} else {
			subItemID, err = createOrUpdateItem(ctx, r, input.Item)
			if err != nil {
				return
			}
		}
		slotValues["itemId"] = subItemID
	}

	var slot *gen.ConfiguratorSlot
	if input.ID != nil {
		slot, err = r.Handlers.UpdateConfiguratorSlot(ctx, r, *input.ID, slotValues)
	} else {
		slot, err = r.Handlers.CreateConfiguratorSlot(ctx, r, slotValues)
	}
	if err == nil {
		slotID = slot.ID
	}
	// fmt.Println("!!!", err)

	return
}
