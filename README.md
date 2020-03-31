# Crawler-Storage-API

Interface for saving and accessing crawled website to/from S3 storage.

The idea behind it is that you use this as an access point to your S3 storage which is used for you web crawler to store website data.
To access the data on an efficent way, there will be also an interface for accessing and searching for the correct data in fast response time.
It is planned to do transactions through an gRPC layer, which than can be used by other applications.

Feel free to constribute.
