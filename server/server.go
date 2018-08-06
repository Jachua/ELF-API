package main

import (
	"fmt"
	"log"
	"net"

	pb "ELF-API"

	uuid "github.com/nu7hatch/gouuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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

type room struct {
	human      *player
	AI         *player
	nextPlayer uint32
	hasChosen  bool
	resumed    []string
	ID         string
	isAssigned bool
	endGame    bool
	userID     string
}

type server struct {
	rooms map[string]*room
}

func (r *room) IsHuman(in *pb.Player) bool {
	if in.Color == r.human.color {
		return true
	}
	return false
}

func checkNil(in *pb.Step) *pb.Step {
	switch in {
	case nil:
		log.Println("The player has not made a move yet.")
		return &pb.Step{X: 20, Y: 20, Player: &pb.Player{Color: 3}}
	default:
		return in
	}
}

func (s *server) NewRoom(ctx context.Context, in *pb.State) (*pb.State, error) {
	u, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error: ", err)
		return &pb.State{Status: false, ID: ""}, err
	}
	id := u.String()
	s.rooms[id] = &room{
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
		ID:         id,
		isAssigned: false,
		nextPlayer: 1,
		hasChosen:  false,
		endGame:    false,
		// resumed:    []string{"BKD", "WFB", "BGA", "BAA", "WBB", "BCC", "WDD", "BEE", "WFF", "BGG", "WHH"},
	}
	return &pb.State{Status: true, ID: id}, nil
}

func (s *server) GetID(ctx context.Context, in *pb.State) (*pb.State, error) {
	for id, room := range s.rooms {
		if !room.isAssigned {
			room.isAssigned = true
			log.Println("Assigned to player with ID ", id)
			return &pb.State{Status: true, ID: id}, nil
		}
	}
	log.Println("Failed to assign a room to player.")
	// return &pb.State{Status: false, ID: ""}, nil
	err := status.Error(codes.Unavailable, "failed to assign a room to player.")
	return nil, err
}

func (s *server) GetMove(ctx context.Context, in *pb.Player) (*pb.Step, error) {
	r := s.rooms[in.ID]
	if r.IsHuman(in) {
		return checkNil(r.human.prevMove), nil
	}
	return checkNil(r.AI.prevMove), nil
}

func (s *server) SetMove(ctx context.Context, in *pb.Step) (*pb.State, error) {
	log.Println("Registering move from next player with ID ", in.Player.ID, " color ", in.Player.Color, "...")
	r := s.rooms[in.Player.ID]
	if r.IsHuman(in.Player) {
		r.human.prevMove = in
		r.human.hasMoved = true
	} else {
		r.AI.prevMove = in
		r.AI.hasMoved = true
	}
	return &pb.State{Status: true}, nil
}

func (s *server) HasMoved(ctx context.Context, in *pb.Player) (*pb.State, error) {
	return nil, nil
}

func (s *server) UpdateNext(ctx context.Context, in *pb.State) (*pb.State, error) {
	r := s.rooms[in.ID]
	r.nextPlayer = r.nextPlayer%2 + 1
	r.human.hasMoved = false
	r.AI.hasMoved = false
	return &pb.State{Status: true}, nil
}

func (s *server) IsNextPlayer(ctx context.Context, in *pb.Player) (*pb.State, error) {
	r := s.rooms[in.ID]
	return &pb.State{Status: r.nextPlayer == in.Color}, nil
}

func (s *server) SetPlayer(ctx context.Context, in *pb.Player) (*pb.State, error) {
	r := s.rooms[in.ID]
	r.human.color = in.Color
	r.AI.color = in.Color%2 + 1
	r.hasChosen = true
	// log.Println("Set player has chosen: ", r.hasChosen)
	return &pb.State{Status: true}, nil
}

func (s *server) GetAIPlayer(ctx context.Context, in *pb.State) (*pb.Player, error) {
	r := s.rooms[in.ID]
	return &pb.Player{Color: r.AI.color}, nil
}

func (s *server) HasChosen(ctx context.Context, in *pb.State) (*pb.State, error) {
	r := s.rooms[in.ID]
	// log.Println("Console has chosen, ", r.hasChosen)
	return &pb.State{Status: r.hasChosen}, nil
}

func (s *server) SetResumed(ctx context.Context, in *pb.Resumed) (*pb.State, error) {
	r := s.rooms[in.ID]
	r.resumed = in.Move
	return &pb.State{Status: true}, nil
}

func (s *server) GetResumed(ctx context.Context, in *pb.State) (*pb.Resumed, error) {
	r := s.rooms[in.ID]
	switch r.resumed {
	case nil:
		return &pb.Resumed{Move: make([]string, 0)}, nil
	default:
		return &pb.Resumed{Move: r.resumed}, nil
	}
}

func (s *server) SetExit(ctx context.Context, in *pb.State) (*pb.State, error) {
	r := s.rooms[in.ID]
	r.endGame = true
	return &pb.State{Status: true}, nil
}

func (s *server) CheckExit(ctx context.Context, in *pb.State) (*pb.State, error) {
	r := s.rooms[in.ID]
	if goExit := r.endGame; goExit {
		delete(s.rooms, in.ID)
		return &pb.State{Status: true}, nil
	}
	return &pb.State{Status: false}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gameServer := &server{
		rooms: make(map[string]*room),
	}
	pb.RegisterTurnServer(s, gameServer)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
