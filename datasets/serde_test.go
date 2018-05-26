package datasets

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
	"os"
	"testing"
)

func strPtr(s string) *string {
	return &s
}

func TestSerDe(t *testing.T) {
	c, err := os.Create("testserde")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	enc := msgpack.NewEncoder(c)
	if err := enc.Encode(&ActorEntry{MovieId: "123", Name: []*string{strPtr("one")}}); err != nil {
		panic(err)
	}
	if err := enc.Encode(&ActorEntry{MovieId: "1234", Name: []*string{strPtr("two")}}); err != nil {
		panic(err)
	}

	o, err := os.Open("testserde")
	if err != nil {
		panic(err)
	}
	defer o.Close()

	dec := msgpack.NewDecoder(o)

	entry := &ActorEntry{}

	if err := dec.Decode(entry); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", *entry)

	if err := dec.Decode(entry); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", *entry)
}

func TestActorsDe(t *testing.T) {
	o, err := os.Open("../actors.segment.00")
	if err != nil {
		panic(err)
	}
	defer o.Close()

	dec := msgpack.NewDecoder(o)

	entry := &ActorEntry{}

	if err := dec.Decode(entry); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", *entry)

	if err := dec.Decode(entry); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", *entry)
}

func TestTitlesDe(t *testing.T) {
	o, err := os.Open("../titles.segment.00")
	if err != nil {
		panic(err)
	}
	defer o.Close()

	dec := msgpack.NewDecoder(o)

	entry := &TitleEntry{}

	if err := dec.Decode(entry); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", *entry)

	if err := dec.Decode(entry); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", *entry)
}
