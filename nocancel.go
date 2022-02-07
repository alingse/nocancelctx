package nocancelctx

import (
	"context"
	"time"
)

type noCancelCtx struct {
	ctx context.Context
}

func WithoutCancel(ctx context.Context) context.Context {
	return &noCancelCtx{ctx: ctx}
}

func WithoutDeadline(ctx context.Context) context.Context {
	return &noCancelCtx{ctx: ctx}
}

func (c *noCancelCtx) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (c *noCancelCtx) Done() <-chan struct{} {
	return nil
}

func (c *noCancelCtx) Err() error {
	return nil
}

func (c *noCancelCtx) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}
