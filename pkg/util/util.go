package util

import (
	"context"
	"encoding/json"
)

func UserId(ctx context.Context) (int64, error) {
	return ctx.Value("user_id").(json.Number).Int64()
}
