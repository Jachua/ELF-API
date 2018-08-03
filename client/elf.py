# def run():
#     address = addrs["game_server"]
#     if address != "":
#         channel = grpc.insecure_channel(address + ":50051")
#     else: 
#         channel = grpc.insecure_channel('localhost:50051')
#     stub = play_pb2_grpc.TurnStub(channel)
#     if input("Enter 'B' for Black, 'W' for White \n>").upper() == "B":
#         human_color = 1
#     else:
#         human_color = 2
#     # ID_state = stub.GetID(play_pb2.State(status = True, ID = ""))
#     # while not ID_state.status:
#     #     print("Waiting for AI to connect to server...")
#     #     ID_state = stub.GetID(play_pb2.State(status = True, ID = ""))
#     # ID = ID_state.ID
#     id = ""
#     _ = stub.SetPlayer(play_pb2.Player(color = human_color, ID = id))
#     AI_color = human_color % 2 + 1
#     while True: 
#         if stub.IsNextPlayer(play_pb2.Player(color = human_color, ID = id)).status:
#             x, y = get_move()
#             _ = stub.SetMove(play_pb2.Step(x = x, y = y, player = play_pb2.Player(color =  human_color, ID = id)))
#             _ = stub.UpdateNext(play_pb2.State(status = True, ID = id))
#             human_move = stub.GetMove(play_pb2.Player(color = human_color, ID = id))
#             print("You placed a stone at coordinates: ", human_move.x, ",", human_move.y)
#         else: 
#             while stub.IsNextPlayer(play_pb2.Player(color = AI_color, ID = id)).status:
#                 pass
#             AI_move = stub.GetMove(play_pb2.Player(color = AI_color, ID = id))
#             # print("OpenGo placed a stone at coordinates: ", AI_move.x, ",", AI_move.y)


# if __name__ == '__main__':
#     run()

import grpc

import play_pb2
import play_pb2_grpc

from server_addrs import addrs

import sys

def run():
    address = addrs["game_server"]
    if address != "":
        channel = grpc.insecure_channel(address + ":50051")
    else: 
        channel = grpc.insecure_channel('localhost:50051')
    stub = play_pb2_grpc.TurnStub(channel)

    def exit(c):
        if c.strip().lower() == "exit":
            print("Ending game...")
            # _ = stub.Exit(play_pb2.State(status = True))
            sys.exit(0)
        return

    """x = -1 && y = -1 indicates passing this round"""
    def get_move():
        x = input("Enter coordiante x \n>")
        exit(x)
        x = int(x)
        y = input("Enter coorindate y \n>")
        exit(y)
        y = int(y)
        return x, y

    color = input("Enter 'B' for Black, 'W' for White \n>").upper()
    exit(color)
    if color == "B":
        human_color = 1
    else:
        human_color = 2
    # ID_state = ""
    # while True:
    #     print("Waiting for AI to connect to server...")
    #     ID_state = stub.GetID(play_pb2.State(status = True))
    #     if ID_state.status:
    #         break
    # ID = ID_state.ID
    while True:
        try:
            ID = stub.GetID(play_pb2.State(status = True)).ID
            break
        except:
            retry = input("AI is unavailable at the moment. Press enter to try again," 
            "or input 'exit' to end game.\n")
            exit(retry)
            continue
    print("You are assigned with ID ", ID)
    _ = stub.SetPlayer(play_pb2.Player(color = human_color, ID = ID))
    AI_color = human_color % 2 + 1
    while True: 
        if stub.IsNextPlayer(play_pb2.Player(color = human_color, ID = ID)).status:
            try:
                x, y = get_move()
                _ = stub.SetMove(play_pb2.Step(x = x, y = y, player = play_pb2.Player(color =  human_color, ID = ID)))
                _ = stub.UpdateNext(play_pb2.State(status = True, ID = ID))
                human_move = stub.GetMove(play_pb2.Player(color = human_color, ID = ID))
                print("You placed a stone at coordinates: ", human_move.x, ",", human_move.y)
            except ValueError:
                print("Invalid coordinates")
        else: 
            print("It's AI's turn to play.")
            while stub.IsNextPlayer(play_pb2.Player(color = AI_color, ID = ID)).status:
                # print("Waiting for AI to make a move...")
                pass
            ("Receiving next move from AI...")
            AI_move = stub.GetMove(play_pb2.Player(color = AI_color, ID = ID))
            if AI_move.x == -1 and AI_move.y == -1:
                print("AI chose to pass this round.")
            else:
                print("AI placed a stone at coordinates: ", AI_move.x, ",", AI_move.y)


if __name__ == '__main__':
    run()
