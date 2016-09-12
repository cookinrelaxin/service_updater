package main

import (
    "log"
    "net"
    "os/exec"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "github.com/cookinrelaxin/service_updater/protocol"
)

type server struct {}

func (s *server) UpdateService(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
    serviceName := in.ServiceName
    imageName := in.ImageName
    log.Printf("Attemping to update service %s with image %s", serviceName, imageName)

    updateCmd := exec.Command("docker", "service", "update", "--image", imageName, serviceName)
    out, err := updateCmd.CombinedOutput()
    if err != nil {
        log.Printf("Failed to update: %v", string(out))
        return &pb.UpdateResponse{Status: pb.UpdateResponse_FAILURE, Message: string(out)}, nil
    }
    log.Printf("Successfully updated service")

    return &pb.UpdateResponse{Status: pb.UpdateResponse_SUCCESS, Message: string(out)}, nil
}

func main() {
    lis, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterServiceUpdaterServer(grpcServer, &server{})
    grpcServer.Serve(lis)
}
