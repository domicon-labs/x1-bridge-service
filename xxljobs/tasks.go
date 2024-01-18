package xxljobs

import (
	"context"

	"github.com/xxl-job/xxl-job-executor-go"
)

func Task1(ctx context.Context, param *xxl.RunReq) string {
	return "This is task 1"
}

func Task2(ctx context.Context, param *xxl.RunReq) string {
	return "This is task 2"
}

func TaskPanic(ctx context.Context, param *xxl.RunReq) string {
	panic("This task panic!")
	return "This task always panics"
}
