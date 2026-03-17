package handler

type IService interface {
	IServiceGrpc
	IServiceRest
}

type AuthHandler struct {
	*HandlerRest
	*HandlerGrpc
}

func NewHandler(service IService) *AuthHandler {
	return &AuthHandler{
		HandlerRest: NewHandlerRest(service),
		HandlerGrpc: NewHandlerGrpc(service),
	}
}
