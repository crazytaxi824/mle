package conv

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

func toggle(b *bool) {
	if *b {
		*b = false
	} else {
		*b = true
	}
}

// multiLineComment 说明是否在多行注释中 /* */
func lastValidChatInJSONCline(src []byte, start int) (lastValidCharIndex int, multiLineComment bool, err error) {
	l := len(src)

	// NOTE lastIndex = -1, 说明该行是空行。
	lastValidCharIndex = -1 // 最后一位有效 char 的 index。

	var quote bool    // 在引号内还是引号外
	var transfer bool // 是否在转义状态

	// 逐字判断
ForLoop:
	for i := start; i < l; i++ {
		if transfer { // 如果转义了，忽略后面一个char
			transfer = false
			lastValidCharIndex = i
			continue
		}

		switch src[i] {
		case ' ', '\t':
			// 如果是空则不移动 lastIndex
			break
		case '"':
			toggle(&quote)
			lastValidCharIndex = i // 移动 lastIndex
		case '\\':
			// 如果是转义符，则标记转义
			if !quote {
				// 如果转义符在引号外面，ERROR
				return 0, false, errors.New("error: '\\' out side quote")
			}
			transfer = true
		case '/':
			if quote {
				// 如果 ‘/’ 在引号内，不需要特殊处理
				lastValidCharIndex = i
				break
			}

			// 如果 '/' 在引号外面，判断后一位是否也是 '/'，说明后面是 comments.
			if i+1 < l && src[i+1] == '/' {
				break ForLoop // 结束循环 break for loop
			} else if i+1 < l && src[i+1] == '*' { // /* */ 多行注释的情况
				// NOTE /* */ 多行注释问题
				ci := bytes.Index(src[i+2:], []byte("*/")) // 查看该 line 有没有 */
				if ci == -1 {
					multiLineComment = true
					break ForLoop // 结束循环 break for loop
				} else {
					i = i + 2 + ci + 1 // NOTE 跳过检查
					lastValidCharIndex = i
					break
				}
			}

			// 如果 ‘/’ 在引号外面而且后面不是 ‘/’ ，ERROR
			return 0, false, fmt.Errorf("error: '/' out side quote %s", string(src))
		default:
			// 其他正常情况下直接向后处理。
			lastValidCharIndex = i
		}
	}

	// 如果 line 结束，单引号没有关闭则，Error
	if quote {
		return 0, false, errors.New("error: statement is Unquoted")
	}

	return lastValidCharIndex, multiLineComment, nil
}

// multiLineComment 说明是否在多行注释中 /* */
func jsoncLineTojson(src []byte, start int, buf *bytes.Buffer) (multiLineComment bool, err error) {
	l := len(src)

	var quote bool    // 在引号内还是引号外
	var transfer bool // 是否在转义状态

	// 逐字判断
ForLoop:
	for i := start; i < l; i++ {
		if transfer { // 如果转义了，后面一个char不做特殊处理，// TODO 判断是否合法
			transfer = false
			buf.WriteByte(src[i])
			continue
		}

		switch src[i] {
		case ' ', '\t':
			break // break switch
		case '"':
			toggle(&quote)
			buf.WriteByte(src[i])
		case '\\':
			// 如果是转义符，则标记转义
			if !quote {
				// 如果转义符在引号外面，ERROR
				return false, errors.New("format error: '\\' out side quote")
			}
			buf.WriteByte(src[i])
			transfer = true
		case '/':
			if quote {
				// 如果 ‘/’ 在引号内，不需要特殊处理
				buf.WriteByte(src[i])
				break // break switch
			}

			// 如果 '/' 在引号外面，判断后一位是否也是 '/'，说明后面是 comments.
			if i+1 < l && src[i+1] == '/' {
				break ForLoop // 结束循环 break for loop
			} else if i+1 < l && src[i+1] == '*' { // /* */ 多行注释的情况
				// NOTE /* */ 多行注释问题
				ci := bytes.Index(src[i+2:], []byte("*/")) // 查看该 line 有没有 */
				if ci == -1 {
					// 如果不存在 */
					multiLineComment = true // 标记多行注释
					break ForLoop           // 结束循环 break for loop
				} else {
					// 如果 */ 存在，直接移动读取位置到 */ 后面
					i = i + 2 + ci + 1 // NOTE 跳过检查
					break              // break switch
				}
			}

			// 如果 ‘/’ 在引号外面而且后面不是 ‘/’ ，ERROR
			return false, fmt.Errorf("format error: '/' out side quote %s", string(src))
		default:
			// 其他正常情况下直接向后处理。
			buf.WriteByte(src[i])
		}
	}

	// 如果 line 结束，单引号没有关闭则，Error
	if quote {
		return false, errors.New("format error: statement is Unquoted")
	}

	return multiLineComment, nil
}

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

		multiComment, er = jsoncLineTojson(line, start, &buf)
		if er != nil {
			return nil, er
		}
	}

	if !json.Valid(buf.Bytes()) {
		return nil, errors.New("not a legal json format")
	}

	return buf.Bytes(), nil
}

type JSONCstatment struct {
	LineIndex          int // 行号
	LastValidCharIndex int // 最后一个有效字符的 index，后面的 // comments 不算在内
}

// find second last line and lastCharIndex
func FindSecondLastLine(jsonc []byte) (secondLastLine, lastCharIndex int, err error) {
	lines := bytes.Split(jsonc, []byte("\n"))

	var (
		result       []JSONCstatment
		lastIndex    int
		multiComment bool
		er           error
	)

	for i, line := range lines {
		start := 0
		if multiComment {
			ci := bytes.Index(line, []byte("*/"))
			if ci == -1 {
				continue
			} else {
				start = ci + 2
			}
		}

		lastIndex, multiComment, er = lastValidChatInJSONCline(line, start)
		if er != nil {
			return 0, 0, er
		}

		// lastIndex == -1, 表示整行都是 comment, 或者是空行
		if lastIndex != -1 {
			result = append(result, JSONCstatment{i, lastIndex})
		}
	}

	l := len(result)
	var r JSONCstatment
	if l > 1 {
		r = result[l-2]
	}

	return r.LineIndex, r.LastValidCharIndex, nil
}
