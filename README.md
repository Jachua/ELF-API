# API
An interface that sets up a server to commicate with the [OpenGo AI](https://github.com/pytorch/ELF) developed by the Facebook team.

## **Installation**

* Go 1.10.3 darwin/amd64
* Python 3.6.5

Download Go [here](https://golang.org/dl/)

After installing Go by the instructions, run ```export PATH=$PATH:$(go env GOPATH)/bin``` and ```export GOPATH=$(go env GOPATH)```.

After installing Go by the instructions, run 
```
$ export PATH=$PATH:$(go env GOPATH)/bin
$ export GOPATH=$(go env GOPATH)
```

Installing dependencies for go:
```
$ go get -u google.golang.org/grpc 
```
or follow the instructions on [this page](https://grpc.io/docs/quickstart/go.html).
```
$ go get github.com/nu7hatch/gouuid

$ go get -u github.com/go-sql-driver/mysql
```

Installing dependencies for python:
```
python -m pip install grpcio
python -m pip install grpcio-tools googleapis-common-protos
```
or follow the instructions on [this page](https://grpc.io/docs/quickstart/python.html).


## **Setup**

Run ```git clone https://github.com/Jachua/ELF-API.git $(go env GOPATH)/src/ELF-API``` to create a local copy of the API server.

Set up the environment for the AI engine by the instructions posted on [this repo](https://github.com/Jachua/ELF).

## **Deployment**

The API constitues 3 major components: the game server, the client console CLI, and the AI engine.
The game server must first be loaded for the client console and the AI engine to send out and receive messages from each other. 

**Server**

Navigate to the API project root and run ```go run server/server.go```.

**Client console**


Edit the ```game_server``` field in ```client/server_addrs.py``` to be the IP address where you want to run the game server. If unspecified, the game will run on local address. 

Navigate to the API project root and run ```python client/elf.py```. The client console is where the prompts for entering commands and the moves received from OpenGo will be displayed. 

Make sure that the server is up and running before starting the client session. 

**AI engine**

Download the AI engine from [here](https://github.com/Jachua/ELF). 

Edit the ```game_server``` field in ```scripts/elfgames/go/server_addrs.py``` to be the same IP address that you have provided for the client console. 

At the ELF project root, run ```source scripts/devmode_set_pythonpath.sh```. Navigate to ```scripts/elfgames/go``` and run ```./gtp.sh``` followed by 1 shell argument to set the difficulty level ```EASY```, ```MEDIUM```, or ```HARD```. The default difficulty level is ```MEDIUM```.

You should now be able to play the game with OpenGo via the client console. 

## **How it works**

Once the server and the clients are properly set up, you would be prompted to reply whether you would like to resume from a previous session, and if so, enter the ID for that session. The client console will ask you to choose the side you would like to play on. If you choose black, it will then prompt you to enther the coordinate where you want to make the first move. If you choose white, it will display where OpenGo has placed a stone, and you will see messages such as "AI placed on stone on coordinates 15, 15". The game thus progresses by logging the moves received from the client console and OpenGo in turn. 

You may also observe how the game proceeds through the original CLI for the AI engine, which has not been disabled. For instance, if you are playing black and you have entered 18, 0 for your next move, you will see a black stone placed at position S1 on the AI console. And if the AI decides to place a stone at coordiates 3, 15, youb will notice something in the form of "Proposed move: D16" on the AI console.

If you want to restart the game, input "exit" in the console, which will also clear the current progress for the AI and assign it with a different ID. Once the AI becomes available, you will be able to start a new game.

*Note:

The x coordinate [0-18] in the client console translates to a character from [A-S] in the AI console, and the y coordinate [0-18] translates to an integer from range [1-19].

## **Citations**
@misc{ELFOpenGo2018,
  author = {Yuandong Tian and {Jerry Ma*} and {Qucheng Gong*} and Shubho Sengupta and Zhuoyuan Chen and C. Lawrence Zitnick},
  title = {ELF OpenGo},
  year = {2018},
  journal = {GitHub repository},
  howpublished = {\url{https://github.com/pytorch/ELF}}
}
