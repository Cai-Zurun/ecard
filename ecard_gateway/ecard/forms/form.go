package forms

type AddFamilyNumberRequest struct {
	Imei	string	`json:"imei"`
	Name1	string	`json:"name1"`
	Fn1	    string	`json:"fn1"`
	Name2	string	`json:"name2"`
	Fn2	    string	`json:"fn2"`
	Name3	string	`json:"name3"`
	Fn3	    string	`json:"fn3"`
}

type AddSOSNumberRequest struct {
	Imei	string	`json:"imei"`
	Sos1	string	`json:"sos1"`
	Sos2	string	`json:"sos2"`
	Sos3	string	`json:"sos3"`
}

type SetGPSTimerRequest struct {
	Imei	        string	`json:"imei"`
	TimeInterval	int	    `json:"timeInterval,string"`
}

type SetTiming0Request struct {
	Imei	string	`json:"imei"`
}

type Response struct {
	StatusCode	int	        `json:"statusCode"`
	Message	    interface{}	`json:"message"`
	Data	    interface{}	`json:"data"`
}
