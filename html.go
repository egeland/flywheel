package main

/*
This could probably be smarter, but it'll do for now.
*/

const HTML_STOPPED = `
	<html>
		<body style="color: #333333; background: #f5f5f5;">
			<h1 style="text-align: center; margin-top: 50px; font-size: larger;">Your service is currently powered down</h1>
			<p style="text-align: center;"><a href="%s">Click here</a> to start.</p>
		</body>
	</html>`

const HTML_STARTING = `
	<html>
		<script>
			setTimeout(function() {
				window.location.reload(1);
			}, 5000);
		</script>
		<body style="color: #333333; background: #f5f5f5">
			<h1 style="text-align: center; margin-top: 50px; font-size: larger;">Your service is starting, please wait.</h1>
			<p style="text-align: center;">Your site will be loaded once startup is complete.</p>
		</body>
	</html>`

const HTML_STOPPING = `
	<html>
		<script>
			setTimeout(function() {
				window.location.reload(1);
			}, 5000);
		</script>
		<body style="color: #333333; background: #f5f5f5">
			<h1 style="text-align: center; margin-top: 50px; font-size: larger;">Your service is being powered down.</h1>
			<p style="text-align: center;">Please wait for shutdown to complete before restarting.</p>
		</body>
	</html>`

const HTML_UNHEALTHY = `
	<html>
		<body style="color: #333333; background: #f5f5f5">
			<h1 style="text-align: center; margin-top: 50px; font-size: larger;">Your service appears to be in an unhealthy or inconsistent state</h1>
			<p style="text-align: center;">This may be a temporary error, or may require manual intervention.</p>
		</body>
	</html>`

const HTML_ERROR = `
	<html>
		<body style="color: #333333; background: #f5f5f5">
			<h1 style="text-align: center; margin-top: 50px; font-size: larger;">An error occured processing your request</h1>
			<p style="text-align: center;">%v</p>
		</body>
	</html>`
