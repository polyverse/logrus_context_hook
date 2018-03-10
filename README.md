# logrus_context_hook
A Logrus Hook to pull fields from Context (when passed) and convert into fields.

# Use

## Easiest

```
package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/polyverse/logrus_context_hook"
)

func main() {
  // Add the Context Hook
	log.AddHook(logrus_context_hook.NewContextHook("*", // * is special for ANY field of time context.Context.
    []string{"ServerId", "RequestId", "HostId"} //Second field is the list of keys in that context to look for.
  ))
  
  // Then log with as you do
  log.Infof("Hello from Logger...");

  // OR Log with Context
  ctx := log.WithValue(context.Background(), "ServerId", "DemoServer")
  log.WithField("Context", ctx).Infof("Hello from Logger with context under key 'Context'")

  // OR Log with Context Fields under any key name
  ctx = log.WithValue(context.Background(), "RequestId", "Want this logged anyway")
  log.WithField("NotContextKey", ctx).Infof("Hello from Logger with context under key 'NotContextKey'")
}
```

## Only a specific context Key
```
package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/polyverse/logrus_context_hook"
)

func main() {
  // Add the Context Hook
	log.AddHook(logrus_context_hook.NewContextHook("Context", // First field is the key to look for context at
    []string{"ServerId", "RequestId", "HostId"} //Second field is the list of keys in that context to look for.
  ))
  
  // Then log with as you do
  log.Infof("Hello from Logger...");

  // OR Log with Context Key
  ctx := log.WithValue(context.Background(), "ServerId", "DemoServer")
  log.WithField("Context", ctx).Infof("Hello from Logger with context under key 'Context'")
  
  // EXCEPT Log without Context Key, and hook does nothing
  ctx = log.WithValue(context.Background(), "RequestId", "Don't Want this logged")
  log.WithField("NotContextKey", ctx).Infof("Hello from Logger with context under key 'NotContextKey'")

}
```

## Change Context Key or Fields to look for

```
package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/polyverse/logrus_context_hook"
)

func main() {

  // Hold reference to the hook
  ctxHook := logrus_context_hook.NewContextHook("Context", // First field is the key to look for context at
    []string{"ServerId", "RequestId", "HostId"} //Second field is the list of keys in that context to look for.
  )

  // Add the Context Hook
	log.AddHook(ctxHook)
  
  // Then log with as you do
  log.Infof("Hello from Logger...");

  // OR Log with Context Key
  ctx := log.WithValue(context.Background(), "ServerId", "DemoServer")
  log.WithField("Context", ctx).Infof("Hello from Logger with context under key 'Context'")
  
  // EXCEPT Log without Context Key, and hook does nothing
  ctx = log.WithValue(context.Background(), "RequestId", "Don't Want this logged")
  log.WithField("NotContextKey", ctx).Infof("Hello from Logger with context under key 'NotContextKey'")
  
  // BUT we can change that at runtime
  ctxHook.SetContextKey("*")

  // AND it works now...
  ctx = log.WithValue(context.Background(), "RequestId", "Changed our mind at runtime. We want this logged.")
  log.WithField("NotContextKey", ctx).Infof("Hello from Logger with context under key 'NotContextKey'")
  
  // EXCEPT we just need a new Field called UserId
  ctx = log.WithValue(context.Background(), "UserId", "This won't field wasn't being looked for")
  log.WithField("NotContextKey", ctx).Infof("Hello from Logger with context having new key 'UserId'")

  // BUT we can change that at runtime TOO
  ctxHook.SetContextFields(append(ctxHook.GetContextFields(), "UserId"))

  // AMD it works now
  ctx = log.WithValue(context.Background(), "UserId", "This won't field was being looked for. Changed our mind at runtime again.")
  log.WithField("NotContextKey", ctx).Infof("Hello from Logger with context having new key 'UserId'")

}
```
