# Microservices Demo Application

This application demonstrates the use of microservices for a fictional e-commerce application. The application is composed of multiple microservices written in different languages that talk to each other over gRPC.

## Architecture

**Online Boutique** is composed of 12 microservices that talk to each other over gRPC.

[![Architecture diagram](./docs/img/architecture-diagram.png)](./docs/img/architecture-diagram.png)

| Service | Language | Description |
| --- | --- | --- |
| [frontend](./src/frontend) | Go | Exposes an HTTP server to serve the website |
| [cartservice](./src/cartservice) | C# | Stores the items in the user's shopping cart |
| [productcatalogservice](./src/productcatalogservice) | Go | Provides the list of products from a JSON file and ability to search products and get individual products |
| [currencyservice](./src/currencyservice) | Node.js | Converts one money amount to another currency |
| [paymentservice](./src/paymentservice) | Node.js | Charges the given credit card info (mock) |
| [shippingservice](./src/shippingservice) | Go | Gives shipping cost estimates based on the shopping cart |
| [emailservice](./src/emailservice) | Python | Sends users an order confirmation email (mock) |
| [checkoutservice](./src/checkoutservice) | Go | Retrieves user cart, prepares order and orchestrates the payment, shipping and the email notification |
| [recommendationservice](./src/recommendationservice) | Python | Recommends other products based on what's in the user's cart |
| [adservice](./src/adservice) | Java | Provides text ads based on given context words |
| [loadgenerator](./src/loadgenerator) | Python | Continuously sends requests imitating realistic user shopping flows |
| [shoppingassistantservice](./src/shoppingassistantservice) | Python | AI assistant for product suggestions |

## CI/CD Pipeline

This project includes a comprehensive CI/CD pipeline using GitHub Actions:

- **CI Pipeline**: Runs on every push and PR to validate code quality, security, and functionality
- **Security Scanning**: Weekly security scans for vulnerabilities
- **Build and Publish**: Creates versioned Docker images for all microservices

## Local Development

See [LOCAL_DEVELOPMENT.md](./LOCAL_DEVELOPMENT.md) for instructions on how to run the application locally.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](./LICENSE) file for details.