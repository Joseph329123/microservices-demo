# generates the csv with its header (columns) inside ./output directory

rm -f results_unformatted.csv
touch results_unformatted.csv

echo "service, best case response time, worst case response time, average response time, std dev response time, availability" >> results_unformatted.csv
