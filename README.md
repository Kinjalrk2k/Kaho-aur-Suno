# Kahu-aur-Suno

### [Read the blog](https://blog.kinjal.dev/grpc-based-message-broker)

The blog will provide insights about this project!

"Kaho aur Suno" (Hindi: कहो और सुनो, English: "Say and Listen") is a delightful showcase of a gRPC-based message broker. It provides examples of both the broker server and clients, all elegantly written in Go. It's important to note that this application is intended as a Proof of Concept and not suitable for production use.

## Structure

This project has been divided into multiple sub-directories and is managed by a vscode workspace.

- `proto`: Contains protobuf files
- `lib`: Here's the generated proto packages goes in
  - `go`: The proto package for go. Designed as a local package and can be installed in standalone projects
- `scripts/protogen.sh`: Used to generate the proto package
- `server`: Server side implementation. To run the server, hit the command: `go run main.go` in it's directory
- `client`: Client side implementations
  - `go`: Implementation in Go (Root)
    - `publisher`: Publish messages, run: `go run publisher/main.go`
    - `subsciber`: Subscribe messages, run: `go run subscibe/main.go <topic-name>`

## Extending

Right now the proto package and client implementations are only done in Go, however the directories are set up in a way that they can be later extended for other programming languages and/or use cases too!

<hr />

This Message Broker is just a hobby project and not a production ready application at all. There are still multiple sectors in which we can improve it - like persisting messages, improved error handlings, message delivery acknowlegements, delayed delivery, etc. There are various industry grade Message Brokers like [RabbitMQ](https://www.rabbitmq.com/), [ActiveMQ](https://activemq.apache.org/), [Google Pub/Sub](https://cloud.google.com/pubsub?hl=en) etc. which implements such features!
