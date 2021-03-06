package main

import (
    "golang.org/x/net/context"
    "log"
    "net"
    "sync"

    "google.golang.org/grpc"

    pb "github.com/KoganezawaRyouta/grpc_test/stream/pb"
    "fmt"
)

type customerService struct {
    customers []*pb.Person
    m         sync.Mutex
}


func remove(persons []*pb.Person, person *pb.Person) []*pb.Person {
    result := []*pb.Person{}
    for _, v := range persons {
        if v.Name != person.Name {
            result = append(result, v)
        }
    }
    return result
}

func (cs *customerService) ListPerson(p *pb.RequestType, stream pb.CustomerService_ListPersonServer) error {
    cs.m.Lock()
    defer cs.m.Unlock()
    for _, p := range cs.customers {
        if err := stream.Send(p); err != nil {
            return err
        }
        cs.customers = remove(cs.customers, p)
    }
    return nil
}

func (cs *customerService) AddPerson(c context.Context, p *pb.Person) (*pb.ResponseType, error) {
    cs.m.Lock()
    defer cs.m.Unlock()
    cs.customers = append(cs.customers, p)
    fmt.Printf("result:%#v \n", p)
    return new(pb.ResponseType), nil
}

func main() {
    lis, err := net.Listen("tcp", ":11111")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    server := grpc.NewServer()

    pb.RegisterCustomerServiceServer(server, new(customerService))
    server.Serve(lis)
}