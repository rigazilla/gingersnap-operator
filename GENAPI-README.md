# Generate API from proto definition

## Requisite
dnf install protoc
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install github.com/protobuf-tools/protoc-gen-deepcopy@v0.0.3


In the operator root dir:

git clone git@github.com:gingersnap-project/api.git gingersnap-api
protoc --proto_path=gingersnap-api --go_out ./gen --include_source_info --descriptor_set_out=gen/gingersnap-api/descriptor --deepcopy_out=gen config/cache/v1alpha/cache.proto config/cache/v1alpha/region.proto


