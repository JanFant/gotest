package netsite

var (
	form = `<form action="/" method="POST">
			<label for="numbers">Numbers (comma or space-separated):</label><br /> 
			<input type="text" name="numbers" size="30"><br /> 
			<input type="submit" value="Calculate"> 
			</form>`
	buttonVchs   = `<button formaction="http://127.0.0.1:8080/vchs">vchs</button>`
	buttonVas    = `<button formaction="http://127.0.0.1:8080/vas">vas</button>`
	mainpageinfo = `<h3>Обработка данных ВЧС и ВАС</h3>
					<p>Выберите что хотите отображать</p>
					<form method="post" enctype="application/x-www-form-urlencoded">`
	pageTop = `<!DOCTYPE HTML><html><head>
			   <style>.error{color:#FF0000;}</style></head><title>VCHS AND VAS</title>
			   <body>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)
