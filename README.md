# Basic Go Rest API
Uses the Go package Gin to implement a simple Rest API.

This example creates a part of a meeting app where users can arrange a meeting with a date, start and end time, as well as meeting title and description.

To start the Rest API, use ``go run .`` on the terminal.

# Demonstration
All demonstrations were done on Postman.

## Get all meetings
Send a GET request to ``localhost:8080/meetings/`` to get all meetings.

![image](https://github.com/user-attachments/assets/54b1ebfd-ce16-4bfe-b6cf-217818f5e90a)

## Get a certain meeting
Send a GET request to ``localhost:8080/meetings/{id}`` to get a certain meeting.

![image](https://github.com/user-attachments/assets/773420b7-4433-4e3f-82e3-6cc00e518446)

## Create a new meeting
Send a POST request to ``localhost:8080/meetings/`` to create a new meeting.

![image](https://github.com/user-attachments/assets/0eaabec2-222c-4554-8f6a-c6ce1fef07d3)

## Edit an existing meeting
Send a PUT request to ``localhost:8080/meetings/{id}`` to edit an existing meeting.

![image](https://github.com/user-attachments/assets/6247fe45-98fd-432f-a708-931377f1908a)

## Delete a meeting
Send a DELETE request to ``localhost:8080/meetings/{id}`` to delete a meeting.

![image](https://github.com/user-attachments/assets/df52dc1e-1e77-436e-be94-4a7420b5fdcd)

![image](https://github.com/user-attachments/assets/c8bd69ec-4ebb-4bad-998c-21e4b48648a2)
