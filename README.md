# Gin&Tonic

Documentation | [中文文档](README_zh.md)

## Introduction

Gin&Tonic is a `LemonSqueezy` webhook library for `gin` framework that lets you build `LemonSqueezy` webhooks quickly and elegantly

## Quick start

Installation

```bash
go get -u github.com/YianAndCode/gintonic
```

Import

```go
import "github.com/YianAndCode/gintonic"
```

Write your bussines code

```go
import (
	"context"
	"fmt"

	"github.com/YianAndCode/gintonic"
	"github.com/YianAndCode/gintonic/lmdata"
)

var _ gintonic.OrderCreatedHandler = OrderCreated

func OrderCreated(ctx context.Context, meta lmdata.Meta, data lmdata.Order) error {
	fmt.Printf("You made a sale! Total amount is %s\n", data.Attributes.TotalFormatted)
	return nil
}
```

Register routes to `gin` and run

```go
func main() {
	r := gin.New()
	r.Use(gin.Recovery())

	gt := gintonic.New(
		"[Load secret from somewhere]",
		gintonic.WithOrderCreatedHandler(OrderCreated),
	)
	r.POST("/lemonsqueezy/webhook", gt.LemonSqueezyWebhook)

	r.Run()
}
```

## Definition

### Instantiate

You can get an instantce of `gintonic` just by the `New` function, and the first parameter is fixed as `secret`, then you can set any handlers as needed:

```go
gt := gintonic.New(
    "[YOUR_SECRET]",
    gintonic.WithOrderCreatedHandler(OrderCreated), // The handler for orcer_created event
    gintonic.WithDefaultHandler(DefaultHandler),    // The default handler
)
```

Setting handler is done using the `gintonic.With[EventName]Handler()` option functions, the specific functions can be found in `option.go`

### Handler

Handler definitions are in `handler.go`, divided into two types: `Event Handler` and `DefaultHandler`.

Their relationship is like `case` and `default` in a `switch` statement. When instantiating a gintonic instance with `New`, the `Handler` registered using `gintonic.WithXXXHandler(XXX)` can be considered a `case`. When receiving a `webhook` event, `gintonic` will judge which `case` it belongs to and then call the corresponding `handler`. If it cannot find one, it will look for a `DefaultHandler`. If it cannot find the corresponding `Event Handler` and no `DefaultHandler` is set, then it will return an error.

`DefaultHandler` is defined as:

```go
type DefaultHandler func(ctx context.Context, eventName string, meta lmdata.Meta, data interface{}) error
```

`Event Handler` is like:

```go
type OrderCreatedHandler func(ctx context.Context, meta lmdata.Meta, data lmdata.Order) error
```

The only diffrence is the type of `data`.

### Core

`gintonic` interacts with `gin` using the `GinTonic.LemonSqueezyWebhook` method. This method will read `HTTP Headers` and `POST body` from `gin`, then verify the signature. After passing verification, it dispatches events to the registered `handler`.

So we just need to register the `LemonSqueezyWebhook` method of the `GinTonic` instance to a route in `gin`.