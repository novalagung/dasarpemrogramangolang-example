# Dockerize Hello World

## Usage

```bash
# build image
docker build -t my-image-hello-world .
# create container from above image
docker container create --name my-container-hello-world -e PORT=8080 -e INSTANCE_ID="my first instance" -p 8080:8080 my-image-hello-world
# run newly created container
docker container start my-container-hello-world
```

Or

```bash
# create container then run it, then auto delete if stopped
docker container run --name my-container-hello-world --rm -itd -e PORT=8080 -e INSTANCE_ID="my first instance" -p 8080:8080 my-image-hello-world
```