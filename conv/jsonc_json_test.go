package conv

import (
	"testing"
)

const jsoncTest = `{
  "editor.formatOnSave": true, // save 的时候格式化文件
  "[go]": {
    "editor.defaultFormatter": "golang.go", // go 不使用 prettier 格式化代码
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
    // Optional: Disable snippets, as they conflict with completion ranking.
    // "editor.snippetSuggestions": "none"
  },
  "[go.mod]": {
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },

  // gopls 设置, https://github.com/golang/tools/tree/master/gopls
  "go.useLanguageServer": true, // 使用 gopls
  "gopls": {
    // Add parameter placeholders when completing a function.
    "usePlaceholders": true,

    // If true, enable additional analyses with staticcheck.
    // Warning: This will significantly increase memory usage.
    "staticcheck": false
  },

  // DEBUG 神器，可以通过访问 http://localhost:16060 查看性能参数
  "go.languageServerFlags": [
    "-rpc.trace", // for more detailed debug logging
    "serve",
    "--debug=localhost:16060" // to investigate memory usage, see profiles
  ],

  // gotests 打印详请
  "go.testFlags": ["-v"],

  // gomodifytags json 配置, 改之前看文档
  // https://github.com/golang/vscode-go/blob/HEAD/docs/settings.md
  "go.addTags": {
    "tags": "json",
    "options": "json=omitempty",
    "promptForTags": true, // user tags
    "template": "",
    "transform": "snakecase"
  },

  // golangci-lint 设置
  "go.lintTool": "golangci-lint",

  // NOTE save 时 golangci-lint 整个 package，使用 'file' 时，
  // 如果变量定义在别的文件中会造成 undeclared 错误。
  "go.lintOnSave": "package",

  "go.lintFlags": [
    "--fast", // without --fast can freeze your editor.

    // golangci-lint 配置文件地址
    // "--config=${workspaceRoot}/golangci.yml" // 本地
    // "--config=~/.golangci/prod-ci.yml"
    "--config=~/.golangci/dev-ci.yml"
  ],

  // search.exclude 用来忽略搜索的文件夹
  // files.exclude 用来忽略工程打开的文件夹
  // 直接写文件/文件夹名字就实在项目根路径下进行匹配，不要用 / ./ 开头，
  // **/所有路径下进行匹配
  "search.exclude": {
    ".idea": true,
    // "**/pkg": true,
    "*.iml": true,
    "**/vendor": true,
    ".history": true
  },

  // files.exclude 不显示文件，
  // 直接写文件/文件夹名字就实在项目根路径下进行匹配，不要用 / ./ 开头，
  // **/所有路径下进行匹配
  "files.exclude": {
    ".idea": true
    // "**/pkg": true,
  },
  "a":1,
  "b":"ok",
  "c":true
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
