package main
import (
    "golang.org/x/net/context"
    "fmt"
    "log"
    pb "github.com/KoganezawaRyouta/grpc_test/quick/pb"
    "google.golang.org/grpc"
    "time"
)
func main() {
    conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
    if err != nil {
        log.Fatal("client connection error:", err)
    }
    defer conn.Close()
    client := pb.NewHelloClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    message := &pb.GetMyHelloMessage{TargetCat: "mike"}
    res, err := client.GetMyHello(ctx, message)
    if err != nil {
        fmt.Printf("error::%#v \n", err)
    }
    fmt.Printf("result:%#v \n", res)

}
