# This program merges environments.csv and results.csv into one csv

import csv

merged = []

with open('generated_csv/environments.csv') as csvfile:
	readCSV = csv.reader(csvfile, delimiter=',')
	for row in readCSV:
		latency = row[0]
		traffic = row[1]
		shipping_service_get_quote_response = row[2]
		cart_service_cart_return_type = row[3]
		merged.append([latency, traffic, shipping_service_get_quote_response, cart_service_cart_return_type])

k = 0

with open('output/results_unformatted.csv') as csvfile:
	readCSV = csv.reader(csvfile, delimiter=',')
	for row in readCSV:
		service = row[0]
		best_case_response_time = row[1]
		worst_case_response_time = row[2]
		average_response_time = row[3]
		std_dev_response_time = row[4]
		availability = row[5]
		merged[k].insert(0, service)
		merged[k].append(best_case_response_time)
		merged[k].append(worst_case_response_time)
		merged[k].append(average_response_time)
		merged[k].append(std_dev_response_time)
		merged[k].append(availability)
		k += 1

with open("./output/results_formatted.csv", "wb") as f:
    writer = csv.writer(f)
    writer.writerows(merged)



