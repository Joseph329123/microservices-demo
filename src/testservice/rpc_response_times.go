package main

import (
	"context"
	"time"
	"fmt"

	pb "github.com/Joseph329123/microservices-demo/src/testservice/genproto"

	"github.com/pkg/errors"
)

func runResponseTimeTests(ctx context.Context, svc *frontendServer, samples int) {
	fmt.Println("runResponseTimeTests()")
	fmt.Println("# of samples", samples)

	data := make([][]float64, NUMBER_OF_REQUESTS)

	for i := 0; i < NUMBER_OF_REQUESTS; i++ {
		data[i] = make([]float64, samples)
	}

	for i := 0; i < samples; i++ {
		fmt.Println(float64(i)/float64(samples))

		/* ProductCatalogueService */
		data[PRODUCT_CATALOGUE_EMPTY_REQUEST_INDEX][i], _, _ = svc.timeProductCatalogueServiceEmptyRequest(ctx)

		data[PRODUCT_CATALOGUE_GET_PRODUCT_REQUEST_INDEX][i], _, _ = svc.timeProductCatalogueServiceGetProductRequest(ctx, VINTAGE_TYPEWRITER_ID)
		
		data[PRODUCT_CATALOGUE_SEARCH_PRODUCT_REQUEST_INDEX][i], _, _ = svc.timeProductCatalogueServiceSearchProductsRequest(ctx, VINTAGE_TYPEWRITER)

		/* RecommendationService */
		data[RECOMMENDATION_LIST_RECOMMENDATIONS_REQUEST_INDEX][i], _, _ = svc.timeRecommendationServiceListRecommendationsRequest(ctx, USER_ID, PRODUCT_IDS)

		/* CheckoutService */
		data[CHECKOUT_PLACE_ORDER_REQUEST_INDEX][i], _, _ = svc.timeCheckoutServicePlaceOrderRequest(ctx)

		/* ShippingService */ 
		data[SHIPPING_GET_QUOTE_REQUEST_INDEX][i], _, _ = svc.timeShippingServiceGetQuoteRequest(ctx, ITEMS, USD)

		data[SHIPPING_SHIP_ORDER_REQUEST_INDEX][i], _, _ = svc.timeShippingServiceShipOrderRequest(ctx, ADDRESS, ITEMS)
		
		/* CurrencyService */
		data[CURRENCY_CURRENCY_CONVERSION_REQUEST_INDEX][i], _, _ = svc.timeCurrencyServiceCurrencyConversionRequest(ctx, MONEY, USD)
		
		data[CURRENCY_EMPTY_REQUEST_INDEX][i], _, _ = svc.timeCurrencyServiceEmptyRequest(ctx)

		/* CartService */
		data[CART_ADD_ITEM_REQUEST_INDEX][i], _, _ = svc.timeCartServiceAddItemRequest(ctx, USER_ID, VINTAGE_TYPEWRITER_ID, 1)
	
		data[CART_GET_CART_REQUEST_INDEX][i], _, _ = svc.timeCartServiceGetCartRequest(ctx, USER_ID)
		
		data[CART_EMPTY_CART_REQUEST_INDEX][i], _, _ = svc.timeCartServiceEmptyCartRequest(ctx, USER_ID)
		
		/* AdService */
		data[AD_AD_REQUEST_INDEX][i], _, _ = svc.timeAdServiceAdRequest(ctx, CTX_KEYS)

		/* PaymentService */
		data[PAYMENT_CHARGE_REQUEST_INDEX][i], _, _ = svc.timePaymentServiceChargeRequest(ctx, MONEY, CREDITCARDINFO)

		/* EmailService */
		data[EMAIL_SEND_ORDER_CONFIRMATION_REQUEST_INDEX][i], _, _ = svc.timeEmailServiceSendOrderConfirmationRequest(ctx, EMAIL, ORDER_RESULT)
	}

	output(data)
}

func output(data[][] float64) {
	for i := 0; i < NUMBER_OF_REQUESTS; i++ {
		fmt.Println(find_min(data[i]),",", find_max(data[i]),",", mean(data[i]),",", std_dev(data[i]))
	}
}


/********************************************************************/
/* ProductCatalogueService */
/********************************************************************/

/* Send 'Empty' to ProductCatalogueService, receive 'ListProductsResponse' */
func (fe *frontendServer) timeProductCatalogueServiceEmptyRequest(ctx context.Context) (float64, []*pb.Product, error) {
	//fmt.Println("timeProductCatalogueServiceEmptyRequest()")

	productCatalogServiceClient := pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn)

	start := time.Now()
	resp, err := productCatalogServiceClient.ListProducts(ctx, &pb.Empty{})
	elapsed := time.Since(start)

	for err != nil {
		//fmt.Println("error: timeProductCatalogueServiceEmptyRequest:", err)
		
		start = time.Now()
		resp, err = productCatalogServiceClient.ListProducts(ctx, &pb.Empty{})
		elapsed = time.Since(start)
	}

	//fmt.Println(elapsed.Seconds())
	return elapsed.Seconds(), resp.GetProducts(), err
}

/* Send 'GetProductRequest' to ProductCatalogueService, receive 'Product' */
func (fe *frontendServer) timeProductCatalogueServiceGetProductRequest(ctx context.Context, id string) (float64, *pb.Product, error) {
	//fmt.Println("timeProductCatalogueServiceGetProductRequest()")

	productCatalogServiceClient := pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn)

	start := time.Now()
	resp, err := productCatalogServiceClient.GetProduct(ctx, &pb.GetProductRequest{Id: id})
	elapsed := time.Since(start)

	for err != nil {
		//fmt.Println("error: timeProductCatalogueServiceGetProductRequest:", err)
		
		start = time.Now()
		resp, err = productCatalogServiceClient.GetProduct(ctx, &pb.GetProductRequest{Id: id})
		elapsed = time.Since(start)
	}

	//fmt.Println(elapsed.Seconds())
	return elapsed.Seconds(), resp, err
}

/* Send 'SearchProductsRequest' to ProductCatalogueService, receive 'SearchProductsResponse' */
func (fe *frontendServer) timeProductCatalogueServiceSearchProductsRequest(ctx context.Context, query string) (float64, *pb.SearchProductsResponse, error) {
	//fmt.Println("timeProductCatalogueServiceSearchProductsRequest()")

	productCatalogServiceClient := pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn)	

	start := time.Now()
	resp, err := productCatalogServiceClient.SearchProducts(ctx, &pb.SearchProductsRequest{Query: query})
	elapsed := time.Since(start)

	for err != nil {
		//fmt.Println("error: timeProductCatalogueServiceSearchProductsRequest:", err)
		
		start = time.Now()
		resp, err = productCatalogServiceClient.SearchProducts(ctx, &pb.SearchProductsRequest{Query: query})
		elapsed = time.Since(start)
	}

	//fmt.Println(elapsed.Seconds())
	return elapsed.Seconds(), resp, err
}

/********************************************************************/
/* RecommendationService */
/********************************************************************/

/* Send 'ListRecommendationsRequest' to RecommendationService, receive 'ListRecommendationsResponse' */
func (fe *frontendServer) timeRecommendationServiceListRecommendationsRequest(ctx context.Context, userID string, productIDs []string) (float64, []*pb.Product, error) {
	//fmt.Println("timeRecommendationServiceListRecommendationsRequest()")

	recommendationServiceClient := pb.NewRecommendationServiceClient(fe.recommendationSvcConn)

	start := time.Now()
	resp, err := recommendationServiceClient.ListRecommendations(ctx, 
		&pb.ListRecommendationsRequest{UserId: userID, ProductIds: productIDs})
	elapsed := time.Since(start)

	for err != nil {
		//fmt.Println("error: timeRecommendationServiceListRecommendationsRequest:", err)
		
		start = time.Now()
		resp, err = recommendationServiceClient.ListRecommendations(ctx, 
			&pb.ListRecommendationsRequest{UserId: userID, ProductIds: productIDs})	
		elapsed = time.Since(start)
	}
	
	out := make([]*pb.Product, len(resp.GetProductIds()))
	for i, v := range resp.GetProductIds() {
		p, err := fe.getProduct(ctx, v)
		if err != nil {
			return elapsed.Seconds(), nil, errors.Wrapf(err, "failed to get recommended product info (#%s)", v)
		}
		out[i] = p
	}
	if len(out) > 4 {
		out = out[:4] // take only first four to fit the UI
	}

	//fmt.Println(elapsed.Seconds())
	return elapsed.Seconds(), out, err
}

/********************************************************************/
/* CheckoutService */
/********************************************************************/

/* Send 'PlaceOrderRequest' to CheckoutService, receive 'PlaceOrderResponse' */
func (fe *frontendServer) timeCheckoutServicePlaceOrderRequest(ctx context.Context) (float64, *pb.PlaceOrderResponse, error) {
	//fmt.Println("timeCheckoutServicePlaceOrderRequest()")

	checkoutServiceClient := pb.NewCheckoutServiceClient(fe.checkoutSvcConn)

	start := time.Now()
	order, err := checkoutServiceClient.
		PlaceOrder(ctx, &pb.PlaceOrderRequest{
			Email: EMAIL,
			CreditCard: CREDITCARDINFO,
			UserId: USER_ID,
			UserCurrency: USD,
			Address: ADDRESS,
		})	
	elapsed := time.Since(start)

	for err != nil {
		//fmt.Println("error: timeCheckoutServicePlaceOrderRequest:", err)
		
		start = time.Now()
		order, err = checkoutServiceClient.
			PlaceOrder(ctx, &pb.PlaceOrderRequest{
				Email: EMAIL,
				CreditCard: CREDITCARDINFO,
				UserId: USER_ID,
				UserCurrency: USD,
				Address: ADDRESS,
			})
		elapsed = time.Since(start)
	}
	
	//fmt.Println(elapsed.Seconds())
	return elapsed.Seconds(), order, err 
}

/********************************************************************/
/* ShippingService */
/********************************************************************/

/* Send 'GetQuoteRequest' to ShippingService, receive 'GetQuoteResponse' */
func (fe *frontendServer) timeShippingServiceGetQuoteRequest(ctx context.Context, items []*pb.CartItem, currency string) (float64, *pb.Money, error) {
	//fmt.Println("timeShippingServiceGetQuoteRequest()")

	shippingServiceClient := pb.NewShippingServiceClient(fe.shippingSvcConn)

	start := time.Now()
	quote, err := shippingServiceClient.GetQuote(ctx,
		&pb.GetQuoteRequest{
			Address: nil,
			Items:   items})
	elapsed := time.Since(start)

	for err != nil {
		//fmt.Println("error: timeShippingServiceGetQuoteRequest:", err)
		
		start = time.Now()
		quote, err = shippingServiceClient.GetQuote(ctx,
			&pb.GetQuoteRequest{
				Address: nil,
				Items:   items})
		elapsed = time.Since(start)
	}

	localized, err := fe.convertCurrency(ctx, quote.GetCostUsd(), currency)

	//fmt.Println(elapsed.Seconds())
	return elapsed.Seconds(), localized, errors.Wrap(err, "failed to convert currency for shipping cost")
}

/* Send 'ShipOrderRequest' to ShippingService, receive 'ShipOrderResponse' */
func (fe *frontendServer) timeShippingServiceShipOrderRequest(ctx context.Context, address *pb.Address, items []*pb.CartItem) (float64, string, error) {
	//fmt.Println("timeShippingServiceShipOrderRequest()")

	shippingServiceClient := pb.NewShippingServiceClient(fe.shippingSvcConn)

	start := time.Now()
	resp, err := shippingServiceClient.ShipOrder(ctx, &pb.ShipOrderRequest{
		Address: address,
		Items:   items})
	elapsed := time.Since(start)

	for err != nil {
		//fmt.Println("error: timeShippingServiceShipOrderRequest:", err)
		
		start = time.Now()
		resp, err = shippingServiceClient.ShipOrder(ctx, &pb.ShipOrderRequest{
			Address: address,
			Items:   items})
		elapsed = time.Since(start)
	}

	//fmt.Println(elapsed.Seconds())
	return elapsed.Seconds(), resp.GetTrackingId(), err
}

/********************************************************************/
/* CurrencyService */
/********************************************************************/

// Send 'Empty' to CurrencyService, receive 'GetSupportedCurrenciesResponse'
func (fe *frontendServer) timeCurrencyServiceEmptyRequest(ctx context.Context) (float64, []string, error) {
	//fmt.Println("timeCurrencyServiceEmptyRequest()")

	currencyServiceClient := pb.NewCurrencyServiceClient(fe.currencySvcConn)

	start := time.Now()
	resp, err := currencyServiceClient.
		GetSupportedCurrencies(ctx, &pb.Empty{})
	elapsed := time.Since(start)

	for err != nil {
		//fmt.Println("error: timeCurrencyServiceEmptyRequest:", err)
		
		start = time.Now()
		resp, err = currencyServiceClient.
			GetSupportedCurrencies(ctx, &pb.Empty{})
		elapsed = time.Since(start)
	}

	var out []string
	for _, c := range resp.CurrencyCodes {
		if _, ok := whitelistedCurrencies[c]; ok {
			out = append(out, c)
		}
	}
	
	//fmt.Println(elapsed.Seconds())
	return elapsed.Seconds(), out, err
}

// Send 'CurrencyConversionRequest' to CurrencyService, receive 'Money'
func (fe *frontendServer) timeCurrencyServiceCurrencyConversionRequest(ctx context.Context, money *pb.Money, currency string) (float64, *pb.Money, error) {
	//fmt.Println("timeCurrencyServiceCurrencyConversionRequest()")

	currencyServiceClient := pb.NewCurrencyServiceClient(fe.currencySvcConn)

	start := time.Now()
	resp, err := currencyServiceClient.
		Convert(ctx, &pb.CurrencyConversionRequest{
			From:   money,
			ToCode: currency})
	elapsed := time.Since(start)

	for err != nil {
		//fmt.Println("error: timeCurrencyServiceCurrencyConversionRequest", err)
		
		start = time.Now()
		resp, err = currencyServiceClient.
				Convert(ctx, &pb.CurrencyConversionRequest{
					From:   money,
					ToCode: currency})	
		elapsed = time.Since(start)
	}

	//fmt.Println(elapsed.Seconds())
	return elapsed.Seconds(), resp, err
}

/********************************************************************/
/* CartService */
/********************************************************************/

// Send 'AddItemRequest' to CartService, receive 'Empty'
func (fe *frontendServer) timeCartServiceAddItemRequest(ctx context.Context, userID, productID string, quantity int32) (float64, *pb.Empty, error) {
	//fmt.Println("timeCartServiceAddItemRequest()")

	cartServiceClient := pb.NewCartServiceClient(fe.cartSvcConn)

	start := time.Now()
	resp, err := cartServiceClient.AddItem(ctx, &pb.AddItemRequest{
		UserId: userID,
		Item: &pb.CartItem{
			ProductId: productID,
			Quantity:  quantity},
	})
	elapsed := time.Since(start)

	for err != nil {
		//fmt.Println("error: timeCartServiceAddItemRequest:", err)
		
		start = time.Now()
		resp, err = cartServiceClient.AddItem(ctx, &pb.AddItemRequest{
			UserId: userID,
			Item: &pb.CartItem{
				ProductId: productID,
				Quantity:  quantity},
		})		
		elapsed = time.Since(start)
	}

	//fmt.Println(elapsed.Seconds())
	return elapsed.Seconds(), resp, err
} 

// Send 'GetCartRequest' to CartService, receive 'Cart'
func (fe *frontendServer) timeCartServiceGetCartRequest(ctx context.Context, userID string) (float64, []*pb.CartItem, error) {
	//fmt.Println("timeCartServiceGetCartRequest()")

	cartServiceClient := pb.NewCartServiceClient(fe.cartSvcConn)
	
	start := time.Now()
	resp, err := cartServiceClient.GetCart(ctx, &pb.GetCartRequest{UserId: userID})
	elapsed := time.Since(start)

	for err != nil {
		//fmt.Println("error: timeCartServiceGetCartRequest:", err)
		
		start = time.Now()
		resp, err = cartServiceClient.GetCart(ctx, &pb.GetCartRequest{UserId: userID})
		elapsed = time.Since(start)
	}

	//fmt.Println(elapsed.Seconds())
	return elapsed.Seconds(), resp.GetItems(), err
} 

// Send 'EmptyCartRequest' to CartService, receive 'Empty'
func (fe *frontendServer) timeCartServiceEmptyCartRequest(ctx context.Context, userID string) (float64, *pb.Empty, error) {
	//fmt.Println("timeCartServiceEmptyCartRequest()")

	cartServiceClient := pb.NewCartServiceClient(fe.cartSvcConn)

	start := time.Now()	
	resp, err := cartServiceClient.EmptyCart(ctx, &pb.EmptyCartRequest{UserId: userID})
	elapsed := time.Since(start)

	for err != nil {
		//fmt.Println("error: timeCartServiceEmptyCartRequest:", err)
		
		start = time.Now()
		resp, err = cartServiceClient.EmptyCart(ctx, &pb.EmptyCartRequest{UserId: userID})
		elapsed = time.Since(start)
	}

	//fmt.Println(elapsed.Seconds())
	return elapsed.Seconds(), resp, err
} 

/********************************************************************/
/* AdService */
/********************************************************************/

// Send 'AdRequest' to AdService, receive 'AdResponse'
func (fe *frontendServer) timeAdServiceAdRequest(ctx context.Context, ctxKeys []string) (float64, []*pb.Ad, error) {
	//fmt.Println("timeAdServiceAdRequest()")

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	adServiceClient := pb.NewAdServiceClient(fe.adSvcConn)

	start := time.Now()	
	resp, err := adServiceClient.GetAds(ctx, &pb.AdRequest{
		ContextKeys: ctxKeys,
	})
	elapsed := time.Since(start)

	for i := 0; i < 5 && err != nil; i++ {
		//fmt.Println("error: timeAdServiceAdRequest:", err)
		
		start = time.Now()
		resp, err = adServiceClient.GetAds(ctx, &pb.AdRequest{
			ContextKeys: ctxKeys,
		})	
		elapsed = time.Since(start)
	}

	//fmt.Println(elapsed.Seconds())
	return elapsed.Seconds(), resp.GetAds(), err
}

/********************************************************************/
/* PaymentService */
/********************************************************************/

// Send 'ChargeRequest' to PaymentService, receive 'ChargeResponse'
func (fe *frontendServer) timePaymentServiceChargeRequest(ctx context.Context, amount *pb.Money, paymentInfo *pb.CreditCardInfo) (float64, string, error) {
	//fmt.Println("timePaymentServiceChargeRequest()")

	paymentServiceClient := pb.NewPaymentServiceClient(fe.paymentSvcConn)

	start := time.Now()	
	resp, err := paymentServiceClient.Charge(ctx, &pb.ChargeRequest{
		Amount:     amount,
		CreditCard: paymentInfo})
	elapsed := time.Since(start)

	for i := 0; i < 5 && err != nil; i++  {
		//fmt.Println("error: timePaymentServiceChargeRequest:", err)
		
		start = time.Now()
		resp, err = paymentServiceClient.Charge(ctx, &pb.ChargeRequest{
			Amount:     amount,
			CreditCard: paymentInfo})
		elapsed = time.Since(start)
	}

	//fmt.Println(elapsed.Seconds())
	return elapsed.Seconds(), resp.GetTransactionId(), err
}

/********************************************************************/
/* EmailService */
/********************************************************************/

// Send 'SendOrderConfirmationRequest' to EmailService, receive 'Empty'
func (fe *frontendServer) timeEmailServiceSendOrderConfirmationRequest(ctx context.Context, email string, order *pb.OrderResult) (float64, *pb.Empty, error) {
	//fmt.Println("timeEmailServiceSendOrderConfirmationRequest()")

	emailServiceClient := pb.NewEmailServiceClient(fe.emailSvcConn)

	start := time.Now()	
	resp, err := emailServiceClient.SendOrderConfirmation(ctx, &pb.SendOrderConfirmationRequest{
		Email: email,
		Order: order})
	elapsed := time.Since(start)

	for i := 0; i < 5 && err != nil; i++ {
		//fmt.Println("error: timeEmailServiceSendOrderConfirmationRequest", err)
		
		start = time.Now()
		resp, err = emailServiceClient.SendOrderConfirmation(ctx, &pb.SendOrderConfirmationRequest{
			Email: email,
			Order: order})	
		elapsed = time.Since(start)
	}

	//fmt.Println(elapsed.Seconds())
	return elapsed.Seconds(), resp, err
}