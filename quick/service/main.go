package main

import (
    "log"
    "net"
    "golang.org/x/net/context"
    pb "github.com/KoganezawaRyouta/grpc_test/quick/pb"
    "google.golang.org/grpc"
    "errors"
    "google.golang.org/grpc/reflection"
)
type MyHelloService struct {}
func (s *MyHelloService) GetMyHello(ctx context.Context, message *pb.GetMyHelloMessage) (*pb.MyHelloResponse, error) {
    switch message.TargetCat {
    case "tama":
        return &pb.MyHelloResponse{
            Name: "tama",
            Kind: "mainecoon",
        }, nil
    case "mike":
        return &pb.MyHelloResponse{
            Name: "mike",
            Kind: "Norwegian Forest Cat",
        }, nil
    }
    return nil, errors.New("not found your cat")
}

func main() {
    listenPort, err := net.Listen("tcp", ":19003")
    if err != nil {
        log.Fatalln(err)
    }

    s := grpc.NewServer()
    pb.RegisterHelloServer(s, &MyHelloService{})
    // Register reflection service on gRPC server.
    reflection.Register(s)

    if err := s.Serve(listenPort); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
