package main

import (
	"context"
	"log"
	"time"

	pb "ELF-API"

	"google.golang.org/grpc"
)

func setClientColor(ctx context.Context, c pb.TurnClient, color uint32) {
	_, _ = c.SetPlayer(ctx, &pb.Player{Color: color})
	return
}

func setClientMove(ctx context.Context, c pb.TurnClient, x uint32, y uint32, color uint32) {
	_, _ = c.SetMove(ctx, &pb.Step{X: x, Y: y, Player: &pb.Player{Color: color}})
	_, _ = c.UpdateNext(ctx, &pb.State{Status: true})
	return
}

func getAIMove(ctx context.Context, c pb.TurnClient, color uint32) (uint32, uint32) {
	move, _ := c.GetMove(ctx, &pb.Player{Color: color})
	return move.X, move.Y
}

func isNext(ctx context.Context, c pb.TurnClient, color uint32) bool {
	isNext, err := c.IsNextPlayer(ctx, &pb.Player{Color: color})
	if err != nil {
		log.Println("Error at isNext: ", err)
		return false
	}
	return isNext.Status
}

func getClientMove() (uint32, uint32) {
	return uint32(3), uint32(3)
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: &v", err)
	}
	defer conn.Close()
	c := pb.NewTurnClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	clientColor := uint32(1)
	AIColor := clientColor%2 + 1
	setClientColor(ctx, c, clientColor)
	for {
		if isNext(ctx, c, clientColor) {
			// Indicates that it's the client's turn
			x, y := getClientMove()
			setClientMove(ctx, c, x, y, clientColor)
		} else {
			for {
				if !isNext(ctx, c, AIColor) {
					// Indicates that AI has made a move
					break
				}
			}
			x, y := getAIMove(ctx, c, AIColor)
			log.Println("AI places a stone at coordinates, ", x, ", ", y)
		}
	}
}
