# Mediawiki events stream client for Go

Mediawiki server side events client package.

Usage example:
```go
ctx, cancel := context.WithCancel(context.Background())
client := eventstream.NewClient()

stream := client.RevisionCreate(ctx, time.Now(), func(evt *events.RevisionCreate) {
	fmt.Println(evt)
})

go func() {
	time.Sleep(2 * time.Second)
	cancel()
}()

for err := range stream.Sub() {
	fmt.Println(err)
}
```