# Task Management Application

This is a full-stack task management application with a React frontend and a Go backend, using PostgreSQL for the database.

The application is live at the following URLs:

- **Frontend:** [https://tasktrackhub.netlify.app](https://tasktrackhub.netlify.app)
- **Backend:** [https://task-manager-ov06.onrender.com](https://task-manager-ov06.onrender.com)

## Local Setup
  
### Prerequisites

- [Go](https://golang.org/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Node.js and npm](https://nodejs.org/en/download/) (for the React frontend)

### Database Setup

1. Install PostgreSQL and create a new database.

2. Set up your database user and password, ensuring they have the necessary permissions to access and modify the database.

3. Update the backend `.env` file with your database credentials and other environment variables.

### Backend Setup

1. Clone the repository and navigate to the `backend` directory:

    ```bash
    git clone https://github.com/souravkumar0546/TaskManager
    cd task-manager/backend
    ```

2. Install dependencies:

    ```bash
    go mod download
    ```

3. Set up environment variables (e.g., in a `.env` file):

    ```env
    DB_HOST=your_database_host
    DB_USER=your_username
    DB_PASSWORD=your_password
    DB_NAME=your_database_name
    DB_PORT=your_database_port
    API_URL=your_backend_app_url
    APP_URL=your_frontend_app_url
    ```


4. Run the backend server:

    ```bash
    go run main.go
    ```

### Frontend Setup

1. Navigate to the `frontend` directory:

    ```bash
    cd task-manager/frontend
    ```

2. Install dependencies:

    ```bash
    npm install
    ```

3. Update Backend URL:
    - Navigate to `src/services/apiService.js`.
    - Locate and update the `API_URL` constant with your deployed backend URL.

4. Run the frontend development server:

    ```bash
    npm start
    ```

## Features

- User authentication
- Task management with CRUD operations
- Task filtering and searching capabilities
- User profiles with avatars
- Responsive design for both desktop and mobile devices
- Server-side validation and error handling

## Assumptions and Notes

- **New User Default Avatar:** Upon signup, each new user is assigned a default avatar. Users can change their avatars in the profile section.

## Troubleshooting

- **Third-Party Cookies:** If you encounter issues with logging in or fetching tasks, ensure that third-party cookies are enabled in your browser. This is necessary for the application to manage authentication correctly when accessing cross-site resources.
