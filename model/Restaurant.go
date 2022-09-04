package model

type Restaurant struct {
	Id                string         `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string         `json:"name,omitempty"  bson:"name"`
	Address           string         `json:"address,omitempty"  bson:"address"`
	RestaurantPicture []string       `json:"restaurant_picture,omitempty"  bson:"restaurant_picture"`
	RecommendedDish   []string       `json:"recommended_dish,omitempty"  bson:"recommended_dish"`
	Tag               []string       `json:"tag,omitempty"  bson:"tag"`
	Coordinate        Coordinate     `json:"coordinate,omitempty"  bson:"coordinate"`
	Rating            float32        `json:"rating,omitempty"  bson:"rating"`
	DeliveryInfo      []DeliveryInfo `json:"delivery_info,omitempty"  bson:"delivery_info"`
}

type Coordinate struct {
	Lat float64 `json:"lat,omitempty"  bson:"lat"`
	Lng float64 `json:"lng,omitempty"  bson:"lng"`
}

type DeliveryInfo struct {
	Platform string `json:"platform,omitempty"  bson:"platform"`
	Url      string `json:"url,omitempty"  bson:"url"`
}
