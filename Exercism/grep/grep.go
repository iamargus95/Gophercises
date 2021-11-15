package grep

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Option grep options
type Option struct {
	invert, ignoreCase, lineNumber, filename, entireLine bool
}

// Grep grep structure
type Grep struct {
	files   []string
	option  Option
	current int
	pattern *regexp.Regexp
	results []string
}

// ParseOptions parse string and get options
func ParseOptions(flags []string) Option {
	opt := Option{}
	for i := 0; i < len(flags); i++ {
		switch flags[i] {
		case "-n":
			opt.lineNumber = true
		case "-l":
			opt.filename = true
		case "-i":
			opt.ignoreCase = true
		case "-x":
			opt.entireLine = true
		case "-v":
			opt.invert = true
		}
	}
	return opt
}

// Match match two strings using given params
func (opt Option) Match(rx *regexp.Regexp, text string) (bool, bool) {
	match := false
	done := false
	if opt.ignoreCase {
		text = strings.ToLower(text)
	}
	match = rx.MatchString(text)
	if opt.filename && match {
		done = true
	}
	if opt.invert {
		match = !match
	}
	return match, done
}

// New get new grep
func New(files []string, pattern string, opt Option) *Grep {
	if opt.ignoreCase {
		pattern = strings.ToLower(pattern)
	}
	if opt.entireLine {
		pattern = "^" + pattern + "$"
	}
	rx := regexp.MustCompile(pattern)
	r := Grep{option: opt, files: files, pattern: rx, results: make([]string, 0)}
	return &r
}

// Current get current file name
func (g *Grep) Current() string {
	if g.current >= 0 && g.current < len(g.files) {
		return g.files[g.current]
	}
	return ""
}

// Save save this string in results
func (g *Grep) Save(n int, text string) {
	if g.option.filename {
		g.results = append(g.results, g.Current())
	} else {
		s := ""
		if g.option.lineNumber {
			s = strconv.Itoa(n) + ":" + text
		} else {
			s = text
		}
		if len(g.files) == 1 {
			g.results = append(g.results, s)
		} else {
			g.results = append(g.results, g.Current()+":"+s)
		}
	}
}

// Search run a search using grep
func (g *Grep) Search() error {
	for g.current = 0; g.current < len(g.files); g.current++ {
		fname := g.Current()
		f, err := os.Open(fname)
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(f)
		index := 0
		for scanner.Scan() {
			index++
			line := scanner.Text()
			status, done := g.option.Match(g.pattern, line)
			if status {
				g.Save(index, line)
			}
			if done {
				break
			}
		}
		if err := scanner.Err(); err != nil {
			return err
		}
	}
	return nil
}

// Search run grep on all files for pattern with flags
func Search(pattern string, flags, files []string) []string {
	opt := ParseOptions(flags)
	grep := New(files, pattern, opt)
	// Search
	err := grep.Search()
	if err != nil {
		log.Fatal(err)
	}
	return grep.results
}
