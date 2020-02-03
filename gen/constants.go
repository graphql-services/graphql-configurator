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
  configuratorItemDefinitionCategory(id: ID, q: String, filter: ConfiguratorItemDefinitionCategoryFilterType): ConfiguratorItemDefinitionCategory
  configuratorItemDefinitionCategories(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorItemDefinitionCategorySortType!], filter: ConfiguratorItemDefinitionCategoryFilterType): ConfiguratorItemDefinitionCategoryResultType
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
  createConfiguratorItemDefinitionCategory(input: ConfiguratorItemDefinitionCategoryCreateInput!): ConfiguratorItemDefinitionCategory!
  updateConfiguratorItemDefinitionCategory(id: ID!, input: ConfiguratorItemDefinitionCategoryUpdateInput!): ConfiguratorItemDefinitionCategory!
  deleteConfiguratorItemDefinitionCategory(id: ID!): ConfiguratorItemDefinitionCategory!
  deleteAllConfiguratorItemDefinitionCategories: Boolean!
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
  id: ID
  definitionId: ID
  referenceID: ID
  code: String
  name: String
  slots: [ConfiguratorAssemblySlot!]!
  attributes: [ConfiguratorAssemblyAttribute!]!
}

type ConfiguratorAssemblyAttribute {
  id: ID
  definitionId: ID!
  stringValue: String
  intValue: Int
  floatValue: Float
}

type ConfiguratorAssemblySlot {
  id: ID
  definitionId: ID!
  item: ConfiguratorAssemblyItem
}

input ConfiguratorAssemblyCreateInput {
  item: ConfiguratorAssemblyItemInput
}

input ConfiguratorAssemblyUpdateInput {
  item: ConfiguratorAssemblyItemInput
}

input ConfiguratorAssemblyItemInput {
  id: ID
  definitionId: ID
  referenceID: ID
  code: String
  name: String
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
  item: ConfiguratorAssemblyItemInput
}

extend type Query {
  configuratorAssembly(id: ID!): ConfiguratorAssembly
}

extend type Mutation {
  createConfiguratorAssembly(input: ConfiguratorAssemblyCreateInput!): ConfiguratorAssembly
  updateConfiguratorAssembly(id: ID!, input: ConfiguratorAssemblyUpdateInput!): ConfiguratorAssembly
}

type ConfiguratorItemDefinitionCategory {
  id: ID!
  code: String
  name: String
  type: String
  definitions: [ConfiguratorItemDefinition!]!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  definitionsIds: [ID!]!
}

type ConfiguratorItemDefinition {
  id: ID!
  code: String
  name: String
  attributes: [ConfiguratorAttributeDefinition!]!
  slots: [ConfiguratorSlotDefinition!]!
  items: [ConfiguratorItem!]!
  allowedInSlots: [ConfiguratorSlotDefinition!]!
  category: ConfiguratorItemDefinitionCategory
  categoryId: ID
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  attributesIds: [ID!]!
  slotsIds: [ID!]!
  itemsIds: [ID!]!
  allowedInSlotsIds: [ID!]!
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
  definitions: [ConfiguratorItemDefinition!]!
  attributes: [ConfiguratorAttribute!]!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  definitionsIds: [ID!]!
  attributesIds: [ID!]!
}

type ConfiguratorSlotDefinition {
  id: ID!
  name: String
  minCount: Int
  maxCount: Int
  definition: ConfiguratorItemDefinition
  slots: [ConfiguratorSlot!]!
  allowedItemDefinitions: [ConfiguratorItemDefinition!]!
  definitionId: ID
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  slotsIds: [ID!]!
  allowedItemDefinitionsIds: [ID!]!
}

type ConfiguratorItem {
  id: ID!
  code: String
  name: String
  stockItemId: ID
  referenceID: String
  rawData: String
  definition: ConfiguratorItemDefinition
  attributes: [ConfiguratorAttribute!]!
  slots: [ConfiguratorSlot!]!
  parentSlots: [ConfiguratorSlot!]!
  definitionId: ID
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  attributesIds: [ID!]!
  slotsIds: [ID!]!
  parentSlotsIds: [ID!]!
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
  item: ConfiguratorItem
  definition: ConfiguratorSlotDefinition
  parentItem: ConfiguratorItem
  itemId: ID
  definitionId: ID
  parentItemId: ID
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
}

input ConfiguratorItemDefinitionCategoryCreateInput {
  id: ID
  code: String
  name: String
  type: String
  definitionsIds: [ID!]
}

input ConfiguratorItemDefinitionCategoryUpdateInput {
  code: String
  name: String
  type: String
  definitionsIds: [ID!]
}

input ConfiguratorItemDefinitionCategorySortType {
  id: ObjectSortType
  code: ObjectSortType
  name: ObjectSortType
  type: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  definitionsIds: ObjectSortType
  definitions: ConfiguratorItemDefinitionSortType
}

input ConfiguratorItemDefinitionCategoryFilterType {
  AND: [ConfiguratorItemDefinitionCategoryFilterType!]
  OR: [ConfiguratorItemDefinitionCategoryFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  id_null: Boolean
  code: String
  code_ne: String
  code_gt: String
  code_lt: String
  code_gte: String
  code_lte: String
  code_in: [String!]
  code_like: String
  code_prefix: String
  code_suffix: String
  code_null: Boolean
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
  type: String
  type_ne: String
  type_gt: String
  type_lt: String
  type_gte: String
  type_lte: String
  type_in: [String!]
  type_like: String
  type_prefix: String
  type_suffix: String
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
  definitions: ConfiguratorItemDefinitionFilterType
}

type ConfiguratorItemDefinitionCategoryResultType {
  items: [ConfiguratorItemDefinitionCategory!]!
  count: Int!
}

input ConfiguratorItemDefinitionCreateInput {
  id: ID
  code: String
  name: String
  categoryId: ID
  attributesIds: [ID!]
  slotsIds: [ID!]
  itemsIds: [ID!]
  allowedInSlotsIds: [ID!]
}

input ConfiguratorItemDefinitionUpdateInput {
  code: String
  name: String
  categoryId: ID
  attributesIds: [ID!]
  slotsIds: [ID!]
  itemsIds: [ID!]
  allowedInSlotsIds: [ID!]
}

input ConfiguratorItemDefinitionSortType {
  id: ObjectSortType
  code: ObjectSortType
  name: ObjectSortType
  categoryId: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  attributesIds: ObjectSortType
  slotsIds: ObjectSortType
  itemsIds: ObjectSortType
  allowedInSlotsIds: ObjectSortType
  attributes: ConfiguratorAttributeDefinitionSortType
  slots: ConfiguratorSlotDefinitionSortType
  items: ConfiguratorItemSortType
  allowedInSlots: ConfiguratorSlotDefinitionSortType
  category: ConfiguratorItemDefinitionCategorySortType
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
  code: String
  code_ne: String
  code_gt: String
  code_lt: String
  code_gte: String
  code_lte: String
  code_in: [String!]
  code_like: String
  code_prefix: String
  code_suffix: String
  code_null: Boolean
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
  categoryId: ID
  categoryId_ne: ID
  categoryId_gt: ID
  categoryId_lt: ID
  categoryId_gte: ID
  categoryId_lte: ID
  categoryId_in: [ID!]
  categoryId_null: Boolean
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
  allowedInSlots: ConfiguratorSlotDefinitionFilterType
  category: ConfiguratorItemDefinitionCategoryFilterType
}

type ConfiguratorItemDefinitionResultType {
  items: [ConfiguratorItemDefinition!]!
  count: Int!
}

input ConfiguratorAttributeDefinitionCreateInput {
  id: ID
  name: String
  type: ConfiguratorAttributeType
  definitionsIds: [ID!]
  attributesIds: [ID!]
}

input ConfiguratorAttributeDefinitionUpdateInput {
  name: String
  type: ConfiguratorAttributeType
  definitionsIds: [ID!]
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
  definitionsIds: ObjectSortType
  attributesIds: ObjectSortType
  definitions: ConfiguratorItemDefinitionSortType
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
  definitions: ConfiguratorItemDefinitionFilterType
  attributes: ConfiguratorAttributeFilterType
}

type ConfiguratorAttributeDefinitionResultType {
  items: [ConfiguratorAttributeDefinition!]!
  count: Int!
}

input ConfiguratorSlotDefinitionCreateInput {
  id: ID
  name: String
  minCount: Int
  maxCount: Int
  definitionId: ID
  slotsIds: [ID!]
  allowedItemDefinitionsIds: [ID!]
}

input ConfiguratorSlotDefinitionUpdateInput {
  name: String
  minCount: Int
  maxCount: Int
  definitionId: ID
  slotsIds: [ID!]
  allowedItemDefinitionsIds: [ID!]
}

input ConfiguratorSlotDefinitionSortType {
  id: ObjectSortType
  name: ObjectSortType
  minCount: ObjectSortType
  maxCount: ObjectSortType
  definitionId: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  slotsIds: ObjectSortType
  allowedItemDefinitionsIds: ObjectSortType
  definition: ConfiguratorItemDefinitionSortType
  slots: ConfiguratorSlotSortType
  allowedItemDefinitions: ConfiguratorItemDefinitionSortType
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
  minCount: Int
  minCount_ne: Int
  minCount_gt: Int
  minCount_lt: Int
  minCount_gte: Int
  minCount_lte: Int
  minCount_in: [Int!]
  minCount_null: Boolean
  maxCount: Int
  maxCount_ne: Int
  maxCount_gt: Int
  maxCount_lt: Int
  maxCount_gte: Int
  maxCount_lte: Int
  maxCount_in: [Int!]
  maxCount_null: Boolean
  definitionId: ID
  definitionId_ne: ID
  definitionId_gt: ID
  definitionId_lt: ID
  definitionId_gte: ID
  definitionId_lte: ID
  definitionId_in: [ID!]
  definitionId_null: Boolean
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
  slots: ConfiguratorSlotFilterType
  allowedItemDefinitions: ConfiguratorItemDefinitionFilterType
}

type ConfiguratorSlotDefinitionResultType {
  items: [ConfiguratorSlotDefinition!]!
  count: Int!
}

input ConfiguratorItemCreateInput {
  id: ID
  code: String
  name: String
  stockItemId: ID
  referenceID: String
  rawData: String
  definitionId: ID
  attributesIds: [ID!]
  slotsIds: [ID!]
  parentSlotsIds: [ID!]
}

input ConfiguratorItemUpdateInput {
  code: String
  name: String
  stockItemId: ID
  referenceID: String
  rawData: String
  definitionId: ID
  attributesIds: [ID!]
  slotsIds: [ID!]
  parentSlotsIds: [ID!]
}

input ConfiguratorItemSortType {
  id: ObjectSortType
  code: ObjectSortType
  name: ObjectSortType
  stockItemId: ObjectSortType
  referenceID: ObjectSortType
  rawData: ObjectSortType
  definitionId: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  attributesIds: ObjectSortType
  slotsIds: ObjectSortType
  parentSlotsIds: ObjectSortType
  definition: ConfiguratorItemDefinitionSortType
  attributes: ConfiguratorAttributeSortType
  slots: ConfiguratorSlotSortType
  parentSlots: ConfiguratorSlotSortType
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
  code: String
  code_ne: String
  code_gt: String
  code_lt: String
  code_gte: String
  code_lte: String
  code_in: [String!]
  code_like: String
  code_prefix: String
  code_suffix: String
  code_null: Boolean
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
  stockItemId: ID
  stockItemId_ne: ID
  stockItemId_gt: ID
  stockItemId_lt: ID
  stockItemId_gte: ID
  stockItemId_lte: ID
  stockItemId_in: [ID!]
  stockItemId_null: Boolean
  referenceID: String
  referenceID_ne: String
  referenceID_gt: String
  referenceID_lt: String
  referenceID_gte: String
  referenceID_lte: String
  referenceID_in: [String!]
  referenceID_like: String
  referenceID_prefix: String
  referenceID_suffix: String
  referenceID_null: Boolean
  rawData: String
  rawData_ne: String
  rawData_gt: String
  rawData_lt: String
  rawData_gte: String
  rawData_lte: String
  rawData_in: [String!]
  rawData_like: String
  rawData_prefix: String
  rawData_suffix: String
  rawData_null: Boolean
  definitionId: ID
  definitionId_ne: ID
  definitionId_gt: ID
  definitionId_lt: ID
  definitionId_gte: ID
  definitionId_lte: ID
  definitionId_in: [ID!]
  definitionId_null: Boolean
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
  parentSlots: ConfiguratorSlotFilterType
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
  itemId: ID
  definitionId: ID
  parentItemId: ID
}

input ConfiguratorSlotUpdateInput {
  itemId: ID
  definitionId: ID
  parentItemId: ID
}

input ConfiguratorSlotSortType {
  id: ObjectSortType
  itemId: ObjectSortType
  definitionId: ObjectSortType
  parentItemId: ObjectSortType
  updatedAt: ObjectSortType
  createdAt: ObjectSortType
  updatedBy: ObjectSortType
  createdBy: ObjectSortType
  item: ConfiguratorItemSortType
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
  itemId: ID
  itemId_ne: ID
  itemId_gt: ID
  itemId_lt: ID
  itemId_gte: ID
  itemId_lte: ID
  itemId_in: [ID!]
  itemId_null: Boolean
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
  item: ConfiguratorItemFilterType
  definition: ConfiguratorSlotDefinitionFilterType
  parentItem: ConfiguratorItemFilterType
}

type ConfiguratorSlotResultType {
  items: [ConfiguratorSlot!]!
  count: Int!
}`
)
