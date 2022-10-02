module github.com/batchcorp/plumber

go 1.16

replace github.com/v2pro/plz => github.com/batchcorp/plz v0.9.2

require (
	cloud.google.com/go/pubsub v1.6.1
	github.com/Azure/azure-event-hubs-go/v3 v3.3.16
	github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus v1.1.0
	github.com/Azure/go-amqp v0.16.4 // indirect
	github.com/Masterminds/semver v1.5.0
	github.com/apache/pulsar-client-go v0.7.0
	github.com/aws/aws-sdk-go v1.34.28
	github.com/batchcorp/collector-schemas v0.0.16
	github.com/batchcorp/kong v0.2.17-batch-fix
	github.com/batchcorp/natty v0.0.16
	github.com/batchcorp/pgoutput v0.3.2
	github.com/batchcorp/plumber-schemas v0.0.168
	github.com/batchcorp/rabbit v0.1.17
	github.com/batchcorp/thrifty v0.0.10
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/fatih/color v1.12.0 // indirect
	github.com/form3tech-oss/jwt-go v3.2.3+incompatible // indirect
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-stomp/stomp v2.1.4+incompatible
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/hokaccha/go-prettyjson v0.0.0-20210113012101-fb4e108d2519
	github.com/imdario/mergo v0.3.13
	github.com/jackc/pgx v3.2.0+incompatible
	github.com/jhump/protoreflect v1.10.1
	github.com/julienschmidt/httprouter v1.3.0
	github.com/kataras/tablewriter v0.0.0-20180708051242-e063d29b7c23
	github.com/kubemq-io/kubemq-go v1.7.2
	github.com/lensesio/tableprinter v0.0.0-20201125135848-89e81fc956e7
	github.com/linkedin/goavro/v2 v2.9.8
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/maxbrunsfeld/counterfeiter/v6 v6.3.0
	github.com/mcuadros/go-lookup v0.0.0-20200831155250-80f87a4fa5ee
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/nats-io/nats-streaming-server v0.24.1 // indirect
	github.com/nats-io/nats.go v1.13.1-0.20220121202836-972a071d373d
	github.com/nats-io/nkeys v0.3.0
	github.com/nats-io/stan.go v0.10.2
	github.com/nsqio/go-nsq v1.0.8
	github.com/olekukonko/tablewriter v0.0.5
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.20.2
	github.com/pkg/errors v0.9.1
	github.com/posthog/posthog-go v0.0.0-20220817142604-0b0bbf0f9c0f
	github.com/prometheus/client_golang v1.11.0
	github.com/rabbitmq/rabbitmq-stream-go-client v1.0.1-rc.2
	github.com/relistan/go-director v0.0.0-20200406104025-dbbf5d95248d
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/segmentio/kafka-go v0.4.16
	github.com/sirupsen/logrus v1.8.1
	github.com/streadway/amqp v1.0.0
	github.com/tidwall/gjson v1.9.3
	go.mongodb.org/mongo-driver v1.7.3
	golang.org/x/crypto v0.0.0-20220511200225-c6db032c6c88
	google.golang.org/api v0.29.0
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.28.1
	nhooyr.io/websocket v1.8.7 // indirect
)
