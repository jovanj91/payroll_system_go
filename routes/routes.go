package routes

import (
	"github.com/jovanj91/payroll_system_go/controllers"
	"github.com/jovanj91/payroll_system_go/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(middlewares.AuditMiddleware())

	r.POST("/login", controllers.Login)

	employee := r.Group("/")
	employee.Use(middlewares.AuthMiddleware(), middlewares.AuthEmployee())
	{
		employee.POST("/attendances", controllers.SubmitAttendance)
		employee.POST("/overtimes", controllers.SubmitOvertime)
		employee.POST("/reimbursements", controllers.SubmitReimbursement)
		employee.GET("/payslip/:id", controllers.GetPayslip)
	}

	admin := r.Group("/")
	admin.Use(middlewares.AuthMiddleware(), middlewares.AuthAdmin())
	{
		admin.POST("/payroll-periods", controllers.CreatePayrollPeriod)
		admin.POST("/run-payroll/:id", controllers.RunPayroll)
		admin.GET("/payslip-summary/:id", controllers.GetPayslipSummary)
	}
}
