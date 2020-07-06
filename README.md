# FarmSale API

## Setup
Use mongo uri format without the `srv` syntax due to a known issue with [mongo-driver](https://github.com/golang/go/issues/37362) and `srv` formatting with go versions >1.10

## Routes

### Credentials
1. Admin
 ```
 {
 "email": "admin@user.com",
 "password": "user12345"
 }
 ```
 2. User
 ```
 {
  "email": "user@user.com",
  "password": "user12345"
  }
 ```
 3. Agent
 ```
 {
 "email": "agent@user.com",
 "password": "user12345"
 }
 ```
 
 4. Manager
 ```
 {
 "email": "manager@user.com",
 "password": "user12345"
 }
 ```
 
 ### Public Routes
`GET https://farmsaledev.herokuapp.com/api/v1`

### User authenticated route
`GET https://farmsaledev.herokuapp.com/api/v1/auth/products`

### Admin Routes
`GET https://farmsaledev.herokuapp.com/api/v1/admin`

### Manager Routes
`GET https://farmsaledev.herokuapp.com/api/v1/manager`

### AgentRoutes
`GET https://farmsaledev.herokuapp.com/api/v1/agent`

## WORKING DEMO

https://farmsaledev.herokuapp.com/api/v1/
