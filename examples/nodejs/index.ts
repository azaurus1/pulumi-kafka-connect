import * as pulumi from "@pulumi/pulumi";
import * as kafkaconnect from "@pulumi/kafkaconnect";

const defaultProvider = new kafkaconnect.Provider("defaultProvider", {url: "http://localhost:8083"});
const myConnector = new kafkaconnect.connector.Connector("myConnector", {}, {
    provider: defaultProvider,
});
export const output = {
    value: myConnector.result,
};
