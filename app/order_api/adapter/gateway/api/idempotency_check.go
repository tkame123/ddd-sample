package connect

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/idempotency"
	"log"
)

const idempotencyKeyHeader = "Idempotency-Key"

func (s *Server) NewIdempotencyCheckInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if req.Spec().IdempotencyLevel != connect.IdempotencyIdempotent {
				return next(ctx, req)
			}

			idempotencyKey := req.Header().Get(idempotencyKeyHeader)
			if idempotencyKey == "" {
				return nil, connect.NewError(
					connect.CodeInvalidArgument,
					errors.New("no idempotencyKey"),
				)
			}

			key, err := s.repIdempotency.IdempotencyKeyFindByID(ctx, idempotencyKey)
			if err != nil && !idempotency.IsNotFound(err) {
				log.Printf("error caused by %v", err)
				return nil, connect.NewError(
					connect.CodeInternal,
					errors.New("internal error"),
				)
			}

			// Keyが存在する場合は、処理実施しない
			if key != nil {
				switch key.Status {
				case idempotency.IdempotencyKeyStatusProcessing:
					return nil, connect.NewError(
						connect.CodeAlreadyExists,
						errors.New("it is Processing"),
					)
				case idempotency.IdempotencyKeyStatusSuccess:
					// TODO: return cache response(やり方を確認する必要がある）
					return nil, connect.NewError(
						connect.CodeUnimplemented,
						errors.New("exits key is unimplemented"),
					)
				}
			}

			err = s.repIdempotency.IdempotencyKeySave(ctx, idempotency.IdempotencyKey{
				ID:      idempotencyKey,
				Status:  idempotency.IdempotencyKeyStatusProcessing,
				Request: req,
			})
			if err != nil {
				log.Printf("error caused by %v", err)
				return nil, connect.NewError(
					connect.CodeInternal,
					errors.New("internal error"),
				)
			}

			res, err := next(ctx, req)
			if err != nil {
				log.Printf("error caused by %v", err)
				return nil, err
			}

			err = s.repIdempotency.IdempotencyKeySave(ctx, idempotency.IdempotencyKey{
				ID:       idempotencyKey,
				Status:   idempotency.IdempotencyKeyStatusSuccess,
				Request:  req,
				Response: res,
			})
			if err != nil {
				log.Printf("error caused by %v", err)
				return nil, connect.NewError(
					connect.CodeInternal,
					errors.New("internal error"),
				)
			}

			return res, nil
		}
	}
	return interceptor
}
