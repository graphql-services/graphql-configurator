package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/novacloudcz/graphql-orm/events"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

type MutationEvents struct {
	Events []events.Event
}

func EnrichContextWithMutations(ctx context.Context, r *GeneratedResolver) context.Context {
	_ctx := context.WithValue(ctx, KeyMutationTransaction, r.DB.db.Begin())
	_ctx = context.WithValue(_ctx, KeyMutationEvents, &MutationEvents{})
	return _ctx
}
func FinishMutationContext(ctx context.Context, r *GeneratedResolver) (err error) {
	s := GetMutationEventStore(ctx)

	for _, event := range s.Events {
		err = r.Handlers.OnEvent(ctx, r, &event)
		if err != nil {
			return
		}
	}

	tx := GetTransaction(ctx)
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	for _, event := range s.Events {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func GetTransaction(ctx context.Context) *gorm.DB {
	return ctx.Value(KeyMutationTransaction).(*gorm.DB)
}
func GetMutationEventStore(ctx context.Context) *MutationEvents {
	return ctx.Value(KeyMutationEvents).(*MutationEvents)
}
func AddMutationEvent(ctx context.Context, e events.Event) {
	s := GetMutationEventStore(ctx)
	s.Events = append(s.Events, e)
}

func (r *GeneratedMutationResolver) CreateConfiguratorItemDefinition(ctx context.Context, input map[string]interface{}) (item *ConfiguratorItemDefinition, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateConfiguratorItemDefinition(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateConfiguratorItemDefinitionHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ConfiguratorItemDefinition, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &ConfiguratorItemDefinition{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "ConfiguratorItemDefinition",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ConfiguratorItemDefinitionChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["attributesIds"]; exists {
		items := []ConfiguratorAttributeDefinition{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Attributes")
		association.Replace(items)
	}

	if ids, exists := input["slotsIds"]; exists {
		items := []ConfiguratorSlotDefinition{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Slots")
		association.Replace(items)
	}

	if ids, exists := input["itemsIds"]; exists {
		items := []ConfiguratorItem{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Items")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateConfiguratorItemDefinition(ctx context.Context, id string, input map[string]interface{}) (item *ConfiguratorItemDefinition, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateConfiguratorItemDefinition(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateConfiguratorItemDefinitionHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ConfiguratorItemDefinition, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ConfiguratorItemDefinition{}
	now := time.Now()
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "ConfiguratorItemDefinition",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ConfiguratorItemDefinitionChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["attributesIds"]; exists {
		items := []ConfiguratorAttributeDefinition{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Attributes")
		association.Replace(items)
	}

	if ids, exists := input["slotsIds"]; exists {
		items := []ConfiguratorSlotDefinition{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Slots")
		association.Replace(items)
	}

	if ids, exists := input["itemsIds"]; exists {
		items := []ConfiguratorItem{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Items")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteConfiguratorItemDefinition(ctx context.Context, id string) (item *ConfiguratorItemDefinition, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteConfiguratorItemDefinition(ctx, r.GeneratedResolver, id)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func DeleteConfiguratorItemDefinitionHandler(ctx context.Context, r *GeneratedResolver, id string) (item *ConfiguratorItemDefinition, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ConfiguratorItemDefinition{}
	now := time.Now()
	tx := GetTransaction(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "ConfiguratorItemDefinition",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("configurator_item_definitions")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteAllConfiguratorItemDefinitions(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllConfiguratorItemDefinitions(ctx, r.GeneratedResolver)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}
func DeleteAllConfiguratorItemDefinitionsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := GetTransaction(ctx)
	err := tx.Delete(&ConfiguratorItemDefinition{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) CreateConfiguratorAttributeDefinition(ctx context.Context, input map[string]interface{}) (item *ConfiguratorAttributeDefinition, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateConfiguratorAttributeDefinition(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateConfiguratorAttributeDefinitionHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ConfiguratorAttributeDefinition, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &ConfiguratorAttributeDefinition{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "ConfiguratorAttributeDefinition",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ConfiguratorAttributeDefinitionChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["type"]; ok && (item.Type != changes.Type) && (item.Type == nil || changes.Type == nil || *item.Type != *changes.Type) {
		item.Type = changes.Type

		event.AddNewValue("type", changes.Type)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["itemDefinitionsIds"]; exists {
		items := []ConfiguratorItemDefinition{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("ItemDefinitions")
		association.Replace(items)
	}

	if ids, exists := input["attributesIds"]; exists {
		items := []ConfiguratorAttribute{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Attributes")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateConfiguratorAttributeDefinition(ctx context.Context, id string, input map[string]interface{}) (item *ConfiguratorAttributeDefinition, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateConfiguratorAttributeDefinition(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateConfiguratorAttributeDefinitionHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ConfiguratorAttributeDefinition, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ConfiguratorAttributeDefinition{}
	now := time.Now()
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "ConfiguratorAttributeDefinition",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ConfiguratorAttributeDefinitionChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["type"]; ok && (item.Type != changes.Type) && (item.Type == nil || changes.Type == nil || *item.Type != *changes.Type) {
		event.AddOldValue("type", item.Type)
		event.AddNewValue("type", changes.Type)
		item.Type = changes.Type
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["itemDefinitionsIds"]; exists {
		items := []ConfiguratorItemDefinition{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("ItemDefinitions")
		association.Replace(items)
	}

	if ids, exists := input["attributesIds"]; exists {
		items := []ConfiguratorAttribute{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Attributes")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteConfiguratorAttributeDefinition(ctx context.Context, id string) (item *ConfiguratorAttributeDefinition, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteConfiguratorAttributeDefinition(ctx, r.GeneratedResolver, id)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func DeleteConfiguratorAttributeDefinitionHandler(ctx context.Context, r *GeneratedResolver, id string) (item *ConfiguratorAttributeDefinition, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ConfiguratorAttributeDefinition{}
	now := time.Now()
	tx := GetTransaction(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "ConfiguratorAttributeDefinition",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("configurator_attribute_definitions")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteAllConfiguratorAttributeDefinitions(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllConfiguratorAttributeDefinitions(ctx, r.GeneratedResolver)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}
func DeleteAllConfiguratorAttributeDefinitionsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := GetTransaction(ctx)
	err := tx.Delete(&ConfiguratorAttributeDefinition{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) CreateConfiguratorSlotDefinition(ctx context.Context, input map[string]interface{}) (item *ConfiguratorSlotDefinition, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateConfiguratorSlotDefinition(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateConfiguratorSlotDefinitionHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ConfiguratorSlotDefinition, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &ConfiguratorSlotDefinition{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "ConfiguratorSlotDefinition",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ConfiguratorSlotDefinitionChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["itemDefinitionId"]; ok && (item.ItemDefinitionID != changes.ItemDefinitionID) && (item.ItemDefinitionID == nil || changes.ItemDefinitionID == nil || *item.ItemDefinitionID != *changes.ItemDefinitionID) {
		item.ItemDefinitionID = changes.ItemDefinitionID

		event.AddNewValue("itemDefinitionId", changes.ItemDefinitionID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["slotsIds"]; exists {
		items := []ConfiguratorSlot{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Slots")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateConfiguratorSlotDefinition(ctx context.Context, id string, input map[string]interface{}) (item *ConfiguratorSlotDefinition, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateConfiguratorSlotDefinition(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateConfiguratorSlotDefinitionHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ConfiguratorSlotDefinition, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ConfiguratorSlotDefinition{}
	now := time.Now()
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "ConfiguratorSlotDefinition",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ConfiguratorSlotDefinitionChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["itemDefinitionId"]; ok && (item.ItemDefinitionID != changes.ItemDefinitionID) && (item.ItemDefinitionID == nil || changes.ItemDefinitionID == nil || *item.ItemDefinitionID != *changes.ItemDefinitionID) {
		event.AddOldValue("itemDefinitionId", item.ItemDefinitionID)
		event.AddNewValue("itemDefinitionId", changes.ItemDefinitionID)
		item.ItemDefinitionID = changes.ItemDefinitionID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["slotsIds"]; exists {
		items := []ConfiguratorSlot{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Slots")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteConfiguratorSlotDefinition(ctx context.Context, id string) (item *ConfiguratorSlotDefinition, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteConfiguratorSlotDefinition(ctx, r.GeneratedResolver, id)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func DeleteConfiguratorSlotDefinitionHandler(ctx context.Context, r *GeneratedResolver, id string) (item *ConfiguratorSlotDefinition, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ConfiguratorSlotDefinition{}
	now := time.Now()
	tx := GetTransaction(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "ConfiguratorSlotDefinition",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("configurator_slot_definitions")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteAllConfiguratorSlotDefinitions(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllConfiguratorSlotDefinitions(ctx, r.GeneratedResolver)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}
func DeleteAllConfiguratorSlotDefinitionsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := GetTransaction(ctx)
	err := tx.Delete(&ConfiguratorSlotDefinition{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) CreateConfiguratorItem(ctx context.Context, input map[string]interface{}) (item *ConfiguratorItem, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateConfiguratorItem(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateConfiguratorItemHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ConfiguratorItem, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &ConfiguratorItem{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "ConfiguratorItem",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ConfiguratorItemChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["stockItemID"]; ok && (item.StockItemID != changes.StockItemID) && (item.StockItemID == nil || changes.StockItemID == nil || *item.StockItemID != *changes.StockItemID) {
		item.StockItemID = changes.StockItemID

		event.AddNewValue("stockItemID", changes.StockItemID)
	}

	if _, ok := input["definitionId"]; ok && (item.DefinitionID != changes.DefinitionID) && (item.DefinitionID == nil || changes.DefinitionID == nil || *item.DefinitionID != *changes.DefinitionID) {
		item.DefinitionID = changes.DefinitionID

		event.AddNewValue("definitionId", changes.DefinitionID)
	}

	if _, ok := input["parentSlotId"]; ok && (item.ParentSlotID != changes.ParentSlotID) && (item.ParentSlotID == nil || changes.ParentSlotID == nil || *item.ParentSlotID != *changes.ParentSlotID) {
		item.ParentSlotID = changes.ParentSlotID

		event.AddNewValue("parentSlotId", changes.ParentSlotID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["attributesIds"]; exists {
		items := []ConfiguratorAttribute{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Attributes")
		association.Replace(items)
	}

	if ids, exists := input["slotsIds"]; exists {
		items := []ConfiguratorSlot{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Slots")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateConfiguratorItem(ctx context.Context, id string, input map[string]interface{}) (item *ConfiguratorItem, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateConfiguratorItem(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateConfiguratorItemHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ConfiguratorItem, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ConfiguratorItem{}
	now := time.Now()
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "ConfiguratorItem",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ConfiguratorItemChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["stockItemID"]; ok && (item.StockItemID != changes.StockItemID) && (item.StockItemID == nil || changes.StockItemID == nil || *item.StockItemID != *changes.StockItemID) {
		event.AddOldValue("stockItemID", item.StockItemID)
		event.AddNewValue("stockItemID", changes.StockItemID)
		item.StockItemID = changes.StockItemID
	}

	if _, ok := input["definitionId"]; ok && (item.DefinitionID != changes.DefinitionID) && (item.DefinitionID == nil || changes.DefinitionID == nil || *item.DefinitionID != *changes.DefinitionID) {
		event.AddOldValue("definitionId", item.DefinitionID)
		event.AddNewValue("definitionId", changes.DefinitionID)
		item.DefinitionID = changes.DefinitionID
	}

	if _, ok := input["parentSlotId"]; ok && (item.ParentSlotID != changes.ParentSlotID) && (item.ParentSlotID == nil || changes.ParentSlotID == nil || *item.ParentSlotID != *changes.ParentSlotID) {
		event.AddOldValue("parentSlotId", item.ParentSlotID)
		event.AddNewValue("parentSlotId", changes.ParentSlotID)
		item.ParentSlotID = changes.ParentSlotID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["attributesIds"]; exists {
		items := []ConfiguratorAttribute{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Attributes")
		association.Replace(items)
	}

	if ids, exists := input["slotsIds"]; exists {
		items := []ConfiguratorSlot{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Slots")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteConfiguratorItem(ctx context.Context, id string) (item *ConfiguratorItem, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteConfiguratorItem(ctx, r.GeneratedResolver, id)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func DeleteConfiguratorItemHandler(ctx context.Context, r *GeneratedResolver, id string) (item *ConfiguratorItem, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ConfiguratorItem{}
	now := time.Now()
	tx := GetTransaction(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "ConfiguratorItem",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("configurator_items")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteAllConfiguratorItems(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllConfiguratorItems(ctx, r.GeneratedResolver)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}
func DeleteAllConfiguratorItemsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := GetTransaction(ctx)
	err := tx.Delete(&ConfiguratorItem{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) CreateConfiguratorAttribute(ctx context.Context, input map[string]interface{}) (item *ConfiguratorAttribute, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateConfiguratorAttribute(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateConfiguratorAttributeHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ConfiguratorAttribute, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &ConfiguratorAttribute{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "ConfiguratorAttribute",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ConfiguratorAttributeChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["stringValue"]; ok && (item.StringValue != changes.StringValue) && (item.StringValue == nil || changes.StringValue == nil || *item.StringValue != *changes.StringValue) {
		item.StringValue = changes.StringValue

		event.AddNewValue("stringValue", changes.StringValue)
	}

	if _, ok := input["floatValue"]; ok && (item.FloatValue != changes.FloatValue) && (item.FloatValue == nil || changes.FloatValue == nil || *item.FloatValue != *changes.FloatValue) {
		item.FloatValue = changes.FloatValue

		event.AddNewValue("floatValue", changes.FloatValue)
	}

	if _, ok := input["intValue"]; ok && (item.IntValue != changes.IntValue) && (item.IntValue == nil || changes.IntValue == nil || *item.IntValue != *changes.IntValue) {
		item.IntValue = changes.IntValue

		event.AddNewValue("intValue", changes.IntValue)
	}

	if _, ok := input["definitionId"]; ok && (item.DefinitionID != changes.DefinitionID) && (item.DefinitionID == nil || changes.DefinitionID == nil || *item.DefinitionID != *changes.DefinitionID) {
		item.DefinitionID = changes.DefinitionID

		event.AddNewValue("definitionId", changes.DefinitionID)
	}

	if _, ok := input["itemId"]; ok && (item.ItemID != changes.ItemID) && (item.ItemID == nil || changes.ItemID == nil || *item.ItemID != *changes.ItemID) {
		item.ItemID = changes.ItemID

		event.AddNewValue("itemId", changes.ItemID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateConfiguratorAttribute(ctx context.Context, id string, input map[string]interface{}) (item *ConfiguratorAttribute, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateConfiguratorAttribute(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateConfiguratorAttributeHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ConfiguratorAttribute, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ConfiguratorAttribute{}
	now := time.Now()
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "ConfiguratorAttribute",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ConfiguratorAttributeChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["stringValue"]; ok && (item.StringValue != changes.StringValue) && (item.StringValue == nil || changes.StringValue == nil || *item.StringValue != *changes.StringValue) {
		event.AddOldValue("stringValue", item.StringValue)
		event.AddNewValue("stringValue", changes.StringValue)
		item.StringValue = changes.StringValue
	}

	if _, ok := input["floatValue"]; ok && (item.FloatValue != changes.FloatValue) && (item.FloatValue == nil || changes.FloatValue == nil || *item.FloatValue != *changes.FloatValue) {
		event.AddOldValue("floatValue", item.FloatValue)
		event.AddNewValue("floatValue", changes.FloatValue)
		item.FloatValue = changes.FloatValue
	}

	if _, ok := input["intValue"]; ok && (item.IntValue != changes.IntValue) && (item.IntValue == nil || changes.IntValue == nil || *item.IntValue != *changes.IntValue) {
		event.AddOldValue("intValue", item.IntValue)
		event.AddNewValue("intValue", changes.IntValue)
		item.IntValue = changes.IntValue
	}

	if _, ok := input["definitionId"]; ok && (item.DefinitionID != changes.DefinitionID) && (item.DefinitionID == nil || changes.DefinitionID == nil || *item.DefinitionID != *changes.DefinitionID) {
		event.AddOldValue("definitionId", item.DefinitionID)
		event.AddNewValue("definitionId", changes.DefinitionID)
		item.DefinitionID = changes.DefinitionID
	}

	if _, ok := input["itemId"]; ok && (item.ItemID != changes.ItemID) && (item.ItemID == nil || changes.ItemID == nil || *item.ItemID != *changes.ItemID) {
		event.AddOldValue("itemId", item.ItemID)
		event.AddNewValue("itemId", changes.ItemID)
		item.ItemID = changes.ItemID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteConfiguratorAttribute(ctx context.Context, id string) (item *ConfiguratorAttribute, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteConfiguratorAttribute(ctx, r.GeneratedResolver, id)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func DeleteConfiguratorAttributeHandler(ctx context.Context, r *GeneratedResolver, id string) (item *ConfiguratorAttribute, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ConfiguratorAttribute{}
	now := time.Now()
	tx := GetTransaction(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "ConfiguratorAttribute",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("configurator_attributes")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteAllConfiguratorAttributes(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllConfiguratorAttributes(ctx, r.GeneratedResolver)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}
func DeleteAllConfiguratorAttributesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := GetTransaction(ctx)
	err := tx.Delete(&ConfiguratorAttribute{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

func (r *GeneratedMutationResolver) CreateConfiguratorSlot(ctx context.Context, input map[string]interface{}) (item *ConfiguratorSlot, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateConfiguratorSlot(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateConfiguratorSlotHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ConfiguratorSlot, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &ConfiguratorSlot{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "ConfiguratorSlot",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ConfiguratorSlotChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["definitionId"]; ok && (item.DefinitionID != changes.DefinitionID) && (item.DefinitionID == nil || changes.DefinitionID == nil || *item.DefinitionID != *changes.DefinitionID) {
		item.DefinitionID = changes.DefinitionID

		event.AddNewValue("definitionId", changes.DefinitionID)
	}

	if _, ok := input["parentItemId"]; ok && (item.ParentItemID != changes.ParentItemID) && (item.ParentItemID == nil || changes.ParentItemID == nil || *item.ParentItemID != *changes.ParentItemID) {
		item.ParentItemID = changes.ParentItemID

		event.AddNewValue("parentItemId", changes.ParentItemID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["itemsIds"]; exists {
		items := []ConfiguratorItem{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Items")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateConfiguratorSlot(ctx context.Context, id string, input map[string]interface{}) (item *ConfiguratorSlot, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateConfiguratorSlot(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateConfiguratorSlotHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ConfiguratorSlot, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ConfiguratorSlot{}
	now := time.Now()
	tx := GetTransaction(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "ConfiguratorSlot",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ConfiguratorSlotChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["definitionId"]; ok && (item.DefinitionID != changes.DefinitionID) && (item.DefinitionID == nil || changes.DefinitionID == nil || *item.DefinitionID != *changes.DefinitionID) {
		event.AddOldValue("definitionId", item.DefinitionID)
		event.AddNewValue("definitionId", changes.DefinitionID)
		item.DefinitionID = changes.DefinitionID
	}

	if _, ok := input["parentItemId"]; ok && (item.ParentItemID != changes.ParentItemID) && (item.ParentItemID == nil || changes.ParentItemID == nil || *item.ParentItemID != *changes.ParentItemID) {
		event.AddOldValue("parentItemId", item.ParentItemID)
		event.AddNewValue("parentItemId", changes.ParentItemID)
		item.ParentItemID = changes.ParentItemID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["itemsIds"]; exists {
		items := []ConfiguratorItem{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Items")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteConfiguratorSlot(ctx context.Context, id string) (item *ConfiguratorSlot, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteConfiguratorSlot(ctx, r.GeneratedResolver, id)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func DeleteConfiguratorSlotHandler(ctx context.Context, r *GeneratedResolver, id string) (item *ConfiguratorSlot, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ConfiguratorSlot{}
	now := time.Now()
	tx := GetTransaction(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "ConfiguratorSlot",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("configurator_slots")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteAllConfiguratorSlots(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllConfiguratorSlots(ctx, r.GeneratedResolver)
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}
func DeleteAllConfiguratorSlotsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := GetTransaction(ctx)
	err := tx.Delete(&ConfiguratorSlot{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}
