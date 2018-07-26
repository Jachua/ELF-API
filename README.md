# API
An interface that sets up a server to commicate with the OpenGo AI developed by the Facebook team.

## **Dependencies**

* Go 1.10.3 darwin/amd64
* Python 3.6.5

Download Go [here](https://golang.org/dl/)

After installing Go by the instructions, run ```export PATH=$PATH:$(go env GOPATH)/bin``` and ```export GOPATH=$(go env GOPATH)```.

To install grpc, run ```go get -u google.golang.org/grpc```.



## **Setup**

Run ```git clone https://github.com/Jachua/ELF.git $(go env GOPATH)/src/ELF```.

Set up the environment for OpenGo by the instructions posted on [pytorch/ELF](https://github.com/pytorch/ELF)

Open 3 terminals, each navigate to the project root. In the first terminal, run```go run server/server.go``` . In another terminal, run ```source scripts/devmode_set_pythonpath.sh``` and navigate to ```scripts/elfgames/go```. Activate the AI engine by running ```./gtp.sh '../../../pretrained-go-19x19-v0.bin' --verbose --gpu 0 --num_block 20 --dim 224 --mcts_puct 1.50 --batchsize 16 --mcts_rollout_per_batch 16 --mcts_threads 2 --mcts_rollout_per_thread 64 --resign_thres 0.05 --mcts_virtual_loss 1```. 

Run ```python client/elf.py``` and enter commands when prompted. The console will then display in turn the move received from the console and the move by OpenGO. 


## **Citations**
@misc{ELFOpenGo2018,
  author = {Yuandong Tian and {Jerry Ma*} and {Qucheng Gong*} and Shubho Sengupta and Zhuoyuan Chen and C. Lawrence Zitnick},
  title = {ELF OpenGo},
  year = {2018},
  journal = {GitHub repository},
  howpublished = {\url{https://github.com/pytorch/ELF}}
}
