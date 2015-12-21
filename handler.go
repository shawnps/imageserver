package imageserver

// Handler handles an Image and returns an Image.
type Handler interface {
	Handle(*Image, Params) (*Image, error)
}

// HandlerFunc is a Handler func.
type HandlerFunc func(*Image, Params) (*Image, error)

// Handle implements Handler.
func (f HandlerFunc) Handle(im *Image, params Params) (*Image, error) {
	return f(im, params)
}

// IdentityHandler is a Handler that returns the same Image.
type IdentityHandler struct{}

// Handle implements Handler.
func (hdr IdentityHandler) Handle(im *Image, params Params) (*Image, error) {
	return im, nil
}

// HandlerServer is a Handler Server.
type HandlerServer struct {
	Server
	Handler
}

// Get implements Server
func (srv *HandlerServer) Get(params Params) (*Image, error) {
	im, err := srv.Server.Get(params)
	if err != nil {
		return nil, err
	}
	im, err = srv.Handle(im, params)
	if err != nil {
		return nil, err
	}
	return im, nil
}