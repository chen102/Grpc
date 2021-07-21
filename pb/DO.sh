#!/bin/bash
protoc --go_out=plugins=grpc,paths=source_relative:. *.proto

#-I ../google/protobuf/ -I .  #import 直接导入prot名称，编译器会在-I参数指定的目录查找导入的文件
#paths参数有两个选项import和source_relative。默认为import，代表按照生成go代码的包的全路径去创建层级，source_relative代表安装proto源文件的目录层级去创建go代码的目录层级，简单来说，就是使用source_relative则不会使用option go_package,默认则是使用option go_package

