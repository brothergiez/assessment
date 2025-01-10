# Group Anagram, Booking System & Calculate Logic

This project provides utilities for solving common problems including:
- Anagram Grouping: Group words that are anagrams of each other.
- Booking System: Manage room bookings with time conflict checks.
- Calculate Logic: Calculate the total cost of items based on price and quantity.
---

## Prerequisites
Before running the project, ensure you have the following installed:
- Go (version 1.20 or later)
- Git

--- 

## Setup 
### 1. Clone the Repository
```sh
git clone https://github.com/brothergiez/assessment.git
cd assessment
```

### 2. Install Dependency
```sh
go mod tidy
```
---

## How to Use
This project includes three functionalities: Anagram, Booking System, and Calculate Logic. Each feature has its own command-line interface.

### 1. Group Anagram
Description: Groups words that are anagrams of each other.

#### Run the Anagram Utility:
```sh
go run main.go anagram
```

#### Output :
```sh
[["listen","silent","enlist"],["google","gogole"],["cat","act"]]
```

### 2. Booking System
Description: Manage room bookings, checking for conflicts and ensuring room capacity.
#### Run the Booking System:
```sh
go run main.go booking
```

#### Output: 
```sh
Memesan room1 mulai jam 10:00 selama 2 jam
Error: booking conflict detected
```

### 3. Calculate Logic
Description: Calculate the total cost of items based on price and quantity.
#### Run the Calculate Logic:
```sh
go run main.go calculate
```

#### Output: 
```sh
Total Price: 35
```
