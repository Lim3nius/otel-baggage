package main

import (
	"context"
	"fmt"
	"net/url"

	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/propagation"
)

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

// Simple program to demonstrate problem with baggage.
// First is baggage created and set member "ok-key": "messed,up%value" with proper url escaping, and is accepted as valid baggage
// then given baggage is injected into HeaderCarrier and everything is fine, bad characters (",", "%") are escaped
// and finally new baggage is parsed from previously used HeaderCarrier, and fails to parse baggage because of comma and percent sign in string
func main() {
	bag := must(baggage.New())
	bag = must(bag.SetMember(
		must(baggage.NewMember("ok-key", url.QueryEscape("messed,up%value")))))

	fmt.Println("injected baggage")
	for _, m := range bag.Members() {
		fmt.Printf("  %q: %q\n", m.Key(), m.Value())
	}

	ctx := baggage.ContextWithBaggage(context.Background(), bag)
	hc := propagation.HeaderCarrier{}

	propagation.Baggage{}.Inject(ctx, hc)

	fmt.Printf("\nheader carrier after inject: %+v\n\n", hc)

	nctx := propagation.Baggage{}.Extract(context.Background(), hc)
	nbag := baggage.FromContext(nctx)

	fmt.Println("extracted baggage")
	for _, m := range nbag.Members() {
		fmt.Printf("  %q: %q\n", m.Key(), m.Value())
	}

	if len(nbag.Members()) == 0 {
		panic("propagation failed")
	}

	// baggage which is accepted by OTEL lib and propagated isn't accepted by same OTEL lib
	b, err := baggage.Parse(hc.Get("Baggage"))
	fmt.Printf("\n\nparsing err: %s", err)
	if b.Len() != 0 {
		panic("It works!")
	}
}
