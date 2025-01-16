using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Kafkaconnect = Pulumi.Kafkaconnect;

return await Deployment.RunAsync(() => 
{
    var defaultProvider = new Kafkaconnect.Provider("defaultProvider", new()
    {
        Url = "http://localhost:8083",
    });

    var mirrorHeartbeatConnectorYxdw = new Kafkaconnect.Connector.Connector("mirror-heartbeat-connector-yxdw", new()
    {
        Config = 
        {
            { "connector.class", "org.apache.kafka.connect.mirror.MirrorHeartbeatConnector" },
            { "source.cluster.alias", "source" },
            { "heartbeats.topic.replication.factor", "-1" },
            { "name", "mirror-heartbeat-connector-yxdw" },
        },
    }, new CustomResourceOptions
    {
        Provider = defaultProvider,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", mirrorHeartbeatConnectorYxdw.Result },
        },
    };
});

