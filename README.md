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

Here i added the screenshots of the images
 1) without using Multi stage images,
    
    ![image](https://github.com/sathishkumar-2001/Multi_stage_docker_build/assets/126504329/e68a9f6b-7361-4900-b9f1-c4cf5011fcba)
    

    ![image](https://github.com/sathishkumar-2001/Multi_stage_docker_build/assets/126504329/65de0ec0-8db6-4630-b445-aef8b7ef0861)
 2) By using Multi stage docker build,

    ![image](https://github.com/sathishkumar-2001/Multi_stage_docker_build/assets/126504329/be080670-8009-4118-b440-503f7488a2a2)

    ![image](https://github.com/sathishkumar-2001/Multi_stage_docker_build/assets/126504329/aed1836b-293f-4219-8df8-a1bf5899df9e)

we can clearly see the magic of how much the image size has been reduced and also the time taken to build the image also decreases when 
you are using the Multi stage docker build .


