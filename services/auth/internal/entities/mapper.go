package entities

import "github.com/isOdin-l/TinderArt/services/auth/pkg/api"

// Function design:
// From<RequestModel>To<EntityName>  (req *grpc_gen.<RequestModel>) *<EntityName>
// From<EntityName>To<ResponseModel> (req *<EntityName>) *grpc_gen.<ResponseModel>

// --------- Request ---------
func FromRequestToEntity(req *api.Request) *Entity {
	return &Entity{}
}

// --------- Response ---------
func FromEntityToResponse(req *Entity) *api.Response {
	return &api.Response{}
}
