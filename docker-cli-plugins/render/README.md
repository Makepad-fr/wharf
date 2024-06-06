# `docker render`

## Install the Docker plugin

Use the following command to install the Docker plugin in your current Docker CLI installation

```bash
make install-plugins
```

Test the installation using `docker render` from your Terminal.

## Usage

Once you've installed the Docker plugin you can use as following

```bash
docker render [OPTIONS] <context_path>
```

With the following options:

|   Name    |                                                               Description                                                                | Required |    Default value    |
| :-------: | :--------------------------------------------------------------------------------------------------------------------------------------: | :------: | :-----------------: |
| file-name |                                                      The name of the template file                                                       |  False   | Dockerfile.template |
|  output   |                         The name of the outputfile, if it's an empty string the result will be printed on stdout                         |  False   |                     |
|  values   | The path for the values file. If an absolute path is passed context path will be ignored. Otherwise it will be joined to the conext path |  False   | docker-values.yaml  |


## Example

Once you've installed the Docker plugin you can follow the below steps:

- Create a `Dockerfile.template` with the following content
    ```
    # Use an official base image
    FROM {{.BaseImage}}

    # Set the exposed port
    EXPOSE {{.Deployment.ExposedPort}}

    # Set environment variables
    {{range .Deployment.Environment}}
    ENV {{.Name}}="{{.Value}}"
    {{end}}

    # Copy source code
    COPY . {{.CopyDestination}}

    # Run command
    CMD {{.RunCommand}}
    ```
- Create a `docker-values.yaml` file with the following content:
    ```yaml
    BaseImage: "ubuntu:latest"
    Deployment:
    ExposedPort: "9091"
    Environment:
        - Name: "DEBUG"
        Value: "false"
        - Name: "LOG_LEVEL"
        Value: "info"
    CopyDestination: "/app"
    RunCommand: "echo Hello World"
    ```
- Create the Dockerfile with
    ```bash
    docker render -output Dockerfile .
    ```
- Build the Docker image
  ```bash
  docker build . -t rendered-docker-image
  ```
- Run the container
  ```bash
  docker run -it rendered-docker-image
  ```
  This shoould output `Hello World` on your terminal.
