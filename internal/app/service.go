package app

import "github.com/NicholasLiem/AssetFindr_BackendAssignment/internal/service"

type MicroserviceServer struct {
	postService service.PostService
}

func NewMicroservice(
	postService service.PostService,
) *MicroserviceServer {
	return &MicroserviceServer{
		postService: postService,
	}
}
