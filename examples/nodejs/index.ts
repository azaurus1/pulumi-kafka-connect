import * as pulumi from "@pulumi/pulumi";
import * as kafkaconnect from "@pulumi/kafkaconnect";

const myRandomResource = new kafkaconnect.Random("myRandomResource", {length: 24});
const myRandomComponent = new kafkaconnect.RandomComponent("myRandomComponent", {length: 24});
export const output = {
    value: myRandomResource.result,
};
