# API
An interface that sets up a server to commicate with the OpenGo AI developed by the Facebook team.

## **Dependencies**

* Go 1.10.3 darwin/amd64
* Python 3.6.5

Download Go [here](https://golang.org/dl/)

After installing Go by the instructions, run ```export PATH=$PATH:$(go env GOPATH)/bin``` and ```export GOPATH=$(go env GOPATH)```.

Install grpc for go - run ```go get -u google.golang.org/grpc```, or follow the instructions on [this page](https://grpc.io/docs/quickstart/go.html).

```go get github.com/nu7hatch/gouuid```

Install grpc for python - run ```python -m pip install grpcio``` and then install grpc tools with ```python -m pip install grpcio-tools googleapis-common-protos```, or follow the instructions on [this page](https://grpc.io/docs/quickstart/python.html).


## **Setup**

Run ```git clone https://github.com/Jachua/ELF-API.git $(go env GOPATH)/src/ELF-API```.

Set up the environment for OpenGo by the instructions posted on [this repo](https://github.com/Jachua/ELF)

Open 3 terminals, each navigate to the project root. In the first terminal, run```go run server/server.go``` . In another terminal, run ```source scripts/devmode_set_pythonpath.sh``` and navigate to ```scripts/elfgames/go```. Activate the AI engine by running ```./gtp.sh path/to/modelfile.bin --verbose --gpu 0 --num_block 20 --dim 224 --mcts_puct 1.50 --batchsize 16 --mcts_rollout_per_batch 16 --mcts_threads 2 --mcts_rollout_per_thread 8192 --resign_thres 0.05 --mcts_virtual_loss 1```. 

*As noted in the original repo, ```mcts_rollout_per_thread``` can be modified to tune the AI response rate. 

Edit the ```game_server``` field in ```client/server_addrs.py``` to be the IP address where you want to run the game server. If unspecified, the game will run on local address. The address provided must match the server address for ELF.

Run ```python client/elf.py``` and enter commands when prompted. The console will then display in turn the move received from the console and the move by OpenGO. 


## **Citations**
@misc{ELFOpenGo2018,
  author = {Yuandong Tian and {Jerry Ma*} and {Qucheng Gong*} and Shubho Sengupta and Zhuoyuan Chen and C. Lawrence Zitnick},
  title = {ELF OpenGo},
  year = {2018},
  journal = {GitHub repository},
  howpublished = {\url{https://github.com/pytorch/ELF}}
}
