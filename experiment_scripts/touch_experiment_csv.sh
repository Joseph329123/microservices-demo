# generates the csv with its header (columns) inside ./output directory

rm -f ./output/test_service_measurements.csv
touch ./output/test_service_measurements.csv

echo "service, best response time, worst response time, average response time, std dev response time, availability" >> ./output/test_service_measurements.csv 

