# API Specification

## Endpoints

### User

| Method  | Endpoint          | Description    | Jump                    |
| ------- | ----------------- | -------------- | ----------------------- |
| `POST`  | `/users/register` | Register user  | [Here](#register-user)  |
| `POST`  | `/users/login`    | Login user     | [Here](#login-user)     |
| `PATCH` | `/users/topup`    | Top up balance | [Here](#top-up-balance) |

### Category

| Method   | Endpoint                  | Description           | Jump                           |
| -------- | ------------------------- | --------------------- | ------------------------------ |
| `POST`   | `/categories`             | Create category       | [Here](#create-category)       |
| `GET`    | `/categories`             | Get categories        | [Here](#get-categories)        |
| `PATCH`  | `/categories/:categoryId` | Edit category by ID   | [Here](#edit-category-by-id)   |
| `DELETE` | `/categories/:categoryId` | Delete category by ID | [Here](#delete-category-by-id) |

### Product

| Method   | Endpoint               | Description          | Jump                          |
| -------- | ---------------------- | -------------------- | ----------------------------- |
| `POST`   | `/products`            | Create product       | [Here](#create-product)       |
| `GET`    | `/products`            | Get products         | [Here](#get-products)         |
| `PUT`    | `/products/:productId` | Edit product by ID   | [Here](#edit-product-by-id)   |
| `DELETE` | `/products/:productId` | Delete product by ID | [Here](#delete-product-by-id) |

### Transaction

| Method | Endpoint                          | Description           | Jump                           |
| ------ | --------------------------------- | --------------------- | ------------------------------ |
| `POST` | `/transactions`                   | Create transaction    | [Here](#create-transaction)    |
| `GET`  | `/transactions/my-transactions`   | Get my transactions   | [Here](#get-my-transactions)   |
| `GET`  | `/transactions/user-transactions` | Get user transactions | [Here](#get-user-transactions) |

## Endpoints Detail

### User

#### Register User

- Endpoint: `/users/register`
- Method: `POST`
- Request Body:

  ```json
  {
    "full_name": "string",
    "email": "string",
    "password": "string"
  }
  ```

- Response:

  - Code: `201`
  - Body:

    ```json
    {
      "id": "integer",
      "full_name": "string",
      "email": "string",
      "password": "string",
      "balance": "integer",
      "created_at": "date"
    }
    ```

#### Login User

- Endpoint: `/users/login`
- Method: `POST`
- Request Body:

  ```json
  {
    "email": "string",
    "password": "string"
  }
  ```

- Response:

  - Code: `200`
  - Body:

    ```json
    {
      "token": "jwt"
    }
    ```

#### Top Up Balance

- Endpoint: `/users/topup`
- Method: `PATCH`
- Request Body:

  ```json
  {
    "balance": "integer"
  }
  ```

- Response:

  - Code: `200`
  - Body:

    ```json
    {
      "message": "Your balance has been successfully updated to Rp <balance>"
    }
    ```

### Category

#### Create category

- Endpoint: `/categories`
- Method: `POST`
- Request Body:

  ```json
  {
    "type": "string"
  }
  ```

- Response:

  - Code: `201`
  - Body:

    ```json
    {
      "id": "integer",
      "name": "string",
      "created_at": "date"
    }
    ```

#### Get categories

- Endpoint: `/categories`
- Method: `GET`
- Headers:

  - `Authorization`: `Bearer <token>`

- Response:

  - Code: `200`
  - Body:

    ```json
    [
      {
        "id": "integer",
        "type": "string",
        "sold_product_amount": "integer",
        "created_at": "date",
        "updated_at": "date",
        "products": [
          {
            "id": "integer",
            "title": "string",
            "price": "integer",
            "stock": "integer",
            "created_at": "date",
            "updated_at": "date"
          },
          { ... }
        ]
      }
    ]
    ```

#### Edit category by ID

- Endpoint: `/categories/:categoryId`
- Method: `PATCH`
- Headers:
  - `Authorization`: `Bearer <token>`
- Request Body:

  ```json
  {
    "type": "string"
  }
  ```

- Response:

  - Code: `200`
  - Body:

    ```json
    {
      "id": "integer",
      "type": "string",
      "sold_product_amount": "integer",
      "updated_at": "date"
    }
    ```

#### Delete category by ID

- Endpoint: `/categories/:categoryId`
- Method: `DELETE`
- Headers:
  - `Authorization`: `Bearer <token>`
- Response:

  - Code: `200`
  - Body:

    ```json
    {
      "message": "Category has been successfully deleted"
    }
    ```

### Product

#### Create product

- Endpoint: `/products`
- Method: `POST`
- Headers:
  - `Authorization`: `Bearer <token>`
- Request Body:

  ```json
  {
    "title": "string",
    "price": "integer",
    "stock": "integer",
    "category_id": "integer"
  }
  ```

- Response:

  - Code: `201`
  - Body:

    ```json
    {
      "id": "integer",
      "title": "string",
      "price": "integer",
      "stock": "integer",
      "category_id": "integer",
      "created_at": "date"
    }
    ```

#### Get products

- Endpoint: `/products`
- Method: `GET`
- Headers:
  - `Authorization`: `Bearer <token>`
- Response:
  - Code: `200`
  - Body:
    ```json
    [
      {
        "id": "integer",
        "title": "string",
        "price": "integer",
        "stock": "integer",
        "category_id": "integer",
        "created_at": "date",
      },
      { ... }
    ]
    ```

#### Edit product by ID

- Endpoint: `/products/:productId`
- Method: `PUT`
- Headers:
  - `Authorization`: `Bearer <token>`
- Request Body:

  ```json
  {
    "title": "string",
    "price": "integer",
    "stock": "integer",
    "category_id": "integer"
  }
  ```

- Response:
  - Code: `200`
  - Body:
    ```json
    {
      "product": {
        "id": "integer",
        "title": "string",
        "price": "integer",
        "stock": "integer",
        "category_id": "integer",
        "created_at": "date",
        "updated_at": "date"
      }
    }
    ```

#### Delete product by ID

- Endpoint: `/products/:productId`
- Method: `DELETE`
- Headers:
  - `Authorization`: `Bearer <token>`
- Response:
  - Code: `200`
  - Body:
    ```json
    {
      "message": "Product has been successfully deleted"
    }
    ```

### Transaction

#### Create transaction

- Endpoint: `/transactions`
- Method: `POST`
- Headers:
  - `Authorization`: `Bearer <token>`
- Request Body:

  ```json
  {
    "product_id": "integer",
    "quantity": "integer"
  }
  ```

- Response:

  - Code: `201`
  - Body:

    ```json
    {
      "message": "You have successfully purchased the product",
      "transaction_bill": {
        "total_price": "integer",
        "quantity": "integer",
        "product_title": "string"
      }
    }
    ```

#### Get my transactions

- Endpoint: `/transactions/my-transactions`
- Method: `GET`
- Headers:
  - `Authorization`: `Bearer <token>`
- Response:

  - Code: `200`
  - Body:

    ```json
    [
      {
        "id": "integer",
        "product_id": "integer",
        "user_id": "integer",
        "quantity": "integer",
        "total_price": "integer",
        "product": {
          "id": "integer",
          "title": "string",
          "price": "integer",
          "stock": "integer",
          "category_id": "integer",
          "created_at": "date",
          "updated_at": "date"
        }
      }
    ]
    ```

#### Get user transactions

- Endpoint: `/transactions/user-transactions`
- Method: `GET`
- Headers:

  - `Authorization`: `Bearer <token>`

- Response:

  - Code: `200`

  - Body:
    ```json
    [
      {
        "id": "integer",
        "product_id": "integer",
        "user_id": "integer",
        "quantity": "integer",
        "total_price": "integer",
        "product": {
          "id": "integer",
          "title": "string",
          "price": "integer",
          "stock": "integer",
          "category_id": "integer",
          "created_at": "date",
          "updated_at": "date"
        },
        "user": {
          "id": "integer",
          "email": "string",
          "full_name": "string",
          "balance": "integer",
          "created_at": "date"
          "updated_at": "date"
        }
      }
    ]
    ```
