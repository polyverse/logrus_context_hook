package logrus_context_hook

import (
	"context"
	"github.com/sirupsen/logrus"
)

type ContextHook interface {
	logrus.Hook

	GetContextField() string
	SetContextField(string)

	GetContextKeys() []string
	SetContextKeys([]string)
}

type contextHookImpl struct {
	contextField string
	contextKeys  []string
}

func NewContextHook(contextField string, contextKeys []string) ContextHook {
	return &contextHookImpl{
		contextField: contextField,
		contextKeys:  contextKeys,
	}
}

func (hook *contextHookImpl) GetContextField() string {
	return hook.contextField
}

func (hook *contextHookImpl) SetContextField(ctxField string) {
	hook.contextField = ctxField
}

func (hook *contextHookImpl) GetContextKeys() []string {
	return hook.contextKeys
}

func (hook *contextHookImpl) SetContextKeys(ctxKeys []string) {
	hook.contextKeys = ctxKeys
}

func (hook *contextHookImpl) Fire(entry *logrus.Entry) error {
	// Context field
	field := hook.contextField

	if field == "*" { // if field is a wildcard
		for k, v := range entry.Data {
			if _, ok := v.(context.Context); ok { // find first field with value of type context, and use that.
				field = k
				break
			}
		}
	}

	if field == "*" {
		return nil
	}

	// Get value
	ctxCandidate, ok := entry.Data[field]
	if !ok {
		return nil
	}

	// Make sure value is of type context.Context
	ctx, ok := ctxCandidate.(context.Context)
	if !ok {
		return nil
	}

	//Delete original field
	delete(entry.Data, field)

	for _, key := range hook.contextKeys {
		value := ctx.Value(key)
		if value != nil {
			entry.Data[key] = value
		}
	}

	return nil
}

// Operate on ALL levels
func (hook *contextHookImpl) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}
