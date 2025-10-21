# Commontly used API types are mainly as below:

## Request-Response APIs

| API Type          | Example URL                         | Usage Case                      |
| ----------------- | ----------------------------------- | ------------------------------- |
| Path Parameter    | `/users/42`                         | Identify specific resource      |
| Query Parameter   | `/search?keyword=golang&page=2`     | Filters, sorting, pagination    |
| Matrix Parameter  | `/products;color=red;size=large`    | Extra metadata in path (rare)   |
| Subdomain         | `https://api.customer1.localhost:8080` | Multi-tenant SaaS, routing      |
| RESTful Hierarchy | `/users/42/orders`                  | Parent-child resource relations |

## Event Driven APIs

The detailed explanation of four key technologies used to build Event-Driven APIs (EDAs) as follows:

**1. Apache Kafka (Event Streaming/Distributed Commit Log)**<br>
**2. Redis Pub/Sub (In-Memory Messaging)**<br>
**3. Webhooks (HTTP Callback)**<br>
**4. MQTT (Lightweight IoT Protocol)**


