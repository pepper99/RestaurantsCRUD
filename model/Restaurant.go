package model

type Restaurant struct {
	Id                string         `json:"id,omitempty"  bson:"_id,omitempty"`
	Name              string         `json:"name,omitempty"  bson:"name,omitempty"`
	Address           string         `json:"address,omitempty"  bson:"address,omitempty"`
	RestaurantPicture []string       `json:"restaurant_picture,omitempty"  bson:"restaurant_picture,omitempty"`
	RecommendedDish   []string       `json:"recommended_dish,omitempty"  bson:"recommended_dish,omitempty"`
	Tag               []string       `json:"tag,omitempty"  bson:"tag,omitempty"`
	Coordinate        Coordinate     `json:"coordinate,omitempty"  bson:"coordinate,omitempty,inline"`
	Rating            float32        `json:"rating,omitempty"  bson:"rating,omitempty"`
	DeliveryInfo      []DeliveryInfo `json:"delivery_info,omitempty"  bson:"delivery_info,omitempty"`
}

type Coordinate struct {
	Lat float64 `json:"lat,omitempty"  bson:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"  bson:"lng,omitempty"`
}

type DeliveryInfo struct {
	Platform string `json:"platform,omitempty"  bson:"platform,omitempty"`
	Url      string `json:"url,omitempty"  bson:"url,omitempty"`
}
