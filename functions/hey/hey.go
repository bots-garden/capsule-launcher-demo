package main

// TinyGo wasm module
import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
	/* string to json */
	"github.com/tidwall/gjson"
	/* create json string */
	"github.com/tidwall/sjson"
)

// main is required.
func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(req hf.Request) (resp hf.Response, errResp error) {

	// display the body request
	hf.Log("📝 body: " + req.Body)

	// get the data of the body request
	author := gjson.Get(req.Body, "author")
	message := gjson.Get(req.Body, "message")

	hf.Log("🟢 Content-Type: " + req.Headers["Content-Type"])
	hf.Log("🔵 Content-Length: " + req.Headers["Content-Length"])
	hf.Log("🟠 User-Agent: " + req.Headers["User-Agent"])
	hf.Log("🔴 My-Token: " + req.Headers["My-Token"])

	headers := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"YourMessage":  message.String(),
		"MyToken":      req.Headers["My-Token"],
	}

	jsondoc := `{"message": ""}`
	jsondoc, err := sjson.Set(jsondoc, "message", "👋 hey! "+author.String()+" What's up?")

	return hf.Response{Body: jsondoc, Headers: headers}, err
}
