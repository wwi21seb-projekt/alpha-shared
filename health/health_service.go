package health

import (
	"context"
	"github.com/wwi21seb-projekt/alpha-shared/db"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"sync"
	"time"

	healthv1 "github.com/wwi21seb-projekt/alpha-shared/gen/server_alpha/health/v1"
)

type HealthService struct {
	healthv1.UnimplementedHealthServiceServer
	mu     sync.Mutex
	status healthv1.ServingStatus
	subs   map[healthv1.HealthService_WatchServer]struct{}
	logger *zap.SugaredLogger
}

func NewHealthServer(logger *zap.SugaredLogger) *HealthService {
	return &HealthService{
		status: healthv1.ServingStatus_SERVING_STATUS_SERVING,
		subs:   make(map[healthv1.HealthService_WatchServer]struct{}),
		logger: logger,
	}
}

func (hs *HealthService) Check(ctx context.Context, req *healthv1.CheckRequest) (*healthv1.CheckResponse, error) {
	hs.mu.Lock()
	defer hs.mu.Unlock()

	response := &healthv1.CheckResponse{
		Status: hs.status,
	}
	return response, nil
}

func (hs *HealthService) Watch(req *healthv1.WatchRequest, stream healthv1.HealthService_WatchServer) error {
	hs.mu.Lock()
	hs.subs[stream] = struct{}{}
	hs.mu.Unlock()

	<-stream.Context().Done()

	hs.mu.Lock()
	delete(hs.subs, stream)
	hs.mu.Unlock()

	return stream.Context().Err()
}

func (hs *HealthService) updateStatus(status healthv1.ServingStatus) {
	hs.mu.Lock()
	defer hs.mu.Unlock()

	hs.status = status
	response := &healthv1.WatchResponse{
		Status: status,
	}

	for stream := range hs.subs {
		if err := stream.Send(response); err != nil {
			hs.logger.Warnw("Failed to send health status update", "error", err)
		}
	}
}

// CheckDependencies checks the status of various dependencies and updates the health status accordingly
func (hs *HealthService) CheckDependencies(ctx context.Context, database *db.DB, grpcClients map[string]*grpc.ClientConn) {
	go func() {
		for {
			// Check database connection
			if err := database.Ping(ctx); err != nil {
				hs.logger.Warn("Database connection failed")
				hs.updateStatus(healthv1.ServingStatus_SERVING_STATUS_NOT_SERVING)
				continue
			}

			// Check gRPC client connections
			allClientsHealthy := true
			for name, conn := range grpcClients {
				if conn.GetState() != connectivity.Ready {
					hs.logger.Warnw("gRPC client connection not ready", "client", name)
					allClientsHealthy = false
					break
				}
			}

			if allClientsHealthy {
				hs.updateStatus(healthv1.ServingStatus_SERVING_STATUS_SERVING)
			} else {
				hs.updateStatus(healthv1.ServingStatus_SERVING_STATUS_NOT_SERVING)
			}

			time.Sleep(30 * time.Second) // Check every 30 seconds
		}
	}()
}
