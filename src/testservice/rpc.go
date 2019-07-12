// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"time"

	"fmt"

	pb "github.com/Joseph329123/microservices-demo/src/testservice/genproto"

	"github.com/pkg/errors"
)

const (
	avoidNoopCurrencyConversionRPC = false
)

/********************************************************************/
/* Test response times */
/********************************************************************/

/********************************************************************/
/* ProductCatalogueService */
/********************************************************************/

/* Send 'Empty' to ProductCatalogueService, receive 'ListProductsResponse' */
func (fe *frontendServer) timeProductCatalogueServiceEmptyRequest(ctx context.Context) (time.Duration, []*pb.Product, error) {
	fmt.Println("timeProductCatalogueServiceEmptyRequest()")

	start := time.Now()
	resp, err := pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn).
		ListProducts(ctx, &pb.Empty{})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}
	
	for err != nil {
		//fmt.Println("error")
		start = time.Now()
		resp, err = pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn).
		ListProducts(ctx, &pb.Empty{})
		elapsed = time.Since(start)
	}

	fmt.Println(elapsed)
	return elapsed, resp.GetProducts(), err
}

/* Send 'GetProductRequest' to ProductCatalogueService, receive 'Product' */
func (fe *frontendServer) timeProductCatalogueServiceGetProductRequest(ctx context.Context, id string) (time.Duration, *pb.Product, error) {
	fmt.Println("timeProductCatalogueServiceGetProductRequest()")

	start := time.Now()
	resp, err := pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn).
		GetProduct(ctx, &pb.GetProductRequest{Id: id})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}

	for err != nil {
		//fmt.Println("error")
		start = time.Now()
		resp, err = pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn).
		GetProduct(ctx, &pb.GetProductRequest{Id: id})
		elapsed = time.Since(start)
	}
	
	fmt.Println(elapsed)
	fmt.Println(resp.Name)
	return elapsed, resp, err
}

/* Send 'SearchProductsRequest' to ProductCatalogueService, receive 'SearchProductsResponse' */
func (fe *frontendServer) timeProductCatalogueServiceSearchProductsRequest(ctx context.Context, query string) (time.Duration, *pb.SearchProductsResponse, error) {
	fmt.Println("timeProductCatalogueServiceSearchProductsRequest()")

	start := time.Now()
	resp, err := pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn).
		SearchProducts(ctx, &pb.SearchProductsRequest{Query: query})
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println("no errors")
	}
	
	for err != nil {
		//fmt.Println("error")
		start = time.Now()
		resp, err = pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn).
		SearchProducts(ctx, &pb.SearchProductsRequest{Query: query})
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

	start := time.Now()
	resp, err := pb.NewRecommendationServiceClient(fe.recommendationSvcConn).ListRecommendations(ctx,
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
		resp, err = pb.NewRecommendationServiceClient(fe.recommendationSvcConn).ListRecommendations(ctx,
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

	start := time.Now()
	order, err := pb.NewCheckoutServiceClient(fe.checkoutSvcConn).
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
		order, err = pb.NewCheckoutServiceClient(fe.checkoutSvcConn).
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

	start := time.Now()
	quote, err := pb.NewShippingServiceClient(fe.shippingSvcConn).GetQuote(ctx,
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
		quote, err = pb.NewShippingServiceClient(fe.shippingSvcConn).GetQuote(ctx,
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

	start := time.Now()
	resp, err := pb.NewShippingServiceClient(fe.shippingSvcConn).ShipOrder(ctx, &pb.ShipOrderRequest{
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
		resp, err = pb.NewShippingServiceClient(fe.shippingSvcConn).ShipOrder(ctx, &pb.ShipOrderRequest{
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
func (fe *frontendServer) timeCurrencyServiceCurrencyConversionRequest(ctx context.Context, money *pb.Money, currency string) (time.Duration, *pb.Money, error) {
	start := time.Now()
	resp, err := pb.NewCurrencyServiceClient(fe.currencySvcConn).
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
		start = time.Now()
		resp, err = pb.NewCurrencyServiceClient(fe.currencySvcConn).
			Convert(ctx, &pb.CurrencyConversionRequest{
				From:   money,
				ToCode: currency})
		elapsed = time.Since(start)
	}

	fmt.Println(resp.GetCurrencyCode())
	fmt.Println(resp.GetUnits())
	fmt.Println(resp.GetNanos())
	fmt.Println(elapsed)
	return elapsed, resp, err
}

func (fe *frontendServer) getCurrencies(ctx context.Context) ([]string, error) {
	currs, err := pb.NewCurrencyServiceClient(fe.currencySvcConn).
		GetSupportedCurrencies(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}
	var out []string
	for _, c := range currs.CurrencyCodes {
		if _, ok := whitelistedCurrencies[c]; ok {
			out = append(out, c)
		}
	}
	return out, nil
}

func (fe *frontendServer) getProducts(ctx context.Context) ([]*pb.Product, error) {
	resp, err := pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn).
		ListProducts(ctx, &pb.Empty{})
	return resp.GetProducts(), err
}

func (fe *frontendServer) getProduct(ctx context.Context, id string) (*pb.Product, error) {
	resp, err := pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn).
		GetProduct(ctx, &pb.GetProductRequest{Id: id})
	return resp, err
}

func (fe *frontendServer) getCart(ctx context.Context, userID string) ([]*pb.CartItem, error) {
	resp, err := pb.NewCartServiceClient(fe.cartSvcConn).GetCart(ctx, &pb.GetCartRequest{UserId: userID})
	return resp.GetItems(), err
}

func (fe *frontendServer) emptyCart(ctx context.Context, userID string) error {
	_, err := pb.NewCartServiceClient(fe.cartSvcConn).EmptyCart(ctx, &pb.EmptyCartRequest{UserId: userID})
	return err
}

func (fe *frontendServer) insertCart(ctx context.Context, userID, productID string, quantity int32) error {
	_, err := pb.NewCartServiceClient(fe.cartSvcConn).AddItem(ctx, &pb.AddItemRequest{
		UserId: userID,
		Item: &pb.CartItem{
			ProductId: productID,
			Quantity:  quantity},
	})
	return err
}

func (fe *frontendServer) convertCurrency(ctx context.Context, money *pb.Money, currency string) (*pb.Money, error) {
	if avoidNoopCurrencyConversionRPC && money.GetCurrencyCode() == currency {
		return money, nil
	}
	return pb.NewCurrencyServiceClient(fe.currencySvcConn).
		Convert(ctx, &pb.CurrencyConversionRequest{
			From:   money,
			ToCode: currency})
}

func (fe *frontendServer) getShippingQuote(ctx context.Context, items []*pb.CartItem, currency string) (*pb.Money, error) {
	quote, err := pb.NewShippingServiceClient(fe.shippingSvcConn).GetQuote(ctx,
		&pb.GetQuoteRequest{
			Address: nil,
			Items:   items})
	if err != nil {
		return nil, err
	}
	localized, err := fe.convertCurrency(ctx, quote.GetCostUsd(), currency)
	return localized, errors.Wrap(err, "failed to convert currency for shipping cost")
}

func (fe *frontendServer) getRecommendations(ctx context.Context, userID string, productIDs []string) ([]*pb.Product, error) {
	resp, err := pb.NewRecommendationServiceClient(fe.recommendationSvcConn).ListRecommendations(ctx,
		&pb.ListRecommendationsRequest{UserId: userID, ProductIds: productIDs})
	if err != nil {
		return nil, err
	}
	out := make([]*pb.Product, len(resp.GetProductIds()))
	for i, v := range resp.GetProductIds() {
		p, err := fe.getProduct(ctx, v)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get recommended product info (#%s)", v)
		}
		out[i] = p
	}
	if len(out) > 4 {
		out = out[:4] // take only first four to fit the UI
	}
	return out, err
}

func (fe *frontendServer) getAd(ctx context.Context, ctxKeys []string) ([]*pb.Ad, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	resp, err := pb.NewAdServiceClient(fe.adSvcConn).GetAds(ctx, &pb.AdRequest{
		ContextKeys: ctxKeys,
	})
	return resp.GetAds(), errors.Wrap(err, "failed to get ads")
}
