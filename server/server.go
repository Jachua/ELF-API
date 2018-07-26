package main

import (
	"log"
	"net"

	pb "ELF-API"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

//color = 1 Black
//color = 2 White
type player struct {
	prevMove *pb.Step
	color    uint32
	hasMoved bool
}

type server struct {
	human      *player
	AI         *player
	nextPlayer uint32
	hasChosen  bool
}

func (s *server) IsHuman(in *pb.Player) bool {
	if in.Color == s.human.color {
		return true
	}
	return false
}

func checkNil(in *pb.Step) *pb.Step {
	switch in {
	case nil:
		return &pb.Step{X: 20, Y: 20, Player: &pb.Player{Color: 3}}
	default:
		return in
	}
}

func (s *server) GetMove(ctx context.Context, in *pb.Player) (*pb.Step, error) {

	if s.IsHuman(in) {
		return checkNil(s.human.prevMove), nil
	}
	return checkNil(s.AI.prevMove), nil
}

func (s *server) SetMove(ctx context.Context, in *pb.Step) (*pb.State, error) {
	log.Println("Registering move from next player...")
	if s.IsHuman(in.Player) {
		s.human.prevMove = in
		s.human.hasMoved = true
	} else {
		s.AI.prevMove = in
		s.AI.hasMoved = true
	}
	return &pb.State{Status: true}, nil
}

func (s *server) HasMoved(ctx context.Context, in *pb.Player) (*pb.State, error) {
	if s.IsHuman(in) {
		return &pb.State{Status: s.human.hasMoved}, nil
	}
	return &pb.State{Status: s.AI.hasMoved}, nil
}

func (s *server) UpdateNext(ctx context.Context, in *pb.State) (*pb.State, error) {
	s.nextPlayer = s.nextPlayer%2 + 1
	s.human.hasMoved = false
	s.AI.hasMoved = false
	return &pb.State{Status: true}, nil
}

func (s *server) IsNextPlayer(ctx context.Context, in *pb.Player) (*pb.State, error) {
	return &pb.State{Status: s.nextPlayer == in.Color}, nil
}

func (s *server) SetPlayer(ctx context.Context, in *pb.Player) (*pb.State, error) {
	s.human.color = in.Color
	s.AI.color = in.Color%2 + 1
	s.hasChosen = true
	return &pb.State{Status: true}, nil
}

func (s *server) GetAIPlayer(ctx context.Context, in *pb.State) (*pb.Player, error) {
	return &pb.Player{Color: s.AI.color}, nil
}

func (s *server) HasChosen(ctx context.Context, in *pb.State) (*pb.State, error) {
	return &pb.State{Status: s.hasChosen}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	playServer := &server{
		human: &player{
			prevMove: nil,
			hasMoved: false,
			color:    1,
		},
		AI: &player{
			prevMove: nil,
			hasMoved: false,
			color:    2,
		},
		nextPlayer: 1,
		hasChosen:  false,
	}
	pb.RegisterTurnServer(s, playServer)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
