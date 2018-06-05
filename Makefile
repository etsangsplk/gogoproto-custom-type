.PHONY: proto
proto:
	protoc \
		-I jsontest \
		-I vendor/github.com/gogo/googleapis/ \
		-I vendor/ \
		--gogo_out=plugins=grpc,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:\
./jsontest/ \
		test_objects.proto
