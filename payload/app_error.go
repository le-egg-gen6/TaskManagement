package payload

type AppError struct {
	StatusCode int    `json:"statusCode"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"errorKey"`
}
