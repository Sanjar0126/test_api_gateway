package models

type Location struct {
	Long float64 `json:"long" example:"60.123"`
	Lat  float64 `json:"lat" example:"40.123"`
}

type CourierModel struct {
	Phone         string `json:"phone"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	VehicleNumber string `json:"vehicle_number"`
}

type productDemandModel struct {
	Name              string  `json:"name" example:"Choyxona Osh"`
	Quantity          float64 `json:"quantity" example:"2"`
	Price             float64 `json:"price" example:"25000"`
	ExternalProductId int64   `json:"external_product_id,omitempty,string" example:"1234"`
	Description       string  `json:"description"`
}

type stepDemandModel struct {
	BranchName         string   `json:"branch_name"`
	BranchID           string   `json:"branch_id"`
	PhoneNumber        string   `json:"phone_number"`
	Address            string   `json:"address"`
	DestinationAddress string   `json:"destination_address"`
	Location           Location `json:"location"`
	ExternalStepID     int64    `json:"external_step_id,omitempty,string"`
	Description        string   `json:"description"`
}

type productOnDemandModel struct {
	ProductID   string  `json:"product_id" example:"a010f178-da52-4373-aacd-e477d871e27a"`
	Quantity    float64 `json:"quantity" example:"5"`
	Price       int64   `json:"price,string" example:"50000"`
	Description string  `json:"description"`
}

type stepOnDemandModel struct {
	BranchID    string `json:"branch_id"`
	Description string `json:"description"`
}

type orderDemandModel struct {
	ToLocation        Location `json:"to_location"`
	ToAddress         string   `json:"to_address" example:"Hamid Olimjon maydoni 10A dom 40-kvartira"`
	ClientName        string   `json:"client_name" example:"Oybek"`
	ClientPhoneNumber string   `json:"client_phone_number" example:"998998765432"`
	CoDeliveryPrice   float64  `json:"co_delivery_price" example:"10000"`
	Description       string   `json:"description"`
	ExternalOrderId   string   `json:"external_order_id"`
	DeliveryTime      string   `json:"delivery_time" example:"12:00"`
	DeliveryType      string   `json:"delivery_type" example:"delivery"`
}

type orderOnDemandModel struct {
	ToLocation       Location `json:"to_location"`
	ToAddress        string   `json:"to_address" example:"Hamid Olimjon maydoni 10A dom 40-kvartira"`
	ClientID         string   `json:"client_id"`
	CoDeliveryPrice  float64  `json:"co_delivery_price" example:"10000"`
	Description      string   `json:"description"`
	PaymentType      string   `json:"payment_type"`
	Source           string   `json:"source"`
	Apartment        string   `json:"apartment"`
	Building         string   `json:"building"`
	Floor            string   `json:"floor"`
	ExtraPhoneNumber string   `json:"extra_phone_number"`
	DeliveryTime     string   `json:"delivery_time"`
	DeliveryType     string   `json:"delivery_type" example:"delivery"`
	Paid             bool     `json:"paid"`
	StatusID         string   `json:"status_id"`
}

type CreateOnDemandOrderModel struct {
	orderOnDemandModel
	Steps []struct {
		stepOnDemandModel
		Products []productOnDemandModel `json:"products"`
	} `json:"steps"`
}

type CreateOnDemandOrderModelResponse struct {
	OrderID         string `json:"order_id"`
	ExternalOrderID string `json:"external_order_id"`
	PaymentLink     string `json:"payment_link"`
}

type StatusNote struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	StatusID    string `json:"status_id"`
	CreatedAt   string `json:"created_at"`
}

type GetOrderModel struct {
	orderDemandModel
	ID               string       `json:"id"`
	ClientID         string       `json:"client_id"`
	CourierID        string       `json:"courier_id"`
	Courier          CourierModel `json:"courier,omitempty"`
	StatusID         string       `json:"status_id"`
	CreatedAt        string       `json:"created_at"`
	FinishedAt       string       `json:"finished_at"`
	PaymentType      string       `json:"payment_type"`
	OrderAmount      uint64       `json:"order_amount,string"`
	Source           string       `json:"source"`
	Apartment        string       `json:"apartment"`
	Building         string       `json:"building"`
	Floor            string       `json:"floor"`
	ExtraPhoneNumber string       `json:"extra_phone_number"`
	Paid             bool         `json:"paid"`
	Rating           string       `json:"rating"`
	Review           string       `json:"review"`
	Steps            []struct {
		stepDemandModel
		ID         string `json:"id"`
		BranchID   string `json:"branch_id"`
		Status     string `json:"status"`
		StepAmount uint64 `json:"step_amount,string"`
		Products   []struct {
			productDemandModel
			ID          string `json:"id"`
			ProductID   string `json:"product_id"`
			TotalAmount uint64 `json:"total_amount,string"`
		} `json:"products"`
	} `json:"steps"`
	StatusNotes []StatusNote `json:"status_notes"`
	ShipperUser []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"shipper_user"`
}

type GetAllOrderModel struct {
	Orders []struct {
		orderDemandModel
		ID               string            `json:"id"`
		ClientID         string            `json:"client_id"`
		CourierID        string            `json:"courier_id"`
		Courier          CourierModel      `json:"courier,omitempty"`
		StatusID         string            `json:"status_id"`
		CreatedAt        string            `json:"created_at"`
		FinishedAt       string            `json:"finished_at"`
		PaymentType      string            `json:"payment_type"`
		Source           string            `json:"source"`
		Apartment        string            `json:"apartment"`
		Building         string            `json:"building"`
		Floor            string            `json:"floor"`
		ExtraPhoneNumber string            `json:"extra_phone_number"`
		OrderAmount      int64             `json:"order_amount,omitempty"`
		Paid             bool              `json:"paid"`
		Rating           string            `json:"rating"`
		Review           string            `json:"review"`
		Steps            []stepDemandModel `json:"steps"`
		StatusNotes      []StatusNote      `json:"status_notes"`
	} `json:"orders"`
	Count int64 `json:"count,string"`
}

type step struct {
	BranchName         string   `json:"branch_name"`
	Location           Location `json:"location"`
	Address            string   `json:"address"`
	DestinationAddress string   `json:"destination_address"`
	PhoneNumber        string   `json:"phone_number"`
	Description        string   `json:"description"`
}

type stepModel struct {
	step
	Products []productOnDemandModel `json:"products"`
}

type getOrderProductModel struct {
	productOnDemandModel
	ID          string  `json:"id" `
	TotalAmount float64 `json:"total_amount"`
}

type getOrderModel struct {
	productOnDemandModel
	ID       string `json:"id" example:"701dc270-0adc-4d00-ae78-4f2f78d794cc"`
	StatusID string `json:"status_id" example:"52f248b4-23a0-4350-80b7-1704eaff6c8c"`
}

type GetOrder struct {
	getOrderModel
	Products []getOrderProductModel `json:"products"`
}

type GetOrders struct {
	Orders []getOrderModel `json:"orders"`
	Count  uint64          `json:"count"`
}

type ChangeStatusRequest struct {
	StatusID    string `json:"status_id"`
	Description string `json:"description"`
}

type Status struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetStatuses struct {
	Statuses []Status `json:"statuses"`
}

type GetAllBranchOrdersModel struct {
	BranchID string `json:"branch_id"`
	Quantity uint64 `json:"quantity"`
}

type ConfirmFinishOrderModel struct {
	Code string `json:"code"`
}

type CourierPickedUpRequest struct {
	OrderID string `json:"order_id"`
}

type CustomerAddress struct {
	Location         Location `json:"location"`
	Address          string   `json:"address"`
	Apartment        string   `json:"apartment"`
	Building         string   `json:"building"`
	Floor            string   `json:"floor"`
	ExtraPhoneNumber string   `json:"extra_phone_number"`
}

type CustomerAddressesModel struct {
	Addresses []CustomerAddress `json:"addresses"`
}

type OrderReview struct {
	Rating uint64 `json:"rating"`
	Review string `json:"review"`
}

type UpdateOrderReview struct {
	UserID string `json:"user_id"`
	Seen   bool   `json:"seen"`
}
