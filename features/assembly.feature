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
                    input: { id: "wheel", name: "Wheel", minCount: 4, maxCount: 8, defaultCount: 5 }
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
                        id: "tesla"
                        code: "tesla"
                        name: "Tesla"
                        definitionId: "car"
                        attributes:[{
                            definitionId:"weight",stringValue:"aaa"
                        }]
                        slots:[{
                            definitionId:"wheel",
                            count:4,
                            item:{
                                definitionId:"wheel",
                                code:"wheel",
                                name:"Right front"
                            }
                        }]
                    }
                    }
                ) {
                    item {
                    id
                    code
                    name
                    definitionId
                    attributes {
                        definitionId
                        stringValue
                    }
                    slots {
                        count
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
                        "id": "tesla",
                        "code": "tesla",
                        "name": "Tesla",
                        "definitionId": "car",
                        "attributes": [
                            {
                                "definitionId": "weight",
                                "stringValue": "aaa"
                            }
                        ],
                        "slots": [
                            {
                                "count": 4
                            }
                        ]
                    }
                }
            }
            """

# Scenario: Create configuration assembly from template
#     When I send query:
#         """
#         mutation {
#             car1: createConfiguratorAssembly(
#             input: { item: { id: "tesla", code: "tesla", name: "Tesla",definitionId:"car" } }
#         ) {
#             item {
#             id
#             code
#             name
#             definitionId
#             }
#         }
#         car2: createConfiguratorAssembly(
#             input: { item: { id: "new_car", templateId: "tesla" } }
#         ) {
#             item {
#             id
#             code
#             name
#             definitionId
#             attributes {
#                 definitionId
#                 intValue
#             }
#             slots {
#                 definitionId
#                 item {
#                 code
#                 name
#                 }
#             }
#             }
#         }
#         }
#         """
#     Then the response should be:
#         """
#         {
#             "car1": {
#                 "item": {
#                     "id": "tesla",
#                     "code": "tesla",
#                     "name": "Tesla",
#                     "definitionId": "car"
#                 }
#             },
#             "car2": {
#                 "item": {
#                     "id": "tesla",
#                     "code": "tesla",
#                     "name": "Tesla",
#                     "definitionId": "car",
#                     "attributes": [],
#                     "slots": []
#                 }
#             }
#         }
#         """
