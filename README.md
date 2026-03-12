
# Go Boilernone

A simple Golang boilerplate using the Gin Web Framework and GORM (PostgreSQL).

## Usage

1. **Clone this repository** into your new project directory:
    Bash
    
    ```
    git clone https://github.com/nullablenone/go-boilernone.git your-project-name
    cd your-project-name
    ```

2. **Remove the existing Git history** so you can start fresh:
    
    Bash
    
    ```
    rm -rf .git
    ```
    
3. **Re-initialize Git** for your new project (optional but recommended):
    
    Bash
    
    ```
    git init
    ```
    
4. **Set up Environment Variables** Create a `.env` file in the root directory of your project and configure your PostgreSQL database credentials:
    
    Cuplikan kode
    
    ```
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASS=your_database_password
    DB_NAME=your_database_name
    DB_PORT=5432
    DB_SSLMODE=disable
    ```
    
5. **Install Dependencies & Run the Server** Download all required modules and start the application:
    
    Bash
    
    ```
    go mod tidy
    go run main.go
    ```
    

The application will start running at `http://localhost:3333`.
