<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>upload</title>
    <style type="text/css">
    </style>
</head>
<body>
<form enctype="multipart/form-data" name="fileinfo">
    <label>Your email address:</label>
    <input type="email" autocomplete="on" autofocus name="userid" placeholder="email" required size="32"
           maxlength="64"/><br/>
    <label>Custom file label:</label>
    <input type="text" name="filelabel" size="12" maxlength="32"/><br/>
    <label>File to stash:</label>
    <input type="file" name="uploadfile" required/>
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="Stash the file!"/>
</form>
<div></div>
</body>
<script>
    var form = document.forms.namedItem('fileinfo')
    form.addEventListener('submit', function (ev) {
        var oOutput = document.querySelector('div'),
                oData = new FormData(form)
        oData.append('CustomField', 'this is extra data')

        var oReq = new XMLHttpRequest()
        oReq.open('post', 'http://chuans.online:8089/upload', true)
        oReq.onload = function (oEvent) {
            if (oReq.status === 200) {
                oOutput.innerHTML = 'Uploaded!'
            } else {
                oOutput.innerHTML = 'Error' + oReq.status + ' occurred when trying to upload your file.<br \/>'
            }
        }
        oReq.send(oData)
        ev.preventDefault()
    })
</script>
</html>
