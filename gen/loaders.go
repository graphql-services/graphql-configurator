package gen

import (
	"context"

	"github.com/graph-gophers/dataloader"
)

func GetLoaders(db *DB) map[string]*dataloader.Loader {
	loaders := map[string]*dataloader.Loader{}

	configurator_item_definition_categoriesBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]ConfiguratorItemDefinitionCategory{}
		res := db.Query().Find(items, "id IN (?)", ids)
		if res.Error != nil && !res.RecordNotFound() {
			return []*dataloader.Result{
				&dataloader.Result{Error: res.Error},
			}
		}

		itemMap := make(map[string]ConfiguratorItemDefinitionCategory, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("ConfiguratorItemDefinitionCategory with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["ConfiguratorItemDefinitionCategory"] = dataloader.NewBatchedLoader(configurator_item_definition_categoriesBatchFn, dataloader.WithClearCacheOnBatch())

	configurator_item_definitionsBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]ConfiguratorItemDefinition{}
		res := db.Query().Find(items, "id IN (?)", ids)
		if res.Error != nil && !res.RecordNotFound() {
			return []*dataloader.Result{
				&dataloader.Result{Error: res.Error},
			}
		}

		itemMap := make(map[string]ConfiguratorItemDefinition, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("ConfiguratorItemDefinition with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["ConfiguratorItemDefinition"] = dataloader.NewBatchedLoader(configurator_item_definitionsBatchFn, dataloader.WithClearCacheOnBatch())

	configurator_attribute_definitionsBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]ConfiguratorAttributeDefinition{}
		res := db.Query().Find(items, "id IN (?)", ids)
		if res.Error != nil && !res.RecordNotFound() {
			return []*dataloader.Result{
				&dataloader.Result{Error: res.Error},
			}
		}

		itemMap := make(map[string]ConfiguratorAttributeDefinition, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("ConfiguratorAttributeDefinition with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["ConfiguratorAttributeDefinition"] = dataloader.NewBatchedLoader(configurator_attribute_definitionsBatchFn, dataloader.WithClearCacheOnBatch())

	configurator_slot_definitionsBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]ConfiguratorSlotDefinition{}
		res := db.Query().Find(items, "id IN (?)", ids)
		if res.Error != nil && !res.RecordNotFound() {
			return []*dataloader.Result{
				&dataloader.Result{Error: res.Error},
			}
		}

		itemMap := make(map[string]ConfiguratorSlotDefinition, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("ConfiguratorSlotDefinition with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["ConfiguratorSlotDefinition"] = dataloader.NewBatchedLoader(configurator_slot_definitionsBatchFn, dataloader.WithClearCacheOnBatch())

	configurator_itemsBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]ConfiguratorItem{}
		res := db.Query().Find(items, "id IN (?)", ids)
		if res.Error != nil && !res.RecordNotFound() {
			return []*dataloader.Result{
				&dataloader.Result{Error: res.Error},
			}
		}

		itemMap := make(map[string]ConfiguratorItem, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("ConfiguratorItem with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["ConfiguratorItem"] = dataloader.NewBatchedLoader(configurator_itemsBatchFn, dataloader.WithClearCacheOnBatch())

	configurator_attributesBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]ConfiguratorAttribute{}
		res := db.Query().Find(items, "id IN (?)", ids)
		if res.Error != nil && !res.RecordNotFound() {
			return []*dataloader.Result{
				&dataloader.Result{Error: res.Error},
			}
		}

		itemMap := make(map[string]ConfiguratorAttribute, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("ConfiguratorAttribute with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["ConfiguratorAttribute"] = dataloader.NewBatchedLoader(configurator_attributesBatchFn, dataloader.WithClearCacheOnBatch())

	configurator_slotsBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]ConfiguratorSlot{}
		res := db.Query().Find(items, "id IN (?)", ids)
		if res.Error != nil && !res.RecordNotFound() {
			return []*dataloader.Result{
				&dataloader.Result{Error: res.Error},
			}
		}

		itemMap := make(map[string]ConfiguratorSlot, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("ConfiguratorSlot with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["ConfiguratorSlot"] = dataloader.NewBatchedLoader(configurator_slotsBatchFn, dataloader.WithClearCacheOnBatch())

	return loaders
}
