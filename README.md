# Mediawiki events stream client for Go

Mediawiki server side events package.

Usage example:
```go
ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

go func() {
	err := eventstream.RevisionScore(ctx, func(evt *events.RevisionScore) {
		fmt.Println(evt)
	})

	if err != nil {
		log.Panic(err)
	}
}()

select {
case <-ctx.Done():
	fmt.Println("done")
	break
}
```