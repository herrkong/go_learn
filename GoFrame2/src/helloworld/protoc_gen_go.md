

1 protoc --go_out=plugins=grpc:. *.proto

2 option go_package = "path;name";

path 表示生成的go文件的存放地址，会自动生成目录的。
name 表示生成的go文件所属的包名



protoc --proto_path= --go_out=./ --go_opt=paths=source_relative data.proto 