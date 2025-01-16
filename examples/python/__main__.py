import pulumi
import pulumi_kafkaconnect as kafkaconnect

default_provider = kafkaconnect.Provider("defaultProvider", url="http://localhost:8083")
my_connector = kafkaconnect.connector.Connector("myConnector", opts = pulumi.ResourceOptions(provider=default_provider))
pulumi.export("output", {
    "value": my_connector.result,
})
