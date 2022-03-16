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

type ListBidsByProductAndStatusRequest struct {
	ProductID int
	Page      int
	Limit     int
	Status    int
}

type BidProductRequest struct {
	UserID    int
	ProductID int
	Amount    int
}
