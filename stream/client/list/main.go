package main

import (
"fmt"
"io"
pb "github.com/KoganezawaRyouta/grpc_test/stream/pb"
"golang.org/x/net/context"
"google.golang.org/grpc"
"flag"
)

var (
    tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
    caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
    serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
    serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func list() error {
    conn, err := grpc.Dial("127.0.0.1:11111", grpc.WithInsecure())
    if err != nil {
        return err
    }
    defer conn.Close()
    client := pb.NewCustomerServiceClient(conn)

    stream, err := client.ListPerson(context.Background(), new(pb.RequestType))
    if err != nil {
        return err
    }
    for {
        person, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        fmt.Println(person)
    }
    return nil
}

func main() {
    for {
        list()
    }
}