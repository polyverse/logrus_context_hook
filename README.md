# logrus_context_hook
A Logrus Hook to pull keys from Context (when passed) and convert into log fields.

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
	log.AddHook(logrus_context_hook.NewContextHook("*", // * is special for ANY field of type context.Context.
		[]string{"ServerId", "RequestId", "HostId"})) //Second parameter is the list of keys in that context to look for.

	// Then log with as you do
	log.Infof("Test1: Hello from Logger...")

	// OR Log with Context field name
	ctx := context.WithValue(context.Background(), "ServerId", "DemoServer")
	log.WithField("Context", ctx).Infof("Test1: Hello from Logger with context under field 'Context'")

	// OR Log with Context under any field name (because of the "*")
	ctx = context.WithValue(context.Background(), "RequestId", "Want this logged anyway")
	log.WithField("NotContextField", ctx).Infof("Test1: Hello from Logger with context under field 'NotContextField'")
}
```

## Only a specific context field
```
package main

import (
  log "github.com/sirupsen/logrus"
  "github.com/polyverse/logrus_context_hook"
)

func main() {

	// Add the Context Hook
	log.AddHook(logrus_context_hook.NewContextHook("Context", // First parameter is the field name to look for context at
		[]string{"ServerId", "RequestId", "HostId"})) //Second parameter is the list of keys in that context to look for.

	// Then log with as you do
	log.Infof("Test2: Hello from Logger...");

	// OR Log with Context field Name
	ctx := context.WithValue(context.Background(), "ServerId", "DemoServer")
	log.WithField("Context", ctx).Infof("Test2: Hello from Logger with context under field 'Context'")

	// EXCEPT Log without Context field name, and hook does nothing
	ctx = context.WithValue(context.Background(), "RequestId", "Don't Want this logged")
	log.WithField("NotContextField", ctx).Infof("Test2: Hello from Logger with context under field 'NotContextField'")
}
```

## Change Context Field or Keys to look for

```
package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/polyverse/logrus_context_hook"
)

func main() {


	// Hold reference to the hook
	ctxHook := logrus_context_hook.NewContextHook("Context", // First parameter is the field to look for context at
		[]string{"ServerId", "RequestId", "HostId"}) //Second parameter is the list of keys in that context to look for.

	// Add the Context Hook
	log.AddHook(ctxHook)

	// Then log with as you do
	log.Infof("Test3: Hello from Logger...");

	// OR Log with Context field
	ctx := context.WithValue(context.Background(), "ServerId", "DemoServer")
	log.WithField("Context", ctx).Infof("Test3: Hello from Logger with context under field 'Context'")

	// EXCEPT Log without Context field, and hook does nothing
	ctx = context.WithValue(context.Background(), "RequestId", "Don't Want this logged")
	log.WithField("NotContextField", ctx).Infof("Test3: Hello from Logger with context under field 'NotContextField'")

	// BUT we can change that at runtime
	ctxHook.SetContextField("*")

	// AND it works now...
	ctx = context.WithValue(context.Background(), "RequestId", "Changed our mind at runtime. We want this logged.")
	log.WithField("NotContextField", ctx).Infof("Test3: Hello from Logger with context under field 'NotContextField'")

	// EXCEPT we just need a new key called UserId
	ctx = context.WithValue(context.Background(), "UserId", "This key wasn't being looked for")
	log.WithField("NotContextField", ctx).Infof("Test3: Hello from Logger with context having new key 'UserId'")

	// BUT we can change that at runtime TOO
	ctxHook.SetContextKeys(append(ctxHook.GetContextKeys(), "UserId"))

	// AMD it works now
	ctx = context.WithValue(context.Background(), "UserId", "This key is being looked for. Changed our mind at runtime again.")
	log.WithField("NotContextField", ctx).Infof("Test3: Hello from Logger with context having new key 'UserId'")
}
```


