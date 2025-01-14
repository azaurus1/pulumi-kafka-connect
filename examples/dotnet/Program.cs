using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Kafkaconnect = Pulumi.Kafkaconnect;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new Kafkaconnect.Random("myRandomResource", new()
    {
        Length = 24,
    });

    var myRandomComponent = new Kafkaconnect.RandomComponent("myRandomComponent", new()
    {
        Length = 24,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myRandomResource.Result },
        },
    };
});

