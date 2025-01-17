---
title: Kafka Connect Setup
meta_desc: Information on how to install the Kafka Connect Provider for Pulumi.
layout: package
---

## Installation
The Kafka Connect Pulumi provider is available as a package in all Pulumi languages:
- JavaScript/TypeScript: [`@azaurus/kafkaconnect`]
- Python: [`pulumi-kafkaconnect`]
- Go: [`github.com/azaurus1/pulumi-kafka-connect/sdk/go/kafkaconnect`]
- .NET: [Azaurus1.Kafkaconnect]

## Configuration
The following configuration points are available for the `kafka-connect` provider:
- `kafkaconnect:url` - This is the URL of the Connect cluster (Env: `KAFKA_CONNECT_URL`)
