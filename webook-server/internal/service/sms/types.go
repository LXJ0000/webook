package sms

import "context"

type Service interface {
	Send(ctx context.Context, templateId string, args []string, numbers ...string) error
}
