OUTPUT_FOLDER?=./out
DOCKER_PLUGINS_PATH=~/.docker/cli-plugins

.PHONY: build
build: ensure-output-folder-exists
	go build -o ${OUTPUT_FOLDER}/docker-render cmd/docker-render.go

.PHONY: ensure-output-folder-exists
ensure-output-folder-exists:
	mkdir -p ${OUTPUT_FOLDER}

.PHONY: install
install: clean build
	cp ${OUTPUT_FOLDER}/docker-render ${DOCKER_PLUGINS_PATH}/ 
	chmod +x ${DOCKER_PLUGINS_PATH}/docker-render

.PHONY: clean
clean: clean-docker-plugin clean-output-folder

.PHONY: clean-docker-plugin
clean-docker-plugin:
	rm -rf ${DOCKER_PLUGINS_PATH}/docker-render

.PHONY: clean-output-folder
clean-output-folder:
	rm -rf ${OUTPUT_FOLDER}