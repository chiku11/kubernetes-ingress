## SETUP KUBERNETES DEPLOYMENT
sh orchestrate.sh

## CURL COMMAND to test
curl --location --request POST 'http://localhost:80/findlocation' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "siri",
    "age": 30
}'