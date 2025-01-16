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

    var myConnector = new Kafkaconnect.Connector.Connector("myConnector", new()
    {
    }, new CustomResourceOptions
    {
        Provider = defaultProvider,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myConnector.Result },
        },
    };
});

