## Docker Render

### Install the Docker plugin

Use the following Makefile command to install the Docker plugin in your current Docker CLI installation

```bash
make install
```

Test the installation using `docker render` from your Terminal.

### Usage

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
