package gen

type key int

const (
	KeyPrincipalID         key    = iota
	KeyLoaders             key    = iota
	KeyExecutableSchema    key    = iota
	KeyJWTClaims           key    = iota
	KeyMutationTransaction key    = iota
	KeyMutationEvents      key    = iota
	SchemaSDL              string = `scalar Time

type Query {
  configuratorItemDefinition(id: ID, q: String, filter: ConfiguratorItemDefinitionFilterType): ConfiguratorItemDefinition
  configuratorItemDefinitions(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorItemDefinitionSortType!], filter: ConfiguratorItemDefinitionFilterType): ConfiguratorItemDefinitionResultType
  configuratorAttributeDefinition(id: ID, q: String, filter: ConfiguratorAttributeDefinitionFilterType): ConfiguratorAttributeDefinition
  configuratorAttributeDefinitions(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorAttributeDefinitionSortType!], filter: ConfiguratorAttributeDefinitionFilterType): ConfiguratorAttributeDefinitionResultType
  configuratorSlotDefinition(id: ID, q: String, filter: ConfiguratorSlotDefinitionFilterType): ConfiguratorSlotDefinition
  configuratorSlotDefinitions(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorSlotDefinitionSortType!], filter: ConfiguratorSlotDefinitionFilterType): ConfiguratorSlotDefinitionResultType
  configuratorItem(id: ID, q: String, filter: ConfiguratorItemFilterType): ConfiguratorItem
  configuratorItems(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorItemSortType!], filter: ConfiguratorItemFilterType): ConfiguratorItemResultType
  configuratorAttribute(id: ID, q: String, filter: ConfiguratorAttributeFilterType): ConfiguratorAttribute
  configuratorAttributes(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorAttributeSortType!], filter: ConfiguratorAttributeFilterType): ConfiguratorAttributeResultType
  configuratorSlot(id: ID, q: String, filter: ConfiguratorSlotFilterType): ConfiguratorSlot
  configuratorSlots(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorSlotSortType!], filter: ConfiguratorSlotFilterType): ConfiguratorSlotResultType
}

type Mutation {
  createConfiguratorItemDefinition(input: ConfiguratorItemDefinitionCreateInput!): ConfiguratorItemDefinition!
  updateConfiguratorItemDefinition(id: ID!, input: ConfiguratorItemDefinitionUpdateInput!): ConfiguratorItemDefinition!
  deleteConfiguratorItemDefinition(id: ID!): ConfiguratorItemDefinition!
  deleteAllConfiguratorItemDefinitions: Boolean!
  createConfiguratorAttributeDefinition(input: ConfiguratorAttributeDefinitionCreateInput!): ConfiguratorAttributeDefinition!
  updateConfiguratorAttributeDefinition(id: ID!, input: ConfiguratorAttributeDefinitionUpdateInput!): ConfiguratorAttributeDefinition!
  deleteConfiguratorAttributeDefinition(id: ID!): ConfiguratorAttributeDefinition!
  deleteAllConfiguratorAttributeDefinitions: Boolean!
  createConfiguratorSlotDefinition(input: ConfiguratorSlotDefinitionCreateInput!): ConfiguratorSlotDefinition!
  updateConfiguratorSlotDefinition(id: ID!, input: ConfiguratorSlotDefinitionUpdateInput!): ConfiguratorSlotDefinition!
  deleteConfiguratorSlotDefinition(id: ID!): ConfiguratorSlotDefinition!
  deleteAllConfiguratorSlotDefinitions: Boolean!
  createConfiguratorItem(input: ConfiguratorItemCreateInput!): ConfiguratorItem!
  updateConfiguratorItem(id: ID!, input: ConfiguratorItemUpdateInput!): ConfiguratorItem!
  deleteConfiguratorItem(id: ID!): ConfiguratorItem!
  deleteAllConfiguratorItems: Boolean!
  createConfiguratorAttribute(input: ConfiguratorAttributeCreateInput!): ConfiguratorAttribute!
  updateConfiguratorAttribute(id: ID!, input: ConfiguratorAttributeUpdateInput!): ConfiguratorAttribute!
  deleteConfiguratorAttribute(id: ID!): ConfiguratorAttribute!
  deleteAllConfiguratorAttributes: Boolean!
  createConfiguratorSlot(input: ConfiguratorSlotCreateInput!): ConfiguratorSlot!
  updateConfiguratorSlot(id: ID!, input: ConfiguratorSlotUpdateInput!): ConfiguratorSlot!
  deleteConfiguratorSlot(id: ID!): ConfiguratorSlot!
  deleteAllConfiguratorSlots: Boolean!
}

enum ObjectSortType {
  ASC
  DESC
}

type ConfiguratorAssembly {
  id: ID!
  item: ConfiguratorAssemblyItem
}

type ConfiguratorAssemblyItem {
  id: ID!
  definitionId: ID!
  slots: [ConfiguratorAssemblySlot!]!
  attributes: [ConfiguratorAssemblyAttribute!]!
}

type ConfiguratorAssemblyAttribute {
  id: ID!
  definitionId: ID!
  stringValue: String
  intValue: Int
  floatValue: Float
}

type ConfiguratorAssemblySlot {
  id: ID!
  definitionId: ID!
  items: [ConfiguratorAssemblyItem!]!
}

input ConfiguratorAssemblyInput {
  item: ConfiguratorAssemblyItemInput
}

input ConfiguratorAssemblyItemInput {
  id: ID
  definitionId: ID!
  slots: [ConfiguratorAssemblySlotInput!]
  attributes: [ConfiguratorAssemblyAttributeInput!]
}

input ConfiguratorAssemblyAttributeInput {
  id: ID
  definitionId: ID!
  stringValue: String
  intValue: Int
  floatValue: Float
}

input ConfiguratorAssemblySlotInput {
  id: ID
  definitionId: ID!
  items: [ConfiguratorAssemblyItemInput!]!
}

extend type Query {
  configuratorAssembly(id: ID!): ConfiguratorAssembly
}

extend type Mutation {
  createConfiguratorAssembly(input: ConfiguratorAssemblyInput!): ConfiguratorAssembly
}

type ConfiguratorItemDefinition {
  id: ID!
  name: String
  attributes: [ConfiguratorAttributeDefinition!]!
  slots: [ConfiguratorSlotDefinition!]!
  items: [ConfiguratorItem!]!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  attributesIds: [ID!]!
  slotsIds: [ID!]!
  itemsIds: [ID!]!
}

enum ConfiguratorAttributeType {
  STRING
  FLOAT
  INT
}

type ConfiguratorAttributeDefinition {
  id: ID!
  name: String
  type: ConfiguratorAttributeType
  itemDefinitions: [ConfiguratorItemDefinition!]!
  attributes: [ConfiguratorAttribute!]!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  itemDefinitionsIds: [ID!]!
  attributesIds: [ID!]!
}

type ConfiguratorSlotDefinition {
  id: ID!
  name: String
  itemDefinition: ConfiguratorItemDefinition
  slots: [ConfiguratorSlot!]!
  itemDefinitionId: ID
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  slotsIds: [ID!]!
}

type ConfiguratorItem {
  id: ID!
  stockItemID: ID
  definition: ConfiguratorItemDefinition
  attributes: [ConfiguratorAttribute!]!
  slots: [ConfiguratorSlot!]!
  parentSlot: ConfiguratorSlot
  definitionId: ID
  parentSlotId: ID
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  attributesIds: [ID!]!
  slotsIds: [ID!]!
}

type ConfiguratorAttribute {
  id: ID!
  stringValue: String
  floatValue: Float
  intValue: Int
  definition: ConfiguratorAttributeDefinition
  item: ConfiguratorItem
  definitionId: ID
  itemId: ID
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
}

type ConfiguratorSlot {
  id: ID!
  items: [ConfiguratorItem!]!
  definition: ConfiguratorSlotDefinition
  parentItem: ConfiguratorItem
  definitionId: ID
  parentItemId: ID
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  itemsIds: [ID!]!
}

input ConfiguratorItemDefinitionCreateInput {
  id: ID
  name: String
  attributesIds: [ID!]
  slotsIds: [ID!]
  itemsIds: [ID!]
}

input ConfiguratorItemDefinitionUpdateInput {
  name: String
  attributesIds: [ID!]
  slotsIds: [ID!]
  itemsIds: [ID!]
}

input ConfiguratorItemDefinitionSortType {
  id: ObjectSortType
  name: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  attributesIds: ObjectSortType
  slotsIds: ObjectSortType
  itemsIds: ObjectSortType
  attributes: ConfiguratorAttributeDefinitionSortType
  slots: ConfiguratorSlotDefinitionSortType
  items: ConfiguratorItemSortType
}

input ConfiguratorItemDefinitionFilterType {
  AND: [ConfiguratorItemDefinitionFilterType!]
  OR: [ConfiguratorItemDefinitionFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  id_null: Boolean
  name: String
  name_ne: String
  name_gt: String
  name_lt: String
  name_gte: String
  name_lte: String
  name_in: [String!]
  name_like: String
  name_prefix: String
  name_suffix: String
  name_null: Boolean
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  createdBy_null: Boolean
  attributes: ConfiguratorAttributeDefinitionFilterType
  slots: ConfiguratorSlotDefinitionFilterType
  items: ConfiguratorItemFilterType
}

type ConfiguratorItemDefinitionResultType {
  items: [ConfiguratorItemDefinition!]!
  count: Int!
}

input ConfiguratorAttributeDefinitionCreateInput {
  id: ID
  name: String
  type: ConfiguratorAttributeType
  itemDefinitionsIds: [ID!]
  attributesIds: [ID!]
}

input ConfiguratorAttributeDefinitionUpdateInput {
  name: String
  type: ConfiguratorAttributeType
  itemDefinitionsIds: [ID!]
  attributesIds: [ID!]
}

input ConfiguratorAttributeDefinitionSortType {
  id: ObjectSortType
  name: ObjectSortType
  type: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  itemDefinitionsIds: ObjectSortType
  attributesIds: ObjectSortType
  itemDefinitions: ConfiguratorItemDefinitionSortType
  attributes: ConfiguratorAttributeSortType
}

input ConfiguratorAttributeDefinitionFilterType {
  AND: [ConfiguratorAttributeDefinitionFilterType!]
  OR: [ConfiguratorAttributeDefinitionFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  id_null: Boolean
  name: String
  name_ne: String
  name_gt: String
  name_lt: String
  name_gte: String
  name_lte: String
  name_in: [String!]
  name_like: String
  name_prefix: String
  name_suffix: String
  name_null: Boolean
  type: ConfiguratorAttributeType
  type_ne: ConfiguratorAttributeType
  type_gt: ConfiguratorAttributeType
  type_lt: ConfiguratorAttributeType
  type_gte: ConfiguratorAttributeType
  type_lte: ConfiguratorAttributeType
  type_in: [ConfiguratorAttributeType!]
  type_null: Boolean
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  createdBy_null: Boolean
  itemDefinitions: ConfiguratorItemDefinitionFilterType
  attributes: ConfiguratorAttributeFilterType
}

type ConfiguratorAttributeDefinitionResultType {
  items: [ConfiguratorAttributeDefinition!]!
  count: Int!
}

input ConfiguratorSlotDefinitionCreateInput {
  id: ID
  name: String
  itemDefinitionId: ID
  slotsIds: [ID!]
}

input ConfiguratorSlotDefinitionUpdateInput {
  name: String
  itemDefinitionId: ID
  slotsIds: [ID!]
}

input ConfiguratorSlotDefinitionSortType {
  id: ObjectSortType
  name: ObjectSortType
  itemDefinitionId: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  slotsIds: ObjectSortType
  itemDefinition: ConfiguratorItemDefinitionSortType
  slots: ConfiguratorSlotSortType
}

input ConfiguratorSlotDefinitionFilterType {
  AND: [ConfiguratorSlotDefinitionFilterType!]
  OR: [ConfiguratorSlotDefinitionFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  id_null: Boolean
  name: String
  name_ne: String
  name_gt: String
  name_lt: String
  name_gte: String
  name_lte: String
  name_in: [String!]
  name_like: String
  name_prefix: String
  name_suffix: String
  name_null: Boolean
  itemDefinitionId: ID
  itemDefinitionId_ne: ID
  itemDefinitionId_gt: ID
  itemDefinitionId_lt: ID
  itemDefinitionId_gte: ID
  itemDefinitionId_lte: ID
  itemDefinitionId_in: [ID!]
  itemDefinitionId_null: Boolean
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  createdBy_null: Boolean
  itemDefinition: ConfiguratorItemDefinitionFilterType
  slots: ConfiguratorSlotFilterType
}

type ConfiguratorSlotDefinitionResultType {
  items: [ConfiguratorSlotDefinition!]!
  count: Int!
}

input ConfiguratorItemCreateInput {
  id: ID
  stockItemID: ID
  definitionId: ID
  parentSlotId: ID
  attributesIds: [ID!]
  slotsIds: [ID!]
}

input ConfiguratorItemUpdateInput {
  stockItemID: ID
  definitionId: ID
  parentSlotId: ID
  attributesIds: [ID!]
  slotsIds: [ID!]
}

input ConfiguratorItemSortType {
  id: ObjectSortType
  stockItemID: ObjectSortType
  definitionId: ObjectSortType
  parentSlotId: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  attributesIds: ObjectSortType
  slotsIds: ObjectSortType
  definition: ConfiguratorItemDefinitionSortType
  attributes: ConfiguratorAttributeSortType
  slots: ConfiguratorSlotSortType
  parentSlot: ConfiguratorSlotSortType
}

input ConfiguratorItemFilterType {
  AND: [ConfiguratorItemFilterType!]
  OR: [ConfiguratorItemFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  id_null: Boolean
  stockItemID: ID
  stockItemID_ne: ID
  stockItemID_gt: ID
  stockItemID_lt: ID
  stockItemID_gte: ID
  stockItemID_lte: ID
  stockItemID_in: [ID!]
  stockItemID_null: Boolean
  definitionId: ID
  definitionId_ne: ID
  definitionId_gt: ID
  definitionId_lt: ID
  definitionId_gte: ID
  definitionId_lte: ID
  definitionId_in: [ID!]
  definitionId_null: Boolean
  parentSlotId: ID
  parentSlotId_ne: ID
  parentSlotId_gt: ID
  parentSlotId_lt: ID
  parentSlotId_gte: ID
  parentSlotId_lte: ID
  parentSlotId_in: [ID!]
  parentSlotId_null: Boolean
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  createdBy_null: Boolean
  definition: ConfiguratorItemDefinitionFilterType
  attributes: ConfiguratorAttributeFilterType
  slots: ConfiguratorSlotFilterType
  parentSlot: ConfiguratorSlotFilterType
}

type ConfiguratorItemResultType {
  items: [ConfiguratorItem!]!
  count: Int!
}

input ConfiguratorAttributeCreateInput {
  id: ID
  stringValue: String
  floatValue: Float
  intValue: Int
  definitionId: ID
  itemId: ID
}

input ConfiguratorAttributeUpdateInput {
  stringValue: String
  floatValue: Float
  intValue: Int
  definitionId: ID
  itemId: ID
}

input ConfiguratorAttributeSortType {
  id: ObjectSortType
  stringValue: ObjectSortType
  floatValue: ObjectSortType
  intValue: ObjectSortType
  definitionId: ObjectSortType
  itemId: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  definition: ConfiguratorAttributeDefinitionSortType
  item: ConfiguratorItemSortType
}

input ConfiguratorAttributeFilterType {
  AND: [ConfiguratorAttributeFilterType!]
  OR: [ConfiguratorAttributeFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  id_null: Boolean
  stringValue: String
  stringValue_ne: String
  stringValue_gt: String
  stringValue_lt: String
  stringValue_gte: String
  stringValue_lte: String
  stringValue_in: [String!]
  stringValue_like: String
  stringValue_prefix: String
  stringValue_suffix: String
  stringValue_null: Boolean
  floatValue: Float
  floatValue_ne: Float
  floatValue_gt: Float
  floatValue_lt: Float
  floatValue_gte: Float
  floatValue_lte: Float
  floatValue_in: [Float!]
  floatValue_null: Boolean
  intValue: Int
  intValue_ne: Int
  intValue_gt: Int
  intValue_lt: Int
  intValue_gte: Int
  intValue_lte: Int
  intValue_in: [Int!]
  intValue_null: Boolean
  definitionId: ID
  definitionId_ne: ID
  definitionId_gt: ID
  definitionId_lt: ID
  definitionId_gte: ID
  definitionId_lte: ID
  definitionId_in: [ID!]
  definitionId_null: Boolean
  itemId: ID
  itemId_ne: ID
  itemId_gt: ID
  itemId_lt: ID
  itemId_gte: ID
  itemId_lte: ID
  itemId_in: [ID!]
  itemId_null: Boolean
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  createdBy_null: Boolean
  definition: ConfiguratorAttributeDefinitionFilterType
  item: ConfiguratorItemFilterType
}

type ConfiguratorAttributeResultType {
  items: [ConfiguratorAttribute!]!
  count: Int!
}

input ConfiguratorSlotCreateInput {
  id: ID
  definitionId: ID
  parentItemId: ID
  itemsIds: [ID!]
}

input ConfiguratorSlotUpdateInput {
  definitionId: ID
  parentItemId: ID
  itemsIds: [ID!]
}

input ConfiguratorSlotSortType {
  id: ObjectSortType
  definitionId: ObjectSortType
  parentItemId: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  itemsIds: ObjectSortType
  items: ConfiguratorItemSortType
  definition: ConfiguratorSlotDefinitionSortType
  parentItem: ConfiguratorItemSortType
}

input ConfiguratorSlotFilterType {
  AND: [ConfiguratorSlotFilterType!]
  OR: [ConfiguratorSlotFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  id_null: Boolean
  definitionId: ID
  definitionId_ne: ID
  definitionId_gt: ID
  definitionId_lt: ID
  definitionId_gte: ID
  definitionId_lte: ID
  definitionId_in: [ID!]
  definitionId_null: Boolean
  parentItemId: ID
  parentItemId_ne: ID
  parentItemId_gt: ID
  parentItemId_lt: ID
  parentItemId_gte: ID
  parentItemId_lte: ID
  parentItemId_in: [ID!]
  parentItemId_null: Boolean
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  createdBy_null: Boolean
  items: ConfiguratorItemFilterType
  definition: ConfiguratorSlotDefinitionFilterType
  parentItem: ConfiguratorItemFilterType
}

type ConfiguratorSlotResultType {
  items: [ConfiguratorSlot!]!
  count: Int!
}`
)
