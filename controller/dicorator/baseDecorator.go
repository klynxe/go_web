package dicorator

var AllowCors = NewDecorator(AddHeaderFabric("Access-Control-Allow-Origin", "http://127.0.0.1:4200"))

var AllowOptionsCors = NewDecorator(AddHeaderFabric("Access-Control-Allow-Origin", "http://127.0.0.1:4200"),
	AddHeaderFabric("Access-Control-Allow-Methods", "POST, OPTIONS"),
	AddHeaderFabric("Access-Control-Allow-Headers", "access-control-allow-origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-CaptchaToken, Authorization"))
