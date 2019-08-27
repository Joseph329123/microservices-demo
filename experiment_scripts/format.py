# This program takes the output/results_formatted.csv file and "joins" all services
# so the measures of performance and availability for a single environment config
# is on one line

import csv

merged = []

with open('results_formatted.csv') as csvfile:
	readCSV = csv.reader(csvfile, delimiter=',')

	# skip first line
	next(readCSV)

	k = 0

	for row in readCSV:
		latency = row[1]
		traffic = row[2]
		shipping_service_get_quote_response = row[3]
		cart_service_cart_return_type = row[4]
		best_case_response_time = row[5]
		worst_case_response_time = row[6]
		average_response_time = row[7]
		std_dev_response_time = row[8]
		availability = row[9]
		if (k % 15 == 0):
			merged.append([latency, traffic, shipping_service_get_quote_response,
				cart_service_cart_return_type, best_case_response_time,
				worst_case_response_time, average_response_time,
				std_dev_response_time, availability])
		else:
			merged[k//15].append(best_case_response_time)
			merged[k//15].append(worst_case_response_time)
			merged[k//15].append(average_response_time)
			merged[k//15].append(std_dev_response_time)
			merged[k//15].append(availability)
		k += 1

header = ['latency', 'traffic', 'shipping_service_get_quote_response', 'cart_service_cart_return_type',
'Best_Case_Response_Time_Product_Catalogue_Service_Empty_Request',
'Worst_Case_Response_Time_Product_Catalogue_Service_Empty_Request',
'Average_Case_Response_Time_Product_Catalogue_Service_Empty_Request',
'Std_Dev_Response_Time_Product_Catalogue_Service_Empty_Request',
'Availability_Product_Catalogue_Service_Empty_Request',
'Best_Case_Response_Time_Product_Catalogue_Service_Get_Product_Request',
'Worst_Case_Response_Time_Product_Catalogue_Service_Get_Product_Request',
'Average_Case_Response_Time_Product_Catalogue_Service_Get_Product_Request',
'Std_Dev_Response_Time_Product_Catalogue_Service_Get_Product_Request',
'Availability_Product_Catalogue_Service_Get_Product_Request',
'Best_Case_Response_Time_Product_Catalogue_Service_Search_Product_Request',
'Worst_Case_Response_Time_Product_Catalogue_Service_Search_Product_Request',
'Average_Case_Response_Time_Product_Catalogue_Service_Search_Product_Request',
'Std_Dev_Response_Time_Product_Catalogue_Service_Search_Product_Request',
'Availability_Product_Catalogue_Service_Search_Product_Request',
'Best_Case_Response_Time_Recommendation_List_Service_Recommendation_Request',
'Worst_Case_Response_Time_Recommendation_List_Service_Recommendation_Request',
'Average_Case_Response_Time_Recommendation_List_Service_Recommendation_Request',
'Std_Dev_Response_Time_Recommendation_List_Service_Recommendation_Request',
'Availability_Recommendation_List_Service_Recommendation_Request',
'Best_Case_Response_Time_Checkout_Service_Place_Order_Request',
'Worst_Case_Response_Time_Checkout_Service_Place_Order_Request',
'Average_Case_Response_Time_Checkout_Service_Place_Order_Request',
'Std_Dev_Response_Time_Checkout_Service_Place_Order_Request',
'Availability_Checkout_Service_Place_Order_Request',
'Best_Case_Response_Time_Shipping_Service_Get_Quote_Request',
'Worst_Case_Response_Time_Shipping_Service_Get_Quote_Request',
'Average_Case_Response_Time_Shipping_Service_Get_Quote_Request',
'Std_Dev_Response_Time_Shipping_Service_Get_Quote_Request',
'Availability_Shipping_Service_Get_Quote_Request',
'Best_Case_Response_Time_Shipping_Service_Ship_Order_Request',
'Worst_Case_Response_Time_Shipping_Service_Ship_Order_Request',
'Average_Case_Response_Time_Shipping_Service_Ship_Order_Request',
'Std_Dev_Response_Time_Shipping_Service_Ship_Order_Request',
'Availability_Shipping_Service_Ship_Order_Request',
'Best_Case_Response_Time_Currency_Service_Currency_Conversion_Request',
'Worst_Case_Response_Time_Currency_Service_Currency_Conversion_Request',
'Average_Case_Response_Time_Currency_Service_Currency_Conversion_Request',
'Std_Dev_Response_Time_Currency_Service_Currency_Conversion_Request',
'Availability_Currency_Service_Currency_Conversion_Request',
'Best_Case_Response_Time_Currency_Service_Empty_Request',
'Worst_Case_Response_Time_Currency_Service_Empty_Request',
'Average_Case_Response_Time_Currency_Service_Empty_Request',
'Std_Dev_Response_Time_Currency_Service_Empty_Request',
'Availability_Currency_Service_Empty_Request',
'Best_Case_Response_Time_Cart_Service_Add_Item_Request',
'Worst_Case_Response_Time_Cart_Service_Add_Item_Request',
'Average_Case_Response_Time_Cart_Service_Add_Item_Request',
'Std_Dev_Response_Time_Cart_Service_Add_Item_Request',
'Availability_Cart_Service_Add_Item_Request',
'Best_Case_Response_Time_Cart_Service_Get_Cart_Request',
'Worst_Case_Response_Time_Cart_Service_Get_Cart_Request',
'Average_Case_Response_Time_Cart_Service_Get_Cart_Request',
'Std_Dev_Response_Time_Cart_Service_Get_Cart_Request',
'Availability_Cart_Service_Get_Cart_Request',
'Best_Case_Response_Time_Cart_Service_Empty_Cart_Request',
'Worst_Case_Response_Time_Cart_Service_Empty_Cart_Request',
'Average_Case_Response_Time_Cart_Service_Empty_Cart_Request',
'Std_Dev_Response_Time_Cart_Service_Empty_Cart_Request',
'Availability_Cart_Service_Empty_Cart_Request',
'Best_Case_Response_Time_Ad_Service_Ad_Request',
'Worst_Case_Response_Time_Ad_Service_Ad_Request',
'Average_Case_Response_Time_Ad_Service_Ad_Request',
'Std_Dev_Response_Time_Ad_Service_Ad_Request',
'Availability_Ad_Service_Ad_Request',
'Best_Case_Response_Time_Payment_Service_Charge_Request',
'Worst_Case_Response_Time_Payment_Service_Charge_Request',
'Average_Case_Response_Time_Payment_Service_Charge_Request',
'Std_Dev_Response_Time_Payment_Service_Charge_Request',
'Availability_Payment_Service_Charge_Request',
'Best_Case_Response_Time_Email_Service_Send_Order_Confirmation_Request',
'Worst_Case_Response_Time_Email_Service_Send_Order_Confirmation_Request',
'Average_Case_Response_Time_Email_Service_Send_Order_Confirmation_Request',
'Std_Dev_Response_Time_Email_Service_Send_Order_Confirmation_Request',
'Availability_Email_Service_Send_Order_Confirmation_Request']

with open('new_results_formatted.csv','w') as out:
	csv_out=csv.writer(out)
	csv_out.writerow(header)
	for line in merged:
		csv_out.writerow(line)
