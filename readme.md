# log

This is a simple logger that uses `zap/logger` as the core logger

## Docs

```go
package main

func main() {
    // simplest startup, no other setup needed
	l := log.New("service")

    // pass log level and output to configure
	l = New("service-name", WithLevel("INFO"), WithOutput(w))
	
    // use logger with custom fields and level accordingly
    l.WithFields(
		zap.String("somefield", "somevalue"),
		zap.String("somefield2", "somevalue2")).Info("some log")

    l.WithFields(
		zap.String("somefield", "somevalue"),
		zap.String("somefield2", "somevalue2")).Debug("some log")

    l.Error("some error")
    l.Warn("some warning")
}
```

