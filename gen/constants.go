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
  configuratorItemDefinitionCategories(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorItemDefinitionCategorySortType!], filter: ConfiguratorItemDefinitionCategoryFilterType): ConfiguratorItemDefinitionCategoryResultType!
  configuratorItemDefinition(id: ID, q: String, filter: ConfiguratorItemDefinitionFilterType): ConfiguratorItemDefinition
  configuratorItemDefinitions(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorItemDefinitionSortType!], filter: ConfiguratorItemDefinitionFilterType): ConfiguratorItemDefinitionResultType!
  configuratorAttributeDefinition(id: ID, q: String, filter: ConfiguratorAttributeDefinitionFilterType): ConfiguratorAttributeDefinition
  configuratorAttributeDefinitions(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorAttributeDefinitionSortType!], filter: ConfiguratorAttributeDefinitionFilterType): ConfiguratorAttributeDefinitionResultType!
  configuratorSlotDefinition(id: ID, q: String, filter: ConfiguratorSlotDefinitionFilterType): ConfiguratorSlotDefinition
  configuratorSlotDefinitions(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorSlotDefinitionSortType!], filter: ConfiguratorSlotDefinitionFilterType): ConfiguratorSlotDefinitionResultType!
  configuratorItem(id: ID, q: String, filter: ConfiguratorItemFilterType): ConfiguratorItem
  configuratorItems(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorItemSortType!], filter: ConfiguratorItemFilterType): ConfiguratorItemResultType!
  configuratorAttribute(id: ID, q: String, filter: ConfiguratorAttributeFilterType): ConfiguratorAttribute
  configuratorAttributes(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorAttributeSortType!], filter: ConfiguratorAttributeFilterType): ConfiguratorAttributeResultType!
  configuratorSlot(id: ID, q: String, filter: ConfiguratorSlotFilterType): ConfiguratorSlot
  configuratorSlots(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorSlotSortType!], filter: ConfiguratorSlotFilterType): ConfiguratorSlotResultType!
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
  stockItemId: ID
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
  count: Float
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
  stockItemId: ID
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
  count: Float
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
  primary: Boolean
  definitions: [ConfiguratorItemDefinition!]!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  definitionsIds: [ID!]!
  definitionsConnection(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorItemDefinitionSortType!], filter: ConfiguratorItemDefinitionFilterType): ConfiguratorItemDefinitionResultType!
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
  attributesConnection(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorAttributeDefinitionSortType!], filter: ConfiguratorAttributeDefinitionFilterType): ConfiguratorAttributeDefinitionResultType!
  slotsIds: [ID!]!
  slotsConnection(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorSlotDefinitionSortType!], filter: ConfiguratorSlotDefinitionFilterType): ConfiguratorSlotDefinitionResultType!
  itemsIds: [ID!]!
  itemsConnection(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorItemSortType!], filter: ConfiguratorItemFilterType): ConfiguratorItemResultType!
  allowedInSlotsIds: [ID!]!
  allowedInSlotsConnection(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorSlotDefinitionSortType!], filter: ConfiguratorSlotDefinitionFilterType): ConfiguratorSlotDefinitionResultType!
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
  definitionsConnection(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorItemDefinitionSortType!], filter: ConfiguratorItemDefinitionFilterType): ConfiguratorItemDefinitionResultType!
  attributesIds: [ID!]!
  attributesConnection(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorAttributeSortType!], filter: ConfiguratorAttributeFilterType): ConfiguratorAttributeResultType!
}

type ConfiguratorSlotDefinition {
  id: ID!
  name: String
  minCount: Float
  maxCount: Float
  defaultCount: Float
  definition: ConfiguratorItemDefinition
  slots: [ConfiguratorSlot!]!
  allowedItemDefinitions: [ConfiguratorItemDefinition!]!
  definitionId: ID
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  slotsIds: [ID!]!
  slotsConnection(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorSlotSortType!], filter: ConfiguratorSlotFilterType): ConfiguratorSlotResultType!
  allowedItemDefinitionsIds: [ID!]!
  allowedItemDefinitionsConnection(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorItemDefinitionSortType!], filter: ConfiguratorItemDefinitionFilterType): ConfiguratorItemDefinitionResultType!
}

type ConfiguratorItem @key(fields: "id") {
  id: ID!
  code: String
  name: String
  stockItemId: ID
  stockItem: StockItem
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
  attributesConnection(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorAttributeSortType!], filter: ConfiguratorAttributeFilterType): ConfiguratorAttributeResultType!
  slotsIds: [ID!]!
  slotsConnection(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorSlotSortType!], filter: ConfiguratorSlotFilterType): ConfiguratorSlotResultType!
  parentSlotsIds: [ID!]!
  parentSlotsConnection(offset: Int, limit: Int = 30, q: String, sort: [ConfiguratorSlotSortType!], filter: ConfiguratorSlotFilterType): ConfiguratorSlotResultType!
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
  count: Float
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

extend type StockItem @key(fields: "id") {
  id: ID! @external
}

input ConfiguratorItemDefinitionCategoryCreateInput {
  id: ID
  code: String
  name: String
  type: String
  primary: Boolean
  definitionsIds: [ID!]
}

input ConfiguratorItemDefinitionCategoryUpdateInput {
  code: String
  name: String
  type: String
  primary: Boolean
  definitionsIds: [ID!]
}

input ConfiguratorItemDefinitionCategorySortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  code: ObjectSortType
  codeMin: ObjectSortType
  codeMax: ObjectSortType
  name: ObjectSortType
  nameMin: ObjectSortType
  nameMax: ObjectSortType
  type: ObjectSortType
  typeMin: ObjectSortType
  typeMax: ObjectSortType
  primary: ObjectSortType
  primaryMin: ObjectSortType
  primaryMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
  definitionsIds: ObjectSortType
  definitionsIdsMin: ObjectSortType
  definitionsIdsMax: ObjectSortType
  definitions: ConfiguratorItemDefinitionSortType
}

input ConfiguratorItemDefinitionCategoryFilterType {
  AND: [ConfiguratorItemDefinitionCategoryFilterType!]
  OR: [ConfiguratorItemDefinitionCategoryFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_null: Boolean
  code: String
  codeMin: String
  codeMax: String
  code_ne: String
  codeMin_ne: String
  codeMax_ne: String
  code_gt: String
  codeMin_gt: String
  codeMax_gt: String
  code_lt: String
  codeMin_lt: String
  codeMax_lt: String
  code_gte: String
  codeMin_gte: String
  codeMax_gte: String
  code_lte: String
  codeMin_lte: String
  codeMax_lte: String
  code_in: [String!]
  codeMin_in: [String!]
  codeMax_in: [String!]
  code_like: String
  codeMin_like: String
  codeMax_like: String
  code_prefix: String
  codeMin_prefix: String
  codeMax_prefix: String
  code_suffix: String
  codeMin_suffix: String
  codeMax_suffix: String
  code_null: Boolean
  name: String
  nameMin: String
  nameMax: String
  name_ne: String
  nameMin_ne: String
  nameMax_ne: String
  name_gt: String
  nameMin_gt: String
  nameMax_gt: String
  name_lt: String
  nameMin_lt: String
  nameMax_lt: String
  name_gte: String
  nameMin_gte: String
  nameMax_gte: String
  name_lte: String
  nameMin_lte: String
  nameMax_lte: String
  name_in: [String!]
  nameMin_in: [String!]
  nameMax_in: [String!]
  name_like: String
  nameMin_like: String
  nameMax_like: String
  name_prefix: String
  nameMin_prefix: String
  nameMax_prefix: String
  name_suffix: String
  nameMin_suffix: String
  nameMax_suffix: String
  name_null: Boolean
  type: String
  typeMin: String
  typeMax: String
  type_ne: String
  typeMin_ne: String
  typeMax_ne: String
  type_gt: String
  typeMin_gt: String
  typeMax_gt: String
  type_lt: String
  typeMin_lt: String
  typeMax_lt: String
  type_gte: String
  typeMin_gte: String
  typeMax_gte: String
  type_lte: String
  typeMin_lte: String
  typeMax_lte: String
  type_in: [String!]
  typeMin_in: [String!]
  typeMax_in: [String!]
  type_like: String
  typeMin_like: String
  typeMax_like: String
  type_prefix: String
  typeMin_prefix: String
  typeMax_prefix: String
  type_suffix: String
  typeMin_suffix: String
  typeMax_suffix: String
  type_null: Boolean
  primary: Boolean
  primaryMin: Boolean
  primaryMax: Boolean
  primary_ne: Boolean
  primaryMin_ne: Boolean
  primaryMax_ne: Boolean
  primary_gt: Boolean
  primaryMin_gt: Boolean
  primaryMax_gt: Boolean
  primary_lt: Boolean
  primaryMin_lt: Boolean
  primaryMax_lt: Boolean
  primary_gte: Boolean
  primaryMin_gte: Boolean
  primaryMax_gte: Boolean
  primary_lte: Boolean
  primaryMin_lte: Boolean
  primaryMax_lte: Boolean
  primary_in: [Boolean!]
  primaryMin_in: [Boolean!]
  primaryMax_in: [Boolean!]
  primary_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
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
  idMin: ObjectSortType
  idMax: ObjectSortType
  code: ObjectSortType
  codeMin: ObjectSortType
  codeMax: ObjectSortType
  name: ObjectSortType
  nameMin: ObjectSortType
  nameMax: ObjectSortType
  categoryId: ObjectSortType
  categoryIdMin: ObjectSortType
  categoryIdMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
  attributesIds: ObjectSortType
  attributesIdsMin: ObjectSortType
  attributesIdsMax: ObjectSortType
  slotsIds: ObjectSortType
  slotsIdsMin: ObjectSortType
  slotsIdsMax: ObjectSortType
  itemsIds: ObjectSortType
  itemsIdsMin: ObjectSortType
  itemsIdsMax: ObjectSortType
  allowedInSlotsIds: ObjectSortType
  allowedInSlotsIdsMin: ObjectSortType
  allowedInSlotsIdsMax: ObjectSortType
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
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_null: Boolean
  code: String
  codeMin: String
  codeMax: String
  code_ne: String
  codeMin_ne: String
  codeMax_ne: String
  code_gt: String
  codeMin_gt: String
  codeMax_gt: String
  code_lt: String
  codeMin_lt: String
  codeMax_lt: String
  code_gte: String
  codeMin_gte: String
  codeMax_gte: String
  code_lte: String
  codeMin_lte: String
  codeMax_lte: String
  code_in: [String!]
  codeMin_in: [String!]
  codeMax_in: [String!]
  code_like: String
  codeMin_like: String
  codeMax_like: String
  code_prefix: String
  codeMin_prefix: String
  codeMax_prefix: String
  code_suffix: String
  codeMin_suffix: String
  codeMax_suffix: String
  code_null: Boolean
  name: String
  nameMin: String
  nameMax: String
  name_ne: String
  nameMin_ne: String
  nameMax_ne: String
  name_gt: String
  nameMin_gt: String
  nameMax_gt: String
  name_lt: String
  nameMin_lt: String
  nameMax_lt: String
  name_gte: String
  nameMin_gte: String
  nameMax_gte: String
  name_lte: String
  nameMin_lte: String
  nameMax_lte: String
  name_in: [String!]
  nameMin_in: [String!]
  nameMax_in: [String!]
  name_like: String
  nameMin_like: String
  nameMax_like: String
  name_prefix: String
  nameMin_prefix: String
  nameMax_prefix: String
  name_suffix: String
  nameMin_suffix: String
  nameMax_suffix: String
  name_null: Boolean
  categoryId: ID
  categoryIdMin: ID
  categoryIdMax: ID
  categoryId_ne: ID
  categoryIdMin_ne: ID
  categoryIdMax_ne: ID
  categoryId_gt: ID
  categoryIdMin_gt: ID
  categoryIdMax_gt: ID
  categoryId_lt: ID
  categoryIdMin_lt: ID
  categoryIdMax_lt: ID
  categoryId_gte: ID
  categoryIdMin_gte: ID
  categoryIdMax_gte: ID
  categoryId_lte: ID
  categoryIdMin_lte: ID
  categoryIdMax_lte: ID
  categoryId_in: [ID!]
  categoryIdMin_in: [ID!]
  categoryIdMax_in: [ID!]
  categoryId_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
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
  idMin: ObjectSortType
  idMax: ObjectSortType
  name: ObjectSortType
  nameMin: ObjectSortType
  nameMax: ObjectSortType
  type: ObjectSortType
  typeMin: ObjectSortType
  typeMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
  definitionsIds: ObjectSortType
  definitionsIdsMin: ObjectSortType
  definitionsIdsMax: ObjectSortType
  attributesIds: ObjectSortType
  attributesIdsMin: ObjectSortType
  attributesIdsMax: ObjectSortType
  definitions: ConfiguratorItemDefinitionSortType
  attributes: ConfiguratorAttributeSortType
}

input ConfiguratorAttributeDefinitionFilterType {
  AND: [ConfiguratorAttributeDefinitionFilterType!]
  OR: [ConfiguratorAttributeDefinitionFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_null: Boolean
  name: String
  nameMin: String
  nameMax: String
  name_ne: String
  nameMin_ne: String
  nameMax_ne: String
  name_gt: String
  nameMin_gt: String
  nameMax_gt: String
  name_lt: String
  nameMin_lt: String
  nameMax_lt: String
  name_gte: String
  nameMin_gte: String
  nameMax_gte: String
  name_lte: String
  nameMin_lte: String
  nameMax_lte: String
  name_in: [String!]
  nameMin_in: [String!]
  nameMax_in: [String!]
  name_like: String
  nameMin_like: String
  nameMax_like: String
  name_prefix: String
  nameMin_prefix: String
  nameMax_prefix: String
  name_suffix: String
  nameMin_suffix: String
  nameMax_suffix: String
  name_null: Boolean
  type: ConfiguratorAttributeType
  typeMin: ConfiguratorAttributeType
  typeMax: ConfiguratorAttributeType
  type_ne: ConfiguratorAttributeType
  typeMin_ne: ConfiguratorAttributeType
  typeMax_ne: ConfiguratorAttributeType
  type_gt: ConfiguratorAttributeType
  typeMin_gt: ConfiguratorAttributeType
  typeMax_gt: ConfiguratorAttributeType
  type_lt: ConfiguratorAttributeType
  typeMin_lt: ConfiguratorAttributeType
  typeMax_lt: ConfiguratorAttributeType
  type_gte: ConfiguratorAttributeType
  typeMin_gte: ConfiguratorAttributeType
  typeMax_gte: ConfiguratorAttributeType
  type_lte: ConfiguratorAttributeType
  typeMin_lte: ConfiguratorAttributeType
  typeMax_lte: ConfiguratorAttributeType
  type_in: [ConfiguratorAttributeType!]
  typeMin_in: [ConfiguratorAttributeType!]
  typeMax_in: [ConfiguratorAttributeType!]
  type_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
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
  minCount: Float
  maxCount: Float
  defaultCount: Float
  definitionId: ID
  slotsIds: [ID!]
  allowedItemDefinitionsIds: [ID!]
}

input ConfiguratorSlotDefinitionUpdateInput {
  name: String
  minCount: Float
  maxCount: Float
  defaultCount: Float
  definitionId: ID
  slotsIds: [ID!]
  allowedItemDefinitionsIds: [ID!]
}

input ConfiguratorSlotDefinitionSortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  name: ObjectSortType
  nameMin: ObjectSortType
  nameMax: ObjectSortType
  minCount: ObjectSortType
  minCountMin: ObjectSortType
  minCountMax: ObjectSortType
  minCountAvg: ObjectSortType
  maxCount: ObjectSortType
  maxCountMin: ObjectSortType
  maxCountMax: ObjectSortType
  maxCountAvg: ObjectSortType
  defaultCount: ObjectSortType
  defaultCountMin: ObjectSortType
  defaultCountMax: ObjectSortType
  defaultCountAvg: ObjectSortType
  definitionId: ObjectSortType
  definitionIdMin: ObjectSortType
  definitionIdMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
  slotsIds: ObjectSortType
  slotsIdsMin: ObjectSortType
  slotsIdsMax: ObjectSortType
  allowedItemDefinitionsIds: ObjectSortType
  allowedItemDefinitionsIdsMin: ObjectSortType
  allowedItemDefinitionsIdsMax: ObjectSortType
  definition: ConfiguratorItemDefinitionSortType
  slots: ConfiguratorSlotSortType
  allowedItemDefinitions: ConfiguratorItemDefinitionSortType
}

input ConfiguratorSlotDefinitionFilterType {
  AND: [ConfiguratorSlotDefinitionFilterType!]
  OR: [ConfiguratorSlotDefinitionFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_null: Boolean
  name: String
  nameMin: String
  nameMax: String
  name_ne: String
  nameMin_ne: String
  nameMax_ne: String
  name_gt: String
  nameMin_gt: String
  nameMax_gt: String
  name_lt: String
  nameMin_lt: String
  nameMax_lt: String
  name_gte: String
  nameMin_gte: String
  nameMax_gte: String
  name_lte: String
  nameMin_lte: String
  nameMax_lte: String
  name_in: [String!]
  nameMin_in: [String!]
  nameMax_in: [String!]
  name_like: String
  nameMin_like: String
  nameMax_like: String
  name_prefix: String
  nameMin_prefix: String
  nameMax_prefix: String
  name_suffix: String
  nameMin_suffix: String
  nameMax_suffix: String
  name_null: Boolean
  minCount: Float
  minCountMin: Float
  minCountMax: Float
  minCountAvg: Float
  minCount_ne: Float
  minCountMin_ne: Float
  minCountMax_ne: Float
  minCountAvg_ne: Float
  minCount_gt: Float
  minCountMin_gt: Float
  minCountMax_gt: Float
  minCountAvg_gt: Float
  minCount_lt: Float
  minCountMin_lt: Float
  minCountMax_lt: Float
  minCountAvg_lt: Float
  minCount_gte: Float
  minCountMin_gte: Float
  minCountMax_gte: Float
  minCountAvg_gte: Float
  minCount_lte: Float
  minCountMin_lte: Float
  minCountMax_lte: Float
  minCountAvg_lte: Float
  minCount_in: [Float!]
  minCountMin_in: [Float!]
  minCountMax_in: [Float!]
  minCountAvg_in: [Float!]
  minCount_null: Boolean
  maxCount: Float
  maxCountMin: Float
  maxCountMax: Float
  maxCountAvg: Float
  maxCount_ne: Float
  maxCountMin_ne: Float
  maxCountMax_ne: Float
  maxCountAvg_ne: Float
  maxCount_gt: Float
  maxCountMin_gt: Float
  maxCountMax_gt: Float
  maxCountAvg_gt: Float
  maxCount_lt: Float
  maxCountMin_lt: Float
  maxCountMax_lt: Float
  maxCountAvg_lt: Float
  maxCount_gte: Float
  maxCountMin_gte: Float
  maxCountMax_gte: Float
  maxCountAvg_gte: Float
  maxCount_lte: Float
  maxCountMin_lte: Float
  maxCountMax_lte: Float
  maxCountAvg_lte: Float
  maxCount_in: [Float!]
  maxCountMin_in: [Float!]
  maxCountMax_in: [Float!]
  maxCountAvg_in: [Float!]
  maxCount_null: Boolean
  defaultCount: Float
  defaultCountMin: Float
  defaultCountMax: Float
  defaultCountAvg: Float
  defaultCount_ne: Float
  defaultCountMin_ne: Float
  defaultCountMax_ne: Float
  defaultCountAvg_ne: Float
  defaultCount_gt: Float
  defaultCountMin_gt: Float
  defaultCountMax_gt: Float
  defaultCountAvg_gt: Float
  defaultCount_lt: Float
  defaultCountMin_lt: Float
  defaultCountMax_lt: Float
  defaultCountAvg_lt: Float
  defaultCount_gte: Float
  defaultCountMin_gte: Float
  defaultCountMax_gte: Float
  defaultCountAvg_gte: Float
  defaultCount_lte: Float
  defaultCountMin_lte: Float
  defaultCountMax_lte: Float
  defaultCountAvg_lte: Float
  defaultCount_in: [Float!]
  defaultCountMin_in: [Float!]
  defaultCountMax_in: [Float!]
  defaultCountAvg_in: [Float!]
  defaultCount_null: Boolean
  definitionId: ID
  definitionIdMin: ID
  definitionIdMax: ID
  definitionId_ne: ID
  definitionIdMin_ne: ID
  definitionIdMax_ne: ID
  definitionId_gt: ID
  definitionIdMin_gt: ID
  definitionIdMax_gt: ID
  definitionId_lt: ID
  definitionIdMin_lt: ID
  definitionIdMax_lt: ID
  definitionId_gte: ID
  definitionIdMin_gte: ID
  definitionIdMax_gte: ID
  definitionId_lte: ID
  definitionIdMin_lte: ID
  definitionIdMax_lte: ID
  definitionId_in: [ID!]
  definitionIdMin_in: [ID!]
  definitionIdMax_in: [ID!]
  definitionId_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
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
  rawData: String
  definitionId: ID
  attributesIds: [ID!]
  slotsIds: [ID!]
  parentSlotsIds: [ID!]
}

input ConfiguratorItemSortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  code: ObjectSortType
  codeMin: ObjectSortType
  codeMax: ObjectSortType
  name: ObjectSortType
  nameMin: ObjectSortType
  nameMax: ObjectSortType
  stockItemId: ObjectSortType
  stockItemIdMin: ObjectSortType
  stockItemIdMax: ObjectSortType
  rawData: ObjectSortType
  rawDataMin: ObjectSortType
  rawDataMax: ObjectSortType
  definitionId: ObjectSortType
  definitionIdMin: ObjectSortType
  definitionIdMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
  attributesIds: ObjectSortType
  attributesIdsMin: ObjectSortType
  attributesIdsMax: ObjectSortType
  slotsIds: ObjectSortType
  slotsIdsMin: ObjectSortType
  slotsIdsMax: ObjectSortType
  parentSlotsIds: ObjectSortType
  parentSlotsIdsMin: ObjectSortType
  parentSlotsIdsMax: ObjectSortType
  definition: ConfiguratorItemDefinitionSortType
  attributes: ConfiguratorAttributeSortType
  slots: ConfiguratorSlotSortType
  parentSlots: ConfiguratorSlotSortType
}

input ConfiguratorItemFilterType {
  AND: [ConfiguratorItemFilterType!]
  OR: [ConfiguratorItemFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_null: Boolean
  code: String
  codeMin: String
  codeMax: String
  code_ne: String
  codeMin_ne: String
  codeMax_ne: String
  code_gt: String
  codeMin_gt: String
  codeMax_gt: String
  code_lt: String
  codeMin_lt: String
  codeMax_lt: String
  code_gte: String
  codeMin_gte: String
  codeMax_gte: String
  code_lte: String
  codeMin_lte: String
  codeMax_lte: String
  code_in: [String!]
  codeMin_in: [String!]
  codeMax_in: [String!]
  code_like: String
  codeMin_like: String
  codeMax_like: String
  code_prefix: String
  codeMin_prefix: String
  codeMax_prefix: String
  code_suffix: String
  codeMin_suffix: String
  codeMax_suffix: String
  code_null: Boolean
  name: String
  nameMin: String
  nameMax: String
  name_ne: String
  nameMin_ne: String
  nameMax_ne: String
  name_gt: String
  nameMin_gt: String
  nameMax_gt: String
  name_lt: String
  nameMin_lt: String
  nameMax_lt: String
  name_gte: String
  nameMin_gte: String
  nameMax_gte: String
  name_lte: String
  nameMin_lte: String
  nameMax_lte: String
  name_in: [String!]
  nameMin_in: [String!]
  nameMax_in: [String!]
  name_like: String
  nameMin_like: String
  nameMax_like: String
  name_prefix: String
  nameMin_prefix: String
  nameMax_prefix: String
  name_suffix: String
  nameMin_suffix: String
  nameMax_suffix: String
  name_null: Boolean
  stockItemId: ID
  stockItemIdMin: ID
  stockItemIdMax: ID
  stockItemId_ne: ID
  stockItemIdMin_ne: ID
  stockItemIdMax_ne: ID
  stockItemId_gt: ID
  stockItemIdMin_gt: ID
  stockItemIdMax_gt: ID
  stockItemId_lt: ID
  stockItemIdMin_lt: ID
  stockItemIdMax_lt: ID
  stockItemId_gte: ID
  stockItemIdMin_gte: ID
  stockItemIdMax_gte: ID
  stockItemId_lte: ID
  stockItemIdMin_lte: ID
  stockItemIdMax_lte: ID
  stockItemId_in: [ID!]
  stockItemIdMin_in: [ID!]
  stockItemIdMax_in: [ID!]
  stockItemId_null: Boolean
  rawData: String
  rawDataMin: String
  rawDataMax: String
  rawData_ne: String
  rawDataMin_ne: String
  rawDataMax_ne: String
  rawData_gt: String
  rawDataMin_gt: String
  rawDataMax_gt: String
  rawData_lt: String
  rawDataMin_lt: String
  rawDataMax_lt: String
  rawData_gte: String
  rawDataMin_gte: String
  rawDataMax_gte: String
  rawData_lte: String
  rawDataMin_lte: String
  rawDataMax_lte: String
  rawData_in: [String!]
  rawDataMin_in: [String!]
  rawDataMax_in: [String!]
  rawData_like: String
  rawDataMin_like: String
  rawDataMax_like: String
  rawData_prefix: String
  rawDataMin_prefix: String
  rawDataMax_prefix: String
  rawData_suffix: String
  rawDataMin_suffix: String
  rawDataMax_suffix: String
  rawData_null: Boolean
  definitionId: ID
  definitionIdMin: ID
  definitionIdMax: ID
  definitionId_ne: ID
  definitionIdMin_ne: ID
  definitionIdMax_ne: ID
  definitionId_gt: ID
  definitionIdMin_gt: ID
  definitionIdMax_gt: ID
  definitionId_lt: ID
  definitionIdMin_lt: ID
  definitionIdMax_lt: ID
  definitionId_gte: ID
  definitionIdMin_gte: ID
  definitionIdMax_gte: ID
  definitionId_lte: ID
  definitionIdMin_lte: ID
  definitionIdMax_lte: ID
  definitionId_in: [ID!]
  definitionIdMin_in: [ID!]
  definitionIdMax_in: [ID!]
  definitionId_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
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
  idMin: ObjectSortType
  idMax: ObjectSortType
  stringValue: ObjectSortType
  stringValueMin: ObjectSortType
  stringValueMax: ObjectSortType
  floatValue: ObjectSortType
  floatValueMin: ObjectSortType
  floatValueMax: ObjectSortType
  floatValueAvg: ObjectSortType
  intValue: ObjectSortType
  intValueMin: ObjectSortType
  intValueMax: ObjectSortType
  intValueAvg: ObjectSortType
  definitionId: ObjectSortType
  definitionIdMin: ObjectSortType
  definitionIdMax: ObjectSortType
  itemId: ObjectSortType
  itemIdMin: ObjectSortType
  itemIdMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
  definition: ConfiguratorAttributeDefinitionSortType
  item: ConfiguratorItemSortType
}

input ConfiguratorAttributeFilterType {
  AND: [ConfiguratorAttributeFilterType!]
  OR: [ConfiguratorAttributeFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_null: Boolean
  stringValue: String
  stringValueMin: String
  stringValueMax: String
  stringValue_ne: String
  stringValueMin_ne: String
  stringValueMax_ne: String
  stringValue_gt: String
  stringValueMin_gt: String
  stringValueMax_gt: String
  stringValue_lt: String
  stringValueMin_lt: String
  stringValueMax_lt: String
  stringValue_gte: String
  stringValueMin_gte: String
  stringValueMax_gte: String
  stringValue_lte: String
  stringValueMin_lte: String
  stringValueMax_lte: String
  stringValue_in: [String!]
  stringValueMin_in: [String!]
  stringValueMax_in: [String!]
  stringValue_like: String
  stringValueMin_like: String
  stringValueMax_like: String
  stringValue_prefix: String
  stringValueMin_prefix: String
  stringValueMax_prefix: String
  stringValue_suffix: String
  stringValueMin_suffix: String
  stringValueMax_suffix: String
  stringValue_null: Boolean
  floatValue: Float
  floatValueMin: Float
  floatValueMax: Float
  floatValueAvg: Float
  floatValue_ne: Float
  floatValueMin_ne: Float
  floatValueMax_ne: Float
  floatValueAvg_ne: Float
  floatValue_gt: Float
  floatValueMin_gt: Float
  floatValueMax_gt: Float
  floatValueAvg_gt: Float
  floatValue_lt: Float
  floatValueMin_lt: Float
  floatValueMax_lt: Float
  floatValueAvg_lt: Float
  floatValue_gte: Float
  floatValueMin_gte: Float
  floatValueMax_gte: Float
  floatValueAvg_gte: Float
  floatValue_lte: Float
  floatValueMin_lte: Float
  floatValueMax_lte: Float
  floatValueAvg_lte: Float
  floatValue_in: [Float!]
  floatValueMin_in: [Float!]
  floatValueMax_in: [Float!]
  floatValueAvg_in: [Float!]
  floatValue_null: Boolean
  intValue: Int
  intValueMin: Int
  intValueMax: Int
  intValueAvg: Int
  intValue_ne: Int
  intValueMin_ne: Int
  intValueMax_ne: Int
  intValueAvg_ne: Int
  intValue_gt: Int
  intValueMin_gt: Int
  intValueMax_gt: Int
  intValueAvg_gt: Int
  intValue_lt: Int
  intValueMin_lt: Int
  intValueMax_lt: Int
  intValueAvg_lt: Int
  intValue_gte: Int
  intValueMin_gte: Int
  intValueMax_gte: Int
  intValueAvg_gte: Int
  intValue_lte: Int
  intValueMin_lte: Int
  intValueMax_lte: Int
  intValueAvg_lte: Int
  intValue_in: [Int!]
  intValueMin_in: [Int!]
  intValueMax_in: [Int!]
  intValueAvg_in: [Int!]
  intValue_null: Boolean
  definitionId: ID
  definitionIdMin: ID
  definitionIdMax: ID
  definitionId_ne: ID
  definitionIdMin_ne: ID
  definitionIdMax_ne: ID
  definitionId_gt: ID
  definitionIdMin_gt: ID
  definitionIdMax_gt: ID
  definitionId_lt: ID
  definitionIdMin_lt: ID
  definitionIdMax_lt: ID
  definitionId_gte: ID
  definitionIdMin_gte: ID
  definitionIdMax_gte: ID
  definitionId_lte: ID
  definitionIdMin_lte: ID
  definitionIdMax_lte: ID
  definitionId_in: [ID!]
  definitionIdMin_in: [ID!]
  definitionIdMax_in: [ID!]
  definitionId_null: Boolean
  itemId: ID
  itemIdMin: ID
  itemIdMax: ID
  itemId_ne: ID
  itemIdMin_ne: ID
  itemIdMax_ne: ID
  itemId_gt: ID
  itemIdMin_gt: ID
  itemIdMax_gt: ID
  itemId_lt: ID
  itemIdMin_lt: ID
  itemIdMax_lt: ID
  itemId_gte: ID
  itemIdMin_gte: ID
  itemIdMax_gte: ID
  itemId_lte: ID
  itemIdMin_lte: ID
  itemIdMax_lte: ID
  itemId_in: [ID!]
  itemIdMin_in: [ID!]
  itemIdMax_in: [ID!]
  itemId_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
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
  count: Float
  itemId: ID
  definitionId: ID
  parentItemId: ID
}

input ConfiguratorSlotUpdateInput {
  count: Float
  itemId: ID
  definitionId: ID
  parentItemId: ID
}

input ConfiguratorSlotSortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  count: ObjectSortType
  countMin: ObjectSortType
  countMax: ObjectSortType
  countAvg: ObjectSortType
  itemId: ObjectSortType
  itemIdMin: ObjectSortType
  itemIdMax: ObjectSortType
  definitionId: ObjectSortType
  definitionIdMin: ObjectSortType
  definitionIdMax: ObjectSortType
  parentItemId: ObjectSortType
  parentItemIdMin: ObjectSortType
  parentItemIdMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
  item: ConfiguratorItemSortType
  definition: ConfiguratorSlotDefinitionSortType
  parentItem: ConfiguratorItemSortType
}

input ConfiguratorSlotFilterType {
  AND: [ConfiguratorSlotFilterType!]
  OR: [ConfiguratorSlotFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_null: Boolean
  count: Float
  countMin: Float
  countMax: Float
  countAvg: Float
  count_ne: Float
  countMin_ne: Float
  countMax_ne: Float
  countAvg_ne: Float
  count_gt: Float
  countMin_gt: Float
  countMax_gt: Float
  countAvg_gt: Float
  count_lt: Float
  countMin_lt: Float
  countMax_lt: Float
  countAvg_lt: Float
  count_gte: Float
  countMin_gte: Float
  countMax_gte: Float
  countAvg_gte: Float
  count_lte: Float
  countMin_lte: Float
  countMax_lte: Float
  countAvg_lte: Float
  count_in: [Float!]
  countMin_in: [Float!]
  countMax_in: [Float!]
  countAvg_in: [Float!]
  count_null: Boolean
  itemId: ID
  itemIdMin: ID
  itemIdMax: ID
  itemId_ne: ID
  itemIdMin_ne: ID
  itemIdMax_ne: ID
  itemId_gt: ID
  itemIdMin_gt: ID
  itemIdMax_gt: ID
  itemId_lt: ID
  itemIdMin_lt: ID
  itemIdMax_lt: ID
  itemId_gte: ID
  itemIdMin_gte: ID
  itemIdMax_gte: ID
  itemId_lte: ID
  itemIdMin_lte: ID
  itemIdMax_lte: ID
  itemId_in: [ID!]
  itemIdMin_in: [ID!]
  itemIdMax_in: [ID!]
  itemId_null: Boolean
  definitionId: ID
  definitionIdMin: ID
  definitionIdMax: ID
  definitionId_ne: ID
  definitionIdMin_ne: ID
  definitionIdMax_ne: ID
  definitionId_gt: ID
  definitionIdMin_gt: ID
  definitionIdMax_gt: ID
  definitionId_lt: ID
  definitionIdMin_lt: ID
  definitionIdMax_lt: ID
  definitionId_gte: ID
  definitionIdMin_gte: ID
  definitionIdMax_gte: ID
  definitionId_lte: ID
  definitionIdMin_lte: ID
  definitionIdMax_lte: ID
  definitionId_in: [ID!]
  definitionIdMin_in: [ID!]
  definitionIdMax_in: [ID!]
  definitionId_null: Boolean
  parentItemId: ID
  parentItemIdMin: ID
  parentItemIdMax: ID
  parentItemId_ne: ID
  parentItemIdMin_ne: ID
  parentItemIdMax_ne: ID
  parentItemId_gt: ID
  parentItemIdMin_gt: ID
  parentItemIdMax_gt: ID
  parentItemId_lt: ID
  parentItemIdMin_lt: ID
  parentItemIdMax_lt: ID
  parentItemId_gte: ID
  parentItemIdMin_gte: ID
  parentItemIdMax_gte: ID
  parentItemId_lte: ID
  parentItemIdMin_lte: ID
  parentItemIdMax_lte: ID
  parentItemId_in: [ID!]
  parentItemIdMin_in: [ID!]
  parentItemIdMax_in: [ID!]
  parentItemId_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
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
