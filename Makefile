GOTIFY_VERSION:=$(shell curl -sL https://api.github.com/repos/gotify/server/releases/latest | jq -r ".tag_name")
GOSWAGGER_VERSION:=$(shell curl -sL https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r ".tag_name")
BUILD=./build
TEMP_SPEC=${BUILD}/gotify.json
SWAGGER=./.tools/swagger

ifeq ($(OS),Windows_NT)
    DOWNLOAD_SWAGGER=swagger_windows_amd64.exe
else
    DOWNLOAD_SWAGGER=swagger_linux_amd64
endif

clean-tools:
	rm -rf .tools

get-tools: clean-tools
	mkdir -p .tools
	wget -O ${SWAGGER} https://github.com/go-swagger/go-swagger/releases/download/${GOSWAGGER_VERSION}/${DOWNLOAD_SWAGGER}
	chmod u+x .tools/*

install:
	go mod download

obtain-spec:
	mkdir build || true
	wget -O ${TEMP_SPEC} https://raw.githubusercontent.com/gotify/server/${GOTIFY_VERSION}/docs/spec.json

generate: obtain-spec
	${SWAGGER} generate client -f ${TEMP_SPEC} --additional-initialism=rest --skip-models

clean:
	rm -rf client build