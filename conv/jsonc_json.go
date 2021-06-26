package conv

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

func analyseJSONCstatement(src []byte, start int, buf *bytes.Buffer) (multiLineComment bool, err error) {
	l := len(src)

	var quote bool    // 在引号内还是引号外
	var transfer bool // 是否在转义状态

	// 逐字判断
	for i := start; i < l; i++ {
		if transfer { // 如果转义了，忽略后面一个char
			transfer = false
			buf.WriteByte(src[i])
			continue
		}

		if src[i] == ' ' || src[i] == '\t' {
			// 如果是空则不移动 lastIndex
			continue
		} else if src[i] == '"' {
			// 标记是 opening quote 还是 closing quote.
			if quote {
				quote = false
			} else {
				quote = true
			}
			buf.WriteByte(src[i])
		} else if src[i] == '\\' {
			// 如果是转义符，则标记转义
			if !quote {
				// 如果转义符在引号外面，ERROR
				return false, errors.New("error: '\\' out side quote")
			}
			buf.WriteByte(src[i])
			transfer = true
		} else if src[i] == '/' {
			if quote {
				// 如果 ‘/’ 在引号内，不需要特殊处理
				buf.WriteByte(src[i])
				continue
			}

			// 如果 '/' 在引号外面，判断后一位是否也是 '/'，说明后面是 comments.
			if i+1 < l && src[i+1] == '/' {
				break // 结束循环
			} else if i+1 < l && src[i+1] == '*' { // /* */ 多行注释的情况
				// NOTE /* */ 多行注释问题
				ci := bytes.Index(src[i+2:], []byte("*/")) // 查看该 line 有没有 */
				if ci == -1 {
					multiLineComment = true
					break // 结束循环
				} else {
					i = i + 2 + ci + 1 // NOTE 跳过检查
					continue
				}
			}

			// 如果 ‘/’ 在引号外面而且后面不是 ‘/’ ，ERROR
			return false, fmt.Errorf("error: '/' out side quote %s", string(src))
		} else {
			// 其他正常情况下直接向后处理。
			buf.WriteByte(src[i])
		}
	}

	// 如果 line 结束，单引号没有关闭则，Error
	if quote {
		return false, errors.New("error: statement is Unquoted")
	}

	return multiLineComment, nil
}

// NOTE JSONC must be formatted, otherwise cannot be read.
// 传入的 Jsonc 本身必须是合法的，否则也无法被 json 读取。
func JSONCToJSON(jsonc []byte) ([]byte, error) {
	lines := bytes.Split(jsonc, []byte("\n"))

	var (
		multiComment bool
		er           error

		buf bytes.Buffer
	)
	for _, line := range lines {
		start := 0
		if multiComment {
			ci := bytes.Index(line, []byte("*/"))
			if ci == -1 {
				continue
			} else {
				start = ci + 2
			}
		}

		multiComment, er = analyseJSONCstatement(line, start, &buf)
		if er != nil {
			return nil, er
		}
	}

	if !json.Valid(buf.Bytes()) {
		return nil, errors.New("not a legal json format")
	}

	return buf.Bytes(), nil
}
