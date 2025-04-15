package router

import (
	"3d-print-inventory/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// ht3dprint
	ht3dprint := api.Group("/ht3dprint")
	ht3dprint.Post("/files", handler.HT3DPrintAddFile)

	// byUserId
	byUserId := api.Group("/users/:userId")
	byUserId.Get("/printers", handler.GetPrintersByUserID)
	byUserId.Get("/printers/active", handler.GetActivePrintersByUserID)
	byUserId.Post("/printers/active/progress", handler.GetProgressOfActivePrinterByUserID)
	byUserId.Get("/printers/history", handler.GetPrintHistory)
	byUserId.Get("/projects", handler.GetProjectsByUserID)
	byUserId.Post("/printers/activate/:activationKey", handler.ActivatePrinter)
	byUserId.Post("/printers/activate/:activationKey/status", handler.CheckActivationStatus)

	// Project
	project := api.Group("/projects")
	project.Get("/", handler.GetProjects)
	project.Post("/", handler.CreateProject)
	project.Put("/:projectId", handler.UpdateProject)
	project.Delete("/:projectId", handler.DeleteProject)

	projectById := project.Group("/:projectId")
	projectById.Get("/", handler.GetProjectByID)
	projectById.Get("/folders", handler.GetFolderByProjectID)

	// Folder
	folder := api.Group("/folders")
	folder.Post("/", handler.CreateNewFolder)
	folder.Put("/:folderId", handler.UpdateFolder)
	folder.Delete("/:folderId", handler.DeleteFolder)

	// File
	file := api.Group("/files")
	file.Post("/upload", handler.UploadFiles)
	file.Post("/", handler.AddFile)
	file.Put("/:fileId", handler.UpdateFile)
	file.Delete("/:fileId", handler.DeleteFile)

	// Printer
	printer := api.Group("/printers")
	printer.Post("/", handler.AddPrinter)

	printerById := printer.Group("/:printerId")

	// printerById.Get("/schedules/status", handler.PrinterStatus)
	// printerById.Post("/schedules/status", handler.PostStatus)

	printerById.Get("/activities/progress", handler.GetProgress)

	// Activity
	activity := api.Group("/activities")
	activity.Post("/", handler.CreateActivity)

}
