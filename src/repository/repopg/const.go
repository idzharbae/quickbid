package repopg

const (
	InsertProductQuery = "INSERT INTO products(name, initial_price, start_bid_date, end_bid_date, owner_user_id, last_bid_id, image_url, status, bid_increment) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	GetProductQuery = `SELECT 
		COALESCE(id, 0) as id,
		COALESCE(name, '') as name,
		COALESCE(initial_price, 0) as initial_price,
		COALESCE(start_bid_date, '0001-01-01T00:00:00Z'::timestamp) as start_bid_date,
		COALESCE(end_bid_date, '0001-01-01T00:00:00Z'::timestamp) as end_bid_date,
		COALESCE(owner_user_id, 0) as owner_user_id,
		COALESCE(last_bid_id, 0) as last_bid_id,
		COALESCE(image_url, '') as image_url,
		COALESCE(status, 0) as status,
		COALESCE(bid_increment, 0) as bid_increment
	FROM products`
	GetProductByOwnerUserIDQuery = GetProductQuery + " WHERE owner_user_id = $1"
)
