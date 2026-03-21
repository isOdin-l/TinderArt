package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type ICache interface {
	Eval(ctx context.Context, script string, keys []string, args ...any) error
}

type IRepo interface {
	FindMatches(ctx context.Context, id pgtype.UUID) ([]interface{}, error)
	GetAllProfiles(ctx context.Context) ([]pgtype.UUID, error)
}

type Service struct {
	cache ICache
	repo  IRepo
}

func NewService(cache ICache, repo IRepo) *Service {
	return &Service{cache: cache, repo: repo}
}

func (s *Service) GenerateDailyStack(ctx context.Context) error {
	profiles, errAllProfiles := s.repo.GetAllProfiles(ctx)
	if errAllProfiles != nil {
		return errAllProfiles
	}

	for _, profile := range profiles {
		matches, errMatch := s.repo.FindMatches(ctx, profile)
		if errMatch != nil {
			return errMatch
		}

		key := []string{fmt.Sprintf("stack:%s", profile.String())}

		lua := `
		redis.call("DEL", KEYS[1])
		for i=1,#ARGV do
		    redis.call("RPUSH", KEYS[1], ARGV[i])
		end
		return #ARGV
		`

		errEval := s.cache.Eval(ctx, lua, key, matches...)
		if errEval != nil {
			return errEval
		}
	}
	return nil
}
