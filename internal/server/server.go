package server

import (
	"context"
	"fmt"
	"net"
	"os/signal"
	"syscall"

	cardapplication "github.com/akubi0w1/sge-tech-book-vol3-sample/internal/application/card"
	userapplication "github.com/akubi0w1/sge-tech-book-vol3-sample/internal/application/user"
	cardhandler "github.com/akubi0w1/sge-tech-book-vol3-sample/internal/handler/card"
	masterhandler "github.com/akubi0w1/sge-tech-book-vol3-sample/internal/handler/master"
	userhandler "github.com/akubi0w1/sge-tech-book-vol3-sample/internal/handler/user"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/infra/mysql"
	cardrepository "github.com/akubi0w1/sge-tech-book-vol3-sample/internal/infra/mysql/repository/card"
	characterrepository "github.com/akubi0w1/sge-tech-book-vol3-sample/internal/infra/mysql/repository/character"
	userrepository "github.com/akubi0w1/sge-tech-book-vol3-sample/internal/infra/mysql/repository/user"
	usercardrepository "github.com/akubi0w1/sge-tech-book-vol3-sample/internal/infra/mysql/repository/usercard"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/log"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/terror"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/util/closer"
	pb "github.com/akubi0w1/sge-tech-book-vol3-sample/pkg/pb/service"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcvalidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
}

type Config struct {
	Port    string `envconfig:"PORT" default:"50051"`
	Profile string `envconfig:"PROFILE" default:"local"`

	// master db
	MysqlMasterAddr     string `envconfig:"MYSQL_MASTER_ADDR" default:"localhost"`
	MysqlMasterProtocol string `envconfig:"MYSQL_MASTER_PROTOCOL" default:"tcp"`
	MysqlMasterUser     string `envconfig:"MYSQL_MASTER_USER" default:"root"`
	MysqlMasterPassword string `envconfig:"MYSQL_MASTER_PASSWORD" default:"root"`
	MysqlMasterDB       string `envconfig:"MYSQL_MASTER_DB" default:"master"`

	// system db
	MysqlShardAddr     string `envconfig:"MYSQL_SHARD_ADDR" default:"localhost"`
	MysqlShardProtocol string `envconfig:"MYSQL_SHARD_PROTOCOL" default:"tcp"`
	MysqlShardUser     string `envconfig:"MYSQL_SHARD_USER" default:"root"`
	MysqlShardPassword string `envconfig:"MYSQL_SHARD_PASSWORD" default:"root"`
	MysqlShardDB       string `envconfig:"MYSQL_SHARD_DB" default:"shard"`
}

// ListenAndServe
func ListenAndServe(conf *Config) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv, closer, err := newServer(ctx, conf)
	defer closer.Close()
	if err != nil {
		return err
	}

	g, gCtx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Port))
		if err != nil {
			return terror.Wrapf(terror.CodeInternal, err, "failed to listen")
		}

		log.Infof("listen and serve on %s...", conf.Port)
		if err = srv.server.Serve(lis); err != nil {
			return terror.Wrapf(terror.CodeInternal, err, "failed to serve")
		}

		return nil
	})

	g.Go(func() error {
		select {
		case <-gCtx.Done():
		case <-ctx.Done():
			srv.shutdown()
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return terror.Wrapf(terror.CodeInternal, err, "failed to execute")
	}

	return nil
}

// new
func newServer(ctx context.Context, c *Config) (*Server, *closer.Closer, error) {
	closer := &closer.Closer{}

	// connect db
	masterDB, err := mysql.NewMasterDB(&mysql.Config{
		Addr:     c.MysqlMasterAddr,
		Protocol: c.MysqlMasterProtocol,
		User:     c.MysqlMasterUser,
		Password: c.MysqlMasterPassword,
		DB:       c.MysqlMasterDB,
	})
	if err != nil {
		return nil, closer, err
	}
	closer.Add(func() {
		masterDB.Close()
	})

	shardDB, err := mysql.NewShardDB(&mysql.Config{
		Addr:     c.MysqlShardAddr,
		Protocol: c.MysqlShardProtocol,
		User:     c.MysqlShardUser,
		Password: c.MysqlShardPassword,
		DB:       c.MysqlShardDB,
	})
	if err != nil {
		return nil, closer, err
	}
	closer.Add(func() {
		shardDB.Close()
	})

	// new server with interceptor
	server := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		// recoverer.Interceptor(),
		grpcvalidator.UnaryServerInterceptor(),
	))

	// injector
	cardRepository := cardrepository.New(masterDB)
	characterRepository := characterrepository.New(masterDB)
	userRepository := userrepository.New(shardDB)
	userCardRepository := usercardrepository.New(shardDB)

	cardApplication := cardapplication.New(cardRepository, userCardRepository)
	userApplication := userapplication.New(userRepository)

	masterHandler := masterhandler.New(cardRepository, characterRepository)
	userHandler := userhandler.New(userApplication)
	cardHandler := cardhandler.New(cardApplication)

	// register
	pb.RegisterCardServiceServer(server, cardHandler)
	pb.RegisterUserServiceServer(server, userHandler)
	pb.RegisterMasterServiceServer(server, masterHandler)

	return &Server{
		server: server,
	}, closer, nil
}

// shutdown
func (srv *Server) shutdown() {
	log.Infof("shutdown server")
	srv.server.GracefulStop()
}
