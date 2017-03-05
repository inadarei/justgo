# Jump-Start Go (JustGo)

Skeleton project for jump-starting a Go-powered microservice development with Docker and Go best-practices.

To learn more: [https://justgo.rocks](https://justgo.rocks)

## INSTALLATION 

Easiest way to create a new project skeleton is to install JustGo CLI tool. There's no necessity to install Go on your machine, since the setup provides fully functioning Go environment in a Docker container. 

If you already have Go on your machine, you can install the CLI tool with:

```
> go get github.com/inadarei/justgo/cmd/justgo
```

or you can install it using Homebrew, even if you don't have Go:

```
> brew tap inadarei/casks
> brew install justgo
```

## USAGE

After you have installed the CLI tool, to create a skeleton of a new project, just run:

```
> justgo
```

You can see various options by running `justgo -h`

## How to run a project, once created:

### Short version:

```
`docker-compose up -d`
```

### Longer explanation:

1. Install a working Docker environment
    1. Mac: https://docs.docker.com/docker-for-mac/
    2. Windows: https://docs.docker.com/docker-for-windows/
    3. Linux: https://docs.docker.com/engine/installation/linux/
2. cd to the project's root folder and run `docker-compose build --no-cache` (optional but good step)
3. In the same folder, run: `docker-compose up -d`
4. If you get a clean output, you can check which port the server
   attached to by running: `docker-compose ps`
4. For instance, if the output was: `0.0.0.0:32791` under `ports` section then you
   can access your new service at `http://0.0.0.0:32791/`
5. You can edit source files of the application without restarting anything
   since JustGo supports hot code reloading.

## Stopping and removing a container

While in the project folder:

```
> docker-compose stop 
> docker-compose rm -f 
```

## License

[MIT](LICENSE)
