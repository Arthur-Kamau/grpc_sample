protoc \
    --plugin="protoc-gen-ts=./node_modules/.bin/protoc-gen-ts" \
    --js_out="import_style=commonjs,binary:./src/gen" \
    --ts_out="./src/gen" \
    --proto_path=./proto \
    app.proto