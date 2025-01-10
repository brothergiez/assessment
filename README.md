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

---

## Testing
### Run All Tests
This project includes comprehensive unit tests for all functionalities: Anagram, Booking System, and Calculate Logic.

Run all tests using the following command:
```sh
go test ./... -v
```

### Run Tests for Specific Modules
#### 1. Group Anagram :
```sh
go test ./anagram -v
```

#### 2. Booking System :
```sh
go test ./bookingsystem -v
```

### Test Output

```sh
=== RUN   TestMainAnagram
--- PASS: TestMainAnagram (0.41s)
=== RUN   TestMainBooking
--- PASS: TestMainBooking (0.41s)
=== RUN   TestMainCalculate
--- PASS: TestMainCalculate (0.42s)
=== RUN   TestMainInvalidCommand
--- PASS: TestMainInvalidCommand (0.42s)
PASS
ok      github.com/brothergiez/assessment       (cached)
=== RUN   TestGroupAnagram
=== RUN   TestGroupAnagram/Basic_anagram_grouping
=== RUN   TestGroupAnagram/Empty_input
=== RUN   TestGroupAnagram/No_anagrams
=== RUN   TestGroupAnagram/All_identical_words
--- PASS: TestGroupAnagram (0.00s)
    --- PASS: TestGroupAnagram/Basic_anagram_grouping (0.00s)
    --- PASS: TestGroupAnagram/Empty_input (0.00s)
    --- PASS: TestGroupAnagram/No_anagrams (0.00s)
    --- PASS: TestGroupAnagram/All_identical_words (0.00s)
PASS
ok      github.com/brothergiez/assessment/anagram       (cached)
=== RUN   TestAddRoom
--- PASS: TestAddRoom (0.00s)
=== RUN   TestBookRoom
--- PASS: TestBookRoom (0.00s)
=== RUN   TestBookingIntegration
--- PASS: TestBookingIntegration (0.00s)
PASS
ok      github.com/brothergiez/assessment/bookingsystem (cached)
=== RUN   TestCalculateTotal
=== RUN   TestCalculateTotal/Valid_input_with_string_price_and_quantity
=== RUN   TestCalculateTotal/Valid_input_with_integer_price_and_quantity
=== RUN   TestCalculateTotal/Mixed_valid_input_types
=== RUN   TestCalculateTotal/Missing_price_key
=== RUN   TestCalculateTotal/Invalid_price_type
=== RUN   TestCalculateTotal/Invalid_quantity_value
=== RUN   TestCalculateTotal/Empty_input
--- PASS: TestCalculateTotal (0.00s)
    --- PASS: TestCalculateTotal/Valid_input_with_string_price_and_quantity (0.00s)
    --- PASS: TestCalculateTotal/Valid_input_with_integer_price_and_quantity (0.00s)
    --- PASS: TestCalculateTotal/Mixed_valid_input_types (0.00s)
    --- PASS: TestCalculateTotal/Missing_price_key (0.00s)
    --- PASS: TestCalculateTotal/Invalid_price_type (0.00s)
    --- PASS: TestCalculateTotal/Invalid_quantity_value (0.00s)
    --- PASS: TestCalculateTotal/Empty_input (0.00s)
PASS
```