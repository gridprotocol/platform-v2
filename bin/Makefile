GIT_COMMIT = $(shell git rev-parse --short HEAD)
BUILD_TIME = $(shell TZ=Asia/Shanghai date +'%Y-%m-%d.%H:%M:%S%Z')
BUILD_FLAGS = -ldflags "-X 'github.com/rockiecn/platform/cmd.BuildFlag=$(GIT_COMMIT)+$(BUILD_TIME)'"

platform-v2:
	        go build $(BUILD_FLAGS) -o ../bin ../


clean:
	        rm -rf platform-v2

all: clean platform-v2
