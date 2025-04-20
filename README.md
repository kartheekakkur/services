# SERVICE Application

This project provides api's to interact with services. It uses postgres (SQL) as the database keeping consistency and availability as top priority.


### Tools

1. docker-desktop
2. golang-migrate - Database migration tool that reads from sources and applies them in correct order to a database
3. sqlc - sqlc generates type-safe code from SQL
4. viper - reads the configuration files
5. gin - HTTP web framework

### Config

The `app.env` file provides the configuration file that takes in the database source, driver and server address. This configuration file is loaded using viper as an environment variables at runtime. When implementing this in production using k8s this can be passed in as secret using secrets manager(external-secrets) and then be mounted as volume.
