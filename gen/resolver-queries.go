package gen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
	"github.com/vektah/gqlparser/ast"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

type QueryConfiguratorItemDefinitionHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *ConfiguratorItemDefinitionFilterType
}

func (r *GeneratedQueryResolver) ConfiguratorItemDefinition(ctx context.Context, id *string, q *string, filter *ConfiguratorItemDefinitionFilterType) (*ConfiguratorItemDefinition, error) {
	opts := QueryConfiguratorItemDefinitionHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryConfiguratorItemDefinition(ctx, r.GeneratedResolver, opts)
}
func QueryConfiguratorItemDefinitionHandler(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorItemDefinitionHandlerOptions) (*ConfiguratorItemDefinition, error) {
	query := ConfiguratorItemDefinitionQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &ConfiguratorItemDefinitionResultType{
		EntityResultType: EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("configurator_item_definitions")+".id = ?", *opts.ID)
	}

	var items []*ConfiguratorItemDefinition
	giOpts := GetItemsOptions{
		Alias:      TableName("configurator_item_definitions"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "ConfiguratorItemDefinition"}
	}
	return items[0], err
}

type QueryConfiguratorItemDefinitionsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*ConfiguratorItemDefinitionSortType
	Filter *ConfiguratorItemDefinitionFilterType
}

func (r *GeneratedQueryResolver) ConfiguratorItemDefinitions(ctx context.Context, offset *int, limit *int, q *string, sort []*ConfiguratorItemDefinitionSortType, filter *ConfiguratorItemDefinitionFilterType) (*ConfiguratorItemDefinitionResultType, error) {
	opts := QueryConfiguratorItemDefinitionsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryConfiguratorItemDefinitions(ctx, r.GeneratedResolver, opts)
}
func QueryConfiguratorItemDefinitionsHandler(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorItemDefinitionsHandlerOptions) (*ConfiguratorItemDefinitionResultType, error) {
	query := ConfiguratorItemDefinitionQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &ConfiguratorItemDefinitionResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedConfiguratorItemDefinitionResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedConfiguratorItemDefinitionResultTypeResolver) Items(ctx context.Context, obj *ConfiguratorItemDefinitionResultType) (items []*ConfiguratorItemDefinition, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("configurator_item_definitions"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, giOpts, &items)

	return
}

func (r *GeneratedConfiguratorItemDefinitionResultTypeResolver) Count(ctx context.Context, obj *ConfiguratorItemDefinitionResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &ConfiguratorItemDefinition{})
}

type GeneratedConfiguratorItemDefinitionResolver struct{ *GeneratedResolver }

func (r *GeneratedConfiguratorItemDefinitionResolver) Attributes(ctx context.Context, obj *ConfiguratorItemDefinition) (res []*ConfiguratorAttributeDefinition, err error) {
	return r.Handlers.ConfiguratorItemDefinitionAttributes(ctx, r, obj)
}
func ConfiguratorItemDefinitionAttributesHandler(ctx context.Context, r *GeneratedConfiguratorItemDefinitionResolver, obj *ConfiguratorItemDefinition) (res []*ConfiguratorAttributeDefinition, err error) {

	items := []*ConfiguratorAttributeDefinition{}
	err = r.DB.Query().Model(obj).Related(&items, "Attributes").Error
	res = items

	return
}

func (r *GeneratedConfiguratorItemDefinitionResolver) AttributesIds(ctx context.Context, obj *ConfiguratorItemDefinition) (ids []string, err error) {
	ids = []string{}

	items := []*ConfiguratorAttributeDefinition{}
	err = r.DB.Query().Model(obj).Select(TableName("configurator_attribute_definitions")+".id").Related(&items, "Attributes").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

func (r *GeneratedConfiguratorItemDefinitionResolver) Slots(ctx context.Context, obj *ConfiguratorItemDefinition) (res []*ConfiguratorSlotDefinition, err error) {
	return r.Handlers.ConfiguratorItemDefinitionSlots(ctx, r, obj)
}
func ConfiguratorItemDefinitionSlotsHandler(ctx context.Context, r *GeneratedConfiguratorItemDefinitionResolver, obj *ConfiguratorItemDefinition) (res []*ConfiguratorSlotDefinition, err error) {

	items := []*ConfiguratorSlotDefinition{}
	err = r.DB.Query().Model(obj).Related(&items, "Slots").Error
	res = items

	return
}

func (r *GeneratedConfiguratorItemDefinitionResolver) SlotsIds(ctx context.Context, obj *ConfiguratorItemDefinition) (ids []string, err error) {
	ids = []string{}

	items := []*ConfiguratorSlotDefinition{}
	err = r.DB.Query().Model(obj).Select(TableName("configurator_slot_definitions")+".id").Related(&items, "Slots").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

func (r *GeneratedConfiguratorItemDefinitionResolver) Items(ctx context.Context, obj *ConfiguratorItemDefinition) (res []*ConfiguratorItem, err error) {
	return r.Handlers.ConfiguratorItemDefinitionItems(ctx, r, obj)
}
func ConfiguratorItemDefinitionItemsHandler(ctx context.Context, r *GeneratedConfiguratorItemDefinitionResolver, obj *ConfiguratorItemDefinition) (res []*ConfiguratorItem, err error) {

	items := []*ConfiguratorItem{}
	err = r.DB.Query().Model(obj).Related(&items, "Items").Error
	res = items

	return
}

func (r *GeneratedConfiguratorItemDefinitionResolver) ItemsIds(ctx context.Context, obj *ConfiguratorItemDefinition) (ids []string, err error) {
	ids = []string{}

	items := []*ConfiguratorItem{}
	err = r.DB.Query().Model(obj).Select(TableName("configurator_items")+".id").Related(&items, "Items").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

type QueryConfiguratorAttributeDefinitionHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *ConfiguratorAttributeDefinitionFilterType
}

func (r *GeneratedQueryResolver) ConfiguratorAttributeDefinition(ctx context.Context, id *string, q *string, filter *ConfiguratorAttributeDefinitionFilterType) (*ConfiguratorAttributeDefinition, error) {
	opts := QueryConfiguratorAttributeDefinitionHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryConfiguratorAttributeDefinition(ctx, r.GeneratedResolver, opts)
}
func QueryConfiguratorAttributeDefinitionHandler(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorAttributeDefinitionHandlerOptions) (*ConfiguratorAttributeDefinition, error) {
	query := ConfiguratorAttributeDefinitionQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &ConfiguratorAttributeDefinitionResultType{
		EntityResultType: EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("configurator_attribute_definitions")+".id = ?", *opts.ID)
	}

	var items []*ConfiguratorAttributeDefinition
	giOpts := GetItemsOptions{
		Alias:      TableName("configurator_attribute_definitions"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "ConfiguratorAttributeDefinition"}
	}
	return items[0], err
}

type QueryConfiguratorAttributeDefinitionsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*ConfiguratorAttributeDefinitionSortType
	Filter *ConfiguratorAttributeDefinitionFilterType
}

func (r *GeneratedQueryResolver) ConfiguratorAttributeDefinitions(ctx context.Context, offset *int, limit *int, q *string, sort []*ConfiguratorAttributeDefinitionSortType, filter *ConfiguratorAttributeDefinitionFilterType) (*ConfiguratorAttributeDefinitionResultType, error) {
	opts := QueryConfiguratorAttributeDefinitionsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryConfiguratorAttributeDefinitions(ctx, r.GeneratedResolver, opts)
}
func QueryConfiguratorAttributeDefinitionsHandler(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorAttributeDefinitionsHandlerOptions) (*ConfiguratorAttributeDefinitionResultType, error) {
	query := ConfiguratorAttributeDefinitionQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &ConfiguratorAttributeDefinitionResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedConfiguratorAttributeDefinitionResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedConfiguratorAttributeDefinitionResultTypeResolver) Items(ctx context.Context, obj *ConfiguratorAttributeDefinitionResultType) (items []*ConfiguratorAttributeDefinition, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("configurator_attribute_definitions"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, giOpts, &items)

	return
}

func (r *GeneratedConfiguratorAttributeDefinitionResultTypeResolver) Count(ctx context.Context, obj *ConfiguratorAttributeDefinitionResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &ConfiguratorAttributeDefinition{})
}

type GeneratedConfiguratorAttributeDefinitionResolver struct{ *GeneratedResolver }

func (r *GeneratedConfiguratorAttributeDefinitionResolver) ItemDefinitions(ctx context.Context, obj *ConfiguratorAttributeDefinition) (res []*ConfiguratorItemDefinition, err error) {
	return r.Handlers.ConfiguratorAttributeDefinitionItemDefinitions(ctx, r, obj)
}
func ConfiguratorAttributeDefinitionItemDefinitionsHandler(ctx context.Context, r *GeneratedConfiguratorAttributeDefinitionResolver, obj *ConfiguratorAttributeDefinition) (res []*ConfiguratorItemDefinition, err error) {

	items := []*ConfiguratorItemDefinition{}
	err = r.DB.Query().Model(obj).Related(&items, "ItemDefinitions").Error
	res = items

	return
}

func (r *GeneratedConfiguratorAttributeDefinitionResolver) ItemDefinitionsIds(ctx context.Context, obj *ConfiguratorAttributeDefinition) (ids []string, err error) {
	ids = []string{}

	items := []*ConfiguratorItemDefinition{}
	err = r.DB.Query().Model(obj).Select(TableName("configurator_item_definitions")+".id").Related(&items, "ItemDefinitions").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

func (r *GeneratedConfiguratorAttributeDefinitionResolver) Attributes(ctx context.Context, obj *ConfiguratorAttributeDefinition) (res []*ConfiguratorAttribute, err error) {
	return r.Handlers.ConfiguratorAttributeDefinitionAttributes(ctx, r, obj)
}
func ConfiguratorAttributeDefinitionAttributesHandler(ctx context.Context, r *GeneratedConfiguratorAttributeDefinitionResolver, obj *ConfiguratorAttributeDefinition) (res []*ConfiguratorAttribute, err error) {

	items := []*ConfiguratorAttribute{}
	err = r.DB.Query().Model(obj).Related(&items, "Attributes").Error
	res = items

	return
}

func (r *GeneratedConfiguratorAttributeDefinitionResolver) AttributesIds(ctx context.Context, obj *ConfiguratorAttributeDefinition) (ids []string, err error) {
	ids = []string{}

	items := []*ConfiguratorAttribute{}
	err = r.DB.Query().Model(obj).Select(TableName("configurator_attributes")+".id").Related(&items, "Attributes").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

type QueryConfiguratorSlotDefinitionHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *ConfiguratorSlotDefinitionFilterType
}

func (r *GeneratedQueryResolver) ConfiguratorSlotDefinition(ctx context.Context, id *string, q *string, filter *ConfiguratorSlotDefinitionFilterType) (*ConfiguratorSlotDefinition, error) {
	opts := QueryConfiguratorSlotDefinitionHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryConfiguratorSlotDefinition(ctx, r.GeneratedResolver, opts)
}
func QueryConfiguratorSlotDefinitionHandler(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorSlotDefinitionHandlerOptions) (*ConfiguratorSlotDefinition, error) {
	query := ConfiguratorSlotDefinitionQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &ConfiguratorSlotDefinitionResultType{
		EntityResultType: EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("configurator_slot_definitions")+".id = ?", *opts.ID)
	}

	var items []*ConfiguratorSlotDefinition
	giOpts := GetItemsOptions{
		Alias:      TableName("configurator_slot_definitions"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "ConfiguratorSlotDefinition"}
	}
	return items[0], err
}

type QueryConfiguratorSlotDefinitionsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*ConfiguratorSlotDefinitionSortType
	Filter *ConfiguratorSlotDefinitionFilterType
}

func (r *GeneratedQueryResolver) ConfiguratorSlotDefinitions(ctx context.Context, offset *int, limit *int, q *string, sort []*ConfiguratorSlotDefinitionSortType, filter *ConfiguratorSlotDefinitionFilterType) (*ConfiguratorSlotDefinitionResultType, error) {
	opts := QueryConfiguratorSlotDefinitionsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryConfiguratorSlotDefinitions(ctx, r.GeneratedResolver, opts)
}
func QueryConfiguratorSlotDefinitionsHandler(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorSlotDefinitionsHandlerOptions) (*ConfiguratorSlotDefinitionResultType, error) {
	query := ConfiguratorSlotDefinitionQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &ConfiguratorSlotDefinitionResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedConfiguratorSlotDefinitionResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedConfiguratorSlotDefinitionResultTypeResolver) Items(ctx context.Context, obj *ConfiguratorSlotDefinitionResultType) (items []*ConfiguratorSlotDefinition, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("configurator_slot_definitions"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, giOpts, &items)

	return
}

func (r *GeneratedConfiguratorSlotDefinitionResultTypeResolver) Count(ctx context.Context, obj *ConfiguratorSlotDefinitionResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &ConfiguratorSlotDefinition{})
}

type GeneratedConfiguratorSlotDefinitionResolver struct{ *GeneratedResolver }

func (r *GeneratedConfiguratorSlotDefinitionResolver) ItemDefinition(ctx context.Context, obj *ConfiguratorSlotDefinition) (res *ConfiguratorItemDefinition, err error) {
	return r.Handlers.ConfiguratorSlotDefinitionItemDefinition(ctx, r, obj)
}
func ConfiguratorSlotDefinitionItemDefinitionHandler(ctx context.Context, r *GeneratedConfiguratorSlotDefinitionResolver, obj *ConfiguratorSlotDefinition) (res *ConfiguratorItemDefinition, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.ItemDefinitionID != nil {
		item, _err := loaders["ConfiguratorItemDefinition"].Load(ctx, dataloader.StringKey(*obj.ItemDefinitionID))()
		res, _ = item.(*ConfiguratorItemDefinition)

		err = _err
	}

	return
}

func (r *GeneratedConfiguratorSlotDefinitionResolver) Slots(ctx context.Context, obj *ConfiguratorSlotDefinition) (res []*ConfiguratorSlot, err error) {
	return r.Handlers.ConfiguratorSlotDefinitionSlots(ctx, r, obj)
}
func ConfiguratorSlotDefinitionSlotsHandler(ctx context.Context, r *GeneratedConfiguratorSlotDefinitionResolver, obj *ConfiguratorSlotDefinition) (res []*ConfiguratorSlot, err error) {

	items := []*ConfiguratorSlot{}
	err = r.DB.Query().Model(obj).Related(&items, "Slots").Error
	res = items

	return
}

func (r *GeneratedConfiguratorSlotDefinitionResolver) SlotsIds(ctx context.Context, obj *ConfiguratorSlotDefinition) (ids []string, err error) {
	ids = []string{}

	items := []*ConfiguratorSlot{}
	err = r.DB.Query().Model(obj).Select(TableName("configurator_slots")+".id").Related(&items, "Slots").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

type QueryConfiguratorItemHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *ConfiguratorItemFilterType
}

func (r *GeneratedQueryResolver) ConfiguratorItem(ctx context.Context, id *string, q *string, filter *ConfiguratorItemFilterType) (*ConfiguratorItem, error) {
	opts := QueryConfiguratorItemHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryConfiguratorItem(ctx, r.GeneratedResolver, opts)
}
func QueryConfiguratorItemHandler(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorItemHandlerOptions) (*ConfiguratorItem, error) {
	query := ConfiguratorItemQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &ConfiguratorItemResultType{
		EntityResultType: EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("configurator_items")+".id = ?", *opts.ID)
	}

	var items []*ConfiguratorItem
	giOpts := GetItemsOptions{
		Alias:      TableName("configurator_items"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "ConfiguratorItem"}
	}
	return items[0], err
}

type QueryConfiguratorItemsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*ConfiguratorItemSortType
	Filter *ConfiguratorItemFilterType
}

func (r *GeneratedQueryResolver) ConfiguratorItems(ctx context.Context, offset *int, limit *int, q *string, sort []*ConfiguratorItemSortType, filter *ConfiguratorItemFilterType) (*ConfiguratorItemResultType, error) {
	opts := QueryConfiguratorItemsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryConfiguratorItems(ctx, r.GeneratedResolver, opts)
}
func QueryConfiguratorItemsHandler(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorItemsHandlerOptions) (*ConfiguratorItemResultType, error) {
	query := ConfiguratorItemQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &ConfiguratorItemResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedConfiguratorItemResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedConfiguratorItemResultTypeResolver) Items(ctx context.Context, obj *ConfiguratorItemResultType) (items []*ConfiguratorItem, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("configurator_items"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, giOpts, &items)

	return
}

func (r *GeneratedConfiguratorItemResultTypeResolver) Count(ctx context.Context, obj *ConfiguratorItemResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &ConfiguratorItem{})
}

type GeneratedConfiguratorItemResolver struct{ *GeneratedResolver }

func (r *GeneratedConfiguratorItemResolver) Definition(ctx context.Context, obj *ConfiguratorItem) (res *ConfiguratorItemDefinition, err error) {
	return r.Handlers.ConfiguratorItemDefinition(ctx, r, obj)
}
func ConfiguratorItemDefinitionHandler(ctx context.Context, r *GeneratedConfiguratorItemResolver, obj *ConfiguratorItem) (res *ConfiguratorItemDefinition, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.DefinitionID != nil {
		item, _err := loaders["ConfiguratorItemDefinition"].Load(ctx, dataloader.StringKey(*obj.DefinitionID))()
		res, _ = item.(*ConfiguratorItemDefinition)

		err = _err
	}

	return
}

func (r *GeneratedConfiguratorItemResolver) Attributes(ctx context.Context, obj *ConfiguratorItem) (res []*ConfiguratorAttribute, err error) {
	return r.Handlers.ConfiguratorItemAttributes(ctx, r, obj)
}
func ConfiguratorItemAttributesHandler(ctx context.Context, r *GeneratedConfiguratorItemResolver, obj *ConfiguratorItem) (res []*ConfiguratorAttribute, err error) {

	items := []*ConfiguratorAttribute{}
	err = r.DB.Query().Model(obj).Related(&items, "Attributes").Error
	res = items

	return
}

func (r *GeneratedConfiguratorItemResolver) AttributesIds(ctx context.Context, obj *ConfiguratorItem) (ids []string, err error) {
	ids = []string{}

	items := []*ConfiguratorAttribute{}
	err = r.DB.Query().Model(obj).Select(TableName("configurator_attributes")+".id").Related(&items, "Attributes").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

func (r *GeneratedConfiguratorItemResolver) Slots(ctx context.Context, obj *ConfiguratorItem) (res []*ConfiguratorSlot, err error) {
	return r.Handlers.ConfiguratorItemSlots(ctx, r, obj)
}
func ConfiguratorItemSlotsHandler(ctx context.Context, r *GeneratedConfiguratorItemResolver, obj *ConfiguratorItem) (res []*ConfiguratorSlot, err error) {

	items := []*ConfiguratorSlot{}
	err = r.DB.Query().Model(obj).Related(&items, "Slots").Error
	res = items

	return
}

func (r *GeneratedConfiguratorItemResolver) SlotsIds(ctx context.Context, obj *ConfiguratorItem) (ids []string, err error) {
	ids = []string{}

	items := []*ConfiguratorSlot{}
	err = r.DB.Query().Model(obj).Select(TableName("configurator_slots")+".id").Related(&items, "Slots").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

func (r *GeneratedConfiguratorItemResolver) ParentSlot(ctx context.Context, obj *ConfiguratorItem) (res *ConfiguratorSlot, err error) {
	return r.Handlers.ConfiguratorItemParentSlot(ctx, r, obj)
}
func ConfiguratorItemParentSlotHandler(ctx context.Context, r *GeneratedConfiguratorItemResolver, obj *ConfiguratorItem) (res *ConfiguratorSlot, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.ParentSlotID != nil {
		item, _err := loaders["ConfiguratorSlot"].Load(ctx, dataloader.StringKey(*obj.ParentSlotID))()
		res, _ = item.(*ConfiguratorSlot)

		err = _err
	}

	return
}

type QueryConfiguratorAttributeHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *ConfiguratorAttributeFilterType
}

func (r *GeneratedQueryResolver) ConfiguratorAttribute(ctx context.Context, id *string, q *string, filter *ConfiguratorAttributeFilterType) (*ConfiguratorAttribute, error) {
	opts := QueryConfiguratorAttributeHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryConfiguratorAttribute(ctx, r.GeneratedResolver, opts)
}
func QueryConfiguratorAttributeHandler(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorAttributeHandlerOptions) (*ConfiguratorAttribute, error) {
	query := ConfiguratorAttributeQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &ConfiguratorAttributeResultType{
		EntityResultType: EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("configurator_attributes")+".id = ?", *opts.ID)
	}

	var items []*ConfiguratorAttribute
	giOpts := GetItemsOptions{
		Alias:      TableName("configurator_attributes"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "ConfiguratorAttribute"}
	}
	return items[0], err
}

type QueryConfiguratorAttributesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*ConfiguratorAttributeSortType
	Filter *ConfiguratorAttributeFilterType
}

func (r *GeneratedQueryResolver) ConfiguratorAttributes(ctx context.Context, offset *int, limit *int, q *string, sort []*ConfiguratorAttributeSortType, filter *ConfiguratorAttributeFilterType) (*ConfiguratorAttributeResultType, error) {
	opts := QueryConfiguratorAttributesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryConfiguratorAttributes(ctx, r.GeneratedResolver, opts)
}
func QueryConfiguratorAttributesHandler(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorAttributesHandlerOptions) (*ConfiguratorAttributeResultType, error) {
	query := ConfiguratorAttributeQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &ConfiguratorAttributeResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedConfiguratorAttributeResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedConfiguratorAttributeResultTypeResolver) Items(ctx context.Context, obj *ConfiguratorAttributeResultType) (items []*ConfiguratorAttribute, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("configurator_attributes"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, giOpts, &items)

	return
}

func (r *GeneratedConfiguratorAttributeResultTypeResolver) Count(ctx context.Context, obj *ConfiguratorAttributeResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &ConfiguratorAttribute{})
}

type GeneratedConfiguratorAttributeResolver struct{ *GeneratedResolver }

func (r *GeneratedConfiguratorAttributeResolver) Definition(ctx context.Context, obj *ConfiguratorAttribute) (res *ConfiguratorAttributeDefinition, err error) {
	return r.Handlers.ConfiguratorAttributeDefinition(ctx, r, obj)
}
func ConfiguratorAttributeDefinitionHandler(ctx context.Context, r *GeneratedConfiguratorAttributeResolver, obj *ConfiguratorAttribute) (res *ConfiguratorAttributeDefinition, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.DefinitionID != nil {
		item, _err := loaders["ConfiguratorAttributeDefinition"].Load(ctx, dataloader.StringKey(*obj.DefinitionID))()
		res, _ = item.(*ConfiguratorAttributeDefinition)

		err = _err
	}

	return
}

func (r *GeneratedConfiguratorAttributeResolver) Item(ctx context.Context, obj *ConfiguratorAttribute) (res *ConfiguratorItem, err error) {
	return r.Handlers.ConfiguratorAttributeItem(ctx, r, obj)
}
func ConfiguratorAttributeItemHandler(ctx context.Context, r *GeneratedConfiguratorAttributeResolver, obj *ConfiguratorAttribute) (res *ConfiguratorItem, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.ItemID != nil {
		item, _err := loaders["ConfiguratorItem"].Load(ctx, dataloader.StringKey(*obj.ItemID))()
		res, _ = item.(*ConfiguratorItem)

		err = _err
	}

	return
}

type QueryConfiguratorSlotHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *ConfiguratorSlotFilterType
}

func (r *GeneratedQueryResolver) ConfiguratorSlot(ctx context.Context, id *string, q *string, filter *ConfiguratorSlotFilterType) (*ConfiguratorSlot, error) {
	opts := QueryConfiguratorSlotHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryConfiguratorSlot(ctx, r.GeneratedResolver, opts)
}
func QueryConfiguratorSlotHandler(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorSlotHandlerOptions) (*ConfiguratorSlot, error) {
	query := ConfiguratorSlotQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &ConfiguratorSlotResultType{
		EntityResultType: EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("configurator_slots")+".id = ?", *opts.ID)
	}

	var items []*ConfiguratorSlot
	giOpts := GetItemsOptions{
		Alias:      TableName("configurator_slots"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "ConfiguratorSlot"}
	}
	return items[0], err
}

type QueryConfiguratorSlotsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*ConfiguratorSlotSortType
	Filter *ConfiguratorSlotFilterType
}

func (r *GeneratedQueryResolver) ConfiguratorSlots(ctx context.Context, offset *int, limit *int, q *string, sort []*ConfiguratorSlotSortType, filter *ConfiguratorSlotFilterType) (*ConfiguratorSlotResultType, error) {
	opts := QueryConfiguratorSlotsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryConfiguratorSlots(ctx, r.GeneratedResolver, opts)
}
func QueryConfiguratorSlotsHandler(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorSlotsHandlerOptions) (*ConfiguratorSlotResultType, error) {
	query := ConfiguratorSlotQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &ConfiguratorSlotResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedConfiguratorSlotResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedConfiguratorSlotResultTypeResolver) Items(ctx context.Context, obj *ConfiguratorSlotResultType) (items []*ConfiguratorSlot, err error) {
	giOpts := GetItemsOptions{
		Alias:      TableName("configurator_slots"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, giOpts, &items)

	return
}

func (r *GeneratedConfiguratorSlotResultTypeResolver) Count(ctx context.Context, obj *ConfiguratorSlotResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &ConfiguratorSlot{})
}

type GeneratedConfiguratorSlotResolver struct{ *GeneratedResolver }

func (r *GeneratedConfiguratorSlotResolver) Items(ctx context.Context, obj *ConfiguratorSlot) (res []*ConfiguratorItem, err error) {
	return r.Handlers.ConfiguratorSlotItems(ctx, r, obj)
}
func ConfiguratorSlotItemsHandler(ctx context.Context, r *GeneratedConfiguratorSlotResolver, obj *ConfiguratorSlot) (res []*ConfiguratorItem, err error) {

	items := []*ConfiguratorItem{}
	err = r.DB.Query().Model(obj).Related(&items, "Items").Error
	res = items

	return
}

func (r *GeneratedConfiguratorSlotResolver) ItemsIds(ctx context.Context, obj *ConfiguratorSlot) (ids []string, err error) {
	ids = []string{}

	items := []*ConfiguratorItem{}
	err = r.DB.Query().Model(obj).Select(TableName("configurator_items")+".id").Related(&items, "Items").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

func (r *GeneratedConfiguratorSlotResolver) Definition(ctx context.Context, obj *ConfiguratorSlot) (res *ConfiguratorSlotDefinition, err error) {
	return r.Handlers.ConfiguratorSlotDefinition(ctx, r, obj)
}
func ConfiguratorSlotDefinitionHandler(ctx context.Context, r *GeneratedConfiguratorSlotResolver, obj *ConfiguratorSlot) (res *ConfiguratorSlotDefinition, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.DefinitionID != nil {
		item, _err := loaders["ConfiguratorSlotDefinition"].Load(ctx, dataloader.StringKey(*obj.DefinitionID))()
		res, _ = item.(*ConfiguratorSlotDefinition)

		err = _err
	}

	return
}

func (r *GeneratedConfiguratorSlotResolver) ParentItem(ctx context.Context, obj *ConfiguratorSlot) (res *ConfiguratorItem, err error) {
	return r.Handlers.ConfiguratorSlotParentItem(ctx, r, obj)
}
func ConfiguratorSlotParentItemHandler(ctx context.Context, r *GeneratedConfiguratorSlotResolver, obj *ConfiguratorSlot) (res *ConfiguratorItem, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.ParentItemID != nil {
		item, _err := loaders["ConfiguratorItem"].Load(ctx, dataloader.StringKey(*obj.ParentItemID))()
		res, _ = item.(*ConfiguratorItem)

		err = _err
	}

	return
}