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
            "message": "user added successfully",
            "data": null,
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