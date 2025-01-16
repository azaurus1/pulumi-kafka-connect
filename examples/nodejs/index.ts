import * as pulumi from "@pulumi/pulumi";
import * as kafkaconnect from "@pulumi/kafkaconnect";

const defaultProvider = new kafkaconnect.Provider("defaultProvider", {url: "http://localhost:8083"});
const mirrorHeartbeatConnectorYxdw = new kafkaconnect.connector.Connector("mirror-heartbeat-connector-yxdw", {config: {
    "connector.class": "org.apache.kafka.connect.mirror.MirrorHeartbeatConnector",
    "source.cluster.alias": "source",
    "heartbeats.topic.replication.factor": "-1",
    name: "mirror-heartbeat-connector-yxdw",
}}, {
    provider: defaultProvider,
});
export const output = {
    value: mirrorHeartbeatConnectorYxdw.result,
};
