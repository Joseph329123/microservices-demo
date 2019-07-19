package main

import (
	"context"
	"time"
	"fmt"

	pb "github.com/Joseph329123/microservices-demo/src/testservice/genproto"

	"github.com/pkg/errors"
)

func runResponseTimeTests(ctx context.Context, svc *frontendServer, samples int32) {
	fmt.Println("runResponseTimeTests()")
	fmt.Println(samples)

	/* ProductCatalogueService */
	svc.timeProductCatalogueServiceEmptyRequest(ctx)
	svc.timeProductCatalogueServiceGetProductRequest(ctx, VINTAGE_TYPEWRITER_ID)
	svc.timeProductCatalogueServiceSearchProductsRequest(ctx, VINTAGE_TYPEWRITER)

	/* RecommendationService */
	svc.timeRecommendationServiceListRecommendationsRequest(ctx, USER_ID, PRODUCT_IDS)

	/* CheckoutService */
	svc.timeCheckoutServicePlaceOrderRequest(ctx)

	/* ShippingService */
	svc.timeShippingServiceGetQuoteRequest(ctx, ITEMS, USD)
	svc.timeShippingServiceShipOrderRequest(ctx, ADDRESS, ITEMS)

	/* CurrencyService */
	svc.timeCurrencyServiceCurrencyConversionRequest(ctx, MONEY, USD)
	svc.timeCurrencyServiceEmptyRequest(ctx)

	/* CartService */
	svc.timeCartServiceAddItemRequest(ctx, USER_ID, VINTAGE_TYPEWRITER_ID, 1)
	svc.timeCartServiceGetCartRequest(ctx, USER_ID)
	svc.timeCartServiceEmptyCartRequest(ctx, USER_ID)

	/* AdService */
	svc.timeAdServiceAdRequest(ctx, CTX_KEYS)

	/* PaymentService */
	svc.timePaymentServiceChargeRequest(ctx, MONEY, CREDITCARDINFO)

	/* EmailService */
	svc.timeEmailServiceSendOrderConfirmationRequest(ctx, EMAIL, ORDER_RESULT)
}

/********************************************************************/
/* ProductCatalogueService */
/********************************************************************/

/* Send 'Empty' to ProductCatalogueService, receive 'ListProductsResponse' */
func (fe *frontendServer) timeProductCatalogueServiceEmptyRequest(ctx context.Context) (time.Duration, []*pb.Product, error) {
	fmt.Println("timeProductCatalogueServiceEmptyRequest()")

	productCatalogServiceClient := pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn)

	start := time.Now()
	resp, err := productCatalogServiceClient.ListProducts(ctx, &pb.Empty{})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}
	
	for err != nil {
		start = time.Now()
		resp, err = productCatalogServiceClient.ListProducts(ctx, &pb.Empty{})
		elapsed = time.Since(start)
	}

	fmt.Println(elapsed)
	return elapsed, resp.GetProducts(), err
}

/* Send 'GetProductRequest' to ProductCatalogueService, receive 'Product' */
func (fe *frontendServer) timeProductCatalogueServiceGetProductRequest(ctx context.Context, id string) (time.Duration, *pb.Product, error) {
	fmt.Println("timeProductCatalogueServiceGetProductRequest()")

	productCatalogServiceClient := pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn)

	start := time.Now()
	resp, err := productCatalogServiceClient.GetProduct(ctx, &pb.GetProductRequest{Id: id})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}

	for err != nil {
		//fmt.Println("error")
		start = time.Now()
		resp, err = productCatalogServiceClient.GetProduct(ctx, &pb.GetProductRequest{Id: id})
		elapsed = time.Since(start)
	}
	
	fmt.Println(elapsed)
	fmt.Println(resp.Name)
	return elapsed, resp, err
}

/* Send 'SearchProductsRequest' to ProductCatalogueService, receive 'SearchProductsResponse' */
func (fe *frontendServer) timeProductCatalogueServiceSearchProductsRequest(ctx context.Context, query string) (time.Duration, *pb.SearchProductsResponse, error) {
	fmt.Println("timeProductCatalogueServiceSearchProductsRequest()")

	productCatalogServiceClient := pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn)	

	start := time.Now()
	resp, err := productCatalogServiceClient.SearchProducts(ctx, &pb.SearchProductsRequest{Query: query})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}
	
	for err != nil {
		//fmt.Println("error")
		start = time.Now()
		resp, err = productCatalogServiceClient.SearchProducts(ctx, &pb.SearchProductsRequest{Query: query})
		elapsed = time.Since(start)
	}
	
	fmt.Println(elapsed)
	fmt.Println(resp.Results[0].Name)
	return elapsed, resp, err
}

/********************************************************************/
/* RecommendationService */
/********************************************************************/

/* Send 'ListRecommendationsRequest' to RecommendationService, receive 'ListRecommendationsResponse' */
func (fe *frontendServer) timeRecommendationServiceListRecommendationsRequest(ctx context.Context, userID string, productIDs []string) (time.Duration, []*pb.Product, error) {
	fmt.Println("timeRecommendationServiceListRecommendationsRequest()")

	recommendationServiceClient := pb.NewRecommendationServiceClient(fe.recommendationSvcConn)

	start := time.Now()
	resp, err := recommendationServiceClient.ListRecommendations(ctx, 
		&pb.ListRecommendationsRequest{UserId: userID, ProductIds: productIDs})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}

	for err != nil {
		//fmt.Println("error")
		start = time.Now()
		resp, err = recommendationServiceClient.ListRecommendations(ctx,
			&pb.ListRecommendationsRequest{UserId: userID, ProductIds: productIDs})
		elapsed = time.Since(start)
	}
	
	out := make([]*pb.Product, len(resp.GetProductIds()))
	for i, v := range resp.GetProductIds() {
		p, err := fe.getProduct(ctx, v)
		if err != nil {
			return 0, nil, errors.Wrapf(err, "failed to get recommended product info (#%s)", v)
		}
		out[i] = p
	}
	if len(out) > 4 {
		out = out[:4] // take only first four to fit the UI
	}

	fmt.Println(elapsed)
	return elapsed, out, err
}

/********************************************************************/
/* CheckoutService */
/********************************************************************/

/* Send 'PlaceOrderRequest' to CheckoutService, receive 'PlaceOrderResponse' */
func (fe *frontendServer) timeCheckoutServicePlaceOrderRequest(ctx context.Context) (time.Duration, *pb.PlaceOrderResponse, error) {
	fmt.Println("timeCheckoutServicePlaceOrderRequest()")

	checkoutServiceClient := pb.NewCheckoutServiceClient(fe.checkoutSvcConn)

	start := time.Now()
	order, err := checkoutServiceClient.
		PlaceOrder(ctx, &pb.PlaceOrderRequest{
			Email: "someone@example.com",
			CreditCard: &pb.CreditCardInfo{
				CreditCardNumber:          "4432801561520454",
				CreditCardExpirationMonth: 1,
				CreditCardExpirationYear:  2020,
				CreditCardCvv:             672},
			UserId:       "dummy",
			UserCurrency: "USD",
			Address: &pb.Address{
				StreetAddress: "1600 Amphitheatre Parkway",
				City:          "Mountain View",
				State:         "CA",
				ZipCode:       94043,
				Country:       "Mountain View"},
		})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}

	for err != nil {
		//fmt.Println("error")
		start = time.Now()
		order, err = checkoutServiceClient.
			PlaceOrder(ctx, &pb.PlaceOrderRequest{
				Email: "someone@example.com",
				CreditCard: &pb.CreditCardInfo{
					CreditCardNumber:          "4432801561520454",
					CreditCardExpirationMonth: 1,
					CreditCardExpirationYear:  2020,
					CreditCardCvv:             672},
				UserId:       "dummy",
				UserCurrency: "USD",
				Address: &pb.Address{
					StreetAddress: "1600 Amphitheatre Parkway",
					City:          "Mountain View",
					State:         "CA",
					ZipCode:       94043,
					Country:       "Mountain View"},
			})
		elapsed = time.Since(start)
	}

	fmt.Println(elapsed)
	return elapsed, order, err 
}

/********************************************************************/
/* ShippingService */
/********************************************************************/

/* Send 'GetQuoteRequest' to ShippingService, receive 'GetQuoteResponse' */
func (fe *frontendServer) timeShippingServiceGetQuoteRequest(ctx context.Context, items []*pb.CartItem, currency string) (time.Duration, *pb.Money, error) {
	fmt.Println("timeShippingServiceGetQuoteRequest()")

	shippingServiceClient := pb.NewShippingServiceClient(fe.shippingSvcConn)

	start := time.Now()
	quote, err := shippingServiceClient.GetQuote(ctx,
		&pb.GetQuoteRequest{
			Address: nil,
			Items:   items})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}

	for err != nil {
		start = time.Now()
		quote, err = shippingServiceClient.GetQuote(ctx,
			&pb.GetQuoteRequest{
				Address: nil,
				Items:   items})
		elapsed = time.Since(start)

	}
	localized, err := fe.convertCurrency(ctx, quote.GetCostUsd(), currency)

	fmt.Println(elapsed)
	return elapsed, localized, errors.Wrap(err, "failed to convert currency for shipping cost")
}

/* Send 'ShipOrderRequest' to ShippingService, receive 'ShipOrderResponse' */
func (fe *frontendServer) timeShippingServiceShipOrderRequest(ctx context.Context, address *pb.Address, items []*pb.CartItem) (time.Duration, string, error) {
	fmt.Println("timeShippingServiceShipOrderRequest()")

	shippingServiceClient := pb.NewShippingServiceClient(fe.shippingSvcConn)

	start := time.Now()
	resp, err := shippingServiceClient.ShipOrder(ctx, &pb.ShipOrderRequest{
		Address: address,
		Items:   items})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}	

	for err != nil {
		start = time.Now()
		resp, err = shippingServiceClient.ShipOrder(ctx, &pb.ShipOrderRequest{
			Address: address,
			Items:   items})
		elapsed = time.Since(start)
	}

	fmt.Println(elapsed)
	return elapsed, resp.GetTrackingId(), err
}

/********************************************************************/
/* CurrencyService */
/********************************************************************/

// Send 'Empty' to CurrencyService, receive 'GetSupportedCurrenciesResponse'
func (fe *frontendServer) timeCurrencyServiceEmptyRequest(ctx context.Context) (time.Duration, []string, error) {
	fmt.Println("timeCurrencyServiceEmptyRequest()")

	currencyServiceClient := pb.NewCurrencyServiceClient(fe.currencySvcConn)

	start := time.Now()
	resp, err := currencyServiceClient.
		GetSupportedCurrencies(ctx, &pb.Empty{})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}

	for err != nil {
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

	fmt.Println("elapsed:", elapsed)
	return elapsed, out, nil
}

// Send 'CurrencyConversionRequest' to CurrencyService, receive 'Money'
func (fe *frontendServer) timeCurrencyServiceCurrencyConversionRequest(ctx context.Context, money *pb.Money, currency string) (time.Duration, *pb.Money, error) {
	fmt.Println("timeCurrencyServiceCurrencyConversionRequest()")

	currencyServiceClient := pb.NewCurrencyServiceClient(fe.currencySvcConn)

	start := time.Now()
	resp, err := currencyServiceClient.
		Convert(ctx, &pb.CurrencyConversionRequest{
			From:   money,
			ToCode: currency})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}

	for err != nil {
		fmt.Println("for loop")
		start = time.Now()
		resp, err = currencyServiceClient.
			Convert(ctx, &pb.CurrencyConversionRequest{
				From:   money,
				ToCode: currency})
		elapsed = time.Since(start)
	}

	fmt.Println("Curr Code:", resp.GetCurrencyCode())
	fmt.Println("Units:", resp.GetUnits())
	fmt.Println("Nanos:", resp.GetNanos())
	fmt.Println("elapsed:", elapsed)
	return elapsed, resp, err
}

/********************************************************************/
/* CartService */
/********************************************************************/

// Send 'AddItemRequest' to CartService, receive 'Empty'
func (fe *frontendServer) timeCartServiceAddItemRequest(ctx context.Context, userID, productID string, quantity int32) (time.Duration, *pb.Empty, error) {
	fmt.Println("timeCartServiceAddItemRequest()")

	cartServiceClient := pb.NewCartServiceClient(fe.cartSvcConn)

	start := time.Now()
	resp, err := cartServiceClient.AddItem(ctx, &pb.AddItemRequest{
		UserId: userID,
		Item: &pb.CartItem{
			ProductId: productID,
			Quantity:  quantity},
	})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}

	for err != nil {
		start = time.Now()
		resp, err = cartServiceClient.AddItem(ctx, &pb.AddItemRequest{
			UserId: userID,
			Item: &pb.CartItem{
				ProductId: productID,
				Quantity:  quantity},
		})
		elapsed = time.Since(start)		
	}

	fmt.Println(elapsed)
	return elapsed, resp, err
} 

// Send 'GetCartRequest' to CartService, receive 'Cart'
func (fe *frontendServer) timeCartServiceGetCartRequest(ctx context.Context, userID string) (time.Duration, []*pb.CartItem, error) {
	fmt.Println("timeCartServiceGetCartRequest()")

	cartServiceClient := pb.NewCartServiceClient(fe.cartSvcConn)
	
	start := time.Now()
	resp, err := cartServiceClient.GetCart(ctx, &pb.GetCartRequest{UserId: userID})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}

	for err != nil {
		start = time.Now()
		resp, err = cartServiceClient.GetCart(ctx, &pb.GetCartRequest{UserId: userID})
		elapsed = time.Since(start)		
	}

	fmt.Println(elapsed)
	fmt.Println((resp.GetItems())[0].GetProductId())

	return elapsed, resp.GetItems(), err
} 

// Send 'EmptyCartRequest' to CartService, receive 'Empty'
func (fe *frontendServer) timeCartServiceEmptyCartRequest(ctx context.Context, userID string) (time.Duration, *pb.Empty, error) {
	fmt.Println("timeCartServiceEmptyCartRequest()")

	cartServiceClient := pb.NewCartServiceClient(fe.cartSvcConn)

	start := time.Now()	
	resp, err := cartServiceClient.EmptyCart(ctx, &pb.EmptyCartRequest{UserId: userID})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}

	for err != nil {
		start = time.Now()
		resp, err = cartServiceClient.EmptyCart(ctx, &pb.EmptyCartRequest{UserId: userID})
		elapsed = time.Since(start)		
	}

	fmt.Println(elapsed)
	return elapsed, resp, err
} 

/********************************************************************/
/* AdService */
/********************************************************************/

// Send 'AdRequest' to AdService, receive 'AdResponse'
func (fe *frontendServer) timeAdServiceAdRequest(ctx context.Context, ctxKeys []string) (time.Duration, []*pb.Ad, error) {
	fmt.Println("timeAdServiceAdRequest()")

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	adServiceClient := pb.NewAdServiceClient(fe.adSvcConn)

	start := time.Now()	
	resp, err := adServiceClient.GetAds(ctx, &pb.AdRequest{
		ContextKeys: ctxKeys,
	})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}

	for err != nil {
		start = time.Now()
		resp, err = adServiceClient.GetAds(ctx, &pb.AdRequest{
			ContextKeys: ctxKeys,
		})		
		elapsed = time.Since(start)		
	}

	fmt.Println(elapsed)
	return elapsed, resp.GetAds(), err
}

/********************************************************************/
/* PaymentService */
/********************************************************************/

// Send 'ChargeRequest' to PaymentService, receive 'ChargeResponse'
func (fe *frontendServer) timePaymentServiceChargeRequest(ctx context.Context, amount *pb.Money, paymentInfo *pb.CreditCardInfo) (time.Duration, string, error) {
	fmt.Println("timePaymentServiceChargeRequest()")

	paymentServiceClient := pb.NewPaymentServiceClient(fe.paymentSvcConn)

	start := time.Now()	
	resp, err := paymentServiceClient.Charge(ctx, &pb.ChargeRequest{
		Amount:     amount,
		CreditCard: paymentInfo})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}

	for err != nil {
		start = time.Now()
		resp, err = paymentServiceClient.Charge(ctx, &pb.ChargeRequest{
			Amount:     amount,
			CreditCard: paymentInfo})	
		elapsed = time.Since(start)		
	}

	fmt.Println(elapsed)
	fmt.Println(resp.GetTransactionId())
	return elapsed, resp.GetTransactionId(), err
}

/********************************************************************/
/* EmailService */
/********************************************************************/

// Send 'SendOrderConfirmationRequest' to EmailService, receive 'Empty'
func (fe *frontendServer) timeEmailServiceSendOrderConfirmationRequest(ctx context.Context, email string, order *pb.OrderResult) (time.Duration, *pb.Empty, error) {
	fmt.Println("timeEmailServiceSendOrderConfirmationRequest()")

	emailServiceClient := pb.NewEmailServiceClient(fe.emailSvcConn)

	start := time.Now()	
	resp, err := emailServiceClient.SendOrderConfirmation(ctx, &pb.SendOrderConfirmationRequest{
		Email: email,
		Order: order})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}

	for err != nil {
		start = time.Now()
		resp, err = emailServiceClient.SendOrderConfirmation(ctx, &pb.SendOrderConfirmationRequest{
			Email: email,
			Order: order})	
		elapsed = time.Since(start)		
	}

	fmt.Println(elapsed)

	return elapsed, resp, err
}