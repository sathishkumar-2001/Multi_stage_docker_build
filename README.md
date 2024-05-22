# Multi_stage_docker_build
how  we can reduce the size of the docker image by using multi stage builds
In this repo we can see how to reduce the docker image size:
i provided some of the examples on this in the repo.
 1) we need to divide the dockerfile as for build and runtime
 2) we can use very rich base image for build that solves all dependencies for the application
 3) for the runtime we can use minimalastice images like distroless images to run our binaries from the base image
 4) by this we can reduce the image size
    
it's not like we must use only two stages it can go upto n build stages based on the application 

Distroless images offer a streamlined, secure, and efficient way to deploy applications in containers.
By focusing only on the essentials, distroless images minimize the attack surface, reduce image size, and improve performance,
making them an excellent choice for production deployments and environments where security and efficiency are crucial.
