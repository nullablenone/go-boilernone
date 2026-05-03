# Go Boilernone

A simple Golang boilerplate using the Gin Web Framework and GORM (PostgreSQL). 
## Usage 

This boilerplate uses the official `gonew` tool to generate a new project. 

This automatically renames the module paths in `go.mod` and all `.go` files to match your new project name, and provides a clean directory without the original Git history. 

1. **Install `gonew`** (if you haven't already):
   
   ```bash
   go install golang.org/x/tools/cmd/gonew@latest
   
2. **Generate your new project**: Replace `github.com/yourusernamegithub/your-project-name` with your actual desired module name, then move into the new directory.
    
    Bash
    
    ```
    gonew github.com/nullablenone/go-boilernone github.com/yourusernamegithub/your-project-name
    cd your-project-name
    ```
    
3. **Initialize Git** for your new project (optional but recommended):
    
    Bash
    
    ```
    git init
    git add .
    git commit -m "chore: initial commit from boilerplate"
    ```
    
4. **Set up Environment Variables**: Create a `.env` file in the root directory of your project and configure your PostgreSQL database credentials:
    
    Cuplikan kode
    
    ```
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASS=your_database_password
    DB_NAME=your_database_name
    DB_PORT=5432
    DB_SSLMODE=disable
    ```
    
5. **Install Dependencies & Run the Server**: Download all required modules and start the application:
    
    Bash
    
    ```
    go mod tidy
    go run main.go
    ```
    

The application will start running at `http://localhost:3333`.
