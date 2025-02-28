// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Azaurus1.KafkaConnect
{
    [KafkaConnectResourceType("pulumi:providers:kafka-connect")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The url for the kafka connect cluster
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        [Output("user")]
        public Output<string?> User { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs args, CustomResourceOptions? options = null)
            : base("kafka-connect", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/azaurus1/pulumi-kafka-connect",
                AdditionalSecretOutputs =
                {
                    "password",
                    "user",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        [Input("password")]
        private Input<string>? _password;
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The url for the kafka connect cluster
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        [Input("user")]
        private Input<string>? _user;
        public Input<string>? User
        {
            get => _user;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _user = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ProviderArgs()
        {
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
