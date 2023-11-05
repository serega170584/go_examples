1. docker build --tag docker_ubuntu .
2. docker run -dit --privileged --name=docker_ubuntu --volume=/Users/krivobokovsergej/GolandProjects/go_examples/docker_ubuntu:/app docker_ubuntu
3. docker exec -it docker_ubuntu bash    
4. sudo sysctl -w kernel.yama.ptrace_scope=0 in container
