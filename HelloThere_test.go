package hellothere_test

import (
	"fmt"
	"testing"

	hellothere "github.com/Rahul-NITD/SAGo/HelloThere"
)

func TestHelloThere(t *testing.T) {

	cases := []struct {
		name     string
		lang     string
		expected string
	}{
		{"Rahul", "English", "Hello Rahul, have a nice day!"},
		{"Akku", "English", "Hello Akku, have a nice day!"},
		{"", "English", "Hello User, have a nice day!"},

		{"Rahul", "Hinglish", "Namaste Rahul, aapka din shubh ho!"},
		{"Akku", "Hinglish", "Namaste Akku, aapka din shubh ho!"},
		{"", "Hinglish", "Namaste Meri Jaan, aapka din shubh ho!"},

		{"Rahul", "Spanish", "Hola Rahul, ¡que tengas un gran día!"},
		{"Akku", "Spanish", "Hola Akku, ¡que tengas un gran día!"},
		{"", "Spanish", "Hola Usuario, ¡que tengas un gran día!"},

		{"Rahul", "Turkish", "Hello Rahul, have a nice day!"},
		{"", "", "Hello User, have a nice day!"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("return greeting for name: `%s` | lang: `%s`", test.name, test.lang), func(t *testing.T) {
			got := hellothere.SayHello(test.name, test.lang)
			if got != test.expected {
				t.Errorf("got `%s`, want `%s`", got, test.expected)
			}
		})
	}
}
