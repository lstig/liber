package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gofrs/uuid"
	// "github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/swgui/v4emb"

	v1book "github.com/lstig/liber/internal/api/v1/book"
	"github.com/lstig/liber/internal/types"
)

var Version string = "not built correctly"

type Config struct {
	// list of trusted proxies
	TrustedProxies []string

	// The address the server listens on
	ListenAddress string
}

// Returns a configured gin router with all routes registered
func NewServer(config Config) *web.Service {
	s := web.DefaultService()

	s.OpenAPI.Info.Title = "Liber"
	s.OpenAPI.Info.WithDescription("eBook server with OPDS support")
	s.OpenAPI.Info.Version = Version

	s.OpenAPICollector.Reflector().AddTypeMapping(uuid.UUID{}, types.Uuid())
	s.OpenAPICollector.Reflector().InlineDefinition(uuid.UUID{})

	s.Wrap(middleware.Logger)

	s.Route("/api/v1", func(r chi.Router) {
		v1book.RegisterRoutes(r)
	})

	s.Docs("/docs", v4emb.New)

	return s
}
