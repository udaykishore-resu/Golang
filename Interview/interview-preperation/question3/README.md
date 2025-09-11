## Commontly used API types are mainly as below:

- Request-Response APIs

| API Type          | Example URL                         | Usage Case                      |
| ----------------- | ----------------------------------- | ------------------------------- |
| Path Parameter    | `/users/42`                         | Identify specific resource      |
| Query Parameter   | `/search?keyword=golang&page=2`     | Filters, sorting, pagination    |
| Matrix Parameter  | `/products;color=red;size=large`    | Extra metadata in path (rare)   |
| Subdomain         | `https://api.customer1.localhost:8080` | Multi-tenant SaaS, routing      |
| RESTful Hierarchy | `/users/42/orders`                  | Parent-child resource relations |

- Event Driven APIs
brew services start kafka --> to start the kafka
brew services stop kafka --> to stop the kafka