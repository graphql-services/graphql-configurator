Feature: Field xxxId should be automatically used for filling object for xxx field

    Background: We have basic definitions
        Given I send query:
            """
            mutation {
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
                    attributesIds: ["weight"]
                    slotsIds: ["wheel"]
                    }
                ) {
                    id
                }
            }
            """

    Scenario: Fetching country should use the countryId field as id
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
                    }
                    }
                ) {
                    id
                    item {
                    code
                    name
                    definitionId
                    attributes {
                        id
                        definitionId
                        intValue
                    }
                    }
                }
            }
            """
        Then the response should be:
            """
            {
                "company": {
                    "id": "test",
                    "countryId": "xxx",
                    "country": {
                        "id": "xxx"
                    }
                }
            }
            """
