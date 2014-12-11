api_key="INSERT GOOGLE API KEY HERE"
url="https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=-33.8670522,151.1957362&radius=500&types=food&name=cruise&key=$api_key"

curl $url
