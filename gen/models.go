package gen

import (
	"fmt"
	"reflect"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/mapstructure"
)

type NotFoundError struct {
	Entity string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Entity)
}

type ConfiguratorItemDefinitionResultType struct {
	EntityResultType
}

type ConfiguratorItemDefinition struct {
	ID        string     `json:"id" gorm:"column:id;primary_key"`
	Name      *string    `json:"name" gorm:"column:name"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy *string    `json:"createdBy" gorm:"column:createdBy"`

	Attributes []*ConfiguratorAttributeDefinition `json:"attributes" gorm:"many2many:configuratorAttributeDefinition_itemDefinitions;jointable_foreignkey:itemDefinition_id;association_jointable_foreignkey:attribute_id"`

	Slots []*ConfiguratorSlotDefinition `json:"slots" gorm:"foreignkey:ItemDefinitionID"`

	Items []*ConfiguratorItem `json:"items" gorm:"foreignkey:DefinitionID"`
}

func (m *ConfiguratorItemDefinition) Is_Entity() {}

type ConfiguratorItemDefinitionChanges struct {
	ID        string
	Name      *string
	UpdatedAt *time.Time
	CreatedAt time.Time
	UpdatedBy *string
	CreatedBy *string

	AttributesIDs []*string
	SlotsIDs      []*string
	ItemsIDs      []*string
}

type ConfiguratorAttributeDefinitionResultType struct {
	EntityResultType
}

type ConfiguratorAttributeDefinition struct {
	ID        string                     `json:"id" gorm:"column:id;primary_key"`
	Name      *string                    `json:"name" gorm:"column:name"`
	Type      *ConfiguratorAttributeType `json:"type" gorm:"column:type"`
	UpdatedAt *time.Time                 `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time                  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy *string                    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy *string                    `json:"createdBy" gorm:"column:createdBy"`

	ItemDefinitions []*ConfiguratorItemDefinition `json:"itemDefinitions" gorm:"many2many:configuratorAttributeDefinition_itemDefinitions;jointable_foreignkey:attribute_id;association_jointable_foreignkey:itemDefinition_id"`

	Attributes []*ConfiguratorAttribute `json:"attributes" gorm:"foreignkey:DefinitionID"`
}

func (m *ConfiguratorAttributeDefinition) Is_Entity() {}

type ConfiguratorAttributeDefinitionChanges struct {
	ID        string
	Name      *string
	Type      *ConfiguratorAttributeType
	UpdatedAt *time.Time
	CreatedAt time.Time
	UpdatedBy *string
	CreatedBy *string

	ItemDefinitionsIDs []*string
	AttributesIDs      []*string
}

type ConfiguratorSlotDefinitionResultType struct {
	EntityResultType
}

type ConfiguratorSlotDefinition struct {
	ID               string     `json:"id" gorm:"column:id;primary_key"`
	Name             *string    `json:"name" gorm:"column:name"`
	ItemDefinitionID *string    `json:"itemDefinitionId" gorm:"column:itemDefinitionId"`
	UpdatedAt        *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt        time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy        *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy        *string    `json:"createdBy" gorm:"column:createdBy"`

	ItemDefinition *ConfiguratorItemDefinition `json:"itemDefinition"`

	Slots []*ConfiguratorSlot `json:"slots" gorm:"foreignkey:DefinitionID"`
}

func (m *ConfiguratorSlotDefinition) Is_Entity() {}

type ConfiguratorSlotDefinitionChanges struct {
	ID               string
	Name             *string
	ItemDefinitionID *string
	UpdatedAt        *time.Time
	CreatedAt        time.Time
	UpdatedBy        *string
	CreatedBy        *string

	SlotsIDs []*string
}

type ConfiguratorItemResultType struct {
	EntityResultType
}

type ConfiguratorItem struct {
	ID           string     `json:"id" gorm:"column:id;primary_key"`
	Code         *string    `json:"code" gorm:"column:code"`
	Name         *string    `json:"name" gorm:"column:name"`
	StockItemID  *string    `json:"stockItemID" gorm:"column:stockItemID"`
	DefinitionID *string    `json:"definitionId" gorm:"column:definitionId"`
	ParentSlotID *string    `json:"parentSlotId" gorm:"column:parentSlotId"`
	UpdatedAt    *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy    *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy    *string    `json:"createdBy" gorm:"column:createdBy"`

	Definition *ConfiguratorItemDefinition `json:"definition"`

	Attributes []*ConfiguratorAttribute `json:"attributes" gorm:"foreignkey:ItemID"`

	Slots []*ConfiguratorSlot `json:"slots" gorm:"foreignkey:ParentItemID"`

	ParentSlot *ConfiguratorSlot `json:"parentSlot"`
}

func (m *ConfiguratorItem) Is_Entity() {}

type ConfiguratorItemChanges struct {
	ID           string
	Code         *string
	Name         *string
	StockItemID  *string
	DefinitionID *string
	ParentSlotID *string
	UpdatedAt    *time.Time
	CreatedAt    time.Time
	UpdatedBy    *string
	CreatedBy    *string

	AttributesIDs []*string
	SlotsIDs      []*string
}

type ConfiguratorAttributeResultType struct {
	EntityResultType
}

type ConfiguratorAttribute struct {
	ID           string     `json:"id" gorm:"column:id;primary_key"`
	StringValue  *string    `json:"stringValue" gorm:"column:stringValue"`
	FloatValue   *float64   `json:"floatValue" gorm:"column:floatValue"`
	IntValue     *int       `json:"intValue" gorm:"column:intValue"`
	DefinitionID *string    `json:"definitionId" gorm:"column:definitionId"`
	ItemID       *string    `json:"itemId" gorm:"column:itemId"`
	UpdatedAt    *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy    *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy    *string    `json:"createdBy" gorm:"column:createdBy"`

	Definition *ConfiguratorAttributeDefinition `json:"definition"`

	Item *ConfiguratorItem `json:"item"`
}

func (m *ConfiguratorAttribute) Is_Entity() {}

type ConfiguratorAttributeChanges struct {
	ID           string
	StringValue  *string
	FloatValue   *float64
	IntValue     *int
	DefinitionID *string
	ItemID       *string
	UpdatedAt    *time.Time
	CreatedAt    time.Time
	UpdatedBy    *string
	CreatedBy    *string
}

type ConfiguratorSlotResultType struct {
	EntityResultType
}

type ConfiguratorSlot struct {
	ID           string     `json:"id" gorm:"column:id;primary_key"`
	DefinitionID *string    `json:"definitionId" gorm:"column:definitionId"`
	ParentItemID *string    `json:"parentItemId" gorm:"column:parentItemId"`
	UpdatedAt    *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy    *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy    *string    `json:"createdBy" gorm:"column:createdBy"`

	Items []*ConfiguratorItem `json:"items" gorm:"foreignkey:ParentSlotID"`

	Definition *ConfiguratorSlotDefinition `json:"definition"`

	ParentItem *ConfiguratorItem `json:"parentItem"`
}

func (m *ConfiguratorSlot) Is_Entity() {}

type ConfiguratorSlotChanges struct {
	ID           string
	DefinitionID *string
	ParentItemID *string
	UpdatedAt    *time.Time
	CreatedAt    time.Time
	UpdatedBy    *string
	CreatedBy    *string

	ItemsIDs []*string
}

// used to convert map[string]interface{} to EntityChanges struct
func ApplyChanges(changes map[string]interface{}, to interface{}) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
		// This is needed to get mapstructure to call the gqlgen unmarshaler func for custom scalars (eg Date)
		DecodeHook: func(a reflect.Type, b reflect.Type, v interface{}) (interface{}, error) {

			if b == reflect.TypeOf(time.Time{}) {
				switch a.Kind() {
				case reflect.String:
					return time.Parse(time.RFC3339, v.(string))
				case reflect.Float64:
					return time.Unix(0, int64(v.(float64))*int64(time.Millisecond)), nil
				case reflect.Int64:
					return time.Unix(0, v.(int64)*int64(time.Millisecond)), nil
				default:
					return v, fmt.Errorf("Unable to parse date from %v", v)
				}
			}

			if reflect.PtrTo(b).Implements(reflect.TypeOf((*graphql.Unmarshaler)(nil)).Elem()) {
				resultType := reflect.New(b)
				result := resultType.MethodByName("UnmarshalGQL").Call([]reflect.Value{reflect.ValueOf(v)})
				err, _ := result[0].Interface().(error)
				return resultType.Elem().Interface(), err
			}

			return v, nil
		},
	})

	if err != nil {
		return err
	}

	return dec.Decode(changes)
}
