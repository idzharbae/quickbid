package requests

type ListUserBiddedProductsRequest struct {
	UserID int
	Page   int
	Limit  int
}

type ListBidsByProductRequest struct {
	ProductID int
	Page      int
	Limit     int
}
