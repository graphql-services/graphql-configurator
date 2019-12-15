package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

func (s ConfiguratorItemDefinitionSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("configurator_item_definitions"), sorts, joins)
}
func (s ConfiguratorItemDefinitionSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("id")+" "+s.ID.String())
	}

	if s.Name != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("name")+" "+s.Name.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdAt")+" "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdBy")+" "+s.CreatedBy.String())
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

	return nil
}

func (s ConfiguratorAttributeDefinitionSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("configurator_attribute_definitions"), sorts, joins)
}
func (s ConfiguratorAttributeDefinitionSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("id")+" "+s.ID.String())
	}

	if s.Name != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("name")+" "+s.Name.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdAt")+" "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdBy")+" "+s.CreatedBy.String())
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

func (s ConfiguratorSlotDefinitionSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("configurator_slot_definitions"), sorts, joins)
}
func (s ConfiguratorSlotDefinitionSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("id")+" "+s.ID.String())
	}

	if s.Name != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("name")+" "+s.Name.String())
	}

	if s.MinCount != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("minCount")+" "+s.MinCount.String())
	}

	if s.MaxCount != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("maxCount")+" "+s.MaxCount.String())
	}

	if s.DefinitionID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("definitionId")+" "+s.DefinitionID.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdAt")+" "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdBy")+" "+s.CreatedBy.String())
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

	return nil
}

func (s ConfiguratorItemSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("configurator_items"), sorts, joins)
}
func (s ConfiguratorItemSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("id")+" "+s.ID.String())
	}

	if s.Code != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("code")+" "+s.Code.String())
	}

	if s.Name != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("name")+" "+s.Name.String())
	}

	if s.StockItemID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("stockItemID")+" "+s.StockItemID.String())
	}

	if s.TemplateID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("templateId")+" "+s.TemplateID.String())
	}

	if s.DefinitionID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("definitionId")+" "+s.DefinitionID.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdAt")+" "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdBy")+" "+s.CreatedBy.String())
	}

	if s.Template != nil {
		_alias := alias + "_template"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_items"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("templateId"))
		err := s.Template.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.TemplatedChilds != nil {
		_alias := alias + "_templatedChilds"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_items"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("templateId")+" = "+dialect.Quote(alias)+".id")
		err := s.TemplatedChilds.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
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

func (s ConfiguratorAttributeSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("configurator_attributes"), sorts, joins)
}
func (s ConfiguratorAttributeSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("id")+" "+s.ID.String())
	}

	if s.StringValue != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("stringValue")+" "+s.StringValue.String())
	}

	if s.FloatValue != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("floatValue")+" "+s.FloatValue.String())
	}

	if s.IntValue != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("intValue")+" "+s.IntValue.String())
	}

	if s.DefinitionID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("definitionId")+" "+s.DefinitionID.String())
	}

	if s.ItemID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("itemId")+" "+s.ItemID.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdAt")+" "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdBy")+" "+s.CreatedBy.String())
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

func (s ConfiguratorSlotSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("configurator_slots"), sorts, joins)
}
func (s ConfiguratorSlotSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("id")+" "+s.ID.String())
	}

	if s.ItemID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("itemId")+" "+s.ItemID.String())
	}

	if s.DefinitionID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("definitionId")+" "+s.DefinitionID.String())
	}

	if s.ParentItemID != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("parentItemId")+" "+s.ParentItemID.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdAt")+" "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+dialect.Quote("createdBy")+" "+s.CreatedBy.String())
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
