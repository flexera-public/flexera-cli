package cli

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

func TestParseFields(t *testing.T) {
	got, err := ParseFields(" id , owner . email, owner.name , odd-key.$value ")
	if err != nil {
		t.Fatal(err)
	}
	want := []FieldPath{{"id"}, {"owner", "email"}, {"owner", "name"}, {"odd-key", "$value"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}

func TestParseFieldsRejectsInvalid(t *testing.T) {
	for _, input := range []string{"", " ", ",id", "id,", "id,,name", ".id", "id.", "owner..email", "owner. .email", "id,id", "id, id ", "owner,owner.email", "owner.email,owner", "owner.email,owner.email.name"} {
		t.Run(input, func(t *testing.T) {
			paths, err := ParseFields(input)
			var exit *ExitError
			if paths != nil || !errors.As(err, &exit) || exit.Code != 2 {
				t.Fatalf("got paths %#v, error %v; want exit code 2", paths, err)
			}
		})
	}
}

func TestCompileFields(t *testing.T) {
	for _, tc := range []struct {
		name, fields, input, want string
		envelope                  bool
	}{
		{"object", "id,owner.email,owner.name", `{"id":1,"owner":{"email":"a","name":"b","extra":true},"extra":2}`, `{"id":1,"owner":{"email":"a","name":"b"}}`, false},
		{"array", "id", `[{"id":1,"extra":2},{"id":2}]`, `[{"id":1},{"id":2}]`, false},
		{"empty-array", "id", `[]`, `[]`, false},
		{"missing-parent", "owner.email", `{}`, `{"owner":{"email":null}}`, false},
		{"null-parent", "owner.email", `{"owner":null}`, `{"owner":{"email":null}}`, false},
		{"missing-deep-parent", "owner.contact.email", `{"owner":{}}`, `{"owner":{"contact":{"email":null}}}`, false},
		{"envelope", "id", `{"values":[{"id":1,"extra":2}],"nextPage":"next","total":4}`, `{"values":[{"id":1}],"nextPage":"next","total":4}`, true},
		{"null-values", "id", `{"values":null,"total":0}`, `{"values":null,"total":0}`, true},
		{"empty-values", "id", `{"values":[],"total":0}`, `{"values":[],"total":0}`, true},
		{"ordinary-values-key", "id", `{"id":1,"values":[{"id":2}]}`, `{"id":1}`, false},
		{"envelope-fallback", "id", `{"id":1,"extra":2}`, `{"id":1}`, true},
		{"special-keys", "odd-key.$value,quote\"key,back\\slash,雪", `{"odd-key":{"$value":3},"quote\"key":4,"back\\slash":5,"雪":6}`, `{"odd-key":{"$value":3},"quote\"key":4,"back\\slash":5,"雪":6}`, false},
		{"jq-injection", `x"] | error("injected") | ["y`, `{"x\"] | error(\"injected\") | [\"y":7}`, `{"x\"] | error(\"injected\") | [\"y":7}`, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			paths, err := ParseFields(tc.fields)
			if err != nil {
				t.Fatal(err)
			}
			code, err := CompileFields(paths, tc.envelope)
			if err != nil {
				t.Fatal(err)
			}
			var input, want any
			if err := json.Unmarshal([]byte(tc.input), &input); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal([]byte(tc.want), &want); err != nil {
				t.Fatal(err)
			}
			iter := code.Run(input)
			got, ok := iter.Next()
			if !ok || !reflect.DeepEqual(got, want) {
				t.Fatalf("got %#v, want %#v", got, want)
			}
			if extra, ok := iter.Next(); ok {
				t.Fatalf("unexpected additional result: %#v", extra)
			}
		})
	}
}

func TestCompileFieldsRejectsInvalidPaths(t *testing.T) {
	for _, paths := range [][]FieldPath{nil, {{}}, {{""}}, {{"id"}, {"id"}}, {{"owner"}, {"owner", "email"}}} {
		if code, err := CompileFields(paths, false); err == nil || code != nil {
			t.Fatalf("got code %v, error %v for %#v", code, err, paths)
		}
	}
}
