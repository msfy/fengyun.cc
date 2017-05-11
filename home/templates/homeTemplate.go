package templates

import (
    "text/template"
    "github.com/msfy/fengyun.cc/home/utils"
)

var HomeTemplate *template.Template

const tplString = `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Title</title>
</head>
<body>
</body>
<script>
window.__PRELOADED_STATE__ = {{.}}
</script>
</html>`

func init() {
    tpl, err := template.New("HomeTemplate").Parse(tplString)
    utils.CheckError(err)
    HomeTemplate = tpl
}
