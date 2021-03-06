package gen

import (
	"fmt"
	"reflect"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/mapstructure"
)

type ConfiguratorItemDefinitionCategoryResultType struct {
	EntityResultType
}

type ConfiguratorItemDefinitionCategory struct {
	ID        string     `json:"id" gorm:"column:id;primary_key"`
	Code      *string    `json:"code" gorm:"column:code;unique"`
	Name      *string    `json:"name" gorm:"column:name"`
	Type      *string    `json:"type" gorm:"column:type"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy *string    `json:"createdBy" gorm:"column:createdBy"`

	Definitions []*ConfiguratorItemDefinition `json:"definitions" gorm:"foreignkey:CategoryID"`
}

func (m *ConfiguratorItemDefinitionCategory) Is_Entity() {}

type ConfiguratorItemDefinitionCategoryChanges struct {
	ID        string
	Code      *string
	Name      *string
	Type      *string
	UpdatedAt *time.Time
	CreatedAt time.Time
	UpdatedBy *string
	CreatedBy *string

	DefinitionsIDs []*string
}

type ConfiguratorItemDefinitionResultType struct {
	EntityResultType
}

type ConfiguratorItemDefinition struct {
	ID         string     `json:"id" gorm:"column:id;primary_key"`
	Code       *string    `json:"code" gorm:"column:code;unique"`
	Name       *string    `json:"name" gorm:"column:name"`
	CategoryID *string    `json:"categoryId" gorm:"column:categoryId"`
	UpdatedAt  *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt  time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy  *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy  *string    `json:"createdBy" gorm:"column:createdBy"`

	Attributes []*ConfiguratorAttributeDefinition `json:"attributes" gorm:"many2many:configuratorAttributeDefinition_definitions;jointable_foreignkey:definition_id;association_jointable_foreignkey:attribute_id"`

	Slots []*ConfiguratorSlotDefinition `json:"slots" gorm:"foreignkey:DefinitionID"`

	Items []*ConfiguratorItem `json:"items" gorm:"foreignkey:DefinitionID"`

	AllowedInSlots []*ConfiguratorSlotDefinition `json:"allowedInSlots" gorm:"many2many:configuratorSlotDefinition_allowedItemDefinitions;jointable_foreignkey:allowedItemDefinition_id;association_jointable_foreignkey:allowedInSlot_id"`

	Category *ConfiguratorItemDefinitionCategory `json:"category"`
}

func (m *ConfiguratorItemDefinition) Is_Entity() {}

type ConfiguratorItemDefinitionChanges struct {
	ID         string
	Code       *string
	Name       *string
	CategoryID *string
	UpdatedAt  *time.Time
	CreatedAt  time.Time
	UpdatedBy  *string
	CreatedBy  *string

	AttributesIDs     []*string
	SlotsIDs          []*string
	ItemsIDs          []*string
	AllowedInSlotsIDs []*string
}

type ConfiguratorAttributeDefinitionResultType struct {
	EntityResultType
}

type ConfiguratorAttributeDefinition struct {
	ID        string                     `json:"id" gorm:"column:id;primary_key"`
	Name      *string                    `json:"name" gorm:"column:name"`
	Type      *ConfiguratorAttributeType `json:"type" gorm:"column:type"`
	Primary   *bool                      `json:"primary" gorm:"column:primary"`
	UpdatedAt *time.Time                 `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time                  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy *string                    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy *string                    `json:"createdBy" gorm:"column:createdBy"`

	Definitions []*ConfiguratorItemDefinition `json:"definitions" gorm:"many2many:configuratorAttributeDefinition_definitions;jointable_foreignkey:attribute_id;association_jointable_foreignkey:definition_id"`

	Attributes []*ConfiguratorAttribute `json:"attributes" gorm:"foreignkey:DefinitionID"`
}

func (m *ConfiguratorAttributeDefinition) Is_Entity() {}

type ConfiguratorAttributeDefinitionChanges struct {
	ID        string
	Name      *string
	Type      *ConfiguratorAttributeType
	Primary   *bool
	UpdatedAt *time.Time
	CreatedAt time.Time
	UpdatedBy *string
	CreatedBy *string

	DefinitionsIDs []*string
	AttributesIDs  []*string
}

type ConfiguratorAttributeDefinition_definitions struct {
	attribute_id  string
	definition_id string
}

func (ConfiguratorAttributeDefinition_definitions) TableName() string {
	return TableName("configuratorAttributeDefinition_definitions")
}

type ConfiguratorSlotDefinitionResultType struct {
	EntityResultType
}

type ConfiguratorSlotDefinition struct {
	ID           string     `json:"id" gorm:"column:id;primary_key"`
	Name         *string    `json:"name" gorm:"column:name"`
	MinCount     *float64   `json:"minCount" gorm:"column:minCount"`
	MaxCount     *float64   `json:"maxCount" gorm:"column:maxCount"`
	DefaultCount *float64   `json:"defaultCount" gorm:"column:defaultCount"`
	DefinitionID *string    `json:"definitionId" gorm:"column:definitionId"`
	UpdatedAt    *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy    *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy    *string    `json:"createdBy" gorm:"column:createdBy"`

	Definition *ConfiguratorItemDefinition `json:"definition"`

	Slots []*ConfiguratorSlot `json:"slots" gorm:"foreignkey:DefinitionID"`

	AllowedItemDefinitions []*ConfiguratorItemDefinition `json:"allowedItemDefinitions" gorm:"many2many:configuratorSlotDefinition_allowedItemDefinitions;jointable_foreignkey:allowedInSlot_id;association_jointable_foreignkey:allowedItemDefinition_id"`
}

func (m *ConfiguratorSlotDefinition) Is_Entity() {}

type ConfiguratorSlotDefinitionChanges struct {
	ID           string
	Name         *string
	MinCount     *float64
	MaxCount     *float64
	DefaultCount *float64
	DefinitionID *string
	UpdatedAt    *time.Time
	CreatedAt    time.Time
	UpdatedBy    *string
	CreatedBy    *string

	SlotsIDs                  []*string
	AllowedItemDefinitionsIDs []*string
}

type ConfiguratorSlotDefinition_allowedItemDefinitions struct {
	allowedInSlot_id         string
	allowedItemDefinition_id string
}

func (ConfiguratorSlotDefinition_allowedItemDefinitions) TableName() string {
	return TableName("configuratorSlotDefinition_allowedItemDefinitions")
}

type ConfiguratorItemResultType struct {
	EntityResultType
}

type ConfiguratorItem struct {
	ID           string     `json:"id" gorm:"column:id;primary_key"`
	Code         *string    `json:"code" gorm:"column:code"`
	Name         *string    `json:"name" gorm:"column:name"`
	StockItemID  *string    `json:"stockItemId" gorm:"column:stockItemId"`
	RawData      *string    `json:"rawData" gorm:"column:rawData;type:text"`
	DefinitionID *string    `json:"definitionId" gorm:"column:definitionId"`
	UpdatedAt    *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy    *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy    *string    `json:"createdBy" gorm:"column:createdBy"`

	Definition *ConfiguratorItemDefinition `json:"definition"`

	Attributes []*ConfiguratorAttribute `json:"attributes" gorm:"foreignkey:ItemID"`

	Slots []*ConfiguratorSlot `json:"slots" gorm:"foreignkey:ParentItemID"`

	ParentSlots []*ConfiguratorSlot `json:"parentSlots" gorm:"foreignkey:ItemID"`
}

func (m *ConfiguratorItem) Is_Entity() {}

type ConfiguratorItemChanges struct {
	ID           string
	Code         *string
	Name         *string
	StockItemID  *string
	RawData      *string
	DefinitionID *string
	UpdatedAt    *time.Time
	CreatedAt    time.Time
	UpdatedBy    *string
	CreatedBy    *string

	AttributesIDs  []*string
	SlotsIDs       []*string
	ParentSlotsIDs []*string
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
	Count        *float64   `json:"count" gorm:"column:count"`
	ItemID       *string    `json:"itemId" gorm:"column:itemId"`
	DefinitionID *string    `json:"definitionId" gorm:"column:definitionId"`
	ParentItemID *string    `json:"parentItemId" gorm:"column:parentItemId"`
	UpdatedAt    *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy    *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy    *string    `json:"createdBy" gorm:"column:createdBy"`

	Item *ConfiguratorItem `json:"item"`

	Definition *ConfiguratorSlotDefinition `json:"definition"`

	ParentItem *ConfiguratorItem `json:"parentItem"`
}

func (m *ConfiguratorSlot) Is_Entity() {}

type ConfiguratorSlotChanges struct {
	ID           string
	Count        *float64
	ItemID       *string
	DefinitionID *string
	ParentItemID *string
	UpdatedAt    *time.Time
	CreatedAt    time.Time
	UpdatedBy    *string
	CreatedBy    *string
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
