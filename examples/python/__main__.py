import pulumi
import pulumi_kafkaconnect as kafkaconnect

my_random_resource = kafkaconnect.Random("myRandomResource", length=24)
my_random_component = kafkaconnect.RandomComponent("myRandomComponent", length=24)
pulumi.export("output", {
    "value": my_random_resource.result,
})
