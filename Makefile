OUTPUT_FOLDER?=./out
PLUGINS_OUTPUT_FOLDER?=${OUTPUT_FOLDER}/docker-plugins
DOCKER_PLUGINS_PATH=~/.docker/cli-plugins

.PHONY: build
build: ensure-output-folder-exists
	go build -o ${OUTPUT_FOLDER}/wharf cli/wharf.go

.PHONY: build-plugins
build-plugins: ensure-plugins-output-folder-exists build-docker-render 

.PHONY: build-docker-render
build-docker-render:
	go build -o ${PLUGINS_OUTPUT_FOLDER}/docker-render docker-plugins/render/docker_render.go

.PHONY: ensure-output-folder-exists
ensure-output-folder-exists:
	mkdir -p ${OUTPUT_FOLDER}

.PHONY: ensure-plugins-output-folder-exists
ensure-plugins-output-folder-exists:
	mkdir -p ${PLUGINS_OUTPUT_FOLDER}
.PHONY: install-plugins
install-plugins: clean build-plugins
	chmod +x ${PLUGINS_OUTPUT_FOLDER}/docker-*
	cp ${PLUGINS_OUTPUT_FOLDER}/* ${DOCKER_PLUGINS_PATH}/

.PHONY: test
test:
	go test -v ./core

.PHONY: clean
clean: clean-output-folder

.PHONY: clean-output-folder
clean-output-folder:
	rm -rf ${OUTPUT_FOLDER}