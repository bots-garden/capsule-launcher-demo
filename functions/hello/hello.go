package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

//export OnLoad
func OnLoad() {
	hf.Log("👋 hello from `OnLoad` function")
	hf.MemorySet("message", "I 💜 WebAssembly")
}

//export OnExit
func OnExit() {
	hf.Log("👋 hello from `OnExit` function")
	message, _ := hf.MemoryGet("message")
	hf.Log("The message is: " + message)
}

func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(req hf.Request) (resp hf.Response, errResp error) {
	message, _ := hf.GetEnv("MESSAGE")
	html := `
    <html>
			<head>
				<meta charset="utf-8">
				<title>Wasm is fantastic 😍</title>

				<meta name="viewport" content="width=device-width, initial-scale=1">
				
				<style>
					.container { min-height: 100vh; display: flex; justify-content: center; align-items: center; text-align: center; }
					.title { font-family: "Source Sans Pro", "Helvetica Neue", Arial, sans-serif; display: block; font-weight: 300; font-size: 100px; color: #35495e; letter-spacing: 1px; }
					.subtitle { font-family: "Source Sans Pro", "Helvetica Neue", Arial, sans-serif; font-weight: 300; font-size: 42px; color: #526488; word-spacing: 5px; padding-bottom: 15px; }
				</style>

			</head>

			<body>
				<section class="container">
					<div>
						<h1 class="title">` + message + `</h1>
						<h2 class="subtitle">Served with 💜 by Capsule 💊</h2>
					</div>
				</section>
			</body>

    </html>
    `

	headers := map[string]string{
		"Content-Type": "text/html; charset=utf-8",
	}

	return hf.Response{Body: html, Headers: headers}, nil
}
