package gen

import (
	"context"
	"strings"

	"github.com/vektah/gqlparser/ast"

	"github.com/jinzhu/gorm"
)

func GetItem(ctx context.Context, db *gorm.DB, out interface{}, id *string) error {
	return db.Find(out, "id = ?", id).Error
}

func GetItemForRelation(ctx context.Context, db *gorm.DB, obj interface{}, relation string, out interface{}) error {
	return db.Model(obj).Related(out, relation).Error
}

type EntityFilter interface {
	Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error
}
type EntityFilterQuery interface {
	Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error
}
type EntitySort interface {
	Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error
}

type EntityResultType struct {
	Offset       *int
	Limit        *int
	Query        EntityFilterQuery
	Sort         []EntitySort
	Filter       EntityFilter
	Fields       []*ast.Field
	SelectionSet *ast.SelectionSet
}

type GetItemsOptions struct {
	Alias      string
	Preloaders []string
}

type CountResult struct {
	Count int
}

// GetResultTypeItems ...
func (r *EntityResultType) GetItems(ctx context.Context, db *gorm.DB, opts GetItemsOptions, out interface{}) error {
	q := db
	filterQuery := db.Table(opts.Alias).Select("DISTINCT " + opts.Alias + ".id")

	dialect := q.Dialect()

	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	sorts := []string{}

	err := r.Query.Apply(ctx, dialect, r.SelectionSet, &wheres, &values, &joins)
	if err != nil {
		return err
	}

	for _, sort := range r.Sort {
		sort.Apply(ctx, dialect, &sorts, &joins)
	}

	if r.Filter != nil {
		err = r.Filter.Apply(ctx, dialect, &wheres, &values, &joins)
		if err != nil {
			return err
		}
	}

	if len(sorts) > 0 {
		q = q.Order(strings.Join(sorts, ", "))
	}

	if r.Limit != nil {
		filterQuery = filterQuery.Limit(*r.Limit)
	}
	if r.Offset != nil {
		filterQuery = filterQuery.Offset(*r.Offset)
	}
	if len(wheres) > 0 {
		filterQuery = filterQuery.Where(strings.Join(wheres, " AND "), values...)
	}

	uniqueJoinsMap := map[string]bool{}
	uniqueJoins := []string{}
	for _, join := range joins {
		if !uniqueJoinsMap[join] {
			uniqueJoinsMap[join] = true
			uniqueJoins = append(uniqueJoins, join)
		}
	}

	for _, join := range uniqueJoins {
		filterQuery = filterQuery.Joins(join)
		q = q.Joins(join)
	}

	if len(opts.Preloaders) > 0 {
		for _, p := range opts.Preloaders {
			q = q.Preload(p)
		}
	}
	// q = q.Group(opts.Alias + ".id")
	return q.Joins("INNER JOIN (?) as filter ON "+opts.Alias+".id = filter.id", filterQuery.QueryExpr()).Find(out).Error
}

// GetCount ...
func (r *EntityResultType) GetCount(ctx context.Context, db *gorm.DB, out interface{}) (count int, err error) {
	q := db.Model(out).Select(db.NewScope(out).TableName() + ".id")

	dialect := q.Dialect()
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}

	err = r.Query.Apply(ctx, dialect, r.SelectionSet, &wheres, &values, &joins)
	if err != nil {
		return 0, err
	}

	if r.Filter != nil {
		err = r.Filter.Apply(ctx, dialect, &wheres, &values, &joins)
		if err != nil {
			return 0, err
		}
	}

	if len(wheres) > 0 {
		q = q.Where(strings.Join(wheres, " AND "), values...)
	}

	uniqueJoinsMap := map[string]bool{}
	uniqueJoins := []string{}
	for _, join := range joins {
		if !uniqueJoinsMap[join] {
			uniqueJoinsMap[join] = true
			uniqueJoins = append(uniqueJoins, join)
		}
	}

	for _, join := range uniqueJoins {
		q = q.Joins(join)
	}

	var result CountResult
	err = q.Select("DISTINCT COUNT(" + db.NewScope(out).TableName() + ".id) as count").Scan(&result).Error
	count = result.Count
	return
}

func (r *EntityResultType) GetSortStrings() []string {
	return []string{}
}
