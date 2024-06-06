## Using Wharf as Docker CLI Plugin

Docker CLI plugins is used to extend the cappabilities of the Docker CLI. Wharf provide a small list of Docker CLI plugins. You can use them as `docker <plugin>`

## Install
To install all plugins, you can use the following command from [the project's root.](../)

```bash
make install-plugins
```

## List of avaialble plugins:
- [render](./render/README.md): Create a Dockerfile from a template