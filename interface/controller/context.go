package controller

//Context custom interface
type Context interface {
	JSON(code int, i interface{}) error
}
