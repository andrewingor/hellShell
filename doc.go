package main

//HTML form
var (
	Tail string = "</pre></body></html>"

	Head string = `
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

	Form string = `
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
)

//EOF
