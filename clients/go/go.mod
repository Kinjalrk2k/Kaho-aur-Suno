module github.com/kinjalrk2k/Kaho-aur-Suno/clients/go

go 1.21.4

replace github.com/kinjalrk2k/Kaho-aur-Suno/lib/go => ../../lib/go

require (
	github.com/kinjalrk2k/Kaho-aur-Suno/lib/go v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.62.1
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
