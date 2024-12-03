GIT_COMMIT = $(shell git rev-parse --short HEAD)
BUILD_TIME = $(shell TZ=Asia/Shanghai date +'%Y-%m-%d.%H:%M:%S%Z')
BUILD_FLAGS = -ldflags "-X 'github.com/gridprotocol/platform-v2/cmd.BuildFlag=$(GIT_COMMIT)+$(BUILD_TIME)'"

all: clean platform-v2

platform-v2:
	go build $(BUILD_FLAGS) -o ./ ../

clean:
	rm -rf platform-v2