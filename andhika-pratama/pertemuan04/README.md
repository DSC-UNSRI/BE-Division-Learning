# Iftar Dashboard System

## Overview
This system is a simple backend dashboard application built with Go. It is designed for managing user data for iftar events. The application provides a console-based interface where a user can perform several tasks after authentication.

## Features
- **Authentication:**  
  The system uses environment variables stored in a `.env` file to verify credentials (Name, Email, Password). Users must enter the correct email and password to access the dashboard.

- **Vehicle Selection:**  
  The user can choose one vehicle from four available options:
  - Private Vehicle
  - Budget Bus
  - Hitch a Ride
  - Travel Car

- **Item Input:**  
  Users can input an unlimited number of items they intend to bring to iftar. The input process continues until the user types a specific keyword (e.g., "done") to finish.

- **Recommendation:**  
  Users can provide recommendations for the iftar event. Each recommendation requires a category (such as Film, Music, etc.) and a corresponding description.

- **Friend Registration:**  
  Users can register friends who will attend the iftar. For each friend, the system collects the friend's name and division.

- **View All Data:**  
  The dashboard allows the user to view all the recorded information including the selected vehicle, list of items, recommendations, and friends.

## Implementation Details
- The system is structured into multiple packages:
  - **controllers:** Handles the main dashboard logic.
  - **models:** Contains the data structures and operations for vehicles, items, recommendations, and friends.
  - **main.go:** Entry point of the application, managing authentication and invoking the dashboard loop.
  
- **Error Handling & Input Validation:**  
  The code includes basic error handling and input validation to ensure a smooth user experience.

- **Commit Conventions:**  
  The project follows Conventional Commits. For example:
  - `feat(vehicle): update selection to choose 1 vehicle from 4 options`
  - `docs(readme): add system explanation readme file`

## Getting Started
1. **Clone the Repository:**  
   Clone the forked repository to your local machine.

2. **Setup Environment Variables:**  
   Create a `.env` file in the root directory with the following keys:
   ```env
   NAMA="Aristel"
   EMAIL="aristel@gmail.com"
   PASSWORD="inipasswordcuy"
