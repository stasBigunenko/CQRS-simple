# CQRS-simple

It is a pet project with implementing CQRS and Event-sourcing architect patterns.

The app using RabbitMQ as a messages broker and PostgreSQL as a storage with two tables and Redis as a cashe-storage.

Can be tried in the Docker-compose.

Scheme of the app:

Writing the data to db:

http: request for writing -> sending the message to the RabbitMQ
dbConsumer: reading from queue -> writing to PostgreSQL
                               -> creating queue to change the data in cache storage, if cache storage have these data
cacheConsumer: reading from queue -> and update data in cache

Reading data:

http: request for reading -> read from cache db (Redis) if there is no data than read the data from PostgreSQL and write to cache -> return data to the user.