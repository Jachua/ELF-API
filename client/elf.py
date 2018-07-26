

import grpc

import play_pb2
import play_pb2_grpc

from server_addrs import addrs


def get_move():
    x = int(input("Enter coordiante x \n>"))
    y = int(input("Enter coorindate y \n>"))
    return x, y

def run():
    address = addrs["game_server"]
    if address != "":
        channel = grpc.insecure_channel(address + ":50051")
    else: 
        channel = grpc.insecure_channel('localhost:50051')
    stub = play_pb2_grpc.TurnStub(channel)
    if input("Enter 'B' for Black, 'W' for White \n>").upper() == "B":
        human_color = 1
    else:
        human_color = 2
    _ = stub.SetPlayer(play_pb2.Player(color = human_color))
    AI_color = human_color % 2 + 1
    while True: 
        if stub.IsNextPlayer(play_pb2.Player(color = human_color)).status:
            x, y = get_move()
           
            _ = stub.SetMove(play_pb2.Step(x = x, y = y, player = play_pb2.Player(color =  human_color)))
            _ = stub.UpdateNext(play_pb2.State(status = True))
            human_move = stub.GetMove(play_pb2.Player(color = human_color))
            print("You placed a stone at coordinates: ", human_move.x, ",", human_move.y)
        else: 
            while stub.IsNextPlayer(play_pb2.Player(color = AI_color)).status:
                pass
            AI_move = stub.GetMove(play_pb2.Player(color = AI_color))
            print("AI places a stone at coordinates: ", AI_move.x, ",", AI_move.y)


if __name__ == '__main__':
    run()
