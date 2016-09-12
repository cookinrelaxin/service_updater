package main

import (
    "log"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "./service_updater_protocol"
)

func main() {
    conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewServiceUpdaterClient(conn)

    r, err := c.UpdateService(context.Background(), &pb.UpdateRequest{ServiceName: "polymer-demo", ImageName: "cookinrelaxin/polymer-demo:18"})
    if err != nil {
        log.Fatalf("could not update service: %v", err)
    }
    // log.Printf("Response: %s", r.Message)
    log.Printf("%#v", r)
}

