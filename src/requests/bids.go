package requests

type ListUserBiddedProductsRequest struct {
	UserID int
	Page   int
	Limit  int
}
