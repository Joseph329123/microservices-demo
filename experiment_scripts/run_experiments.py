import os
import csv
import subprocess

with open('generated_csv/environments.csv') as csvfile:
    readCSV = csv.reader(csvfile, delimiter=',')

    # skip first line
    next(readCSV)

    prev_latency = ""
    prev_traffic = ""
    prev_shipping_service_get_quote_response = ""
    prev_cart_service_cart_return_type = ""

    for row in readCSV:
		latency = row[0]
		traffic = row[1]
		shipping_service_get_quote_response = row[2]
		cart_service_cart_return_type = row[3]

		if (latency == prev_latency and traffic == prev_traffic and
				shipping_service_get_quote_response == prev_shipping_service_get_quote_response and
				cart_service_cart_return_type == prev_cart_service_cart_return_type):
			continue

		print("(1/7) beginning to update latency")
		os.system("python update_latency.py " + latency)
		print("(2/7) beginning to update traffic")
		os.system("python update_traffic.py " + traffic)
		print("(3/7) beginning to update shipping service")
		os.system("python update_shipping_service_yaml.py " + shipping_service_get_quote_response)
		print("(4/7) beginning to update cart service")
		os.system("python update_cart_service_yaml.py " + cart_service_cart_return_type)
		print("(5/7) beginning to deploy pods")
		os.system("./deploy_pods_gke.sh")
		print("(6/7) beginning to copy csv")
		os.system("./copy_csv_testservice.sh")
		print("(7/7) beginning to destroy pods")
		os.system("./destroy_pods.sh")

		# read every 15th line
		prev_latency = latency
		prev_traffic = traffic
	 	prev_shipping_service_get_quote_response = shipping_service_get_quote_response
		prev_cart_service_cart_return_type = cart_service_cart_return_type

os.system("python merge.py")
