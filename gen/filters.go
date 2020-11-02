package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func (f *ConfiguratorItemDefinitionCategoryFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}
func (f *ConfiguratorItemDefinitionCategoryFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("configurator_item_definition_categories"), wheres, whereValues, havings, havingValues, joins)
}
func (f *ConfiguratorItemDefinitionCategoryFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Definitions != nil {
		_alias := alias + "_definitions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_item_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("categoryId")+" = "+dialect.Quote(alias)+".id")
		err := f.Definitions.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
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

	if f.IDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" NOT IN (?)")
		values = append(values, f.IDNotIn)
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

	if f.CodeNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" NOT IN (?)")
		values = append(values, f.CodeNotIn)
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

	if f.NameNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" NOT IN (?)")
		values = append(values, f.NameNotIn)
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

	if f.TypeNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" NOT IN (?)")
		values = append(values, f.TypeNotIn)
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

	if f.Primary != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("primary")+" = ?")
		values = append(values, f.Primary)
	}

	if f.PrimaryNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("primary")+" != ?")
		values = append(values, f.PrimaryNe)
	}

	if f.PrimaryGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("primary")+" > ?")
		values = append(values, f.PrimaryGt)
	}

	if f.PrimaryLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("primary")+" < ?")
		values = append(values, f.PrimaryLt)
	}

	if f.PrimaryGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("primary")+" >= ?")
		values = append(values, f.PrimaryGte)
	}

	if f.PrimaryLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("primary")+" <= ?")
		values = append(values, f.PrimaryLte)
	}

	if f.PrimaryIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("primary")+" IN (?)")
		values = append(values, f.PrimaryIn)
	}

	if f.PrimaryNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("primary")+" NOT IN (?)")
		values = append(values, f.PrimaryNotIn)
	}

	if f.PrimaryNull != nil {
		if *f.PrimaryNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("primary")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("primary")+" IS NOT NULL")
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

	if f.UpdatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" NOT IN (?)")
		values = append(values, f.UpdatedAtNotIn)
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

	if f.CreatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" NOT IN (?)")
		values = append(values, f.CreatedAtNotIn)
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

	if f.UpdatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" NOT IN (?)")
		values = append(values, f.UpdatedByNotIn)
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

	if f.CreatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" NOT IN (?)")
		values = append(values, f.CreatedByNotIn)
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
func (f *ConfiguratorItemDefinitionCategoryFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.IDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMinNotIn)
	}

	if f.IDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMaxNotIn)
	}

	if f.CodeMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") = ?")
		values = append(values, f.CodeMin)
	}

	if f.CodeMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") = ?")
		values = append(values, f.CodeMax)
	}

	if f.CodeMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") != ?")
		values = append(values, f.CodeMinNe)
	}

	if f.CodeMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") != ?")
		values = append(values, f.CodeMaxNe)
	}

	if f.CodeMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") > ?")
		values = append(values, f.CodeMinGt)
	}

	if f.CodeMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") > ?")
		values = append(values, f.CodeMaxGt)
	}

	if f.CodeMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") < ?")
		values = append(values, f.CodeMinLt)
	}

	if f.CodeMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") < ?")
		values = append(values, f.CodeMaxLt)
	}

	if f.CodeMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") >= ?")
		values = append(values, f.CodeMinGte)
	}

	if f.CodeMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") >= ?")
		values = append(values, f.CodeMaxGte)
	}

	if f.CodeMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") <= ?")
		values = append(values, f.CodeMinLte)
	}

	if f.CodeMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") <= ?")
		values = append(values, f.CodeMaxLte)
	}

	if f.CodeMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") IN (?)")
		values = append(values, f.CodeMinIn)
	}

	if f.CodeMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") IN (?)")
		values = append(values, f.CodeMaxIn)
	}

	if f.CodeMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") NOT IN (?)")
		values = append(values, f.CodeMinNotIn)
	}

	if f.CodeMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") NOT IN (?)")
		values = append(values, f.CodeMaxNotIn)
	}

	if f.CodeMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.CodeMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.CodeMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.CodeMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.CodeMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CodeMinPrefix))
	}

	if f.CodeMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CodeMaxPrefix))
	}

	if f.CodeMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CodeMinSuffix))
	}

	if f.CodeMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CodeMaxSuffix))
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") NOT IN (?)")
		values = append(values, f.NameMinNotIn)
	}

	if f.NameMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") NOT IN (?)")
		values = append(values, f.NameMaxNotIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.TypeMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") = ?")
		values = append(values, f.TypeMin)
	}

	if f.TypeMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") = ?")
		values = append(values, f.TypeMax)
	}

	if f.TypeMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") != ?")
		values = append(values, f.TypeMinNe)
	}

	if f.TypeMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") != ?")
		values = append(values, f.TypeMaxNe)
	}

	if f.TypeMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") > ?")
		values = append(values, f.TypeMinGt)
	}

	if f.TypeMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") > ?")
		values = append(values, f.TypeMaxGt)
	}

	if f.TypeMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") < ?")
		values = append(values, f.TypeMinLt)
	}

	if f.TypeMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") < ?")
		values = append(values, f.TypeMaxLt)
	}

	if f.TypeMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") >= ?")
		values = append(values, f.TypeMinGte)
	}

	if f.TypeMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") >= ?")
		values = append(values, f.TypeMaxGte)
	}

	if f.TypeMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") <= ?")
		values = append(values, f.TypeMinLte)
	}

	if f.TypeMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") <= ?")
		values = append(values, f.TypeMaxLte)
	}

	if f.TypeMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") IN (?)")
		values = append(values, f.TypeMinIn)
	}

	if f.TypeMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") IN (?)")
		values = append(values, f.TypeMaxIn)
	}

	if f.TypeMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") NOT IN (?)")
		values = append(values, f.TypeMinNotIn)
	}

	if f.TypeMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") NOT IN (?)")
		values = append(values, f.TypeMaxNotIn)
	}

	if f.TypeMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TypeMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TypeMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TypeMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TypeMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TypeMinPrefix))
	}

	if f.TypeMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TypeMaxPrefix))
	}

	if f.TypeMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TypeMinSuffix))
	}

	if f.TypeMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TypeMaxSuffix))
	}

	if f.PrimaryMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("primary")+") = ?")
		values = append(values, f.PrimaryMin)
	}

	if f.PrimaryMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("primary")+") = ?")
		values = append(values, f.PrimaryMax)
	}

	if f.PrimaryMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("primary")+") != ?")
		values = append(values, f.PrimaryMinNe)
	}

	if f.PrimaryMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("primary")+") != ?")
		values = append(values, f.PrimaryMaxNe)
	}

	if f.PrimaryMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("primary")+") > ?")
		values = append(values, f.PrimaryMinGt)
	}

	if f.PrimaryMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("primary")+") > ?")
		values = append(values, f.PrimaryMaxGt)
	}

	if f.PrimaryMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("primary")+") < ?")
		values = append(values, f.PrimaryMinLt)
	}

	if f.PrimaryMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("primary")+") < ?")
		values = append(values, f.PrimaryMaxLt)
	}

	if f.PrimaryMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("primary")+") >= ?")
		values = append(values, f.PrimaryMinGte)
	}

	if f.PrimaryMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("primary")+") >= ?")
		values = append(values, f.PrimaryMaxGte)
	}

	if f.PrimaryMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("primary")+") <= ?")
		values = append(values, f.PrimaryMinLte)
	}

	if f.PrimaryMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("primary")+") <= ?")
		values = append(values, f.PrimaryMaxLte)
	}

	if f.PrimaryMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("primary")+") IN (?)")
		values = append(values, f.PrimaryMinIn)
	}

	if f.PrimaryMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("primary")+") IN (?)")
		values = append(values, f.PrimaryMaxIn)
	}

	if f.PrimaryMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("primary")+") NOT IN (?)")
		values = append(values, f.PrimaryMinNotIn)
	}

	if f.PrimaryMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("primary")+") NOT IN (?)")
		values = append(values, f.PrimaryMaxNotIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.UpdatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMinNotIn)
	}

	if f.UpdatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMaxNotIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.CreatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMinNotIn)
	}

	if f.CreatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMaxNotIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.UpdatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMinNotIn)
	}

	if f.UpdatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMaxNotIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	if f.CreatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMinNotIn)
	}

	if f.CreatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMaxNotIn)
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
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}
func (f *ConfiguratorItemDefinitionFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("configurator_item_definitions"), wheres, whereValues, havings, havingValues, joins)
}
func (f *ConfiguratorItemDefinitionFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Attributes != nil {
		_alias := alias + "_attributes"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorAttributeDefinition_definitions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("definition_id")+" LEFT JOIN "+dialect.Quote(TableName("configurator_attribute_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("attribute_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.Attributes.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Slots != nil {
		_alias := alias + "_slots"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_slot_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("definitionId")+" = "+dialect.Quote(alias)+".id")
		err := f.Slots.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Items != nil {
		_alias := alias + "_items"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_items"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("definitionId")+" = "+dialect.Quote(alias)+".id")
		err := f.Items.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.AllowedInSlots != nil {
		_alias := alias + "_allowedInSlots"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorSlotDefinition_allowedItemDefinitions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("allowedItemDefinition_id")+" LEFT JOIN "+dialect.Quote(TableName("configurator_slot_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("allowedInSlot_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.AllowedInSlots.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Category != nil {
		_alias := alias + "_category"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_item_definition_categories"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("categoryId"))
		err := f.Category.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
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

	if f.IDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" NOT IN (?)")
		values = append(values, f.IDNotIn)
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

	if f.CodeNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" NOT IN (?)")
		values = append(values, f.CodeNotIn)
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

	if f.NameNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" NOT IN (?)")
		values = append(values, f.NameNotIn)
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

	if f.CategoryIDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("categoryId")+" NOT IN (?)")
		values = append(values, f.CategoryIDNotIn)
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

	if f.UpdatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" NOT IN (?)")
		values = append(values, f.UpdatedAtNotIn)
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

	if f.CreatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" NOT IN (?)")
		values = append(values, f.CreatedAtNotIn)
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

	if f.UpdatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" NOT IN (?)")
		values = append(values, f.UpdatedByNotIn)
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

	if f.CreatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" NOT IN (?)")
		values = append(values, f.CreatedByNotIn)
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
func (f *ConfiguratorItemDefinitionFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.IDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMinNotIn)
	}

	if f.IDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMaxNotIn)
	}

	if f.CodeMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") = ?")
		values = append(values, f.CodeMin)
	}

	if f.CodeMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") = ?")
		values = append(values, f.CodeMax)
	}

	if f.CodeMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") != ?")
		values = append(values, f.CodeMinNe)
	}

	if f.CodeMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") != ?")
		values = append(values, f.CodeMaxNe)
	}

	if f.CodeMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") > ?")
		values = append(values, f.CodeMinGt)
	}

	if f.CodeMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") > ?")
		values = append(values, f.CodeMaxGt)
	}

	if f.CodeMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") < ?")
		values = append(values, f.CodeMinLt)
	}

	if f.CodeMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") < ?")
		values = append(values, f.CodeMaxLt)
	}

	if f.CodeMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") >= ?")
		values = append(values, f.CodeMinGte)
	}

	if f.CodeMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") >= ?")
		values = append(values, f.CodeMaxGte)
	}

	if f.CodeMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") <= ?")
		values = append(values, f.CodeMinLte)
	}

	if f.CodeMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") <= ?")
		values = append(values, f.CodeMaxLte)
	}

	if f.CodeMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") IN (?)")
		values = append(values, f.CodeMinIn)
	}

	if f.CodeMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") IN (?)")
		values = append(values, f.CodeMaxIn)
	}

	if f.CodeMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") NOT IN (?)")
		values = append(values, f.CodeMinNotIn)
	}

	if f.CodeMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") NOT IN (?)")
		values = append(values, f.CodeMaxNotIn)
	}

	if f.CodeMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.CodeMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.CodeMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.CodeMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.CodeMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CodeMinPrefix))
	}

	if f.CodeMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CodeMaxPrefix))
	}

	if f.CodeMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CodeMinSuffix))
	}

	if f.CodeMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CodeMaxSuffix))
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") NOT IN (?)")
		values = append(values, f.NameMinNotIn)
	}

	if f.NameMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") NOT IN (?)")
		values = append(values, f.NameMaxNotIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.CategoryIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("categoryId")+") = ?")
		values = append(values, f.CategoryIDMin)
	}

	if f.CategoryIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("categoryId")+") = ?")
		values = append(values, f.CategoryIDMax)
	}

	if f.CategoryIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("categoryId")+") != ?")
		values = append(values, f.CategoryIDMinNe)
	}

	if f.CategoryIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("categoryId")+") != ?")
		values = append(values, f.CategoryIDMaxNe)
	}

	if f.CategoryIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("categoryId")+") > ?")
		values = append(values, f.CategoryIDMinGt)
	}

	if f.CategoryIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("categoryId")+") > ?")
		values = append(values, f.CategoryIDMaxGt)
	}

	if f.CategoryIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("categoryId")+") < ?")
		values = append(values, f.CategoryIDMinLt)
	}

	if f.CategoryIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("categoryId")+") < ?")
		values = append(values, f.CategoryIDMaxLt)
	}

	if f.CategoryIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("categoryId")+") >= ?")
		values = append(values, f.CategoryIDMinGte)
	}

	if f.CategoryIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("categoryId")+") >= ?")
		values = append(values, f.CategoryIDMaxGte)
	}

	if f.CategoryIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("categoryId")+") <= ?")
		values = append(values, f.CategoryIDMinLte)
	}

	if f.CategoryIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("categoryId")+") <= ?")
		values = append(values, f.CategoryIDMaxLte)
	}

	if f.CategoryIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("categoryId")+") IN (?)")
		values = append(values, f.CategoryIDMinIn)
	}

	if f.CategoryIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("categoryId")+") IN (?)")
		values = append(values, f.CategoryIDMaxIn)
	}

	if f.CategoryIDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("categoryId")+") NOT IN (?)")
		values = append(values, f.CategoryIDMinNotIn)
	}

	if f.CategoryIDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("categoryId")+") NOT IN (?)")
		values = append(values, f.CategoryIDMaxNotIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.UpdatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMinNotIn)
	}

	if f.UpdatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMaxNotIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.CreatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMinNotIn)
	}

	if f.CreatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMaxNotIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.UpdatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMinNotIn)
	}

	if f.UpdatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMaxNotIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	if f.CreatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMinNotIn)
	}

	if f.CreatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMaxNotIn)
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
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}
func (f *ConfiguratorAttributeDefinitionFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("configurator_attribute_definitions"), wheres, whereValues, havings, havingValues, joins)
}
func (f *ConfiguratorAttributeDefinitionFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Definitions != nil {
		_alias := alias + "_definitions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorAttributeDefinition_definitions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("attribute_id")+" LEFT JOIN "+dialect.Quote(TableName("configurator_item_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("definition_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.Definitions.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Attributes != nil {
		_alias := alias + "_attributes"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_attributes"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("definitionId")+" = "+dialect.Quote(alias)+".id")
		err := f.Attributes.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
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

	if f.IDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" NOT IN (?)")
		values = append(values, f.IDNotIn)
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

	if f.NameNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" NOT IN (?)")
		values = append(values, f.NameNotIn)
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

	if f.TypeNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" NOT IN (?)")
		values = append(values, f.TypeNotIn)
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

	if f.UpdatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" NOT IN (?)")
		values = append(values, f.UpdatedAtNotIn)
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

	if f.CreatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" NOT IN (?)")
		values = append(values, f.CreatedAtNotIn)
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

	if f.UpdatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" NOT IN (?)")
		values = append(values, f.UpdatedByNotIn)
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

	if f.CreatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" NOT IN (?)")
		values = append(values, f.CreatedByNotIn)
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
func (f *ConfiguratorAttributeDefinitionFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.IDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMinNotIn)
	}

	if f.IDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMaxNotIn)
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") NOT IN (?)")
		values = append(values, f.NameMinNotIn)
	}

	if f.NameMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") NOT IN (?)")
		values = append(values, f.NameMaxNotIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.TypeMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") = ?")
		values = append(values, f.TypeMin)
	}

	if f.TypeMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") = ?")
		values = append(values, f.TypeMax)
	}

	if f.TypeMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") != ?")
		values = append(values, f.TypeMinNe)
	}

	if f.TypeMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") != ?")
		values = append(values, f.TypeMaxNe)
	}

	if f.TypeMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") > ?")
		values = append(values, f.TypeMinGt)
	}

	if f.TypeMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") > ?")
		values = append(values, f.TypeMaxGt)
	}

	if f.TypeMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") < ?")
		values = append(values, f.TypeMinLt)
	}

	if f.TypeMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") < ?")
		values = append(values, f.TypeMaxLt)
	}

	if f.TypeMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") >= ?")
		values = append(values, f.TypeMinGte)
	}

	if f.TypeMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") >= ?")
		values = append(values, f.TypeMaxGte)
	}

	if f.TypeMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") <= ?")
		values = append(values, f.TypeMinLte)
	}

	if f.TypeMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") <= ?")
		values = append(values, f.TypeMaxLte)
	}

	if f.TypeMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") IN (?)")
		values = append(values, f.TypeMinIn)
	}

	if f.TypeMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") IN (?)")
		values = append(values, f.TypeMaxIn)
	}

	if f.TypeMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") NOT IN (?)")
		values = append(values, f.TypeMinNotIn)
	}

	if f.TypeMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") NOT IN (?)")
		values = append(values, f.TypeMaxNotIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.UpdatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMinNotIn)
	}

	if f.UpdatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMaxNotIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.CreatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMinNotIn)
	}

	if f.CreatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMaxNotIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.UpdatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMinNotIn)
	}

	if f.UpdatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMaxNotIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	if f.CreatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMinNotIn)
	}

	if f.CreatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMaxNotIn)
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
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}
func (f *ConfiguratorSlotDefinitionFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("configurator_slot_definitions"), wheres, whereValues, havings, havingValues, joins)
}
func (f *ConfiguratorSlotDefinitionFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Definition != nil {
		_alias := alias + "_definition"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_item_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("definitionId"))
		err := f.Definition.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Slots != nil {
		_alias := alias + "_slots"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_slots"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("definitionId")+" = "+dialect.Quote(alias)+".id")
		err := f.Slots.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.AllowedItemDefinitions != nil {
		_alias := alias + "_allowedItemDefinitions"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configuratorSlotDefinition_allowedItemDefinitions"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("allowedInSlot_id")+" LEFT JOIN "+dialect.Quote(TableName("configurator_item_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("allowedItemDefinition_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.AllowedItemDefinitions.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
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

	if f.IDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" NOT IN (?)")
		values = append(values, f.IDNotIn)
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

	if f.NameNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" NOT IN (?)")
		values = append(values, f.NameNotIn)
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

	if f.MinCountNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("minCount")+" NOT IN (?)")
		values = append(values, f.MinCountNotIn)
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

	if f.MaxCountNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("maxCount")+" NOT IN (?)")
		values = append(values, f.MaxCountNotIn)
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

	if f.DefaultCountNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("defaultCount")+" NOT IN (?)")
		values = append(values, f.DefaultCountNotIn)
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

	if f.DefinitionIDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" NOT IN (?)")
		values = append(values, f.DefinitionIDNotIn)
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

	if f.UpdatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" NOT IN (?)")
		values = append(values, f.UpdatedAtNotIn)
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

	if f.CreatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" NOT IN (?)")
		values = append(values, f.CreatedAtNotIn)
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

	if f.UpdatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" NOT IN (?)")
		values = append(values, f.UpdatedByNotIn)
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

	if f.CreatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" NOT IN (?)")
		values = append(values, f.CreatedByNotIn)
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
func (f *ConfiguratorSlotDefinitionFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.IDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMinNotIn)
	}

	if f.IDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMaxNotIn)
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") NOT IN (?)")
		values = append(values, f.NameMinNotIn)
	}

	if f.NameMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") NOT IN (?)")
		values = append(values, f.NameMaxNotIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.MinCountMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("minCount")+") = ?")
		values = append(values, f.MinCountMin)
	}

	if f.MinCountMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("minCount")+") = ?")
		values = append(values, f.MinCountMax)
	}

	if f.MinCountAvg != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("minCount")+") = ?")
		values = append(values, f.MinCountAvg)
	}

	if f.MinCountMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("minCount")+") != ?")
		values = append(values, f.MinCountMinNe)
	}

	if f.MinCountMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("minCount")+") != ?")
		values = append(values, f.MinCountMaxNe)
	}

	if f.MinCountAvgNe != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("minCount")+") != ?")
		values = append(values, f.MinCountAvgNe)
	}

	if f.MinCountMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("minCount")+") > ?")
		values = append(values, f.MinCountMinGt)
	}

	if f.MinCountMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("minCount")+") > ?")
		values = append(values, f.MinCountMaxGt)
	}

	if f.MinCountAvgGt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("minCount")+") > ?")
		values = append(values, f.MinCountAvgGt)
	}

	if f.MinCountMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("minCount")+") < ?")
		values = append(values, f.MinCountMinLt)
	}

	if f.MinCountMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("minCount")+") < ?")
		values = append(values, f.MinCountMaxLt)
	}

	if f.MinCountAvgLt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("minCount")+") < ?")
		values = append(values, f.MinCountAvgLt)
	}

	if f.MinCountMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("minCount")+") >= ?")
		values = append(values, f.MinCountMinGte)
	}

	if f.MinCountMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("minCount")+") >= ?")
		values = append(values, f.MinCountMaxGte)
	}

	if f.MinCountAvgGte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("minCount")+") >= ?")
		values = append(values, f.MinCountAvgGte)
	}

	if f.MinCountMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("minCount")+") <= ?")
		values = append(values, f.MinCountMinLte)
	}

	if f.MinCountMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("minCount")+") <= ?")
		values = append(values, f.MinCountMaxLte)
	}

	if f.MinCountAvgLte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("minCount")+") <= ?")
		values = append(values, f.MinCountAvgLte)
	}

	if f.MinCountMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("minCount")+") IN (?)")
		values = append(values, f.MinCountMinIn)
	}

	if f.MinCountMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("minCount")+") IN (?)")
		values = append(values, f.MinCountMaxIn)
	}

	if f.MinCountAvgIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("minCount")+") IN (?)")
		values = append(values, f.MinCountAvgIn)
	}

	if f.MinCountMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("minCount")+") NOT IN (?)")
		values = append(values, f.MinCountMinNotIn)
	}

	if f.MinCountMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("minCount")+") NOT IN (?)")
		values = append(values, f.MinCountMaxNotIn)
	}

	if f.MinCountAvgNotIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("minCount")+") NOT IN (?)")
		values = append(values, f.MinCountAvgNotIn)
	}

	if f.MaxCountMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("maxCount")+") = ?")
		values = append(values, f.MaxCountMin)
	}

	if f.MaxCountMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("maxCount")+") = ?")
		values = append(values, f.MaxCountMax)
	}

	if f.MaxCountAvg != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("maxCount")+") = ?")
		values = append(values, f.MaxCountAvg)
	}

	if f.MaxCountMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("maxCount")+") != ?")
		values = append(values, f.MaxCountMinNe)
	}

	if f.MaxCountMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("maxCount")+") != ?")
		values = append(values, f.MaxCountMaxNe)
	}

	if f.MaxCountAvgNe != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("maxCount")+") != ?")
		values = append(values, f.MaxCountAvgNe)
	}

	if f.MaxCountMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("maxCount")+") > ?")
		values = append(values, f.MaxCountMinGt)
	}

	if f.MaxCountMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("maxCount")+") > ?")
		values = append(values, f.MaxCountMaxGt)
	}

	if f.MaxCountAvgGt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("maxCount")+") > ?")
		values = append(values, f.MaxCountAvgGt)
	}

	if f.MaxCountMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("maxCount")+") < ?")
		values = append(values, f.MaxCountMinLt)
	}

	if f.MaxCountMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("maxCount")+") < ?")
		values = append(values, f.MaxCountMaxLt)
	}

	if f.MaxCountAvgLt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("maxCount")+") < ?")
		values = append(values, f.MaxCountAvgLt)
	}

	if f.MaxCountMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("maxCount")+") >= ?")
		values = append(values, f.MaxCountMinGte)
	}

	if f.MaxCountMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("maxCount")+") >= ?")
		values = append(values, f.MaxCountMaxGte)
	}

	if f.MaxCountAvgGte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("maxCount")+") >= ?")
		values = append(values, f.MaxCountAvgGte)
	}

	if f.MaxCountMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("maxCount")+") <= ?")
		values = append(values, f.MaxCountMinLte)
	}

	if f.MaxCountMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("maxCount")+") <= ?")
		values = append(values, f.MaxCountMaxLte)
	}

	if f.MaxCountAvgLte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("maxCount")+") <= ?")
		values = append(values, f.MaxCountAvgLte)
	}

	if f.MaxCountMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("maxCount")+") IN (?)")
		values = append(values, f.MaxCountMinIn)
	}

	if f.MaxCountMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("maxCount")+") IN (?)")
		values = append(values, f.MaxCountMaxIn)
	}

	if f.MaxCountAvgIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("maxCount")+") IN (?)")
		values = append(values, f.MaxCountAvgIn)
	}

	if f.MaxCountMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("maxCount")+") NOT IN (?)")
		values = append(values, f.MaxCountMinNotIn)
	}

	if f.MaxCountMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("maxCount")+") NOT IN (?)")
		values = append(values, f.MaxCountMaxNotIn)
	}

	if f.MaxCountAvgNotIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("maxCount")+") NOT IN (?)")
		values = append(values, f.MaxCountAvgNotIn)
	}

	if f.DefaultCountMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("defaultCount")+") = ?")
		values = append(values, f.DefaultCountMin)
	}

	if f.DefaultCountMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("defaultCount")+") = ?")
		values = append(values, f.DefaultCountMax)
	}

	if f.DefaultCountAvg != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("defaultCount")+") = ?")
		values = append(values, f.DefaultCountAvg)
	}

	if f.DefaultCountMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("defaultCount")+") != ?")
		values = append(values, f.DefaultCountMinNe)
	}

	if f.DefaultCountMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("defaultCount")+") != ?")
		values = append(values, f.DefaultCountMaxNe)
	}

	if f.DefaultCountAvgNe != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("defaultCount")+") != ?")
		values = append(values, f.DefaultCountAvgNe)
	}

	if f.DefaultCountMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("defaultCount")+") > ?")
		values = append(values, f.DefaultCountMinGt)
	}

	if f.DefaultCountMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("defaultCount")+") > ?")
		values = append(values, f.DefaultCountMaxGt)
	}

	if f.DefaultCountAvgGt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("defaultCount")+") > ?")
		values = append(values, f.DefaultCountAvgGt)
	}

	if f.DefaultCountMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("defaultCount")+") < ?")
		values = append(values, f.DefaultCountMinLt)
	}

	if f.DefaultCountMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("defaultCount")+") < ?")
		values = append(values, f.DefaultCountMaxLt)
	}

	if f.DefaultCountAvgLt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("defaultCount")+") < ?")
		values = append(values, f.DefaultCountAvgLt)
	}

	if f.DefaultCountMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("defaultCount")+") >= ?")
		values = append(values, f.DefaultCountMinGte)
	}

	if f.DefaultCountMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("defaultCount")+") >= ?")
		values = append(values, f.DefaultCountMaxGte)
	}

	if f.DefaultCountAvgGte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("defaultCount")+") >= ?")
		values = append(values, f.DefaultCountAvgGte)
	}

	if f.DefaultCountMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("defaultCount")+") <= ?")
		values = append(values, f.DefaultCountMinLte)
	}

	if f.DefaultCountMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("defaultCount")+") <= ?")
		values = append(values, f.DefaultCountMaxLte)
	}

	if f.DefaultCountAvgLte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("defaultCount")+") <= ?")
		values = append(values, f.DefaultCountAvgLte)
	}

	if f.DefaultCountMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("defaultCount")+") IN (?)")
		values = append(values, f.DefaultCountMinIn)
	}

	if f.DefaultCountMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("defaultCount")+") IN (?)")
		values = append(values, f.DefaultCountMaxIn)
	}

	if f.DefaultCountAvgIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("defaultCount")+") IN (?)")
		values = append(values, f.DefaultCountAvgIn)
	}

	if f.DefaultCountMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("defaultCount")+") NOT IN (?)")
		values = append(values, f.DefaultCountMinNotIn)
	}

	if f.DefaultCountMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("defaultCount")+") NOT IN (?)")
		values = append(values, f.DefaultCountMaxNotIn)
	}

	if f.DefaultCountAvgNotIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("defaultCount")+") NOT IN (?)")
		values = append(values, f.DefaultCountAvgNotIn)
	}

	if f.DefinitionIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") = ?")
		values = append(values, f.DefinitionIDMin)
	}

	if f.DefinitionIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") = ?")
		values = append(values, f.DefinitionIDMax)
	}

	if f.DefinitionIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") != ?")
		values = append(values, f.DefinitionIDMinNe)
	}

	if f.DefinitionIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") != ?")
		values = append(values, f.DefinitionIDMaxNe)
	}

	if f.DefinitionIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") > ?")
		values = append(values, f.DefinitionIDMinGt)
	}

	if f.DefinitionIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") > ?")
		values = append(values, f.DefinitionIDMaxGt)
	}

	if f.DefinitionIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") < ?")
		values = append(values, f.DefinitionIDMinLt)
	}

	if f.DefinitionIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") < ?")
		values = append(values, f.DefinitionIDMaxLt)
	}

	if f.DefinitionIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") >= ?")
		values = append(values, f.DefinitionIDMinGte)
	}

	if f.DefinitionIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") >= ?")
		values = append(values, f.DefinitionIDMaxGte)
	}

	if f.DefinitionIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") <= ?")
		values = append(values, f.DefinitionIDMinLte)
	}

	if f.DefinitionIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") <= ?")
		values = append(values, f.DefinitionIDMaxLte)
	}

	if f.DefinitionIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") IN (?)")
		values = append(values, f.DefinitionIDMinIn)
	}

	if f.DefinitionIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") IN (?)")
		values = append(values, f.DefinitionIDMaxIn)
	}

	if f.DefinitionIDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") NOT IN (?)")
		values = append(values, f.DefinitionIDMinNotIn)
	}

	if f.DefinitionIDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") NOT IN (?)")
		values = append(values, f.DefinitionIDMaxNotIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.UpdatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMinNotIn)
	}

	if f.UpdatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMaxNotIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.CreatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMinNotIn)
	}

	if f.CreatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMaxNotIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.UpdatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMinNotIn)
	}

	if f.UpdatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMaxNotIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	if f.CreatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMinNotIn)
	}

	if f.CreatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMaxNotIn)
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
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}
func (f *ConfiguratorItemFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("configurator_items"), wheres, whereValues, havings, havingValues, joins)
}
func (f *ConfiguratorItemFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Definition != nil {
		_alias := alias + "_definition"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_item_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("definitionId"))
		err := f.Definition.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Attributes != nil {
		_alias := alias + "_attributes"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_attributes"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("itemId")+" = "+dialect.Quote(alias)+".id")
		err := f.Attributes.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Slots != nil {
		_alias := alias + "_slots"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_slots"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("parentItemId")+" = "+dialect.Quote(alias)+".id")
		err := f.Slots.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.ParentSlots != nil {
		_alias := alias + "_parentSlots"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_slots"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("itemId")+" = "+dialect.Quote(alias)+".id")
		err := f.ParentSlots.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
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

	if f.IDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" NOT IN (?)")
		values = append(values, f.IDNotIn)
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

	if f.CodeNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" NOT IN (?)")
		values = append(values, f.CodeNotIn)
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

	if f.NameNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" NOT IN (?)")
		values = append(values, f.NameNotIn)
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

	if f.StockItemIDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stockItemId")+" NOT IN (?)")
		values = append(values, f.StockItemIDNotIn)
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

	if f.RawDataNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("rawData")+" NOT IN (?)")
		values = append(values, f.RawDataNotIn)
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

	if f.DefinitionIDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" NOT IN (?)")
		values = append(values, f.DefinitionIDNotIn)
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

	if f.UpdatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" NOT IN (?)")
		values = append(values, f.UpdatedAtNotIn)
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

	if f.CreatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" NOT IN (?)")
		values = append(values, f.CreatedAtNotIn)
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

	if f.UpdatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" NOT IN (?)")
		values = append(values, f.UpdatedByNotIn)
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

	if f.CreatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" NOT IN (?)")
		values = append(values, f.CreatedByNotIn)
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
func (f *ConfiguratorItemFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.IDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMinNotIn)
	}

	if f.IDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMaxNotIn)
	}

	if f.CodeMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") = ?")
		values = append(values, f.CodeMin)
	}

	if f.CodeMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") = ?")
		values = append(values, f.CodeMax)
	}

	if f.CodeMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") != ?")
		values = append(values, f.CodeMinNe)
	}

	if f.CodeMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") != ?")
		values = append(values, f.CodeMaxNe)
	}

	if f.CodeMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") > ?")
		values = append(values, f.CodeMinGt)
	}

	if f.CodeMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") > ?")
		values = append(values, f.CodeMaxGt)
	}

	if f.CodeMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") < ?")
		values = append(values, f.CodeMinLt)
	}

	if f.CodeMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") < ?")
		values = append(values, f.CodeMaxLt)
	}

	if f.CodeMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") >= ?")
		values = append(values, f.CodeMinGte)
	}

	if f.CodeMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") >= ?")
		values = append(values, f.CodeMaxGte)
	}

	if f.CodeMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") <= ?")
		values = append(values, f.CodeMinLte)
	}

	if f.CodeMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") <= ?")
		values = append(values, f.CodeMaxLte)
	}

	if f.CodeMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") IN (?)")
		values = append(values, f.CodeMinIn)
	}

	if f.CodeMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") IN (?)")
		values = append(values, f.CodeMaxIn)
	}

	if f.CodeMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") NOT IN (?)")
		values = append(values, f.CodeMinNotIn)
	}

	if f.CodeMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") NOT IN (?)")
		values = append(values, f.CodeMaxNotIn)
	}

	if f.CodeMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.CodeMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.CodeMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.CodeMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.CodeMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CodeMinPrefix))
	}

	if f.CodeMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.CodeMaxPrefix))
	}

	if f.CodeMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CodeMinSuffix))
	}

	if f.CodeMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.CodeMaxSuffix))
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") NOT IN (?)")
		values = append(values, f.NameMinNotIn)
	}

	if f.NameMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") NOT IN (?)")
		values = append(values, f.NameMaxNotIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.StockItemIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stockItemId")+") = ?")
		values = append(values, f.StockItemIDMin)
	}

	if f.StockItemIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stockItemId")+") = ?")
		values = append(values, f.StockItemIDMax)
	}

	if f.StockItemIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stockItemId")+") != ?")
		values = append(values, f.StockItemIDMinNe)
	}

	if f.StockItemIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stockItemId")+") != ?")
		values = append(values, f.StockItemIDMaxNe)
	}

	if f.StockItemIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stockItemId")+") > ?")
		values = append(values, f.StockItemIDMinGt)
	}

	if f.StockItemIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stockItemId")+") > ?")
		values = append(values, f.StockItemIDMaxGt)
	}

	if f.StockItemIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stockItemId")+") < ?")
		values = append(values, f.StockItemIDMinLt)
	}

	if f.StockItemIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stockItemId")+") < ?")
		values = append(values, f.StockItemIDMaxLt)
	}

	if f.StockItemIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stockItemId")+") >= ?")
		values = append(values, f.StockItemIDMinGte)
	}

	if f.StockItemIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stockItemId")+") >= ?")
		values = append(values, f.StockItemIDMaxGte)
	}

	if f.StockItemIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stockItemId")+") <= ?")
		values = append(values, f.StockItemIDMinLte)
	}

	if f.StockItemIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stockItemId")+") <= ?")
		values = append(values, f.StockItemIDMaxLte)
	}

	if f.StockItemIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stockItemId")+") IN (?)")
		values = append(values, f.StockItemIDMinIn)
	}

	if f.StockItemIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stockItemId")+") IN (?)")
		values = append(values, f.StockItemIDMaxIn)
	}

	if f.StockItemIDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stockItemId")+") NOT IN (?)")
		values = append(values, f.StockItemIDMinNotIn)
	}

	if f.StockItemIDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stockItemId")+") NOT IN (?)")
		values = append(values, f.StockItemIDMaxNotIn)
	}

	if f.RawDataMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("rawData")+") = ?")
		values = append(values, f.RawDataMin)
	}

	if f.RawDataMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("rawData")+") = ?")
		values = append(values, f.RawDataMax)
	}

	if f.RawDataMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("rawData")+") != ?")
		values = append(values, f.RawDataMinNe)
	}

	if f.RawDataMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("rawData")+") != ?")
		values = append(values, f.RawDataMaxNe)
	}

	if f.RawDataMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("rawData")+") > ?")
		values = append(values, f.RawDataMinGt)
	}

	if f.RawDataMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("rawData")+") > ?")
		values = append(values, f.RawDataMaxGt)
	}

	if f.RawDataMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("rawData")+") < ?")
		values = append(values, f.RawDataMinLt)
	}

	if f.RawDataMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("rawData")+") < ?")
		values = append(values, f.RawDataMaxLt)
	}

	if f.RawDataMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("rawData")+") >= ?")
		values = append(values, f.RawDataMinGte)
	}

	if f.RawDataMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("rawData")+") >= ?")
		values = append(values, f.RawDataMaxGte)
	}

	if f.RawDataMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("rawData")+") <= ?")
		values = append(values, f.RawDataMinLte)
	}

	if f.RawDataMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("rawData")+") <= ?")
		values = append(values, f.RawDataMaxLte)
	}

	if f.RawDataMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("rawData")+") IN (?)")
		values = append(values, f.RawDataMinIn)
	}

	if f.RawDataMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("rawData")+") IN (?)")
		values = append(values, f.RawDataMaxIn)
	}

	if f.RawDataMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("rawData")+") NOT IN (?)")
		values = append(values, f.RawDataMinNotIn)
	}

	if f.RawDataMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("rawData")+") NOT IN (?)")
		values = append(values, f.RawDataMaxNotIn)
	}

	if f.RawDataMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("rawData")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.RawDataMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.RawDataMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("rawData")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.RawDataMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.RawDataMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("rawData")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.RawDataMinPrefix))
	}

	if f.RawDataMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("rawData")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.RawDataMaxPrefix))
	}

	if f.RawDataMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("rawData")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.RawDataMinSuffix))
	}

	if f.RawDataMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("rawData")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.RawDataMaxSuffix))
	}

	if f.DefinitionIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") = ?")
		values = append(values, f.DefinitionIDMin)
	}

	if f.DefinitionIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") = ?")
		values = append(values, f.DefinitionIDMax)
	}

	if f.DefinitionIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") != ?")
		values = append(values, f.DefinitionIDMinNe)
	}

	if f.DefinitionIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") != ?")
		values = append(values, f.DefinitionIDMaxNe)
	}

	if f.DefinitionIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") > ?")
		values = append(values, f.DefinitionIDMinGt)
	}

	if f.DefinitionIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") > ?")
		values = append(values, f.DefinitionIDMaxGt)
	}

	if f.DefinitionIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") < ?")
		values = append(values, f.DefinitionIDMinLt)
	}

	if f.DefinitionIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") < ?")
		values = append(values, f.DefinitionIDMaxLt)
	}

	if f.DefinitionIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") >= ?")
		values = append(values, f.DefinitionIDMinGte)
	}

	if f.DefinitionIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") >= ?")
		values = append(values, f.DefinitionIDMaxGte)
	}

	if f.DefinitionIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") <= ?")
		values = append(values, f.DefinitionIDMinLte)
	}

	if f.DefinitionIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") <= ?")
		values = append(values, f.DefinitionIDMaxLte)
	}

	if f.DefinitionIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") IN (?)")
		values = append(values, f.DefinitionIDMinIn)
	}

	if f.DefinitionIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") IN (?)")
		values = append(values, f.DefinitionIDMaxIn)
	}

	if f.DefinitionIDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") NOT IN (?)")
		values = append(values, f.DefinitionIDMinNotIn)
	}

	if f.DefinitionIDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") NOT IN (?)")
		values = append(values, f.DefinitionIDMaxNotIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.UpdatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMinNotIn)
	}

	if f.UpdatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMaxNotIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.CreatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMinNotIn)
	}

	if f.CreatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMaxNotIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.UpdatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMinNotIn)
	}

	if f.UpdatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMaxNotIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	if f.CreatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMinNotIn)
	}

	if f.CreatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMaxNotIn)
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
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}
func (f *ConfiguratorAttributeFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("configurator_attributes"), wheres, whereValues, havings, havingValues, joins)
}
func (f *ConfiguratorAttributeFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Definition != nil {
		_alias := alias + "_definition"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_attribute_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("definitionId"))
		err := f.Definition.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Item != nil {
		_alias := alias + "_item"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_items"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("itemId"))
		err := f.Item.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
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

	if f.IDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" NOT IN (?)")
		values = append(values, f.IDNotIn)
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

	if f.StringValueNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("stringValue")+" NOT IN (?)")
		values = append(values, f.StringValueNotIn)
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

	if f.FloatValueNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("floatValue")+" NOT IN (?)")
		values = append(values, f.FloatValueNotIn)
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

	if f.IntValueNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("intValue")+" NOT IN (?)")
		values = append(values, f.IntValueNotIn)
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

	if f.DefinitionIDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" NOT IN (?)")
		values = append(values, f.DefinitionIDNotIn)
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

	if f.ItemIDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" NOT IN (?)")
		values = append(values, f.ItemIDNotIn)
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

	if f.UpdatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" NOT IN (?)")
		values = append(values, f.UpdatedAtNotIn)
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

	if f.CreatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" NOT IN (?)")
		values = append(values, f.CreatedAtNotIn)
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

	if f.UpdatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" NOT IN (?)")
		values = append(values, f.UpdatedByNotIn)
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

	if f.CreatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" NOT IN (?)")
		values = append(values, f.CreatedByNotIn)
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
func (f *ConfiguratorAttributeFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.IDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMinNotIn)
	}

	if f.IDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMaxNotIn)
	}

	if f.StringValueMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stringValue")+") = ?")
		values = append(values, f.StringValueMin)
	}

	if f.StringValueMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stringValue")+") = ?")
		values = append(values, f.StringValueMax)
	}

	if f.StringValueMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stringValue")+") != ?")
		values = append(values, f.StringValueMinNe)
	}

	if f.StringValueMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stringValue")+") != ?")
		values = append(values, f.StringValueMaxNe)
	}

	if f.StringValueMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stringValue")+") > ?")
		values = append(values, f.StringValueMinGt)
	}

	if f.StringValueMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stringValue")+") > ?")
		values = append(values, f.StringValueMaxGt)
	}

	if f.StringValueMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stringValue")+") < ?")
		values = append(values, f.StringValueMinLt)
	}

	if f.StringValueMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stringValue")+") < ?")
		values = append(values, f.StringValueMaxLt)
	}

	if f.StringValueMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stringValue")+") >= ?")
		values = append(values, f.StringValueMinGte)
	}

	if f.StringValueMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stringValue")+") >= ?")
		values = append(values, f.StringValueMaxGte)
	}

	if f.StringValueMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stringValue")+") <= ?")
		values = append(values, f.StringValueMinLte)
	}

	if f.StringValueMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stringValue")+") <= ?")
		values = append(values, f.StringValueMaxLte)
	}

	if f.StringValueMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stringValue")+") IN (?)")
		values = append(values, f.StringValueMinIn)
	}

	if f.StringValueMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stringValue")+") IN (?)")
		values = append(values, f.StringValueMaxIn)
	}

	if f.StringValueMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stringValue")+") NOT IN (?)")
		values = append(values, f.StringValueMinNotIn)
	}

	if f.StringValueMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stringValue")+") NOT IN (?)")
		values = append(values, f.StringValueMaxNotIn)
	}

	if f.StringValueMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stringValue")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.StringValueMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.StringValueMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stringValue")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.StringValueMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.StringValueMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stringValue")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.StringValueMinPrefix))
	}

	if f.StringValueMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stringValue")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.StringValueMaxPrefix))
	}

	if f.StringValueMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("stringValue")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.StringValueMinSuffix))
	}

	if f.StringValueMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("stringValue")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.StringValueMaxSuffix))
	}

	if f.FloatValueMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("floatValue")+") = ?")
		values = append(values, f.FloatValueMin)
	}

	if f.FloatValueMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("floatValue")+") = ?")
		values = append(values, f.FloatValueMax)
	}

	if f.FloatValueAvg != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("floatValue")+") = ?")
		values = append(values, f.FloatValueAvg)
	}

	if f.FloatValueMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("floatValue")+") != ?")
		values = append(values, f.FloatValueMinNe)
	}

	if f.FloatValueMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("floatValue")+") != ?")
		values = append(values, f.FloatValueMaxNe)
	}

	if f.FloatValueAvgNe != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("floatValue")+") != ?")
		values = append(values, f.FloatValueAvgNe)
	}

	if f.FloatValueMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("floatValue")+") > ?")
		values = append(values, f.FloatValueMinGt)
	}

	if f.FloatValueMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("floatValue")+") > ?")
		values = append(values, f.FloatValueMaxGt)
	}

	if f.FloatValueAvgGt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("floatValue")+") > ?")
		values = append(values, f.FloatValueAvgGt)
	}

	if f.FloatValueMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("floatValue")+") < ?")
		values = append(values, f.FloatValueMinLt)
	}

	if f.FloatValueMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("floatValue")+") < ?")
		values = append(values, f.FloatValueMaxLt)
	}

	if f.FloatValueAvgLt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("floatValue")+") < ?")
		values = append(values, f.FloatValueAvgLt)
	}

	if f.FloatValueMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("floatValue")+") >= ?")
		values = append(values, f.FloatValueMinGte)
	}

	if f.FloatValueMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("floatValue")+") >= ?")
		values = append(values, f.FloatValueMaxGte)
	}

	if f.FloatValueAvgGte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("floatValue")+") >= ?")
		values = append(values, f.FloatValueAvgGte)
	}

	if f.FloatValueMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("floatValue")+") <= ?")
		values = append(values, f.FloatValueMinLte)
	}

	if f.FloatValueMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("floatValue")+") <= ?")
		values = append(values, f.FloatValueMaxLte)
	}

	if f.FloatValueAvgLte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("floatValue")+") <= ?")
		values = append(values, f.FloatValueAvgLte)
	}

	if f.FloatValueMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("floatValue")+") IN (?)")
		values = append(values, f.FloatValueMinIn)
	}

	if f.FloatValueMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("floatValue")+") IN (?)")
		values = append(values, f.FloatValueMaxIn)
	}

	if f.FloatValueAvgIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("floatValue")+") IN (?)")
		values = append(values, f.FloatValueAvgIn)
	}

	if f.FloatValueMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("floatValue")+") NOT IN (?)")
		values = append(values, f.FloatValueMinNotIn)
	}

	if f.FloatValueMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("floatValue")+") NOT IN (?)")
		values = append(values, f.FloatValueMaxNotIn)
	}

	if f.FloatValueAvgNotIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("floatValue")+") NOT IN (?)")
		values = append(values, f.FloatValueAvgNotIn)
	}

	if f.IntValueMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("intValue")+") = ?")
		values = append(values, f.IntValueMin)
	}

	if f.IntValueMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("intValue")+") = ?")
		values = append(values, f.IntValueMax)
	}

	if f.IntValueAvg != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("intValue")+") = ?")
		values = append(values, f.IntValueAvg)
	}

	if f.IntValueMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("intValue")+") != ?")
		values = append(values, f.IntValueMinNe)
	}

	if f.IntValueMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("intValue")+") != ?")
		values = append(values, f.IntValueMaxNe)
	}

	if f.IntValueAvgNe != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("intValue")+") != ?")
		values = append(values, f.IntValueAvgNe)
	}

	if f.IntValueMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("intValue")+") > ?")
		values = append(values, f.IntValueMinGt)
	}

	if f.IntValueMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("intValue")+") > ?")
		values = append(values, f.IntValueMaxGt)
	}

	if f.IntValueAvgGt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("intValue")+") > ?")
		values = append(values, f.IntValueAvgGt)
	}

	if f.IntValueMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("intValue")+") < ?")
		values = append(values, f.IntValueMinLt)
	}

	if f.IntValueMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("intValue")+") < ?")
		values = append(values, f.IntValueMaxLt)
	}

	if f.IntValueAvgLt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("intValue")+") < ?")
		values = append(values, f.IntValueAvgLt)
	}

	if f.IntValueMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("intValue")+") >= ?")
		values = append(values, f.IntValueMinGte)
	}

	if f.IntValueMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("intValue")+") >= ?")
		values = append(values, f.IntValueMaxGte)
	}

	if f.IntValueAvgGte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("intValue")+") >= ?")
		values = append(values, f.IntValueAvgGte)
	}

	if f.IntValueMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("intValue")+") <= ?")
		values = append(values, f.IntValueMinLte)
	}

	if f.IntValueMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("intValue")+") <= ?")
		values = append(values, f.IntValueMaxLte)
	}

	if f.IntValueAvgLte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("intValue")+") <= ?")
		values = append(values, f.IntValueAvgLte)
	}

	if f.IntValueMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("intValue")+") IN (?)")
		values = append(values, f.IntValueMinIn)
	}

	if f.IntValueMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("intValue")+") IN (?)")
		values = append(values, f.IntValueMaxIn)
	}

	if f.IntValueAvgIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("intValue")+") IN (?)")
		values = append(values, f.IntValueAvgIn)
	}

	if f.IntValueMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("intValue")+") NOT IN (?)")
		values = append(values, f.IntValueMinNotIn)
	}

	if f.IntValueMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("intValue")+") NOT IN (?)")
		values = append(values, f.IntValueMaxNotIn)
	}

	if f.IntValueAvgNotIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("intValue")+") NOT IN (?)")
		values = append(values, f.IntValueAvgNotIn)
	}

	if f.DefinitionIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") = ?")
		values = append(values, f.DefinitionIDMin)
	}

	if f.DefinitionIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") = ?")
		values = append(values, f.DefinitionIDMax)
	}

	if f.DefinitionIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") != ?")
		values = append(values, f.DefinitionIDMinNe)
	}

	if f.DefinitionIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") != ?")
		values = append(values, f.DefinitionIDMaxNe)
	}

	if f.DefinitionIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") > ?")
		values = append(values, f.DefinitionIDMinGt)
	}

	if f.DefinitionIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") > ?")
		values = append(values, f.DefinitionIDMaxGt)
	}

	if f.DefinitionIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") < ?")
		values = append(values, f.DefinitionIDMinLt)
	}

	if f.DefinitionIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") < ?")
		values = append(values, f.DefinitionIDMaxLt)
	}

	if f.DefinitionIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") >= ?")
		values = append(values, f.DefinitionIDMinGte)
	}

	if f.DefinitionIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") >= ?")
		values = append(values, f.DefinitionIDMaxGte)
	}

	if f.DefinitionIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") <= ?")
		values = append(values, f.DefinitionIDMinLte)
	}

	if f.DefinitionIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") <= ?")
		values = append(values, f.DefinitionIDMaxLte)
	}

	if f.DefinitionIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") IN (?)")
		values = append(values, f.DefinitionIDMinIn)
	}

	if f.DefinitionIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") IN (?)")
		values = append(values, f.DefinitionIDMaxIn)
	}

	if f.DefinitionIDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") NOT IN (?)")
		values = append(values, f.DefinitionIDMinNotIn)
	}

	if f.DefinitionIDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") NOT IN (?)")
		values = append(values, f.DefinitionIDMaxNotIn)
	}

	if f.ItemIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") = ?")
		values = append(values, f.ItemIDMin)
	}

	if f.ItemIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") = ?")
		values = append(values, f.ItemIDMax)
	}

	if f.ItemIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") != ?")
		values = append(values, f.ItemIDMinNe)
	}

	if f.ItemIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") != ?")
		values = append(values, f.ItemIDMaxNe)
	}

	if f.ItemIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") > ?")
		values = append(values, f.ItemIDMinGt)
	}

	if f.ItemIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") > ?")
		values = append(values, f.ItemIDMaxGt)
	}

	if f.ItemIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") < ?")
		values = append(values, f.ItemIDMinLt)
	}

	if f.ItemIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") < ?")
		values = append(values, f.ItemIDMaxLt)
	}

	if f.ItemIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") >= ?")
		values = append(values, f.ItemIDMinGte)
	}

	if f.ItemIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") >= ?")
		values = append(values, f.ItemIDMaxGte)
	}

	if f.ItemIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") <= ?")
		values = append(values, f.ItemIDMinLte)
	}

	if f.ItemIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") <= ?")
		values = append(values, f.ItemIDMaxLte)
	}

	if f.ItemIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") IN (?)")
		values = append(values, f.ItemIDMinIn)
	}

	if f.ItemIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") IN (?)")
		values = append(values, f.ItemIDMaxIn)
	}

	if f.ItemIDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") NOT IN (?)")
		values = append(values, f.ItemIDMinNotIn)
	}

	if f.ItemIDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") NOT IN (?)")
		values = append(values, f.ItemIDMaxNotIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.UpdatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMinNotIn)
	}

	if f.UpdatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMaxNotIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.CreatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMinNotIn)
	}

	if f.CreatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMaxNotIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.UpdatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMinNotIn)
	}

	if f.UpdatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMaxNotIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	if f.CreatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMinNotIn)
	}

	if f.CreatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMaxNotIn)
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
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}
func (f *ConfiguratorSlotFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("configurator_slots"), wheres, whereValues, havings, havingValues, joins)
}
func (f *ConfiguratorSlotFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Item != nil {
		_alias := alias + "_item"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_items"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("itemId"))
		err := f.Item.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Definition != nil {
		_alias := alias + "_definition"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_slot_definitions"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("definitionId"))
		err := f.Definition.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.ParentItem != nil {
		_alias := alias + "_parentItem"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("configurator_items"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("parentItemId"))
		err := f.ParentItem.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
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

	if f.IDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" NOT IN (?)")
		values = append(values, f.IDNotIn)
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

	if f.CountNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("count")+" NOT IN (?)")
		values = append(values, f.CountNotIn)
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

	if f.ItemIDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("itemId")+" NOT IN (?)")
		values = append(values, f.ItemIDNotIn)
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

	if f.DefinitionIDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("definitionId")+" NOT IN (?)")
		values = append(values, f.DefinitionIDNotIn)
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

	if f.ParentItemIDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("parentItemId")+" NOT IN (?)")
		values = append(values, f.ParentItemIDNotIn)
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

	if f.UpdatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" NOT IN (?)")
		values = append(values, f.UpdatedAtNotIn)
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

	if f.CreatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" NOT IN (?)")
		values = append(values, f.CreatedAtNotIn)
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

	if f.UpdatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" NOT IN (?)")
		values = append(values, f.UpdatedByNotIn)
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

	if f.CreatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" NOT IN (?)")
		values = append(values, f.CreatedByNotIn)
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
func (f *ConfiguratorSlotFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.IDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMinNotIn)
	}

	if f.IDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMaxNotIn)
	}

	if f.CountMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("count")+") = ?")
		values = append(values, f.CountMin)
	}

	if f.CountMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("count")+") = ?")
		values = append(values, f.CountMax)
	}

	if f.CountAvg != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("count")+") = ?")
		values = append(values, f.CountAvg)
	}

	if f.CountMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("count")+") != ?")
		values = append(values, f.CountMinNe)
	}

	if f.CountMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("count")+") != ?")
		values = append(values, f.CountMaxNe)
	}

	if f.CountAvgNe != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("count")+") != ?")
		values = append(values, f.CountAvgNe)
	}

	if f.CountMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("count")+") > ?")
		values = append(values, f.CountMinGt)
	}

	if f.CountMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("count")+") > ?")
		values = append(values, f.CountMaxGt)
	}

	if f.CountAvgGt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("count")+") > ?")
		values = append(values, f.CountAvgGt)
	}

	if f.CountMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("count")+") < ?")
		values = append(values, f.CountMinLt)
	}

	if f.CountMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("count")+") < ?")
		values = append(values, f.CountMaxLt)
	}

	if f.CountAvgLt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("count")+") < ?")
		values = append(values, f.CountAvgLt)
	}

	if f.CountMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("count")+") >= ?")
		values = append(values, f.CountMinGte)
	}

	if f.CountMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("count")+") >= ?")
		values = append(values, f.CountMaxGte)
	}

	if f.CountAvgGte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("count")+") >= ?")
		values = append(values, f.CountAvgGte)
	}

	if f.CountMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("count")+") <= ?")
		values = append(values, f.CountMinLte)
	}

	if f.CountMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("count")+") <= ?")
		values = append(values, f.CountMaxLte)
	}

	if f.CountAvgLte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("count")+") <= ?")
		values = append(values, f.CountAvgLte)
	}

	if f.CountMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("count")+") IN (?)")
		values = append(values, f.CountMinIn)
	}

	if f.CountMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("count")+") IN (?)")
		values = append(values, f.CountMaxIn)
	}

	if f.CountAvgIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("count")+") IN (?)")
		values = append(values, f.CountAvgIn)
	}

	if f.CountMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("count")+") NOT IN (?)")
		values = append(values, f.CountMinNotIn)
	}

	if f.CountMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("count")+") NOT IN (?)")
		values = append(values, f.CountMaxNotIn)
	}

	if f.CountAvgNotIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("count")+") NOT IN (?)")
		values = append(values, f.CountAvgNotIn)
	}

	if f.ItemIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") = ?")
		values = append(values, f.ItemIDMin)
	}

	if f.ItemIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") = ?")
		values = append(values, f.ItemIDMax)
	}

	if f.ItemIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") != ?")
		values = append(values, f.ItemIDMinNe)
	}

	if f.ItemIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") != ?")
		values = append(values, f.ItemIDMaxNe)
	}

	if f.ItemIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") > ?")
		values = append(values, f.ItemIDMinGt)
	}

	if f.ItemIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") > ?")
		values = append(values, f.ItemIDMaxGt)
	}

	if f.ItemIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") < ?")
		values = append(values, f.ItemIDMinLt)
	}

	if f.ItemIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") < ?")
		values = append(values, f.ItemIDMaxLt)
	}

	if f.ItemIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") >= ?")
		values = append(values, f.ItemIDMinGte)
	}

	if f.ItemIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") >= ?")
		values = append(values, f.ItemIDMaxGte)
	}

	if f.ItemIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") <= ?")
		values = append(values, f.ItemIDMinLte)
	}

	if f.ItemIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") <= ?")
		values = append(values, f.ItemIDMaxLte)
	}

	if f.ItemIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") IN (?)")
		values = append(values, f.ItemIDMinIn)
	}

	if f.ItemIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") IN (?)")
		values = append(values, f.ItemIDMaxIn)
	}

	if f.ItemIDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("itemId")+") NOT IN (?)")
		values = append(values, f.ItemIDMinNotIn)
	}

	if f.ItemIDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("itemId")+") NOT IN (?)")
		values = append(values, f.ItemIDMaxNotIn)
	}

	if f.DefinitionIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") = ?")
		values = append(values, f.DefinitionIDMin)
	}

	if f.DefinitionIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") = ?")
		values = append(values, f.DefinitionIDMax)
	}

	if f.DefinitionIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") != ?")
		values = append(values, f.DefinitionIDMinNe)
	}

	if f.DefinitionIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") != ?")
		values = append(values, f.DefinitionIDMaxNe)
	}

	if f.DefinitionIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") > ?")
		values = append(values, f.DefinitionIDMinGt)
	}

	if f.DefinitionIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") > ?")
		values = append(values, f.DefinitionIDMaxGt)
	}

	if f.DefinitionIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") < ?")
		values = append(values, f.DefinitionIDMinLt)
	}

	if f.DefinitionIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") < ?")
		values = append(values, f.DefinitionIDMaxLt)
	}

	if f.DefinitionIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") >= ?")
		values = append(values, f.DefinitionIDMinGte)
	}

	if f.DefinitionIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") >= ?")
		values = append(values, f.DefinitionIDMaxGte)
	}

	if f.DefinitionIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") <= ?")
		values = append(values, f.DefinitionIDMinLte)
	}

	if f.DefinitionIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") <= ?")
		values = append(values, f.DefinitionIDMaxLte)
	}

	if f.DefinitionIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") IN (?)")
		values = append(values, f.DefinitionIDMinIn)
	}

	if f.DefinitionIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") IN (?)")
		values = append(values, f.DefinitionIDMaxIn)
	}

	if f.DefinitionIDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("definitionId")+") NOT IN (?)")
		values = append(values, f.DefinitionIDMinNotIn)
	}

	if f.DefinitionIDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("definitionId")+") NOT IN (?)")
		values = append(values, f.DefinitionIDMaxNotIn)
	}

	if f.ParentItemIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("parentItemId")+") = ?")
		values = append(values, f.ParentItemIDMin)
	}

	if f.ParentItemIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("parentItemId")+") = ?")
		values = append(values, f.ParentItemIDMax)
	}

	if f.ParentItemIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("parentItemId")+") != ?")
		values = append(values, f.ParentItemIDMinNe)
	}

	if f.ParentItemIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("parentItemId")+") != ?")
		values = append(values, f.ParentItemIDMaxNe)
	}

	if f.ParentItemIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("parentItemId")+") > ?")
		values = append(values, f.ParentItemIDMinGt)
	}

	if f.ParentItemIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("parentItemId")+") > ?")
		values = append(values, f.ParentItemIDMaxGt)
	}

	if f.ParentItemIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("parentItemId")+") < ?")
		values = append(values, f.ParentItemIDMinLt)
	}

	if f.ParentItemIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("parentItemId")+") < ?")
		values = append(values, f.ParentItemIDMaxLt)
	}

	if f.ParentItemIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("parentItemId")+") >= ?")
		values = append(values, f.ParentItemIDMinGte)
	}

	if f.ParentItemIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("parentItemId")+") >= ?")
		values = append(values, f.ParentItemIDMaxGte)
	}

	if f.ParentItemIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("parentItemId")+") <= ?")
		values = append(values, f.ParentItemIDMinLte)
	}

	if f.ParentItemIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("parentItemId")+") <= ?")
		values = append(values, f.ParentItemIDMaxLte)
	}

	if f.ParentItemIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("parentItemId")+") IN (?)")
		values = append(values, f.ParentItemIDMinIn)
	}

	if f.ParentItemIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("parentItemId")+") IN (?)")
		values = append(values, f.ParentItemIDMaxIn)
	}

	if f.ParentItemIDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("parentItemId")+") NOT IN (?)")
		values = append(values, f.ParentItemIDMinNotIn)
	}

	if f.ParentItemIDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("parentItemId")+") NOT IN (?)")
		values = append(values, f.ParentItemIDMaxNotIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.UpdatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMinNotIn)
	}

	if f.UpdatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMaxNotIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.CreatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMinNotIn)
	}

	if f.CreatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMaxNotIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.UpdatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMinNotIn)
	}

	if f.UpdatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMaxNotIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	if f.CreatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMinNotIn)
	}

	if f.CreatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMaxNotIn)
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
