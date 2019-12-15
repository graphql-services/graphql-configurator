Feature: Creating configuration items using assemblies

    Background: We have basic definitions
        Given I send query:
            """
            mutation {
                deleteAllConfiguratorItems
                deleteAllConfiguratorSlots
                deleteAllConfiguratorAttributes
                deleteAllConfiguratorItemDefinitions
                deleteAllConfiguratorAttributeDefinitions
                deleteAllConfiguratorSlotDefinitions
                createConfiguratorAttributeDefinition(
                    input: { id: "weight", name: "Weight", type: INT }
                ) {
                    id
                }
                createConfiguratorSlotDefinition(
                    input: { id: "wheel", name: "Wheel", minCount: 4, maxCount: 8 }
                ) {
                    id
                }
                car: createConfiguratorItemDefinition(
                    input: {
                    id: "car"
                    name: "Car"
                    attributesIds: ["weight"]
                    slotsIds: ["wheel"]
                    }
                ) {
                    id
                }
                wheel: createConfiguratorItemDefinition(
                    input: {
                    id: "wheel"
                    name: "Wheel"
                    }
                ) {
                    id
                }
            }
            """

    Scenario: Create configuration assembly
        When I send query:
            """
            mutation {
                createConfiguratorAssembly(
                    input: {
                    item: {
                        code: "tesla"
                        name: "Tesla"
                        definitionId: "car"
                        attributes: [{ definitionId: "weight", intValue: 1234 }]
                        slots: [{ definitionId: "wheel", item: { definitionId:"wheel", code: "lf" } }]
                    }
                    }
                ) {
                    item {
                    code
                    name
                    definitionId
                    attributes {
                        definitionId
                        intValue
                    }
                    slots {
                        definitionId
                        item {
                        code
                        name
                        }
                    }
                    }
                }
            }
            """
        Then the response should be:
            """
            {
                "createConfiguratorAssembly": {
                    "item": {
                        "code": "tesla",
                        "name": "Tesla",
                        "definitionId": "car",
                        "attributes": [
                            {
                                "definitionId": "weight",
                                "intValue": 1234
                            }
                        ],
                        "slots": [
                            {
                                "definitionId": "wheel",
                                "item": {
                                    "code": "lf",
                                    "name": null
                                }
                            }
                        ]
                    }
                }
            }
            """
