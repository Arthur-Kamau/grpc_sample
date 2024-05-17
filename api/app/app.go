package app

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/bugsnag/bugsnag-go/v2"

	"os"

	"github.com/go-redis/redis"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/lib/pq"
	"github.com/logrusorgru/aurora"
	"github.com/sirupsen/logrus"

	context "context"
	"log"
	"net"

	. "grpc_sample/app/models"
	 "grpc_sample/app/logger"

	"google.golang.org/grpc"
)

type Server struct {

	// router and DB instance
	DB     *sql.DB
	Redis  *redis.Client
	Logger *logrus.Logger
	UnimplementedAppApiServer
}

var sr Server

// Initialize initializes configs and variable
func (s *Server) Initialize(dbInstance *sql.DB) {

	s.Logger = logger.GetLog()
	s.DB = dbInstance

	s.Redis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST_PORT"), //"localhost:6379",
		Password: os.Getenv("REDIS_HOST_PASSWORD"),
		DB:       13,
	})

	pong, err := s.Redis.Ping().Result()
	s.Logger.Info("Redis ping result ", pong, err)

	// db
	sr.DB = dbInstance
	// bug snag
	key := os.Getenv("BUGSNAG_KEY")
	bugsnag.Configure(bugsnag.Configuration{
		// Your Bugsnag project API key, required unless set as environment
		// variable $BUGSNAG_API_KEY
		APIKey: key,
		// The development stage of your application build, like "alpha" or
		// "production"
		ReleaseStage: "production",
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{"main", "github.com/org/myapp"},
		// more configuration options
	})

	s.setUpGRPCApi()

}





// setUpGRPCApi sets the all required router
func (s *Server) setUpGRPCApi() {

	port := os.Getenv("SYSTEM_PORT")

	if len(port) == 0 {
		port = "9000"
	}

	fmt.Println(aurora.Bold(aurora.Green(fmt.Sprintf("Starting server on port %s", port))))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	RegisterAppApiServer(grpcServer, s) //&Server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	// if logic is large ue sample_controller
	// if has db connection or therid part api etc use functions in sample_controller_util
	return &HelloResponse{
		IsSuccess: true,
		Sentence: fmt.Sprintf("Hello %s from api", in.Name),
	}, nil
}


func (s *Server) GuessNumber(ctx context.Context, in *empty.Empty) (*HelloResponse, error) {

	rand.Seed(time.Now().UnixNano())
	min := 1111
	max := 9999
	res := rand.Intn(max-min+1) + min

	return &HelloResponse{
		IsSuccess: true,
		Sentence: fmt.Sprintf("Hello %d", res),
	}, nil
}
