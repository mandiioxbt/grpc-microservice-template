package main
import ("fmt"; "net"; "google.golang.org/grpc")
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil { panic(err) }
    s := grpc.NewServer()
    fmt.Println("gRPC on :50051")
    _ = s.Serve(lis)
}
