#! /bin/bash
echo "$(cd "$(dirname "$0")";pwd)"
cd $(dirname "$0") # SHELL_FOLDER
for file in $(ls)
do
  if [ "${file##*.}" = "proto" ]; then
    if [ ! -d "../grpc/${file%.*}" ]; then
      mkdir -p ../grpc/${file%.*}
    fi
    #if [ ! -d "../web/grpc/${file%.*}" ]; then
    #  mkdir -p ../web/grpc/${file%.*}
    #fi
    OUT_GO_DIR=../grpc/${file%.*}
    #OUT_WEB_DIR=../web/grpc/${file%.*}
    protoc $file --go_out=plugins=grpc:$OUT_GO_DIR
    #protoc $file --gofast_out=plugins=grpc:$OUT_GO_DIR
    #protoc $file --js_out=import_style=commonjs:$OUT_WEB_DIR --grpc-web_out=import_style=commonjs,mode=grpcwebtext:$OUT_WEB_DIR
    echo "generate $file"
  fi
done