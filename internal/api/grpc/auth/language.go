package auth

import (
	"context"
	"github.com/zitadel/zitadel/internal/domain"

	auth_pb "github.com/zitadel/zitadel/pkg/grpc/auth"
)

func (s *Server) GetSupportedLanguages(ctx context.Context, req *auth_pb.GetSupportedLanguagesRequest) (*auth_pb.GetSupportedLanguagesResponse, error) {
	langs, err := s.query.Languages(ctx)
	if err != nil {
		return nil, err
	}
	return &auth_pb.GetSupportedLanguagesResponse{Languages: domain.LanguagesToStrings(langs)}, nil
}
