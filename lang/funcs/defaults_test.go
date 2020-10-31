package funcs

import (
	"fmt"
	"testing"

	"github.com/zclconf/go-cty/cty"
)

func TestDefaults(t *testing.T) {
	tests := []struct {
		Input, Defaults cty.Value
		Want            cty.Value
		WantErr         string
	}{
		{
			Input: cty.ObjectVal(map[string]cty.Value{
				"a": cty.NullVal(cty.String),
			}),
			Defaults: cty.ObjectVal(map[string]cty.Value{
				"a": cty.StringVal("hello"),
			}),
			Want: cty.ObjectVal(map[string]cty.Value{
				"a": cty.StringVal("hello"),
			}),
		},
		{
			Input: cty.ObjectVal(map[string]cty.Value{
				"a": cty.StringVal("hey"),
			}),
			Defaults: cty.ObjectVal(map[string]cty.Value{
				"a": cty.StringVal("hello"),
			}),
			Want: cty.ObjectVal(map[string]cty.Value{
				"a": cty.StringVal("hey"),
			}),
		},
		{
			Input: cty.ObjectVal(map[string]cty.Value{
				"a": cty.NullVal(cty.String),
			}),
			Defaults: cty.ObjectVal(map[string]cty.Value{
				"a": cty.NullVal(cty.String),
			}),
			Want: cty.ObjectVal(map[string]cty.Value{
				"a": cty.NullVal(cty.String),
			}),
		},
		{
			Input: cty.ObjectVal(map[string]cty.Value{
				"a": cty.NullVal(cty.String),
			}),
			Defaults: cty.ObjectVal(map[string]cty.Value{}),
			Want: cty.ObjectVal(map[string]cty.Value{
				"a": cty.NullVal(cty.String),
			}),
		},
		{
			Input: cty.ObjectVal(map[string]cty.Value{}),
			Defaults: cty.ObjectVal(map[string]cty.Value{
				"a": cty.NullVal(cty.String),
			}),
			WantErr: `.a: target type does not expect an attribute named "a"`,
		},

		{
			Input: cty.ObjectVal(map[string]cty.Value{
				"a": cty.ListVal([]cty.Value{
					cty.NullVal(cty.String),
				}),
			}),
			Defaults: cty.ObjectVal(map[string]cty.Value{
				"a": cty.StringVal("hello"),
			}),
			Want: cty.ObjectVal(map[string]cty.Value{
				"a": cty.ListVal([]cty.Value{
					cty.StringVal("hello"),
				}),
			}),
		},
		{
			Input: cty.ObjectVal(map[string]cty.Value{
				"a": cty.ListVal([]cty.Value{
					cty.NullVal(cty.String),
					cty.StringVal("hey"),
					cty.NullVal(cty.String),
				}),
			}),
			Defaults: cty.ObjectVal(map[string]cty.Value{
				"a": cty.StringVal("hello"),
			}),
			Want: cty.ObjectVal(map[string]cty.Value{
				"a": cty.ListVal([]cty.Value{
					cty.StringVal("hello"),
					cty.StringVal("hey"),
					cty.StringVal("hello"),
				}),
			}),
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("defaults(%#v, %#v)", test.Input, test.Defaults), func(t *testing.T) {
			got, gotErr := Defaults(test.Input, test.Defaults)

			if test.WantErr != "" {
				if gotErr == nil {
					t.Fatalf("unexpected success\nwant error: %s", test.WantErr)
				}
				if got, want := gotErr.Error(), test.WantErr; got != want {
					t.Fatalf("wrong error\ngot:  %s\nwant: %s", got, want)
				}
				return
			} else if gotErr != nil {
				t.Fatalf("unexpected error\ngot:  %s", gotErr.Error())
			}

			if !test.Want.RawEquals(got) {
				t.Errorf("wrong result\ngot:  %#v\nwant: %#v", got, test.Want)
			}
		})
	}
}
