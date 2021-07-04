package main

import (
	"context"
	"errors"

	pb "github.com/tung238/shippy/shippy-service-user/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type handler struct {
	repository   Repository
	tokenService authable
}

func (s *handler) Create(ctx context.Context, in *pb.User, out *pb.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	in.Password = string(hashedPass)
	if err := s.repository.Create(ctx, MarshalUser(in)); err != nil {
		return err
	}
	in.Password = ""
	out.User = in
	return nil
}

func (c *handler) Get(ctx context.Context, in *pb.User, out *pb.Response) error {
	result, err := c.repository.Get(ctx, in.Id)
	if err != nil {
		return err
	}
	out.User = UnmarshalUser(result)
	return nil
}

func (c *handler) GetAll(ctx context.Context, in *pb.Request, out *pb.Response) error {
	result, err := c.repository.GetAll(ctx)
	if err != nil {
		return nil
	}
	out.Users = UnmarshalUserCollection(result)
	return nil
}

func (c *handler) Auth(ctx context.Context, in *pb.User, out *pb.Token) error {
	user, err := c.repository.GetByEmail(ctx, in.Email)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return err
	}
	token, err := c.tokenService.Encode(in)
	if err != nil {
		return err
	}
	out.Token = token
	return nil
}

func (c *handler) ValidateToken(ctx context.Context, in *pb.Token, out *pb.Token) error {
	claims, err := c.tokenService.Decode(in.Token)
	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	out.Valid = true
	return nil
}
