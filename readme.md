# Payroll System API

A backend payroll system written in Go using the Gin framework and PostgreSQL. This API enables employees to submit attendance, overtime, and reimbursement requests, and allows admins to generate and retrieve payslips based on a defined payroll period.

---

## Features

- JWT-based authentication for Admin and Employee roles
- Attendance submission with one-entry-per-day constraint
- Overtime submission (up to 3 hours/day)
- Reimbursement request handling
- ayroll period creation and automated payslip generation
- Employee payslip retrieval and admin summary reports
- Audit logging of create/update actions


---

## Tech Stack

| Layer         | Technology            |
|---------------|------------------------|
| Language      | Go 1.21+               |
| Framework     | [Gin](https://github.com/gin-gonic/gin) |
| ORM           | [GORM](https://gorm.io) |
| Database      | PostgreSQL             |
| Auth          | JWT                    |
| Docs/Test     | Postman, Swagger YAML  |

---

## Setup Instructions

### 1. Clone the Project
```bash
git clone https://github.com/yourusername/payroll-system.git
cd payroll-system

```

### 2. Database setup (psql)
```bash
CREATE DATABASE payroll;
\c payroll
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

### 3. Install Dependencies
```bash
go mod tidy
```

### 4. Run the Application
```bash
go run main.go
```

### 5. Access the API
Server runs on: http://localhost:8080

