package conv

import (
	"bytes"
	"encoding/json"
	"testing"
)

const jsoncTest = `{
  // this is comment
  "s": "abc\" def\" \\\t \n /abc //notComment { } ~/omg/gg.json", // comment
  /* haha */ "a": /*dfsd*/ 1,
  "b */":/* omg */ "ok",

  "c": true, /* abc
  sfdsfsfs
  */ "arr": [
    "a", // this is a
    "b",
    "c" /*
	sfsd
	*/ // dfhskjf*/
  ], /* dfhsf
  fsdfdsf
  */
  "sf":{
	"d":"k",  // haha
	"e":"o"
  }  // hahaha
} // omg

// this is the {comment} after setting
`

func Test_UnmarshalJSONC(t *testing.T) {
	r, err := JSONCToJSON([]byte(jsoncTest))
	if err != nil {
		t.Error(err)
	}

	t.Log(string(r))

	var buf bytes.Buffer
	err = json.Indent(&buf, r, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(buf.String())
}
