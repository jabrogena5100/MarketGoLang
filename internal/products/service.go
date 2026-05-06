package products

import (
	"context"

	repo "github.com/jabrogena5100/MarketGoLang.git/internal/adapters/postgresql/sqlc"
)

type Service interface { 
	//slice of Repo
	ListProducts(ctx context.Context) ([]repo.Product,error)
}

type svc struct { 
	repo repo.Querier

}

func NewService(repo repo.Querier) Service { 
	return &svc { 
		repo: repo,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product,error) { 
	return s.repo.ListProducts(ctx)
}