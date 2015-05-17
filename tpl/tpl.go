package tpl

import (
	"bytes"
)

func Index() []byte {
	_buffer := new(bytes.Buffer)
	_buffer.WriteString("<html lang=\"zh-CN\">")
	head(_buffer, "")
	_buffer.WriteString("<body><div id=\"sidebar\"><div id=\"logo\"></div><div id=\"sidebar-content\"></div></div><div id=\"content\"><div id=\"messages\"></div><textarea id=\"input-box\"></textarea><button id=\"submit\" class=\"button\">发送</button></div><div id=\"toolbar\"></div><script type=\"text/javascript\" charset=\"utf-8\" src=\"./static/js/main.js\"></script></body></html>")
	return _buffer.Bytes()
}

func head(_buffer *bytes.Buffer, title string) {
	_buffer.WriteString("<head><meta charset=\"utf-8\"><title>")
	_buffer.WriteString(title)
	_buffer.WriteString("</title><link href=\"./static/css/main.css\" media=\"all\" rel=\"stylesheet\"></head>")
}
