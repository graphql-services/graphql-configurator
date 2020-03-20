package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func (f *ConfiguratorItemDefinitionCategoryFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *ConfiguratorItemDefinitionCategoryFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("configuratorItemDefinitionCategories"), wheres, values, joins)
}
func (f *ConfiguratorItemDefinitionCategoryFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Definitions != nil {
		_alias := alias + "_definitions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorItemDefinitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("categoryId")+" = "+dialect.Quote(alias)+".id")
		err := f.Definitions.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *ConfiguratorItemDefinitionCategoryFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Code != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" = ?")
		values = append(values, f.Code)
	}

	if f.CodeNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" != ?")
		values = append(values, f.CodeNe)
	}

	if f.CodeGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" > ?")
		values = append(values, f.CodeGt)
	}

	if f.CodeLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" < ?")
		values = append(values, f.CodeLt)
	}

	if f.CodeGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" >= ?")
		values = append(values, f.CodeGte)
	}

	if f.CodeLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" <= ?")
		values = append(values, f.CodeLte)
	}

	if f.CodeIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" IN (?)")
		values = append(values, f.CodeIn)
	}

	if f.CodeLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.CodeLike, "?", "_", -1), "*", "%", -1))
	}

	if f.CodePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CodePrefix))
	}

	if f.CodeSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CodeSuffix))
	}

	if f.CodeNull != nil {
		if *f.CodeNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.Type != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" = ?")
		values = append(values, f.Type)
	}

	if f.TypeNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" != ?")
		values = append(values, f.TypeNe)
	}

	if f.TypeGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" > ?")
		values = append(values, f.TypeGt)
	}

	if f.TypeLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" < ?")
		values = append(values, f.TypeLt)
	}

	if f.TypeGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" >= ?")
		values = append(values, f.TypeGte)
	}

	if f.TypeLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" <= ?")
		values = append(values, f.TypeLte)
	}

	if f.TypeIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" IN (?)")
		values = append(values, f.TypeIn)
	}

	if f.TypeLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TypeLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TypePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TypePrefix))
	}

	if f.TypeSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TypeSuffix))
	}

	if f.TypeNull != nil {
		if *f.TypeNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *ConfiguratorItemDefinitionCategoryFilterType) AndWith(f2 ...*ConfiguratorItemDefinitionCategoryFilterType) *ConfiguratorItemDefinitionCategoryFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ConfiguratorItemDefinitionCategoryFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *ConfiguratorItemDefinitionCategoryFilterType) OrWith(f2 ...*ConfiguratorItemDefinitionCategoryFilterType) *ConfiguratorItemDefinitionCategoryFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ConfiguratorItemDefinitionCategoryFilterType{
		Or: append(_f2, f),
	}
}

func (f *ConfiguratorItemDefinitionFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *ConfiguratorItemDefinitionFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("configuratorItemDefinitions"), wheres, values, joins)
}
func (f *ConfiguratorItemDefinitionFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Attributes != nil {
		_alias := alias + "_attributes"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorAttributeDefinition_definitions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("definition_id")+" LEFT JOIN "+dialect.Quote(TableName("configuratorAttributeDefinitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("attribute_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.Attributes.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Slots != nil {
		_alias := alias + "_slots"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorSlotDefinitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("definitionId")+" = "+dialect.Quote(alias)+".id")
		err := f.Slots.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Items != nil {
		_alias := alias + "_items"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorItems"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("definitionId")+" = "+dialect.Quote(alias)+".id")
		err := f.Items.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.AllowedInSlots != nil {
		_alias := alias + "_allowedInSlots"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorSlotDefinition_allowedItemDefinitions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("allowedItemDefinition_id")+" LEFT JOIN "+dialect.Quote(TableName("configuratorSlotDefinitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("allowedInSlot_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.AllowedInSlots.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Category != nil {
		_alias := alias + "_category"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorItemDefinitionCategories"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("categoryId"))
		err := f.Category.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *ConfiguratorItemDefinitionFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Code != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" = ?")
		values = append(values, f.Code)
	}

	if f.CodeNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" != ?")
		values = append(values, f.CodeNe)
	}

	if f.CodeGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" > ?")
		values = append(values, f.CodeGt)
	}

	if f.CodeLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" < ?")
		values = append(values, f.CodeLt)
	}

	if f.CodeGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" >= ?")
		values = append(values, f.CodeGte)
	}

	if f.CodeLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" <= ?")
		values = append(values, f.CodeLte)
	}

	if f.CodeIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" IN (?)")
		values = append(values, f.CodeIn)
	}

	if f.CodeLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.CodeLike, "?", "_", -1), "*", "%", -1))
	}

	if f.CodePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CodePrefix))
	}

	if f.CodeSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CodeSuffix))
	}

	if f.CodeNull != nil {
		if *f.CodeNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.CategoryID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("categoryId")+" = ?")
		values = append(values, f.CategoryID)
	}

	if f.CategoryIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("categoryId")+" != ?")
		values = append(values, f.CategoryIDNe)
	}

	if f.CategoryIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("categoryId")+" > ?")
		values = append(values, f.CategoryIDGt)
	}

	if f.CategoryIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("categoryId")+" < ?")
		values = append(values, f.CategoryIDLt)
	}

	if f.CategoryIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("categoryId")+" >= ?")
		values = append(values, f.CategoryIDGte)
	}

	if f.CategoryIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("categoryId")+" <= ?")
		values = append(values, f.CategoryIDLte)
	}

	if f.CategoryIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("categoryId")+" IN (?)")
		values = append(values, f.CategoryIDIn)
	}

	if f.CategoryIDNull != nil {
		if *f.CategoryIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("categoryId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("categoryId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *ConfiguratorItemDefinitionFilterType) AndWith(f2 ...*ConfiguratorItemDefinitionFilterType) *ConfiguratorItemDefinitionFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ConfiguratorItemDefinitionFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *ConfiguratorItemDefinitionFilterType) OrWith(f2 ...*ConfiguratorItemDefinitionFilterType) *ConfiguratorItemDefinitionFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ConfiguratorItemDefinitionFilterType{
		Or: append(_f2, f),
	}
}

func (f *ConfiguratorAttributeDefinitionFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *ConfiguratorAttributeDefinitionFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("configuratorAttributeDefinitions"), wheres, values, joins)
}
func (f *ConfiguratorAttributeDefinitionFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Definitions != nil {
		_alias := alias + "_definitions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorAttributeDefinition_definitions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("attribute_id")+" LEFT JOIN "+dialect.Quote(TableName("configuratorItemDefinitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("definition_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.Definitions.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Attributes != nil {
		_alias := alias + "_attributes"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorAttributes"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("definitionId")+" = "+dialect.Quote(alias)+".id")
		err := f.Attributes.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *ConfiguratorAttributeDefinitionFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.Type != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" = ?")
		values = append(values, f.Type)
	}

	if f.TypeNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" != ?")
		values = append(values, f.TypeNe)
	}

	if f.TypeGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" > ?")
		values = append(values, f.TypeGt)
	}

	if f.TypeLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" < ?")
		values = append(values, f.TypeLt)
	}

	if f.TypeGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" >= ?")
		values = append(values, f.TypeGte)
	}

	if f.TypeLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" <= ?")
		values = append(values, f.TypeLte)
	}

	if f.TypeIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" IN (?)")
		values = append(values, f.TypeIn)
	}

	if f.TypeNull != nil {
		if *f.TypeNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *ConfiguratorAttributeDefinitionFilterType) AndWith(f2 ...*ConfiguratorAttributeDefinitionFilterType) *ConfiguratorAttributeDefinitionFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ConfiguratorAttributeDefinitionFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *ConfiguratorAttributeDefinitionFilterType) OrWith(f2 ...*ConfiguratorAttributeDefinitionFilterType) *ConfiguratorAttributeDefinitionFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ConfiguratorAttributeDefinitionFilterType{
		Or: append(_f2, f),
	}
}

func (f *ConfiguratorSlotDefinitionFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *ConfiguratorSlotDefinitionFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("configuratorSlotDefinitions"), wheres, values, joins)
}
func (f *ConfiguratorSlotDefinitionFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Definition != nil {
		_alias := alias + "_definition"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorItemDefinitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("definitionId"))
		err := f.Definition.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Slots != nil {
		_alias := alias + "_slots"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorSlots"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("definitionId")+" = "+dialect.Quote(alias)+".id")
		err := f.Slots.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.AllowedItemDefinitions != nil {
		_alias := alias + "_allowedItemDefinitions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorSlotDefinition_allowedItemDefinitions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("allowedInSlot_id")+" LEFT JOIN "+dialect.Quote(TableName("configuratorItemDefinitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("allowedItemDefinition_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.AllowedItemDefinitions.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *ConfiguratorSlotDefinitionFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.MinCount != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("minCount")+" = ?")
		values = append(values, f.MinCount)
	}

	if f.MinCountNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("minCount")+" != ?")
		values = append(values, f.MinCountNe)
	}

	if f.MinCountGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("minCount")+" > ?")
		values = append(values, f.MinCountGt)
	}

	if f.MinCountLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("minCount")+" < ?")
		values = append(values, f.MinCountLt)
	}

	if f.MinCountGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("minCount")+" >= ?")
		values = append(values, f.MinCountGte)
	}

	if f.MinCountLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("minCount")+" <= ?")
		values = append(values, f.MinCountLte)
	}

	if f.MinCountIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("minCount")+" IN (?)")
		values = append(values, f.MinCountIn)
	}

	if f.MinCountNull != nil {
		if *f.MinCountNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("minCount")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("minCount")+" IS NOT NULL")
		}
	}

	if f.MaxCount != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("maxCount")+" = ?")
		values = append(values, f.MaxCount)
	}

	if f.MaxCountNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("maxCount")+" != ?")
		values = append(values, f.MaxCountNe)
	}

	if f.MaxCountGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("maxCount")+" > ?")
		values = append(values, f.MaxCountGt)
	}

	if f.MaxCountLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("maxCount")+" < ?")
		values = append(values, f.MaxCountLt)
	}

	if f.MaxCountGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("maxCount")+" >= ?")
		values = append(values, f.MaxCountGte)
	}

	if f.MaxCountLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("maxCount")+" <= ?")
		values = append(values, f.MaxCountLte)
	}

	if f.MaxCountIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("maxCount")+" IN (?)")
		values = append(values, f.MaxCountIn)
	}

	if f.MaxCountNull != nil {
		if *f.MaxCountNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("maxCount")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("maxCount")+" IS NOT NULL")
		}
	}

	if f.DefaultCount != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("defaultCount")+" = ?")
		values = append(values, f.DefaultCount)
	}

	if f.DefaultCountNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("defaultCount")+" != ?")
		values = append(values, f.DefaultCountNe)
	}

	if f.DefaultCountGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("defaultCount")+" > ?")
		values = append(values, f.DefaultCountGt)
	}

	if f.DefaultCountLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("defaultCount")+" < ?")
		values = append(values, f.DefaultCountLt)
	}

	if f.DefaultCountGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("defaultCount")+" >= ?")
		values = append(values, f.DefaultCountGte)
	}

	if f.DefaultCountLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("defaultCount")+" <= ?")
		values = append(values, f.DefaultCountLte)
	}

	if f.DefaultCountIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("defaultCount")+" IN (?)")
		values = append(values, f.DefaultCountIn)
	}

	if f.DefaultCountNull != nil {
		if *f.DefaultCountNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("defaultCount")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("defaultCount")+" IS NOT NULL")
		}
	}

	if f.DefinitionID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" = ?")
		values = append(values, f.DefinitionID)
	}

	if f.DefinitionIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" != ?")
		values = append(values, f.DefinitionIDNe)
	}

	if f.DefinitionIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" > ?")
		values = append(values, f.DefinitionIDGt)
	}

	if f.DefinitionIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" < ?")
		values = append(values, f.DefinitionIDLt)
	}

	if f.DefinitionIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" >= ?")
		values = append(values, f.DefinitionIDGte)
	}

	if f.DefinitionIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" <= ?")
		values = append(values, f.DefinitionIDLte)
	}

	if f.DefinitionIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" IN (?)")
		values = append(values, f.DefinitionIDIn)
	}

	if f.DefinitionIDNull != nil {
		if *f.DefinitionIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *ConfiguratorSlotDefinitionFilterType) AndWith(f2 ...*ConfiguratorSlotDefinitionFilterType) *ConfiguratorSlotDefinitionFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ConfiguratorSlotDefinitionFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *ConfiguratorSlotDefinitionFilterType) OrWith(f2 ...*ConfiguratorSlotDefinitionFilterType) *ConfiguratorSlotDefinitionFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ConfiguratorSlotDefinitionFilterType{
		Or: append(_f2, f),
	}
}

func (f *ConfiguratorItemFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *ConfiguratorItemFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("configuratorItems"), wheres, values, joins)
}
func (f *ConfiguratorItemFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Definition != nil {
		_alias := alias + "_definition"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorItemDefinitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("definitionId"))
		err := f.Definition.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Attributes != nil {
		_alias := alias + "_attributes"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorAttributes"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("itemId")+" = "+dialect.Quote(alias)+".id")
		err := f.Attributes.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Slots != nil {
		_alias := alias + "_slots"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorSlots"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("parentItemId")+" = "+dialect.Quote(alias)+".id")
		err := f.Slots.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.ParentSlots != nil {
		_alias := alias + "_parentSlots"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorSlots"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("itemId")+" = "+dialect.Quote(alias)+".id")
		err := f.ParentSlots.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *ConfiguratorItemFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Code != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" = ?")
		values = append(values, f.Code)
	}

	if f.CodeNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" != ?")
		values = append(values, f.CodeNe)
	}

	if f.CodeGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" > ?")
		values = append(values, f.CodeGt)
	}

	if f.CodeLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" < ?")
		values = append(values, f.CodeLt)
	}

	if f.CodeGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" >= ?")
		values = append(values, f.CodeGte)
	}

	if f.CodeLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" <= ?")
		values = append(values, f.CodeLte)
	}

	if f.CodeIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" IN (?)")
		values = append(values, f.CodeIn)
	}

	if f.CodeLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.CodeLike, "?", "_", -1), "*", "%", -1))
	}

	if f.CodePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CodePrefix))
	}

	if f.CodeSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CodeSuffix))
	}

	if f.CodeNull != nil {
		if *f.CodeNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.StockItemID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stockItemId")+" = ?")
		values = append(values, f.StockItemID)
	}

	if f.StockItemIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stockItemId")+" != ?")
		values = append(values, f.StockItemIDNe)
	}

	if f.StockItemIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stockItemId")+" > ?")
		values = append(values, f.StockItemIDGt)
	}

	if f.StockItemIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stockItemId")+" < ?")
		values = append(values, f.StockItemIDLt)
	}

	if f.StockItemIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stockItemId")+" >= ?")
		values = append(values, f.StockItemIDGte)
	}

	if f.StockItemIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stockItemId")+" <= ?")
		values = append(values, f.StockItemIDLte)
	}

	if f.StockItemIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stockItemId")+" IN (?)")
		values = append(values, f.StockItemIDIn)
	}

	if f.StockItemIDNull != nil {
		if *f.StockItemIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("stockItemId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("stockItemId")+" IS NOT NULL")
		}
	}

	if f.RawData != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("rawData")+" = ?")
		values = append(values, f.RawData)
	}

	if f.RawDataNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("rawData")+" != ?")
		values = append(values, f.RawDataNe)
	}

	if f.RawDataGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("rawData")+" > ?")
		values = append(values, f.RawDataGt)
	}

	if f.RawDataLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("rawData")+" < ?")
		values = append(values, f.RawDataLt)
	}

	if f.RawDataGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("rawData")+" >= ?")
		values = append(values, f.RawDataGte)
	}

	if f.RawDataLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("rawData")+" <= ?")
		values = append(values, f.RawDataLte)
	}

	if f.RawDataIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("rawData")+" IN (?)")
		values = append(values, f.RawDataIn)
	}

	if f.RawDataLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("rawData")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.RawDataLike, "?", "_", -1), "*", "%", -1))
	}

	if f.RawDataPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("rawData")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.RawDataPrefix))
	}

	if f.RawDataSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("rawData")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.RawDataSuffix))
	}

	if f.RawDataNull != nil {
		if *f.RawDataNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("rawData")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("rawData")+" IS NOT NULL")
		}
	}

	if f.DefinitionID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" = ?")
		values = append(values, f.DefinitionID)
	}

	if f.DefinitionIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" != ?")
		values = append(values, f.DefinitionIDNe)
	}

	if f.DefinitionIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" > ?")
		values = append(values, f.DefinitionIDGt)
	}

	if f.DefinitionIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" < ?")
		values = append(values, f.DefinitionIDLt)
	}

	if f.DefinitionIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" >= ?")
		values = append(values, f.DefinitionIDGte)
	}

	if f.DefinitionIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" <= ?")
		values = append(values, f.DefinitionIDLte)
	}

	if f.DefinitionIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" IN (?)")
		values = append(values, f.DefinitionIDIn)
	}

	if f.DefinitionIDNull != nil {
		if *f.DefinitionIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *ConfiguratorItemFilterType) AndWith(f2 ...*ConfiguratorItemFilterType) *ConfiguratorItemFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ConfiguratorItemFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *ConfiguratorItemFilterType) OrWith(f2 ...*ConfiguratorItemFilterType) *ConfiguratorItemFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ConfiguratorItemFilterType{
		Or: append(_f2, f),
	}
}

func (f *ConfiguratorAttributeFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *ConfiguratorAttributeFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("configuratorAttributes"), wheres, values, joins)
}
func (f *ConfiguratorAttributeFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Definition != nil {
		_alias := alias + "_definition"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorAttributeDefinitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("definitionId"))
		err := f.Definition.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Item != nil {
		_alias := alias + "_item"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorItems"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("itemId"))
		err := f.Item.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *ConfiguratorAttributeFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.StringValue != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stringValue")+" = ?")
		values = append(values, f.StringValue)
	}

	if f.StringValueNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stringValue")+" != ?")
		values = append(values, f.StringValueNe)
	}

	if f.StringValueGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stringValue")+" > ?")
		values = append(values, f.StringValueGt)
	}

	if f.StringValueLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stringValue")+" < ?")
		values = append(values, f.StringValueLt)
	}

	if f.StringValueGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stringValue")+" >= ?")
		values = append(values, f.StringValueGte)
	}

	if f.StringValueLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stringValue")+" <= ?")
		values = append(values, f.StringValueLte)
	}

	if f.StringValueIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stringValue")+" IN (?)")
		values = append(values, f.StringValueIn)
	}

	if f.StringValueLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stringValue")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.StringValueLike, "?", "_", -1), "*", "%", -1))
	}

	if f.StringValuePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stringValue")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.StringValuePrefix))
	}

	if f.StringValueSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stringValue")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.StringValueSuffix))
	}

	if f.StringValueNull != nil {
		if *f.StringValueNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("stringValue")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("stringValue")+" IS NOT NULL")
		}
	}

	if f.FloatValue != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("floatValue")+" = ?")
		values = append(values, f.FloatValue)
	}

	if f.FloatValueNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("floatValue")+" != ?")
		values = append(values, f.FloatValueNe)
	}

	if f.FloatValueGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("floatValue")+" > ?")
		values = append(values, f.FloatValueGt)
	}

	if f.FloatValueLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("floatValue")+" < ?")
		values = append(values, f.FloatValueLt)
	}

	if f.FloatValueGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("floatValue")+" >= ?")
		values = append(values, f.FloatValueGte)
	}

	if f.FloatValueLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("floatValue")+" <= ?")
		values = append(values, f.FloatValueLte)
	}

	if f.FloatValueIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("floatValue")+" IN (?)")
		values = append(values, f.FloatValueIn)
	}

	if f.FloatValueNull != nil {
		if *f.FloatValueNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("floatValue")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("floatValue")+" IS NOT NULL")
		}
	}

	if f.IntValue != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("intValue")+" = ?")
		values = append(values, f.IntValue)
	}

	if f.IntValueNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("intValue")+" != ?")
		values = append(values, f.IntValueNe)
	}

	if f.IntValueGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("intValue")+" > ?")
		values = append(values, f.IntValueGt)
	}

	if f.IntValueLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("intValue")+" < ?")
		values = append(values, f.IntValueLt)
	}

	if f.IntValueGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("intValue")+" >= ?")
		values = append(values, f.IntValueGte)
	}

	if f.IntValueLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("intValue")+" <= ?")
		values = append(values, f.IntValueLte)
	}

	if f.IntValueIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("intValue")+" IN (?)")
		values = append(values, f.IntValueIn)
	}

	if f.IntValueNull != nil {
		if *f.IntValueNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("intValue")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("intValue")+" IS NOT NULL")
		}
	}

	if f.DefinitionID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" = ?")
		values = append(values, f.DefinitionID)
	}

	if f.DefinitionIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" != ?")
		values = append(values, f.DefinitionIDNe)
	}

	if f.DefinitionIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" > ?")
		values = append(values, f.DefinitionIDGt)
	}

	if f.DefinitionIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" < ?")
		values = append(values, f.DefinitionIDLt)
	}

	if f.DefinitionIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" >= ?")
		values = append(values, f.DefinitionIDGte)
	}

	if f.DefinitionIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" <= ?")
		values = append(values, f.DefinitionIDLte)
	}

	if f.DefinitionIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" IN (?)")
		values = append(values, f.DefinitionIDIn)
	}

	if f.DefinitionIDNull != nil {
		if *f.DefinitionIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" IS NOT NULL")
		}
	}

	if f.ItemID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" = ?")
		values = append(values, f.ItemID)
	}

	if f.ItemIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" != ?")
		values = append(values, f.ItemIDNe)
	}

	if f.ItemIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" > ?")
		values = append(values, f.ItemIDGt)
	}

	if f.ItemIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" < ?")
		values = append(values, f.ItemIDLt)
	}

	if f.ItemIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" >= ?")
		values = append(values, f.ItemIDGte)
	}

	if f.ItemIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" <= ?")
		values = append(values, f.ItemIDLte)
	}

	if f.ItemIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" IN (?)")
		values = append(values, f.ItemIDIn)
	}

	if f.ItemIDNull != nil {
		if *f.ItemIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *ConfiguratorAttributeFilterType) AndWith(f2 ...*ConfiguratorAttributeFilterType) *ConfiguratorAttributeFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ConfiguratorAttributeFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *ConfiguratorAttributeFilterType) OrWith(f2 ...*ConfiguratorAttributeFilterType) *ConfiguratorAttributeFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ConfiguratorAttributeFilterType{
		Or: append(_f2, f),
	}
}

func (f *ConfiguratorSlotFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *ConfiguratorSlotFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("configuratorSlots"), wheres, values, joins)
}
func (f *ConfiguratorSlotFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Item != nil {
		_alias := alias + "_item"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorItems"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("itemId"))
		err := f.Item.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Definition != nil {
		_alias := alias + "_definition"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorSlotDefinitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("definitionId"))
		err := f.Definition.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.ParentItem != nil {
		_alias := alias + "_parentItem"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorItems"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("parentItemId"))
		err := f.ParentItem.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *ConfiguratorSlotFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Count != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("count")+" = ?")
		values = append(values, f.Count)
	}

	if f.CountNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("count")+" != ?")
		values = append(values, f.CountNe)
	}

	if f.CountGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("count")+" > ?")
		values = append(values, f.CountGt)
	}

	if f.CountLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("count")+" < ?")
		values = append(values, f.CountLt)
	}

	if f.CountGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("count")+" >= ?")
		values = append(values, f.CountGte)
	}

	if f.CountLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("count")+" <= ?")
		values = append(values, f.CountLte)
	}

	if f.CountIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("count")+" IN (?)")
		values = append(values, f.CountIn)
	}

	if f.CountNull != nil {
		if *f.CountNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("count")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("count")+" IS NOT NULL")
		}
	}

	if f.ItemID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" = ?")
		values = append(values, f.ItemID)
	}

	if f.ItemIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" != ?")
		values = append(values, f.ItemIDNe)
	}

	if f.ItemIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" > ?")
		values = append(values, f.ItemIDGt)
	}

	if f.ItemIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" < ?")
		values = append(values, f.ItemIDLt)
	}

	if f.ItemIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" >= ?")
		values = append(values, f.ItemIDGte)
	}

	if f.ItemIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" <= ?")
		values = append(values, f.ItemIDLte)
	}

	if f.ItemIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" IN (?)")
		values = append(values, f.ItemIDIn)
	}

	if f.ItemIDNull != nil {
		if *f.ItemIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" IS NOT NULL")
		}
	}

	if f.DefinitionID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" = ?")
		values = append(values, f.DefinitionID)
	}

	if f.DefinitionIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" != ?")
		values = append(values, f.DefinitionIDNe)
	}

	if f.DefinitionIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" > ?")
		values = append(values, f.DefinitionIDGt)
	}

	if f.DefinitionIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" < ?")
		values = append(values, f.DefinitionIDLt)
	}

	if f.DefinitionIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" >= ?")
		values = append(values, f.DefinitionIDGte)
	}

	if f.DefinitionIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" <= ?")
		values = append(values, f.DefinitionIDLte)
	}

	if f.DefinitionIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" IN (?)")
		values = append(values, f.DefinitionIDIn)
	}

	if f.DefinitionIDNull != nil {
		if *f.DefinitionIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" IS NOT NULL")
		}
	}

	if f.ParentItemID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("parentItemId")+" = ?")
		values = append(values, f.ParentItemID)
	}

	if f.ParentItemIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("parentItemId")+" != ?")
		values = append(values, f.ParentItemIDNe)
	}

	if f.ParentItemIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("parentItemId")+" > ?")
		values = append(values, f.ParentItemIDGt)
	}

	if f.ParentItemIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("parentItemId")+" < ?")
		values = append(values, f.ParentItemIDLt)
	}

	if f.ParentItemIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("parentItemId")+" >= ?")
		values = append(values, f.ParentItemIDGte)
	}

	if f.ParentItemIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("parentItemId")+" <= ?")
		values = append(values, f.ParentItemIDLte)
	}

	if f.ParentItemIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("parentItemId")+" IN (?)")
		values = append(values, f.ParentItemIDIn)
	}

	if f.ParentItemIDNull != nil {
		if *f.ParentItemIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("parentItemId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("parentItemId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *ConfiguratorSlotFilterType) AndWith(f2 ...*ConfiguratorSlotFilterType) *ConfiguratorSlotFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ConfiguratorSlotFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *ConfiguratorSlotFilterType) OrWith(f2 ...*ConfiguratorSlotFilterType) *ConfiguratorSlotFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ConfiguratorSlotFilterType{
		Or: append(_f2, f),
	}
}
