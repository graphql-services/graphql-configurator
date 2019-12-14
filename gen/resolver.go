package gen

import (
	"context"

	"github.com/novacloudcz/graphql-orm/events"
)

type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *events.Event) error

	CreateConfiguratorItemDefinition     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ConfiguratorItemDefinition, err error)
	UpdateConfiguratorItemDefinition     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ConfiguratorItemDefinition, err error)
	DeleteConfiguratorItemDefinition     func(ctx context.Context, r *GeneratedResolver, id string) (item *ConfiguratorItemDefinition, err error)
	DeleteAllConfiguratorItemDefinitions func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryConfiguratorItemDefinition      func(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorItemDefinitionHandlerOptions) (*ConfiguratorItemDefinition, error)
	QueryConfiguratorItemDefinitions     func(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorItemDefinitionsHandlerOptions) (*ConfiguratorItemDefinitionResultType, error)

	ConfiguratorItemDefinitionAttributes func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorItemDefinition) (res []*ConfiguratorAttributeDefinition, err error)

	ConfiguratorItemDefinitionSlots func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorItemDefinition) (res []*ConfiguratorSlotDefinition, err error)

	ConfiguratorItemDefinitionItems func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorItemDefinition) (res []*ConfiguratorItem, err error)

	CreateConfiguratorAttributeDefinition     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ConfiguratorAttributeDefinition, err error)
	UpdateConfiguratorAttributeDefinition     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ConfiguratorAttributeDefinition, err error)
	DeleteConfiguratorAttributeDefinition     func(ctx context.Context, r *GeneratedResolver, id string) (item *ConfiguratorAttributeDefinition, err error)
	DeleteAllConfiguratorAttributeDefinitions func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryConfiguratorAttributeDefinition      func(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorAttributeDefinitionHandlerOptions) (*ConfiguratorAttributeDefinition, error)
	QueryConfiguratorAttributeDefinitions     func(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorAttributeDefinitionsHandlerOptions) (*ConfiguratorAttributeDefinitionResultType, error)

	ConfiguratorAttributeDefinitionDefinitions func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorAttributeDefinition) (res []*ConfiguratorItemDefinition, err error)

	ConfiguratorAttributeDefinitionAttributes func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorAttributeDefinition) (res []*ConfiguratorAttribute, err error)

	CreateConfiguratorSlotDefinition     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ConfiguratorSlotDefinition, err error)
	UpdateConfiguratorSlotDefinition     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ConfiguratorSlotDefinition, err error)
	DeleteConfiguratorSlotDefinition     func(ctx context.Context, r *GeneratedResolver, id string) (item *ConfiguratorSlotDefinition, err error)
	DeleteAllConfiguratorSlotDefinitions func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryConfiguratorSlotDefinition      func(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorSlotDefinitionHandlerOptions) (*ConfiguratorSlotDefinition, error)
	QueryConfiguratorSlotDefinitions     func(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorSlotDefinitionsHandlerOptions) (*ConfiguratorSlotDefinitionResultType, error)

	ConfiguratorSlotDefinitionDefinition func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorSlotDefinition) (res *ConfiguratorItemDefinition, err error)

	ConfiguratorSlotDefinitionSlots func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorSlotDefinition) (res []*ConfiguratorSlot, err error)

	CreateConfiguratorItem     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ConfiguratorItem, err error)
	UpdateConfiguratorItem     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ConfiguratorItem, err error)
	DeleteConfiguratorItem     func(ctx context.Context, r *GeneratedResolver, id string) (item *ConfiguratorItem, err error)
	DeleteAllConfiguratorItems func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryConfiguratorItem      func(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorItemHandlerOptions) (*ConfiguratorItem, error)
	QueryConfiguratorItems     func(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorItemsHandlerOptions) (*ConfiguratorItemResultType, error)

	ConfiguratorItemDefinition func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorItem) (res *ConfiguratorItemDefinition, err error)

	ConfiguratorItemAttributes func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorItem) (res []*ConfiguratorAttribute, err error)

	ConfiguratorItemSlots func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorItem) (res []*ConfiguratorSlot, err error)

	ConfiguratorItemParentSlot func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorItem) (res *ConfiguratorSlot, err error)

	CreateConfiguratorAttribute     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ConfiguratorAttribute, err error)
	UpdateConfiguratorAttribute     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ConfiguratorAttribute, err error)
	DeleteConfiguratorAttribute     func(ctx context.Context, r *GeneratedResolver, id string) (item *ConfiguratorAttribute, err error)
	DeleteAllConfiguratorAttributes func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryConfiguratorAttribute      func(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorAttributeHandlerOptions) (*ConfiguratorAttribute, error)
	QueryConfiguratorAttributes     func(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorAttributesHandlerOptions) (*ConfiguratorAttributeResultType, error)

	ConfiguratorAttributeDefinition func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorAttribute) (res *ConfiguratorAttributeDefinition, err error)

	ConfiguratorAttributeItem func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorAttribute) (res *ConfiguratorItem, err error)

	CreateConfiguratorSlot     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ConfiguratorSlot, err error)
	UpdateConfiguratorSlot     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ConfiguratorSlot, err error)
	DeleteConfiguratorSlot     func(ctx context.Context, r *GeneratedResolver, id string) (item *ConfiguratorSlot, err error)
	DeleteAllConfiguratorSlots func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryConfiguratorSlot      func(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorSlotHandlerOptions) (*ConfiguratorSlot, error)
	QueryConfiguratorSlots     func(ctx context.Context, r *GeneratedResolver, opts QueryConfiguratorSlotsHandlerOptions) (*ConfiguratorSlotResultType, error)

	ConfiguratorSlotItems func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorSlot) (res []*ConfiguratorItem, err error)

	ConfiguratorSlotDefinition func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorSlot) (res *ConfiguratorSlotDefinition, err error)

	ConfiguratorSlotParentItem func(ctx context.Context, r *GeneratedResolver, obj *ConfiguratorSlot) (res *ConfiguratorItem, err error)
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *events.Event) error { return nil },

		CreateConfiguratorItemDefinition:     CreateConfiguratorItemDefinitionHandler,
		UpdateConfiguratorItemDefinition:     UpdateConfiguratorItemDefinitionHandler,
		DeleteConfiguratorItemDefinition:     DeleteConfiguratorItemDefinitionHandler,
		DeleteAllConfiguratorItemDefinitions: DeleteAllConfiguratorItemDefinitionsHandler,
		QueryConfiguratorItemDefinition:      QueryConfiguratorItemDefinitionHandler,
		QueryConfiguratorItemDefinitions:     QueryConfiguratorItemDefinitionsHandler,

		ConfiguratorItemDefinitionAttributes: ConfiguratorItemDefinitionAttributesHandler,

		ConfiguratorItemDefinitionSlots: ConfiguratorItemDefinitionSlotsHandler,

		ConfiguratorItemDefinitionItems: ConfiguratorItemDefinitionItemsHandler,

		CreateConfiguratorAttributeDefinition:     CreateConfiguratorAttributeDefinitionHandler,
		UpdateConfiguratorAttributeDefinition:     UpdateConfiguratorAttributeDefinitionHandler,
		DeleteConfiguratorAttributeDefinition:     DeleteConfiguratorAttributeDefinitionHandler,
		DeleteAllConfiguratorAttributeDefinitions: DeleteAllConfiguratorAttributeDefinitionsHandler,
		QueryConfiguratorAttributeDefinition:      QueryConfiguratorAttributeDefinitionHandler,
		QueryConfiguratorAttributeDefinitions:     QueryConfiguratorAttributeDefinitionsHandler,

		ConfiguratorAttributeDefinitionDefinitions: ConfiguratorAttributeDefinitionDefinitionsHandler,

		ConfiguratorAttributeDefinitionAttributes: ConfiguratorAttributeDefinitionAttributesHandler,

		CreateConfiguratorSlotDefinition:     CreateConfiguratorSlotDefinitionHandler,
		UpdateConfiguratorSlotDefinition:     UpdateConfiguratorSlotDefinitionHandler,
		DeleteConfiguratorSlotDefinition:     DeleteConfiguratorSlotDefinitionHandler,
		DeleteAllConfiguratorSlotDefinitions: DeleteAllConfiguratorSlotDefinitionsHandler,
		QueryConfiguratorSlotDefinition:      QueryConfiguratorSlotDefinitionHandler,
		QueryConfiguratorSlotDefinitions:     QueryConfiguratorSlotDefinitionsHandler,

		ConfiguratorSlotDefinitionDefinition: ConfiguratorSlotDefinitionDefinitionHandler,

		ConfiguratorSlotDefinitionSlots: ConfiguratorSlotDefinitionSlotsHandler,

		CreateConfiguratorItem:     CreateConfiguratorItemHandler,
		UpdateConfiguratorItem:     UpdateConfiguratorItemHandler,
		DeleteConfiguratorItem:     DeleteConfiguratorItemHandler,
		DeleteAllConfiguratorItems: DeleteAllConfiguratorItemsHandler,
		QueryConfiguratorItem:      QueryConfiguratorItemHandler,
		QueryConfiguratorItems:     QueryConfiguratorItemsHandler,

		ConfiguratorItemDefinition: ConfiguratorItemDefinitionHandler,

		ConfiguratorItemAttributes: ConfiguratorItemAttributesHandler,

		ConfiguratorItemSlots: ConfiguratorItemSlotsHandler,

		ConfiguratorItemParentSlot: ConfiguratorItemParentSlotHandler,

		CreateConfiguratorAttribute:     CreateConfiguratorAttributeHandler,
		UpdateConfiguratorAttribute:     UpdateConfiguratorAttributeHandler,
		DeleteConfiguratorAttribute:     DeleteConfiguratorAttributeHandler,
		DeleteAllConfiguratorAttributes: DeleteAllConfiguratorAttributesHandler,
		QueryConfiguratorAttribute:      QueryConfiguratorAttributeHandler,
		QueryConfiguratorAttributes:     QueryConfiguratorAttributesHandler,

		ConfiguratorAttributeDefinition: ConfiguratorAttributeDefinitionHandler,

		ConfiguratorAttributeItem: ConfiguratorAttributeItemHandler,

		CreateConfiguratorSlot:     CreateConfiguratorSlotHandler,
		UpdateConfiguratorSlot:     UpdateConfiguratorSlotHandler,
		DeleteConfiguratorSlot:     DeleteConfiguratorSlotHandler,
		DeleteAllConfiguratorSlots: DeleteAllConfiguratorSlotsHandler,
		QueryConfiguratorSlot:      QueryConfiguratorSlotHandler,
		QueryConfiguratorSlots:     QueryConfiguratorSlotsHandler,

		ConfiguratorSlotItems: ConfiguratorSlotItemsHandler,

		ConfiguratorSlotDefinition: ConfiguratorSlotDefinitionHandler,

		ConfiguratorSlotParentItem: ConfiguratorSlotParentItemHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *events.EventController
}
