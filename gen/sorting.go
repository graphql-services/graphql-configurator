package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

func (s ConfiguratorItemDefinitionCategorySortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("configurator_item_definition_categories"), sorts, joins)
}
func (s ConfiguratorItemDefinitionCategorySortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Code != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("code"), Direction: s.Code.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CodeMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("code") + ")", Direction: s.CodeMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CodeMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("code") + ")", Direction: s.CodeMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Type != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("type"), Direction: s.Type.String()}
		*sorts = append(*sorts, sort)
	}

	if s.TypeMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("type") + ")", Direction: s.TypeMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.TypeMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("type") + ")", Direction: s.TypeMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Definitions != nil {
		_alias := alias + "_definitions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_item_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("categoryId")+" = "+dialect.Quote(alias)+".id")
		err := s.Definitions.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s ConfiguratorItemDefinitionSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("configurator_item_definitions"), sorts, joins)
}
func (s ConfiguratorItemDefinitionSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Code != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("code"), Direction: s.Code.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CodeMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("code") + ")", Direction: s.CodeMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CodeMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("code") + ")", Direction: s.CodeMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CategoryID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("categoryId"), Direction: s.CategoryID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CategoryIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("categoryId") + ")", Direction: s.CategoryIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CategoryIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("categoryId") + ")", Direction: s.CategoryIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Attributes != nil {
		_alias := alias + "_attributes"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorAttributeDefinition_definitions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("definition_id")+" LEFT JOIN "+dialect.Quote(TableName("configurator_attribute_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("attribute_id")+" = "+dialect.Quote(_alias)+".id")
		err := s.Attributes.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Slots != nil {
		_alias := alias + "_slots"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_slot_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("definitionId")+" = "+dialect.Quote(alias)+".id")
		err := s.Slots.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Items != nil {
		_alias := alias + "_items"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_items"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("definitionId")+" = "+dialect.Quote(alias)+".id")
		err := s.Items.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.AllowedInSlots != nil {
		_alias := alias + "_allowedInSlots"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorSlotDefinition_allowedItemDefinitions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("allowedItemDefinition_id")+" LEFT JOIN "+dialect.Quote(TableName("configurator_slot_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("allowedInSlot_id")+" = "+dialect.Quote(_alias)+".id")
		err := s.AllowedInSlots.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Category != nil {
		_alias := alias + "_category"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_item_definition_categories"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("categoryId"))
		err := s.Category.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s ConfiguratorAttributeDefinitionSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("configurator_attribute_definitions"), sorts, joins)
}
func (s ConfiguratorAttributeDefinitionSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Primary != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("primary"), Direction: s.Primary.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PrimaryMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("primary") + ")", Direction: s.PrimaryMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PrimaryMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("primary") + ")", Direction: s.PrimaryMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Definitions != nil {
		_alias := alias + "_definitions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorAttributeDefinition_definitions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("attribute_id")+" LEFT JOIN "+dialect.Quote(TableName("configurator_item_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("definition_id")+" = "+dialect.Quote(_alias)+".id")
		err := s.Definitions.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Attributes != nil {
		_alias := alias + "_attributes"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_attributes"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("definitionId")+" = "+dialect.Quote(alias)+".id")
		err := s.Attributes.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s ConfiguratorSlotDefinitionSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("configurator_slot_definitions"), sorts, joins)
}
func (s ConfiguratorSlotDefinitionSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.MinCount != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("minCount"), Direction: s.MinCount.String()}
		*sorts = append(*sorts, sort)
	}

	if s.MinCountMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("minCount") + ")", Direction: s.MinCountMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.MinCountMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("minCount") + ")", Direction: s.MinCountMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.MinCountAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("minCount") + ")", Direction: s.MinCountAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.MaxCount != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("maxCount"), Direction: s.MaxCount.String()}
		*sorts = append(*sorts, sort)
	}

	if s.MaxCountMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("maxCount") + ")", Direction: s.MaxCountMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.MaxCountMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("maxCount") + ")", Direction: s.MaxCountMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.MaxCountAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("maxCount") + ")", Direction: s.MaxCountAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DefaultCount != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("defaultCount"), Direction: s.DefaultCount.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DefaultCountMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("defaultCount") + ")", Direction: s.DefaultCountMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DefaultCountMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("defaultCount") + ")", Direction: s.DefaultCountMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DefaultCountAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("defaultCount") + ")", Direction: s.DefaultCountAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DefinitionID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("definitionId"), Direction: s.DefinitionID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DefinitionIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("definitionId") + ")", Direction: s.DefinitionIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DefinitionIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("definitionId") + ")", Direction: s.DefinitionIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Definition != nil {
		_alias := alias + "_definition"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_item_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("definitionId"))
		err := s.Definition.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Slots != nil {
		_alias := alias + "_slots"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_slots"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("definitionId")+" = "+dialect.Quote(alias)+".id")
		err := s.Slots.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.AllowedItemDefinitions != nil {
		_alias := alias + "_allowedItemDefinitions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorSlotDefinition_allowedItemDefinitions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("allowedInSlot_id")+" LEFT JOIN "+dialect.Quote(TableName("configurator_item_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("allowedItemDefinition_id")+" = "+dialect.Quote(_alias)+".id")
		err := s.AllowedItemDefinitions.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s ConfiguratorItemSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("configurator_items"), sorts, joins)
}
func (s ConfiguratorItemSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Code != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("code"), Direction: s.Code.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CodeMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("code") + ")", Direction: s.CodeMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CodeMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("code") + ")", Direction: s.CodeMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.StockItemID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("stockItemId"), Direction: s.StockItemID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.StockItemIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("stockItemId") + ")", Direction: s.StockItemIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.StockItemIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("stockItemId") + ")", Direction: s.StockItemIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.RawData != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("rawData"), Direction: s.RawData.String()}
		*sorts = append(*sorts, sort)
	}

	if s.RawDataMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("rawData") + ")", Direction: s.RawDataMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.RawDataMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("rawData") + ")", Direction: s.RawDataMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DefinitionID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("definitionId"), Direction: s.DefinitionID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DefinitionIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("definitionId") + ")", Direction: s.DefinitionIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DefinitionIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("definitionId") + ")", Direction: s.DefinitionIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Definition != nil {
		_alias := alias + "_definition"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_item_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("definitionId"))
		err := s.Definition.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Attributes != nil {
		_alias := alias + "_attributes"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_attributes"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("itemId")+" = "+dialect.Quote(alias)+".id")
		err := s.Attributes.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Slots != nil {
		_alias := alias + "_slots"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_slots"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("parentItemId")+" = "+dialect.Quote(alias)+".id")
		err := s.Slots.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.ParentSlots != nil {
		_alias := alias + "_parentSlots"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_slots"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("itemId")+" = "+dialect.Quote(alias)+".id")
		err := s.ParentSlots.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s ConfiguratorAttributeSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("configurator_attributes"), sorts, joins)
}
func (s ConfiguratorAttributeSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.StringValue != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("stringValue"), Direction: s.StringValue.String()}
		*sorts = append(*sorts, sort)
	}

	if s.StringValueMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("stringValue") + ")", Direction: s.StringValueMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.StringValueMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("stringValue") + ")", Direction: s.StringValueMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.FloatValue != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("floatValue"), Direction: s.FloatValue.String()}
		*sorts = append(*sorts, sort)
	}

	if s.FloatValueMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("floatValue") + ")", Direction: s.FloatValueMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.FloatValueMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("floatValue") + ")", Direction: s.FloatValueMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.FloatValueAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("floatValue") + ")", Direction: s.FloatValueAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IntValue != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("intValue"), Direction: s.IntValue.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IntValueMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("intValue") + ")", Direction: s.IntValueMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IntValueMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("intValue") + ")", Direction: s.IntValueMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IntValueAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("intValue") + ")", Direction: s.IntValueAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DefinitionID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("definitionId"), Direction: s.DefinitionID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DefinitionIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("definitionId") + ")", Direction: s.DefinitionIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DefinitionIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("definitionId") + ")", Direction: s.DefinitionIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ItemID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("itemId"), Direction: s.ItemID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ItemIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("itemId") + ")", Direction: s.ItemIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ItemIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("itemId") + ")", Direction: s.ItemIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Definition != nil {
		_alias := alias + "_definition"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_attribute_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("definitionId"))
		err := s.Definition.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Item != nil {
		_alias := alias + "_item"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_items"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("itemId"))
		err := s.Item.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s ConfiguratorSlotSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("configurator_slots"), sorts, joins)
}
func (s ConfiguratorSlotSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Count != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("count"), Direction: s.Count.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CountMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("count") + ")", Direction: s.CountMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CountMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("count") + ")", Direction: s.CountMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CountAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("count") + ")", Direction: s.CountAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ItemID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("itemId"), Direction: s.ItemID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ItemIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("itemId") + ")", Direction: s.ItemIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ItemIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("itemId") + ")", Direction: s.ItemIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DefinitionID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("definitionId"), Direction: s.DefinitionID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DefinitionIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("definitionId") + ")", Direction: s.DefinitionIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DefinitionIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("definitionId") + ")", Direction: s.DefinitionIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ParentItemID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("parentItemId"), Direction: s.ParentItemID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ParentItemIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("parentItemId") + ")", Direction: s.ParentItemIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ParentItemIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("parentItemId") + ")", Direction: s.ParentItemIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Item != nil {
		_alias := alias + "_item"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_items"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("itemId"))
		err := s.Item.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Definition != nil {
		_alias := alias + "_definition"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_slot_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("definitionId"))
		err := s.Definition.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.ParentItem != nil {
		_alias := alias + "_parentItem"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_items"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("parentItemId"))
		err := s.ParentItem.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
