# Toronto_TIme_Store
![](Aspose.Words.9003b68b-6c19-4208-9f71-019bc4c9ffb4.001.png)

![](Aspose.Words.9003b68b-6c19-4208-9f71-019bc4c9ffb4.002.jpeg)

![](Aspose.Words.9003b68b-6c19-4208-9f71-019bc4c9ffb4.003.png) In-Class Lab 6 ![](Aspose.Words.9003b68b-6c19-4208-9f71-019bc4c9ffb4.004.png)![](Aspose.Words.9003b68b-6c19-4208-9f71-019bc4c9ffb4.005.png)

**Student Name: Rahul Raju Pawar** **Student ID: 500227231** 

**Course Code: WINP2000** **Instructor Name: Maziar Sojoudian** 




**Building a Go API for Current Toronto Time with MySQL Database Logging** 

**Objectives:-**  

1. **Create a Go API Endpoint:** 
- Develop an endpoint that returns the current Toronto time in JSON format. 
2. **MySQL Database Integration:** 
- Connect to a MySQL database and store the time data for each API request. 
3. **Time Zone Handling:** 
- Ensure that the time returned is accurately adjusted to Toronto's timezone. 

**Git Hub url:[** https://github.com/rrpawar414/Toronto_TIme_Store ](https://github.com/rrpawar414/Toronto_TIme_Store)**

1. Installation  of MySQL

<img width="468" alt="image" src="https://github.com/user-attachments/assets/72f200a4-138c-484b-ba68-e2284abc5eb6">

![](Aspose.Words.9003b68b-6c19-4208-9f71-019bc4c9ffb4.006.jpeg)

2. Creating Database and Table: 
<img width="468" alt="image" src="https://github.com/user-attachments/assets/b3b4a141-aa76-4d98-9066-9f473a04ce64">

![](Aspose.Words.9003b68b-6c19-4208-9f71-019bc4c9ffb4.007.jpeg)

1\.Localhost:8080/current-time 


![](Aspose.Words.9003b68b-6c19-4208-9f71-019bc4c9ffb4.008.jpeg)
<img width="468" alt="image" src="https://github.com/user-attachments/assets/84e07bc8-36a6-4ccb-8523-bb1bb77c29db">


3. Timestamp logs are saving into the database.

![](Aspose.Words.9003b68b-6c19-4208-9f71-019bc4c9ffb4.009.jpeg)
<img width="468" alt="image" src="https://github.com/user-attachments/assets/8e0edb57-f2b5-471b-9f81-89dd6e95c8c0">

4. localhost:8080/time-logs 

![](Aspose.Words.9003b68b-6c19-4208-9f71-019bc4c9ffb4.010.jpeg)
<img width="469" alt="image" src="https://github.com/user-attachments/assets/303ce675-1423-4b80-8d43-78d1ec9ae47f">


5. Time-logs 

![](Aspose.Words.9003b68b-6c19-4208-9f71-019bc4c9ffb4.011.jpeg)
<img width="468" alt="image" src="https://github.com/user-attachments/assets/2db4b5ef-d6ee-4607-b3e5-9bd89e9f9eab">

**Bonus Challenges** 

- Implement logging in your Go application to log events and errors. 

![](Aspose.Words.9003b68b-6c19-4208-9f71-019bc4c9ffb4.012.jpeg)
<img width="468" alt="image" src="https://github.com/user-attachments/assets/db61e461-2f82-4bc0-b575-30b0336e1bb1">

- Create an additional endpoint to retrive all logged times from the database
<img width="468" alt="image" src="https://github.com/user-attachments/assets/fdcd304a-e642-4706-bb49-92e1cc217c35">


- Dockize your Go application and the MySQL database for easy deployment.
<img width="468" alt="image" src="https://github.com/user-attachments/assets/f15486fb-0221-4db3-8dfe-fc138f2c8baa">

![](Aspose.Words.9003b68b-6c19-4208-9f71-019bc4c9ffb4.014.jpeg)
