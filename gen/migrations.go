package gen

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func Migrate(db *gorm.DB, options *gormigrate.Options, migrations []*gormigrate.Migration) error {
	m := gormigrate.New(db, options, migrations)

	// // it's possible to use this, but in case of any specific keys or columns are created in migrations, they will not be generated by automigrate
	// m.InitSchema(func(tx *gorm.DB) error {
	// 	return AutoMigrate(db)
	// })

	return m.Migrate()
}

func AutoMigrate(db *gorm.DB) error {
	_db := db.AutoMigrate(
		ConfiguratorItemDefinitionCategory{},
		ConfiguratorItemDefinition{},
		ConfiguratorAttributeDefinition{},
		ConfiguratorSlotDefinition{},
		ConfiguratorItem{},
		ConfiguratorAttribute{},
		ConfiguratorSlot{},
	)

	if _db.Dialect().GetName() != "sqlite3" {

		_db.Model(ConfiguratorAttributeDefinition_definitions{}).RemoveForeignKey("definition_id", TableName("configurator_item_definitions")+"(id)")
		_db = _db.Model(ConfiguratorAttributeDefinition_definitions{}).AddForeignKey("definition_id", TableName("configurator_item_definitions")+"(id)", "CASCADE", "CASCADE")

		_db.Model(ConfiguratorSlotDefinition_allowedItemDefinitions{}).RemoveForeignKey("allowedItemDefinition_id", TableName("configurator_item_definitions")+"(id)")
		_db = _db.Model(ConfiguratorSlotDefinition_allowedItemDefinitions{}).AddForeignKey("allowedItemDefinition_id", TableName("configurator_item_definitions")+"(id)", "CASCADE", "CASCADE")

		_db.Model(ConfiguratorItemDefinition{}).RemoveForeignKey("categoryId", TableName("configurator_item_definition_categories")+"(id)")
		_db = _db.Model(ConfiguratorItemDefinition{}).AddForeignKey("categoryId", TableName("configurator_item_definition_categories")+"(id)", "SET NULL", "SET NULL")

		_db.Model(ConfiguratorAttributeDefinition_definitions{}).RemoveForeignKey("attribute_id", TableName("configurator_attribute_definitions")+"(id)")
		_db = _db.Model(ConfiguratorAttributeDefinition_definitions{}).AddForeignKey("attribute_id", TableName("configurator_attribute_definitions")+"(id)", "CASCADE", "CASCADE")

		_db.Model(ConfiguratorSlotDefinition{}).RemoveForeignKey("definitionId", TableName("configurator_item_definitions")+"(id)")
		_db = _db.Model(ConfiguratorSlotDefinition{}).AddForeignKey("definitionId", TableName("configurator_item_definitions")+"(id)", "SET NULL", "SET NULL")

		_db.Model(ConfiguratorSlotDefinition_allowedItemDefinitions{}).RemoveForeignKey("allowedInSlot_id", TableName("configurator_slot_definitions")+"(id)")
		_db = _db.Model(ConfiguratorSlotDefinition_allowedItemDefinitions{}).AddForeignKey("allowedInSlot_id", TableName("configurator_slot_definitions")+"(id)", "CASCADE", "CASCADE")

		_db.Model(ConfiguratorItem{}).RemoveForeignKey("definitionId", TableName("configurator_item_definitions")+"(id)")
		_db = _db.Model(ConfiguratorItem{}).AddForeignKey("definitionId", TableName("configurator_item_definitions")+"(id)", "RESTRICT", "SET NULL")

		_db.Model(ConfiguratorAttribute{}).RemoveForeignKey("definitionId", TableName("configurator_attribute_definitions")+"(id)")
		_db = _db.Model(ConfiguratorAttribute{}).AddForeignKey("definitionId", TableName("configurator_attribute_definitions")+"(id)", "RESTRICT", "SET NULL")

		_db.Model(ConfiguratorAttribute{}).RemoveForeignKey("itemId", TableName("configurator_items")+"(id)")
		_db = _db.Model(ConfiguratorAttribute{}).AddForeignKey("itemId", TableName("configurator_items")+"(id)", "SET NULL", "SET NULL")

		_db.Model(ConfiguratorSlot{}).RemoveForeignKey("itemId", TableName("configurator_items")+"(id)")
		_db = _db.Model(ConfiguratorSlot{}).AddForeignKey("itemId", TableName("configurator_items")+"(id)", "SET NULL", "SET NULL")

		_db.Model(ConfiguratorSlot{}).RemoveForeignKey("definitionId", TableName("configurator_slot_definitions")+"(id)")
		_db = _db.Model(ConfiguratorSlot{}).AddForeignKey("definitionId", TableName("configurator_slot_definitions")+"(id)", "RESTRICT", "SET NULL")

		_db.Model(ConfiguratorSlot{}).RemoveForeignKey("parentItemId", TableName("configurator_items")+"(id)")
		_db = _db.Model(ConfiguratorSlot{}).AddForeignKey("parentItemId", TableName("configurator_items")+"(id)", "SET NULL", "SET NULL")

	}
	return _db.Error
}
