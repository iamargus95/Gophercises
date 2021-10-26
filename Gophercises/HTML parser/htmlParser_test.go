package htmllink

import (
	"reflect"
	"strings"
	"testing"
)

var testCases = []struct {
	desc   string
	input  string
	output []Link
}{
	{
		desc: "single href attr",
		input: `<html>
		<body>
		  <h1>Hello!</h1>
		  <a href="/other-page">A link to another page</a>
		</body>
		</html>`,
		output: []Link{{Href: "/other-page", Text: "A link to another page"}},
	},
	{
		desc: "Multiple attr",
		input: `<html>
		<head>
		  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
		</head>
		<body>
		  <h1>Social stuffs</h1>
		  <div>
			<a href="https://www.twitter.com/joncalhoun">
			  Check me out on twitter
			  <i class="fa fa-twitter" aria-hidden="true"></i>
			</a>
			<a href="https://github.com/gophercises">
			  Gophercises is on <strong>Github</strong>!
			</a>
		  </div>
		</body>
		</html>`,
		output: []Link{{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"}, {Href: "https://github.com/gophercises", Text: "Gophercises is on Github !"}},
	},
}

func TestParse(t *testing.T) {

	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			got, _ := Parse(strings.NewReader(v.input))
			if !reflect.DeepEqual(got, v.output) {
				t.Errorf("Parse(%s) = %s; want %s", v.input, got, v.output)
			}
		})
	}
}
