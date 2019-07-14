package main

import (
 	pb "github.com/Joseph329123/microservices-demo/src/testservice/genproto"
 )

const (
	USER_ID = "dummyID"
	ORDER_ID = "dummyOrderID"
	SHIPPING_TRACKING_ID = "dummyShippingTrackingID"
	EMAIL = "someone@example.com"

	EUR = "EUR"
	USD = "USD"
	JPY = "JPY"
	GBP = "GBP"
	TRY = "TRY"
	CAD = "CAD"

	VINTAGE_TYPEWRITER = "Vintage Typewriter"
	VINTAGE_TYPEWRITER_ID = "OLJCESPC7Z"

	VINTAGE_CAMERA_LENS = "Vintage Camera Lens"
	VINTAGE_CAMERA_LENS_ID = "66VCHSJNUP"

	HOME_BARISTA_KIT = "Home Barista Kit"
	HOME_BARISTA_KIT_ID = "1YMWWN1N4O"

	TERRARIUM = "Terrarium"
	TERRARIUM_ID = "L9ECAV7KIM"

	FILM_CAMERA = "Film Camera"
	FILM_CAMERA_ID = "2ZYFJ3GM2N"

	VINTAGE_RECORD_PLAYER = "Vintage Record Player"
	VINTAGE_RECORD_PLAYER_ID = "0PUK6V6EV0"

	METAL_CAMPING_MUG = "Metal Camping Mug"
	METAL_CAMPING_MUG_ID = "LS4PSXUNUM"

	CITY_BIKE = "City Bike"
	CITY_BIKE_ID = "9SIQT8TOJO"

	AIR_PLANT = "Air Plant"
	AIR_PLANT_ID = "6E92ZMYYFZ"
)

var (
	CTX_KEYS = []string{"camera"}
)

var (
	VINTAGE_TYPEWRITER_CART_ITEM = &pb.CartItem{ProductId: VINTAGE_TYPEWRITER_ID, Quantity: 1}
	
	VINTAGE_CAMERA_LENS_CART_ITEM = &pb.CartItem{ProductId: VINTAGE_CAMERA_LENS_ID, Quantity: 1}
	
	HOME_BARISTA_KIT_CART_ITEM = &pb.CartItem{ProductId: HOME_BARISTA_KIT_ID, Quantity: 1}
	
	TERRARIUM_CART_ITEM = &pb.CartItem{ProductId: TERRARIUM_ID, Quantity: 1}
	
	FILM_CAMERA_CART_ITEM = &pb.CartItem{ProductId: FILM_CAMERA_ID, Quantity: 1}
	
	VINTAGE_RECORD_PLAYER_CART_ITEM = &pb.CartItem{ProductId: VINTAGE_RECORD_PLAYER_ID, Quantity: 1}
	
	METAL_CAMPING_MUG_CART_ITEM = &pb.CartItem{ProductId: METAL_CAMPING_MUG_ID, Quantity: 1}
	
	CITY_BIKE_CART_ITEM = &pb.CartItem{ProductId: CITY_BIKE_ID, Quantity: 1}
	
	AIR_PLANT_CART_ITEM = &pb.CartItem{ProductId: AIR_PLANT_ID, Quantity: 1}
)

var (
	PRODUCT_IDS = []string{VINTAGE_TYPEWRITER_ID, VINTAGE_CAMERA_LENS_ID, HOME_BARISTA_KIT_ID}
)

var (
	ITEMS = []*pb.CartItem{VINTAGE_TYPEWRITER_CART_ITEM, VINTAGE_CAMERA_LENS_CART_ITEM}
)

var (
	ADDRESS = &pb.Address{
					StreetAddress: "1600 Amphitheatre Parkway",
					City:          "Mountain View",
					State:         "CA",
					ZipCode:       94043,
					Country:       "Mountain View"}
)

var (
	MONEY = &pb.Money{
				CurrencyCode: EUR,
				Units: 1,
				Nanos: 0}

	MONEY_VINTAGE_TYPEWRITER = &pb.Money{
									CurrencyCode: USD,
									Units: 67,
									Nanos: 980000000}

	MONEY_VINTAGE_CAMERA_LENS = &pb.Money{
								 	CurrencyCode: USD,
									Units: 12,
									Nanos: 480000000}

	MONEY_HOME_BARISTA_KIT = &pb.Money{
								CurrencyCode: USD,
								Units: 123,
								Nanos: 990000000}

	MONEY_TERRARIUM = &pb.Money{
							CurrencyCode: USD,
							Units: 36,
							Nanos: 440000000}

	MONEY_FILM_CAMERA = &pb.Money{
							CurrencyCode: USD,
							Units: 2244,
							Nanos: 990000000}

	MONEY_VINTAGE_RECORD_PLAYER = &pb.Money{
										CurrencyCode: USD,
										Units: 65,
										Nanos: 500000000}

	MONEY_METAL_CAMPING_MUG = &pb.Money{
									CurrencyCode: USD,
									Units: 24,
									Nanos: 330000000}

	MONEY_CITY_BIKE = &pb.Money{
							CurrencyCode: USD,
							Units: 789,
							Nanos: 500000000}

	MONEY_AIR_PLANT = &pb.Money{
							CurrencyCode: USD,
							Units: 12,
							Nanos: 290000000}
)

var (
	CREDITCARDINFO = &pb.CreditCardInfo{
						CreditCardNumber:          "4432801561520454",
						CreditCardExpirationMonth: 1,
						CreditCardExpirationYear:  2020,
						CreditCardCvv:             672}
)

var (
	ORDER_ITEM_VINTAGE_TYPEWRITER = &pb.OrderItem{
										Item: VINTAGE_TYPEWRITER_CART_ITEM, 
										Cost: MONEY_VINTAGE_TYPEWRITER}

	ORDER_ITEM_VINTAGE_CAMERA_LENS = &pb.OrderItem{
										Item: VINTAGE_CAMERA_LENS_CART_ITEM, 
										Cost: MONEY_VINTAGE_CAMERA_LENS}

	ORDER_ITEM_HOME_BARISTA_KIT = &pb.OrderItem{
										Item: HOME_BARISTA_KIT_CART_ITEM, 
										Cost: MONEY_HOME_BARISTA_KIT}

	ORDER_ITEM_TERRARIUM = &pb.OrderItem{
								Item: TERRARIUM_CART_ITEM, 
								Cost: MONEY_TERRARIUM}

	ORDER_ITEM_FILM_CAMERA = &pb.OrderItem{
									Item: FILM_CAMERA_CART_ITEM, 
									Cost: MONEY_FILM_CAMERA}

	ORDER_ITEM_VINTAGE_RECORD_PLAYER = &pb.OrderItem{
											Item: VINTAGE_RECORD_PLAYER_CART_ITEM, 
											Cost: MONEY_VINTAGE_RECORD_PLAYER}
	
	ORDER_ITEM_METAL_CAMPING_MUG = &pb.OrderItem{
										Item: METAL_CAMPING_MUG_CART_ITEM, 
										Cost: MONEY_METAL_CAMPING_MUG}

	ORDER_ITEM_CITY_BIKE = &pb.OrderItem{
								Item: CITY_BIKE_CART_ITEM, 
								Cost: MONEY_CITY_BIKE}

	ORDER_ITEM_AIR_PLANT = &pb.OrderItem{
								Item: AIR_PLANT_CART_ITEM, 
								Cost: MONEY_AIR_PLANT}
)

var (
	ORDER_ITEMS = []*pb.OrderItem{ORDER_ITEM_VINTAGE_TYPEWRITER, ORDER_ITEM_VINTAGE_CAMERA_LENS}
)

var (
	ORDER_RESULT = &pb.OrderResult {
						OrderId: 			ORDER_ID,
						ShippingTrackingId: SHIPPING_TRACKING_ID,
						ShippingCost: 		MONEY, 
						ShippingAddress:    ADDRESS,
						Items: 				ORDER_ITEMS}
)

