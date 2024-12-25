# Restaurant Management Backend

This project is a backend service for managing restaurant operations, built with Go and the Gin framework.

## Features

- **User Management**: Handle user authentication and authorization.
- **Menu Management**: CRUD operations for menu items.
- **Food Management**: Manage food items and categories.
- **Order Management**: Process and track customer orders.
- **OrderItem Management**: Manage stock levels and supplies.
- **Invoice Management**: Manage payment and billing.
- **Table System**: Handle table reservations.

## Technologies Used

- **Go**: The primary programming language.
- **Gin**: Web framework for building the HTTP server.
- **MongoDB**: Database for storing application data.

## Getting Started

### Prerequisites

- Go (version 1.16 or later)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/0PrashantYadav0/Restaurant-Management-Backend.git
   cd Restaurant-Management-Backend
   ```

2. Install dependencies:
   ```sh
   go mod download
   ```

3. Set up the database:
   - Create a MySQL database named `restaurant_management`.
   - Add the credentials of project in `.env` file as given in [.env.example](./.env.example).

4. Run the application:
   ```sh
   go run main.go
   ```

### API Endpoints

- `POST /users/signup`: Register a new user.
- `POST /users/login`: Authenticate a user.
- `GET /users`: Retrieve all the user.
- `GET /users/:id`: Retrieve user details.

- `GET /menu`: Retrieve all menu items.
- `GET /menu/:id`: Retrieve menu item details.
- `POST /menu`: Add a new menu item.
- `PUT /menu/:id`: Update a menu item.

- `GET /foods`: Retrieve all food items.
- `GET /foods/:id`: Retrieve food item details.
- `POST /foods`: Add a new food item.
- `PUT /foods/:id`: Update a food item.

- `GET /orders`: Retrieve all orders.
- `GET /orders/:id`: Retrieve order details.
- `POST /orders`: Place a new order.
- `PUT /orders/:id`: Update an order.

- `GET /orderItems`: Retrieve all order items.
- `GET /orderItems/:id`: Retrieve order item details.
- `POST /orderItems`: Add a new order item.  
- `PUT /orderItems/:id`: Update an order item.

- `GET /invoices`: Retrieve all invoices.
- `GET /invoices/:id`: Retrieve invoice details.
- `POST /invoices`: Add a new invoice.
- `PUT /invoices/:id`: Update an invoice.

- `GET /tables`: Retrieve all tables.
- `GET /tables/:id`: Retrieve table details.
- `POST /tables`: Add a new table.
- `PUT /tables/:id`: Update a table.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss any changes or improvements.

## License

This project is licensed under the MIT License.