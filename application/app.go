package application

import (
	"context"

	"github.com/dickyadrian/paper-disbursement-system/config"
	"github.com/dickyadrian/paper-disbursement-system/internal/repository"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type App struct {
	AsynqClient   *asynq.Client
	AsynqRedisOpt asynq.RedisConnOpt

	AppConfig config.App

	DB                     *pgxpool.Pool
	UserRepository         *repository.UserRepository
	DisbursementRepository *repository.DisbursementRepository
}

func NewApp(ctx context.Context) (*App, error) {
	appConfig := config.LoadApp()
	dbConfig := config.LoadDatabase()
	redisConfig := config.LoadRedis()

	db, err := repository.InitClient(ctx, dbConfig)
	if err != nil {
		return nil, err
	}

	redisOpt, err := redis.ParseURL(redisConfig.URL)
	if err != nil {
		return nil, err
	}

	asynqRedisOpt := asynq.RedisClientOpt{
		Addr:     redisOpt.Addr,
		Password: redisOpt.Password,
		DB:       redisOpt.DB,
	}

	return &App{
		AsynqClient:            asynq.NewClient(asynqRedisOpt),
		AsynqRedisOpt:          asynqRedisOpt,
		AppConfig:              appConfig,
		DB:                     db,
		UserRepository:         repository.NewUserRepository(db),
		DisbursementRepository: repository.NewDisbursementRepository(db),
	}, nil
}
