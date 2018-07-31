// package main

// import (
// 	"log"
// 	"net"

// 	pb "ELF-API"

// 	"golang.org/x/net/context"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/reflection"
// )

// const (
// 	port = ":50051"
// )

// //color = 1 Black
// //color = 2 White
// type player struct {
// 	prevMove *pb.Step
// 	color    uint32
// 	hasMoved bool
// 	ID       string
// }

// type room struct {
// 	human      *player
// 	AI         *player
// 	nextPlayer uint32
// 	hasChosen  bool
// 	ID         string
// 	isAssigned bool
// }

// type server struct {
// 	rooms map[string]*room
// }

// func (r *room) InRoom(p *player) bool {
// 	return r.ID == p.ID
// }

// func (s *server) GetRoom(id string) *room {
// 	return s.rooms[id]
// }

// func (r *room) IsHuman(in *pb.Player) bool {
// 	if in.Color == r.human.color {
// 		return true
// 	}
// 	return false
// }

// func checkNil(in *pb.Step) *pb.Step {
// 	switch in {
// 	case nil:
// 		return &pb.Step{X: 20, Y: 20, Player: &pb.Player{Color: 3, ID: ""}}
// 	default:
// 		return in
// 	}
// }

// func (s *server) GetMove(ctx context.Context, in *pb.Player) (*pb.Step, error) {
// 	r := s.GetRoom(in.ID)
// 	if r.IsHuman(in) {
// 		return checkNil(r.human.prevMove), nil
// 	}
// 	return checkNil(r.AI.prevMove), nil
// }

// func (s *server) SetMove(ctx context.Context, in *pb.Step) (*pb.State, error) {
// 	log.Println("Registering move from next player...")
// 	r := s.GetRoom(in.Player.ID)
// 	if r.IsHuman(in.Player) {
// 		r.human.prevMove = in
// 		r.human.hasMoved = true
// 	} else {
// 		r.AI.prevMove = in
// 		r.AI.hasMoved = true
// 	}
// 	return &pb.State{Status: true, ID: r.ID}, nil
// }

// func (s *server) HasMoved(ctx context.Context, in *pb.Player) (*pb.State, error) {
// 	r := s.GetRoom(in.ID)
// 	if r.IsHuman(in) {
// 		return &pb.State{Status: r.human.hasMoved, ID: r.ID}, nil
// 	}
// 	return &pb.State{Status: r.AI.hasMoved, ID: r.ID}, nil
// }

// func (s *server) UpdateNext(ctx context.Context, in *pb.State) (*pb.State, error) {
// 	r := s.GetRoom(in.ID)
// 	r.nextPlayer = r.nextPlayer%2 + 1
// 	r.human.hasMoved = false
// 	r.AI.hasMoved = false
// 	return &pb.State{Status: true, ID: in.ID}, nil
// }

// func (s *server) IsNextPlayer(ctx context.Context, in *pb.Player) (*pb.State, error) {
// 	r := s.GetRoom(in.ID)
// 	return &pb.State{Status: r.nextPlayer == in.Color, ID: r.ID}, nil
// }

// func (s *server) SetPlayer(ctx context.Context, in *pb.Player) (*pb.State, error) {
// 	r := s.GetRoom(in.ID)
// 	log.Println("The ID of the room is: ", r.ID)
// 	r.human.color = in.Color
// 	r.AI.color = in.Color%2 + 1
// 	r.hasChosen = true
// 	return &pb.State{Status: true, ID: r.ID}, nil
// }

// func (s *server) GetAIPlayer(ctx context.Context, in *pb.State) (*pb.Player, error) {
// 	r := s.GetRoom(in.ID)
// 	return &pb.Player{Color: r.AI.color, ID: r.ID}, nil
// }

// func (s *server) HasChosen(ctx context.Context, in *pb.State) (*pb.State, error) {
// 	r := s.GetRoom(in.ID)
// 	return &pb.State{Status: r.hasChosen, ID: r.ID}, nil
// }

// func (s *server) NewRoom(ctx context.Context, in *pb.State) (*pb.State, error) {
// 	// u, err := uuid.NewV4()
// 	// if err != nil {
// 	// 	fmt.Println("Error: ", err)
// 	// 	return &pb.State{Status: false, ID: ""}, err
// 	// }
// 	// id := u.String()
// 	id := ""
// 	s.rooms[id] = &room{
// 		human: &player{
// 			prevMove: nil,
// 			hasMoved: false,
// 			color:    1,
// 			ID:       id,
// 		},
// 		AI: &player{
// 			prevMove: nil,
// 			hasMoved: false,
// 			color:    2,
// 			ID:       id,
// 		},
// 		ID:         id,
// 		isAssigned: false,
// 		nextPlayer: 1,
// 		hasChosen:  false,
// 	}
// 	return &pb.State{Status: true, ID: id}, nil
// }

// func (s *server) GetID(ctx context.Context, in *pb.State) (*pb.State, error) {
// 	for id, room := range s.rooms {
// 		if !room.isAssigned {
// 			room.isAssigned = true
// 			return &pb.State{Status: true, ID: id}, nil
// 		}
// 	}
// 	log.Println("Failed to assign a room to player.")
// 	return &pb.State{Status: false, ID: ""}, nil
// }

// func main() {
// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	// playserver := &server{
// 	// 	human: &player{
// 	// 		prevMove: nil,
// 	// 		hasMoved: false,
// 	// 		color:    1,
// 	// 	},
// 	// 	AI: &player{
// 	// 		prevMove: nil,
// 	// 		hasMoved: false,
// 	// 		color:    2,
// 	// 	},
// 	// 	nextPlayer: 1,
// 	// 	hasChosen:  false,
// 	// }
// 	gameServer := &server{
// 		rooms: make(map[string]*room),
// 	}
// 	pb.RegisterTurnServer(s, gameServer)
// 	// Register reflection service on gRPC server.
// 	reflection.Register(s)
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
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
	resumed    []string
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

func (s *server) NewRoom(ctx context.Context, in *pb.State) (*pb.State, error) {
	return in, nil
}

func (s *server) GetID(ctx context.Context, in *pb.State) (*pb.State, error) {
	return in, nil
}

func (s *server) SetResumed(ctx context.Context, in *pb.Resumed) (*pb.State, error) {
	s.resumed = in.Move
	return &pb.State{Status: true}, nil
}

func (s *server) GetResumed(ctx context.Context, in *pb.State) (*pb.Resumed, error) {
	switch s.resumed {
	case nil:
		return &pb.Resumed{Move: make([]string, 0)}, nil
	default:
		return &pb.Resumed{Move: s.resumed}, nil
	}
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
		resumed:    nil,
		// resumed: []string{"BKD", "WFB", "BGA"},
	}
	pb.RegisterTurnServer(s, playServer)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
