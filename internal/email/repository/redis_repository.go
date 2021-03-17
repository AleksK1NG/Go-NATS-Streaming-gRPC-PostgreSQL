package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AleksK1NG/nats-streaming/internal/models"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

const (
	prefix     = "emails"
	expiration = time.Second * 3600
)

type emailRedisRepository struct {
	redis *redis.Client
}

// NewEmailRedisRepository emails redis repository constructor
func NewEmailRedisRepository(redis *redis.Client) *emailRedisRepository {
	return &emailRedisRepository{redis: redis}
}

func (e *emailRedisRepository) SetEmail(ctx context.Context, email *models.Email) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "emailRedisRepository.SetEmail")
	defer span.Finish()

	emailBytes, err := json.Marshal(email)
	if err != nil {
		return errors.Wrap(err, "emailRedisRepository.Marshal")
	}

	return e.redis.SetEX(ctx, e.createKey(email.EmailID), string(emailBytes), expiration).Err()
}

func (e *emailRedisRepository) GetEmailByID(ctx context.Context, emailID uuid.UUID) (*models.Email, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "emailRedisRepository.GetEmailByID")
	defer span.Finish()

	result, err := e.redis.Get(ctx, e.createKey(emailID)).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "emailRedisRepository.redis.Get")
	}

	var res models.Email
	if err := json.Unmarshal(result, &res); err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal")
	}
	return &res, nil
}

func (e *emailRedisRepository) DeleteEmail(ctx context.Context, emailID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "emailRedisRepository.DeleteEmail")
	defer span.Finish()
	return e.redis.Del(ctx, e.createKey(emailID)).Err()
}

func (e *emailRedisRepository) createKey(emailID uuid.UUID) string {
	return fmt.Sprintf("%s: %s", prefix, emailID.String())
}
