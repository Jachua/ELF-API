# TODO: restart the game
#       Client - 
#            Wait for a new ID
#       AI - 
#            NewRoom nil pointer when reset with minimum steps
#            check if is_resumed, res_arr paired with player ID, get resumed after pairing

#       server - if is_resumed assign ID from player to AI

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

    # ID_state = ""
    # while True:
    #     print("Waiting for AI to connect to server...")
    #     ID_state = stub.GetID(play_pb2.State(status = True))
    #     if ID_state.status:
    #         break
    # ID = ID_state.ID

    info = {'ID': "", 'color': 1}

    def reset():
        while True:
            try:
                info['ID'] = stub.GetID(play_pb2.State(status = True)).ID
                break
            except:
                _ = input("AI is unavailable at the moment. Press enter to try again.\n>")
                continue
        print("You are assigned with ID ", info['ID'])

        color = input("Enter 'B' for Black, 'W' for White \n>").upper()
        if color == "B":
            info['color'] = 1
        else:
            info['color'] = 2
        _ = stub.SetPlayer(play_pb2.Player(color = info['color'], ID = info['ID']))
        return

    def exit(c):
        if type(c) == str and c.strip().lower() == "exit":
            print("Ending game...")
            _ = stub.SetExit(play_pb2.State(status = True, ID = info['ID']))
            reset()
            return True
        return False

    """x = -1 && y = -1 indicates passing this round"""
    # def get_move():
    #     x = input("Enter coordinate x \n>")
    #     exit(x)
    #     x = int(x)
    #     y = input("Enter coordindate y \n>")
    #     exit(y)
    #     y = int(y)
    #     return x, y

    # def get_move():
    #     x = input("Enter coordinate x \n>")
    #     y = input("Enter coordinate y \n>")
    #     return x, y

    reset()

    while True: 
        if stub.IsNextPlayer(play_pb2.Player(color = info['color'], ID = info['ID'])).status:
            x, y = "", ""
            try: 
                x = input("Enter coordinate x \n>")
                x = int(x)
                y = input("Enter coordinate y \n>")
                y = int(y)
                _ = stub.SetMove(play_pb2.Step(x = x, y = y, player = play_pb2.Player(color =  info['color'], ID = info['ID'])))
                _ = stub.UpdateNext(play_pb2.State(status = True, ID = info['ID']))
                human_move = stub.GetMove(play_pb2.Player(color = info['color'], ID = info['ID']))
                print("You placed a stone at coordinates: ", human_move.x, ",", human_move.y)
            except:
                if exit(x) or exit(y):
                    continue
                else:
                    print("Invalid coordinates")
        else: 
            print("It's AI's turn to play.")
            while stub.IsNextPlayer(play_pb2.Player(color = info['color'] % 2 + 1, ID = info['ID'])).status:
                # print("Waiting for AI to make a move...")
                pass
            ("Receiving next move from AI...")
            AI_move = stub.GetMove(play_pb2.Player(color = info['color'] % 2 + 1, ID = info['ID']))
            if AI_move.x == -1 and AI_move.y == -1:
                print("AI chose to pass this round.")
            else:
                print("AI placed a stone at coordinates: ", AI_move.x, ",", AI_move.y)


if __name__ == '__main__':
    run()
