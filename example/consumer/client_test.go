package consumer

import "testing"

type stubServer struct {
	// contract test lib
}

func (s *stubServer) AddItem() {
	// return s.lib.handle("ShoppingCart.AddItem", req)
}

func TestClientContract(t *testing.T) {
	// ct lib reads the contracts and store map[rpcMethod]list(req, resp)

}
