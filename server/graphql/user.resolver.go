package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	myContext "github.com/elisalimli/nextsync/server/context"
	"github.com/elisalimli/nextsync/server/domain"
	"github.com/elisalimli/nextsync/server/graphql/models"
)

var (
	ErrInput = errors.New("input errors")
)

func (m *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	isValid, errors := domain.Validation(ctx, input)
	fmt.Println(isValid, errors)
	if !isValid {
		return &models.AuthResponse{Ok: false, Errors: errors}, nil
	}
	return m.Domain.Login(ctx, input)
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	fmt.Println(ctx.Value("currentUser"))
	fmt.Println("recieved from ios")
	return "hello world works!", nil
}

func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	currentUserId, _ := ctx.Value(myContext.CurrentUserIdKey).(string)
	user, err := r.Domain.UsersRepo.GetUserByID(ctx, currentUserId)
	if err != nil {
		return nil, nil
	}
	return user, nil
}

func (m *mutationResolver) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	isValid, errors := domain.Validation(ctx, input)
	if !isValid {
		return &models.AuthResponse{Ok: false, Errors: errors}, nil
	}

	return m.Domain.Register(ctx, input)
}

func (m *mutationResolver) GoogleLogin(ctx context.Context, input models.GoogleLoginInput) (*models.AuthResponse, error) {
	isValid, errors := domain.Validation(ctx, input)
	if !isValid {
		return &models.AuthResponse{Ok: false, Errors: errors}, nil
	}
	return m.Domain.GoogleLogin(ctx, input)
}

func (m *mutationResolver) GoogleSignUp(ctx context.Context, input models.GoogleSignUpInput) (*models.AuthResponse, error) {
	isValid, validationErrors := domain.Validation(ctx, input)
	if !isValid {
		return &models.AuthResponse{Ok: false, Errors: validationErrors}, nil
	}
	return m.Domain.GoogleSignUp(ctx, input)
}

func (m *mutationResolver) RefreshToken(ctx context.Context) (*models.AuthResponse, error) {
	return m.Domain.RefreshToken(ctx)
}

func (m *mutationResolver) Logout(ctx context.Context) (bool, error) {
	rtCookie := http.Cookie{
		Name:   os.Getenv("COOKIE_REFRESH_TOKEN"),
		Path:   "/", // <--- add this line
		Value:  "",
		Secure: false,
	}
	writer, _ := ctx.Value(myContext.HttpWriterKey).(http.ResponseWriter)
	// saving cookie
	http.SetCookie(writer, &rtCookie)
	return true, nil
}

func (m *mutationResolver) SendOtp(ctx context.Context, input models.SendOtpInput) (*models.FormResponse, error) {
	isValid, errors := domain.Validation(ctx, input)
	if !isValid {
		return &models.FormResponse{Ok: false, Errors: errors}, nil
	}

	return m.Domain.SendOtp(ctx, input)
}

func (m *mutationResolver) VerifyOtp(ctx context.Context, input models.VerifyOtpInput) (*models.AuthResponse, error) {
	isValid, errors := domain.Validation(ctx, input)
	if !isValid {
		return &models.AuthResponse{Ok: false, Errors: errors}, nil
	}

	return m.Domain.VerifyOtp(ctx, input)
}
