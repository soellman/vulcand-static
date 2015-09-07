# vulcand-static

Middleware to return direct responses from the middleware configuration.

Set the Status:

    # Return a 503
    etcdctl set /vulcand/frontends/f1/middlewares/s1 '{
       "Type": "static",
       "Middleware":{
           "Status":503}}'

Or the Status with Body:

    # Return a 503 with "sorry!"
    etcdctl set /vulcand/frontends/f1/middlewares/s1 '{
       "Type": "static",
       "Middleware":{
           "Status":503,
           "Body":"sorry!"}}'

Or the Status with a BodyWithHeaders if you'd like to set additional headers:

    # Return a 503 with a json response
    etcdctl set /vulcand/frontends/f1/middlewares/s1 '{
       "Type": "static",
       "Middleware":{
           "Status":503,
           "BodyWithHeaders":"Content-Type: application/json\n\n{\"message\": \"unavailable\"}"}}'

Note, the BodyWithHeaders must be formatted like a raw HTTP response with headers and body separated by two carriage returns.

Some inspiration taken from https://journal.paul.querna.org/articles/2009/08/24/downtime-page-in-apache/
