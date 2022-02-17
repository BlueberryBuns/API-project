# API-project


### Golang applications

In order to start golang application safely with in Go version `1.18` (in beta version as of 02/17/2022) create docker image based on Dockerfile attatched on utils with command.
```bash
docker build [PATH/TO/DOCKERFILE] -t [NAME-OF-YOUR-CONTAINER]
```
The idea behind the command is to crate utility container used later to create project structure and install packages with golang and generally execute go commands in golang `v1.18`.
```bash
docker run -it --rm -v $(pwd):/go [NAME-OF-YOUR-CONTAINER] go [COMMAND(s)]

# This project was generate with the command listed below 
docker run -it --rm -v $(pwd):/go golang-util go mod init \ 
github.com/BlueberryBuns/API-project/backend/go-api
```
To run golang project build image in `backend/go-api` directory and run docker container. The container is based on multi-staged image so it shall take minimal amount of space as it's outcome is based only on the `latest` image version of `alpine` Linux.
```bash
docker build [PATH/TO/DOCKERFILE] -t [NAME-OF-YOUR-CONTAINER]

docekr run -it --rm [NAME-OF-YOUR-CONTAINER]
```