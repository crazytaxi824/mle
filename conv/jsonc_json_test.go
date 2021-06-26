package conv

import (
	"testing"
)

const jsoncTest = `{
  // this is comment
  "s": "abc\" def\" \\\t \n /abc //notComment ~/omg/gg.json", // comment
  "a": 1,
  "b": "ok",
  "c": true,
  "arr": [
    "a", // this is a
    "b",
    "c"
  ]
}

// this is the {comments} after setting.
`

func Test_UnmarshalJSONC(t *testing.T) {
	r, err := JSONCToJSON([]byte(jsoncTest))
	if err != nil {
		t.Error(err)
	}

	t.Log(string(r))
}
