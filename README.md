#  Generation Next: BlackOut Project

Welcome to the BlackOut project! This project aims to tackle communication challenges during power outages or other situations where mobile GSM networks are not available. By leveraging public Wi-Fi networks with UPS (Uninterruptible Power Supply), users can report incidents through our app. When multiple incidents are reported within a certain radius, the system determines that a large-scale disaster has occurred and automatically triggers an email notification.

## Table of Contents

- [Overview](#overview)
- [Implementation](#implementation)
  - [Frontend](#frontend)
  - [Backend](#backend)
  - [Incident Classification Service](#incident-classification-service)
- [Getting Started](#getting-started)
- [Contributing](#contributing)
- [License](#license)

## Overview

The project consists of three main components:

1. A user-friendly mobile app for reporting incidents
2. A backend server for storing and managing incidents
3. A service that classifies incidents based on their proximity to each other

## Implementation

### Frontend

The frontend of our app is developed using [App Inventor](https://appinventor.mit.edu). You can access the App Inventor project files [here](https://mega.nz/file/lPJnEJQS#SHlIX1TRPDJamCv09sjcrANbnw3qH2EqbEWfVhHFZVA).

### Backend

The backend is implemented using the [Gin Web Framework](https://github.com/gin-gonic/gin) in Go, and MySQL is used as the database. The server is responsible for receiving incident reports from the frontend, storing them in the database, and managing the data.

### Incident Classification Service

The Incident Classification Service is a separate microservice implemented in Python using the [FastAPI](https://fastapi.tiangolo.com) framework. It uses the Haversine formula to determine the proximity of reported incidents [repo](https://github.com/avemoi/incident_clustering). If multiple incidents are reported within a specified radius, the service classifies the situation as a large-scale disaster and triggers an email notification.

## Getting Started

To get started with the project, please follow the instructions below:

1. Clone the repository: `git clone https://github.com/your-username/generation-next-communication.git`
2. Run `make prod`.
3. Run `./api`
