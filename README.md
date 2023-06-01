## Pay Pool
#### Pay Pool is an API service designed to facilitate the sharing of expenses and splitting bills among a group of people when dining together.

#### Get started
To get started with the Pay Pool API, you can follow the steps outlined below:
#### API Endpoints
1. **POST /api/v1/user/register**
This endpoint allows you to register a new user account.
    - Request Body:
        ```JSON
        {
            "username": "jose",
            "password": "password",
            "email": "jose@gmail.com"
        }
        ```
    - Response:
        ```JSON
        {
            "success": true,
            "message": "user added successfully",
            "data": null,
        }
        or
        {
            "success": false,
            "message": "email already exists",
            "data": null
        }

        ```
2. **POST /api/v1/user/login**
This endpoint allows you to get the token for the authenticated.
    - Request Body
        ```JSON
        {
            "email": "jose@gmail.com",
            "password": "password"
        }
        ```
    - Response:
        ```JSON
        {
            "success": true,
            "message": "login successfully",
            "data": {
                "token": "JWT token"
            }
        }
    ```
3. **GET auth/api/v1/user**
This endpoint allows you to get the info of the user
    - Request Header:
      "token": {token get from /login}
    - Request Parameters:
      /user?email={email}
    - Response:
        ```JSON
        {
            "success": true,
            "message": "get user successfully",
            "data": {
                "id": 4,
                "username": "adrian",
                "email": "adrian@gmail.com"
            }
        }
        ```