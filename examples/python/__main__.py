import pulumi
import pulumi_kafkaconnect as kafkaconnect

default_provider = kafkaconnect.Provider("defaultProvider", url="http://localhost:8083")
mirror_heartbeat_connector_yxdw = kafkaconnect.connector.Connector("mirror-heartbeat-connector-yxdw", config={
    "connector.class": "org.apache.kafka.connect.mirror.MirrorHeartbeatConnector",
    "source.cluster.alias": "source",
    "heartbeats.topic.replication.factor": "-1",
    "name": "mirror-heartbeat-connector-yxdw",
},
opts = pulumi.ResourceOptions(provider=default_provider))
pulumi.export("output", {
    "value": mirror_heartbeat_connector_yxdw.result,
})
