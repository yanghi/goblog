package serializer

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strings"
	"unicode/utf8"

	"github.com/russross/blackfriday"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

type postSerializer struct {
	extractDescriptionSize int
}

// 通过内容提取markdown描述信息
func (p *postSerializer) ExtractMarkdownDescription(content string) string {
	content = content[:min(p.extractDescriptionSize, len(content))]
	output := blackfriday.MarkdownBasic([]byte(content))

	var text = stripHtmlTags(string(output))

	// return strings.Replace(text, "\n", "&nbsp;", -1)
	return text
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

const (
	htmlTagStart = 60 // Unicode `<`
	htmlTagEnd   = 62 // Unicode `>`
)

func stripHtmlTags(s string) string {
	// Setup a string builder and allocate enough memory for the new string.
	var builder strings.Builder
	builder.Grow(len(s) + utf8.UTFMax)

	in := false // True if we are inside an HTML tag.
	start := 0  // The index of the previous start tag character `<`
	end := 0    // The index of the previous end tag character `>`

	for i, c := range s {
		// If this is the last character and we are not in an HTML tag, save it.
		if (i+1) == len(s) && end >= start {
			builder.WriteString(s[end:])
		}

		// Keep going if the character is not `<` or `>`
		if c != htmlTagStart && c != htmlTagEnd {
			continue
		}

		if c == htmlTagStart {
			// Only update the start if we are not in a tag.
			// This make sure we strip out `<<br>` not just `<br>`
			if !in {
				start = i
			}
			in = true

			// Write the valid string between the close and start of the two tags.
			builder.WriteString(s[end:start])
			continue
		}
		// else c == htmlTagEnd
		in = false
		end = i + 1
	}
	s = builder.String()
	return s
}
func (ps *postSerializer) DecodeContent(s string) ([]byte, error) {
	var reader = DecodeEncoding(bytes.NewReader([]byte(s)))
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}

	return d, nil
}

var Post = postSerializer{extractDescriptionSize: 300}

// 将对应格式文本转换成utf-8
func DecodeEncoding(r io.Reader) *transform.Reader {
	e := determineEncodeing(r)
	return transform.NewReader(r, e.NewDecoder())
}

// 判断传输来的文本的字符集格式是什么
func determineEncodeing(r io.Reader) encoding.Encoding {
	peek, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	determineEncoding, _, _ := charset.DetermineEncoding(peek, "")

	return determineEncoding
}
