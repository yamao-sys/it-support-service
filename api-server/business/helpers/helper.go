package businesshelpers

import "context"

type key string

const (
	ctxSupporterIDKey key = "SupporterID"
	ctxCompanyIDKey key = "CompanyID"
)

func NewWithSupporterIDContext(ctx context.Context, v int) context.Context {
	return context.WithValue(ctx, ctxSupporterIDKey, v)
}

func NewWithCompanyIDContext(ctx context.Context, v int) context.Context {
	return context.WithValue(ctx, ctxCompanyIDKey, v)
}

func ExtractSupporterID(ctx context.Context) (int, bool) {
	v, ok := ctx.Value(ctxSupporterIDKey).(int)
	return int(v), ok
}

func ExtractCompanyID(ctx context.Context) (int, bool) {
	v, ok := ctx.Value(ctxCompanyIDKey).(int)
	return int(v), ok
}
