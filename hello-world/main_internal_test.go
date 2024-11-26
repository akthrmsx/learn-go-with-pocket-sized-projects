package main

import "testing"

func Example() {
	main()
	// Output:
	// Hello world
}

func TestGreet(t *testing.T) {
	var tests = map[string]struct {
		lang language
		want string
	}{
		"Greek": {
			lang: "el",
			want: "Χαίρετε Κόσμε",
		},
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Hebrew": {
			lang: "he",
			want: "שלום עולם",
		},
		"Urdu": {
			lang: "ur",
			want: "ہیلو دنیا",
		},
		"Vietnamese": {
			lang: "vi",
			want: "Xin chào Thế Giới",
		},
		"Akkadian, not supported": {
			lang: "akk",
			want: `unsupported language: "akk"`,
		},
		"Empty": {
			lang: "",
			want: `unsupported language: ""`,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tt.lang)

			if got != tt.want {
				t.Errorf("expected: %q, got: %q", tt.want, got)
			}
		})
	}
}
