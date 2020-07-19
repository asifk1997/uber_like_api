# Uber like Api
Uber like Mini  Api To GetNearByCabs, View all your bookings And Do a Booking.<space><space>

First of All Create a database in postgres and put its credentials in .env file 

My .env file looks like this<space><space>
  <space><space>

API_SECRET= #Used when creating a JWT. It can be anything<space><space>
  
DB_HOST=yourdbip<space><space>
  
DB_DRIVER=postgres<space><space>
  
DB_USER=yourdbusername<space><space>
  
DB_PASSWORD=yourpassword<space><space>
  
DB_NAME=yourdbname<space><space>
  
DB_PORT=5432 #Default postgres port<space><space>
  
<space><space>
The .env file should be in top most directory of repository.<space><space>

To run clone the project and type ![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+) `go run main.go` in terminal.<space><space>

First of All we need to login :<space><space>

![](screenshots/Screenshot_78.png)

Put the JWT Token to the Authorization header and you are logged in.


![](screenshots/Screenshot_80.png)


Now comes fun part you can view all the users 

![](screenshots/Screenshot_81.png)


Create A New User

![](screenshots/Screenshot_83.png)

Update A User

![](screenshots/Screenshot_82.png)

Delete A User

![](screenshots/Screenshot_84.png)


Get Nearby Cabs For <space><space>
This you need to give you need to give your current location and NearByRange in kilometers.
  
![](screenshots/Screenshot_86.png)
![](screenshots/Screenshot_87.png)
![](screenshots/Screenshot_88.png)
![](screenshots/Screenshot_89.png)


Get All Bookings of a User (Only logged in user can access)

![](screenshots/Screenshot_90.png)
![](screenshots/Screenshot_94.png)

Create A New Booking (Only logged in user can book)

![](screenshots/Screenshot_92.png)
