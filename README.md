# Catalog

The catalog application is a demonstration application used for showcasing the Okteto divert feature set.
The application uses canned data to mimic a service catalog.
The service catalog tracks services, their owners, and health information.

The original deployment contains only the most recent health data for each service in the catalog.
This is helpful but it could be better.

A developer decides that health data for each service would be more helpful if it contained historical data.
In this scenario the developer adds a data store and changes the health-checking service to provide more data.
The developer uses the divert feature to develop the new feature side-by-side with the original application.

Head on over to the [tutorial](https://okteto.com/docs/tutorials/getting-started-with-divert) and get started!
