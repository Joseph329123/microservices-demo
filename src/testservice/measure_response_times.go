// This file sends requests to all services and requests
// and measures performance and availability, it then writes these
// results to a csv file

package main

import (
	"context"
	"time"
	"fmt"
	"os"

	pb "github.com/Joseph329123/microservices-demo/src/testservice/genproto"

	"github.com/pkg/errors"
)

func runResponseTimeTests(ctx context.Context, svc *frontendServer, samples int) {
	fmt.Println("runResponseTimeTests()")

	data_response_times := make([][]float64, NUMBER_OF_REQUESTS)
	data_availability := make([][]error, NUMBER_OF_REQUESTS)

	for i := 0; i < NUMBER_OF_REQUESTS; i++ {
		data_response_times[i] = make([]float64, samples)
		data_availability[i] = make([]error, samples)
	}

	// we want the checkout service to receive configured results from the
	// cart service and shipping service (for example,
	// the first cart request the
	// checkoutService sends to the cart service should result in a cart
	// response that contains an item with a nonexistent product ID)
	// this dummy call to checkout service accomplishes this
	if s := os.Getenv("CART_RESPONSE"); (s == "exception") || (s == "bad product id") {
		svc.timeCheckoutServicePlaceOrderRequest(ctx)
	}

	if s := os.Getenv("GET_QUOTE_RESPONSE"); s == "error" {
		svc.timeCheckoutServicePlaceOrderRequest(ctx)
	}

	for i := 0; i < samples; i++ {
		fmt.Println(i)
		/* ProductCatalogueService */
		data_response_times[PRODUCT_CATALOGUE_EMPTY_REQUEST_INDEX][i], _, data_availability[PRODUCT_CATALOGUE_EMPTY_REQUEST_INDEX][i] = svc.timeProductCatalogueServiceEmptyRequest(ctx)

		data_response_times[PRODUCT_CATALOGUE_GET_PRODUCT_REQUEST_INDEX][i], _, data_availability[PRODUCT_CATALOGUE_GET_PRODUCT_REQUEST_INDEX][i] = svc.timeProductCatalogueServiceGetProductRequest(ctx, VINTAGE_TYPEWRITER_ID)

		data_response_times[PRODUCT_CATALOGUE_SEARCH_PRODUCT_REQUEST_INDEX][i], _, data_availability[PRODUCT_CATALOGUE_SEARCH_PRODUCT_REQUEST_INDEX][i] = svc.timeProductCatalogueServiceSearchProductsRequest(ctx, VINTAGE_TYPEWRITER)

		/* RecommendationService */
		data_response_times[RECOMMENDATION_LIST_RECOMMENDATIONS_REQUEST_INDEX][i], _, data_availability[RECOMMENDATION_LIST_RECOMMENDATIONS_REQUEST_INDEX][i] = svc.timeRecommendationServiceListRecommendationsRequest(ctx, USER_ID, PRODUCT_IDS)

		/* CheckoutService */
		data_response_times[CHECKOUT_PLACE_ORDER_REQUEST_INDEX][i], _, data_availability[CHECKOUT_PLACE_ORDER_REQUEST_INDEX][i] = svc.timeCheckoutServicePlaceOrderRequest(ctx)

		/* ShippingService */
		data_response_times[SHIPPING_GET_QUOTE_REQUEST_INDEX][i], _, data_availability[SHIPPING_GET_QUOTE_REQUEST_INDEX][i] = svc.timeShippingServiceGetQuoteRequest(ctx, ITEMS, USD)

		data_response_times[SHIPPING_SHIP_ORDER_REQUEST_INDEX][i], _, data_availability[SHIPPING_SHIP_ORDER_REQUEST_INDEX][i] = svc.timeShippingServiceShipOrderRequest(ctx, ADDRESS, ITEMS)

		/* CurrencyService */
		data_response_times[CURRENCY_CURRENCY_CONVERSION_REQUEST_INDEX][i], _, data_availability[CURRENCY_CURRENCY_CONVERSION_REQUEST_INDEX][i] = svc.timeCurrencyServiceCurrencyConversionRequest(ctx, MONEY, USD)

		data_response_times[CURRENCY_EMPTY_REQUEST_INDEX][i], _, data_availability[CURRENCY_EMPTY_REQUEST_INDEX][i] = svc.timeCurrencyServiceEmptyRequest(ctx)

		/* CartService */
		data_response_times[CART_ADD_ITEM_REQUEST_INDEX][i], _, data_availability[CART_ADD_ITEM_REQUEST_INDEX][i] = svc.timeCartServiceAddItemRequest(ctx, USER_ID, VINTAGE_TYPEWRITER_ID, 1)

		data_response_times[CART_GET_CART_REQUEST_INDEX][i], _, data_availability[CART_GET_CART_REQUEST_INDEX][i] = svc.timeCartServiceGetCartRequest(ctx, USER_ID)

		data_response_times[CART_EMPTY_CART_REQUEST_INDEX][i], _, data_availability[CART_EMPTY_CART_REQUEST_INDEX][i] = svc.timeCartServiceEmptyCartRequest(ctx, USER_ID)

		/* AdService */
		data_response_times[AD_AD_REQUEST_INDEX][i], _, data_availability[AD_AD_REQUEST_INDEX][i] = svc.timeAdServiceAdRequest(ctx, CTX_KEYS)

		/* PaymentService */
		data_response_times[PAYMENT_CHARGE_REQUEST_INDEX][i], _, data_availability[PAYMENT_CHARGE_REQUEST_INDEX][i] = svc.timePaymentServiceChargeRequest(ctx, MONEY, CREDITCARDINFO)

		/* EmailService */
		data_response_times[EMAIL_SEND_ORDER_CONFIRMATION_REQUEST_INDEX][i], _, data_availability[EMAIL_SEND_ORDER_CONFIRMATION_REQUEST_INDEX][i] = svc.timeEmailServiceSendOrderConfirmationRequest(ctx, EMAIL, ORDER_RESULT)
	}

	data_response_times_summary := getResponseTimeSummary(data_response_times)
	data_availability_summary := getAvailabilitySummary(data_availability, samples)
	writeToCSV(data_response_times_summary, data_availability_summary)
	fmt.Println("done")
}

func writeToCSV(data_response_times_summary[][] float64, data_availability_summary[] float64) {
	f, err := os.Create("results_unformatted.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	k := 0
	for _, i := range data_response_times_summary {
		_, err := f.WriteString(ROW_LABELS[k] + ",")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}

		// writing performance stats
		for _, j := range i {
			_, err = f.WriteString(fmt.Sprint(j) + ",")
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
		}

		// writing availability stats
		_, err = f.WriteString(fmt.Sprint(data_availability_summary[k]) + "\n")

		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		k += 1
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getResponseTimeSummary(data[][] float64) ([][] float64) {
	data_response_times_summary := make([][]float64, NUMBER_OF_REQUESTS)

	for i := 0; i < NUMBER_OF_REQUESTS; i++ {
		data_response_times_summary[i] = make([]float64, 4)
		data_response_times_summary[i][BEST_RESPONSE_TIME_INDEX] = find_min(data[i])
		data_response_times_summary[i][WORST_RESPONSE_TIME_INDEX] = find_max(data[i])
		data_response_times_summary[i][MEAN_RESPONSE_TIME_INDEX] = mean(data[i])
		data_response_times_summary[i][STD_DEV_INDEX] = std_dev(data[i])
	}

	return data_response_times_summary
}

func getAvailabilitySummary(data[][] error, samples int) ([] float64) {
	data_availability_summary := make([]float64, NUMBER_OF_REQUESTS)

	for i := 0; i < NUMBER_OF_REQUESTS; i++ {
		availabilityCount := 0
		for j := 0; j < samples; j++ {
			if data[i][j] == nil {
				availabilityCount += 1
			}
		}
		data_availability_summary[i] = float64(availabilityCount)/float64(samples)
	}

	return data_availability_summary
}

/********************************************************************/
/* ProductCatalogueService */
/********************************************************************/

/* Send 'Empty' to ProductCatalogueService, receive 'ListProductsResponse' */
func (fe *frontendServer) timeProductCatalogueServiceEmptyRequest(ctx context.Context) (float64, []*pb.Product, error) {
	productCatalogServiceClient := pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn)

	start := time.Now()
	resp, err := productCatalogServiceClient.ListProducts(ctx, &pb.Empty{})
	elapsed := time.Since(start)

	return elapsed.Seconds(), resp.GetProducts(), err
}

/* Send 'GetProductRequest' to ProductCatalogueService, receive 'Product' */
func (fe *frontendServer) timeProductCatalogueServiceGetProductRequest(ctx context.Context, id string) (float64, *pb.Product, error) {
	productCatalogServiceClient := pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn)

	start := time.Now()
	resp, err := productCatalogServiceClient.GetProduct(ctx, &pb.GetProductRequest{Id: id})
	elapsed := time.Since(start)

	return elapsed.Seconds(), resp, err
}

/* Send 'SearchProductsRequest' to ProductCatalogueService, receive 'SearchProductsResponse' */
func (fe *frontendServer) timeProductCatalogueServiceSearchProductsRequest(ctx context.Context, query string) (float64, *pb.SearchProductsResponse, error) {
	productCatalogServiceClient := pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn)

	start := time.Now()
	resp, err := productCatalogServiceClient.SearchProducts(ctx, &pb.SearchProductsRequest{Query: query})
	elapsed := time.Since(start)

	return elapsed.Seconds(), resp, err
}

/********************************************************************/
/* RecommendationService */
/********************************************************************/

/* Send 'ListRecommendationsRequest' to RecommendationService, receive 'ListRecommendationsResponse' */
func (fe *frontendServer) timeRecommendationServiceListRecommendationsRequest(ctx context.Context, userID string, productIDs []string) (float64, []*pb.Product, error) {
	recommendationServiceClient := pb.NewRecommendationServiceClient(fe.recommendationSvcConn)

	start := time.Now()
	resp, err := recommendationServiceClient.ListRecommendations(ctx,
		&pb.ListRecommendationsRequest{UserId: userID, ProductIds: productIDs})
	elapsed := time.Since(start)

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

	return elapsed.Seconds(), out, err
}

/********************************************************************/
/* CheckoutService */
/********************************************************************/

/* Send 'PlaceOrderRequest' to CheckoutService, receive 'PlaceOrderResponse' */
func (fe *frontendServer) timeCheckoutServicePlaceOrderRequest(ctx context.Context) (float64, *pb.PlaceOrderResponse, error) {
	checkoutServiceClient := pb.NewCheckoutServiceClient(fe.checkoutSvcConn)

	// first add some items to cart
	cartServiceClient := pb.NewCartServiceClient(fe.cartSvcConn)

	cartServiceClient.AddItem(ctx, &pb.AddItemRequest{
		UserId: "dummyIDcheckout",
		Item: &pb.CartItem{
			ProductId: "OLJCESPC7Z",
			Quantity:  10},
	})

	checkout_address := &pb.Address{
					StreetAddress: "1600 Amphitheatre Parkway",
					City:          "checkout",
					State:         "CA",
					ZipCode:       94043,
					Country:       "Mountain View"}

	start := time.Now()
	order, err := checkoutServiceClient.
		PlaceOrder(ctx, &pb.PlaceOrderRequest{
			Email: EMAIL,
			CreditCard: CREDITCARDINFO,
			UserId: "dummyIDcheckout",
			UserCurrency: USD,
			Address: checkout_address,
		})
	elapsed := time.Since(start)

	return elapsed.Seconds(), order, err
}

/********************************************************************/
/* ShippingService */
/********************************************************************/

/* Send 'GetQuoteRequest' to ShippingService, receive 'GetQuoteResponse' */
func (fe *frontendServer) timeShippingServiceGetQuoteRequest(ctx context.Context, items []*pb.CartItem, currency string) (float64, *pb.Money, error) {
	shippingServiceClient := pb.NewShippingServiceClient(fe.shippingSvcConn)

	start := time.Now()
	quote, err := shippingServiceClient.GetQuote(ctx,
		&pb.GetQuoteRequest{
			Address: ADDRESS,
			Items:   items})
	elapsed := time.Since(start)

	localized, err := fe.convertCurrency(ctx, quote.GetCostUsd(), currency)

	return elapsed.Seconds(), localized, errors.Wrap(err, "failed to convert currency for shipping cost")
}

/* Send 'ShipOrderRequest' to ShippingService, receive 'ShipOrderResponse' */
func (fe *frontendServer) timeShippingServiceShipOrderRequest(ctx context.Context, address *pb.Address, items []*pb.CartItem) (float64, string, error) {
	shippingServiceClient := pb.NewShippingServiceClient(fe.shippingSvcConn)

	start := time.Now()
	resp, err := shippingServiceClient.ShipOrder(ctx, &pb.ShipOrderRequest{
		Address: address,
		Items:   items})
	elapsed := time.Since(start)

	return elapsed.Seconds(), resp.GetTrackingId(), err
}

/********************************************************************/
/* CurrencyService */
/********************************************************************/

// Send 'Empty' to CurrencyService, receive 'GetSupportedCurrenciesResponse'
func (fe *frontendServer) timeCurrencyServiceEmptyRequest(ctx context.Context) (float64, []string, error) {
	currencyServiceClient := pb.NewCurrencyServiceClient(fe.currencySvcConn)

	start := time.Now()
	resp, err := currencyServiceClient.
		GetSupportedCurrencies(ctx, &pb.Empty{})
	elapsed := time.Since(start)

	var out []string
	for _, c := range resp.CurrencyCodes {
		if _, ok := whitelistedCurrencies[c]; ok {
			out = append(out, c)
		}
	}

	return elapsed.Seconds(), out, err
}

// Send 'CurrencyConversionRequest' to CurrencyService, receive 'Money'
func (fe *frontendServer) timeCurrencyServiceCurrencyConversionRequest(ctx context.Context, money *pb.Money, currency string) (float64, *pb.Money, error) {
	currencyServiceClient := pb.NewCurrencyServiceClient(fe.currencySvcConn)

	start := time.Now()
	resp, err := currencyServiceClient.
		Convert(ctx, &pb.CurrencyConversionRequest{
			From:   money,
			ToCode: currency})
	elapsed := time.Since(start)

	return elapsed.Seconds(), resp, err
}

/********************************************************************/
/* CartService */
/********************************************************************/

// Send 'AddItemRequest' to CartService, receive 'Empty'
func (fe *frontendServer) timeCartServiceAddItemRequest(ctx context.Context, userID, productID string, quantity int32) (float64, *pb.Empty, error) {
	cartServiceClient := pb.NewCartServiceClient(fe.cartSvcConn)

	start := time.Now()
	resp, err := cartServiceClient.AddItem(ctx, &pb.AddItemRequest{
		UserId: userID,
		Item: &pb.CartItem{
			ProductId: productID,
			Quantity:  quantity},
	})
	elapsed := time.Since(start)

	return elapsed.Seconds(), resp, err
}

// Send 'GetCartRequest' to CartService, receive 'Cart'
func (fe *frontendServer) timeCartServiceGetCartRequest(ctx context.Context, userID string) (float64, []*pb.CartItem, error) {
	cartServiceClient := pb.NewCartServiceClient(fe.cartSvcConn)

	start := time.Now()
	resp, err := cartServiceClient.GetCart(ctx, &pb.GetCartRequest{UserId: userID})
	elapsed := time.Since(start)

	return elapsed.Seconds(), resp.GetItems(), err
}

// Send 'EmptyCartRequest' to CartService, receive 'Empty'
func (fe *frontendServer) timeCartServiceEmptyCartRequest(ctx context.Context, userID string) (float64, *pb.Empty, error) {
	cartServiceClient := pb.NewCartServiceClient(fe.cartSvcConn)

	start := time.Now()
	resp, err := cartServiceClient.EmptyCart(ctx, &pb.EmptyCartRequest{UserId: userID})
	elapsed := time.Since(start)

	return elapsed.Seconds(), resp, err
}

/********************************************************************/
/* AdService */
/********************************************************************/

// Send 'AdRequest' to AdService, receive 'AdResponse'
func (fe *frontendServer) timeAdServiceAdRequest(ctx context.Context, ctxKeys []string) (float64, []*pb.Ad, error) {
	adServiceClient := pb.NewAdServiceClient(fe.adSvcConn)

	start := time.Now()
	resp, err := adServiceClient.GetAds(ctx, &pb.AdRequest{
		ContextKeys: ctxKeys,
	})
	elapsed := time.Since(start)

	return elapsed.Seconds(), resp.GetAds(), err
}

/********************************************************************/
/* PaymentService */
/********************************************************************/

// Send 'ChargeRequest' to PaymentService, receive 'ChargeResponse'
func (fe *frontendServer) timePaymentServiceChargeRequest(ctx context.Context, amount *pb.Money, paymentInfo *pb.CreditCardInfo) (float64, string, error) {
	paymentServiceClient := pb.NewPaymentServiceClient(fe.paymentSvcConn)

	start := time.Now()
	resp, err := paymentServiceClient.Charge(ctx, &pb.ChargeRequest{
		Amount:     amount,
		CreditCard: paymentInfo})
	elapsed := time.Since(start)

	return elapsed.Seconds(), resp.GetTransactionId(), err
}

/********************************************************************/
/* EmailService */
/********************************************************************/

// Send 'SendOrderConfirmationRequest' to EmailService, receive 'Empty'
func (fe *frontendServer) timeEmailServiceSendOrderConfirmationRequest(ctx context.Context, email string, order *pb.OrderResult) (float64, *pb.Empty, error) {
	emailServiceClient := pb.NewEmailServiceClient(fe.emailSvcConn)

	start := time.Now()
	resp, err := emailServiceClient.SendOrderConfirmation(ctx, &pb.SendOrderConfirmationRequest{
		Email: email,
		Order: order})
	elapsed := time.Since(start)

	return elapsed.Seconds(), resp, err
}
