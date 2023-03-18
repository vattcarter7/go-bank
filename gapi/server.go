package gapi

import (
	"fmt"

	db "github.com/vattcarter7/go-bank/db/sqlc"
	pb "github.com/vattcarter7/go-bank/pb"
	"github.com/vattcarter7/go-bank/token"
	"github.com/vattcarter7/go-bank/util"
)

// Server serves gRPC requests for our Banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
