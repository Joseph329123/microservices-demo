# This file generates a csv file in the current directory which
# contains all combinations of
# latency, traffic, shipping_service_get_quote_response_return_types, cart_service_cart_return_types
# e.g. 0.04, 80, currency DNE, exception

import itertools
import csv

# latency in seconds
latency = [0, 0.04, 0.08, 0.4, 0.8, 1.2, 2]

# of simulated users
traffic = [1, 40, 80, 400, 1000]

#---------------------------------------------------------------------------------

# shipping service: "GET QUOTE RESPONSE" (response to "GET QUOTE REQUEST")
shipping_service_get_quote_response_return_types = ['default', 'currency DNE', 'negative shipping quote value', 'error']

#---------------------------------------------------------------------------------

# cart service: "CART" (response to "GET CART REQUEST")
cart_service_cart_return_types = ['default', 'bad product id', 'exception']

#---------------------------------------------------------------------------------

# header of csv file
header = ['latency', 'traffic', 'shipping service get quote response', 'cart service cart return type']

# get all combinations of latency, traffic, shipping_service_get_quote_response_return_types, cart_service_cart_return_types
lines = list(itertools.product(latency, traffic, shipping_service_get_quote_response_return_types, cart_service_cart_return_types))

with open('environments.csv','wb') as out:
    csv_out=csv.writer(out)
    csv_out.writerow(header)
    for line in lines:
    	for i in range(15):
        	csv_out.writerow(line)


# currency service: "MONEY" (response to "CURRENCY CONVERSION REQUEST")

# 0 = default implementation
# 1 = return negative value
# 2 = return incorrect currency
# 3 = return currency code that does not exist
# 4 = return error
# currency_service_money_return_types = [0, 1, 2, 3, 4]
