This is a repo for utilities built around fetching parsing and recording MDT route updates.

http://www.miamidade.gov/transit/WebServices/Buses/?RouteID=1

XML Result Values:

BusID
BusName
Latitude
Longitude
RouteID
TripID
Direction
ServiceDirection
Service
ServiceName
TripHeadsign
LocationUpdated

For the moment, the first round is going to be grabbing the xml and parsing it into separate values.

Then we will both stuff the output into a database table.
And pump out a JSON update.
