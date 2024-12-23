package main

import (
	"flag"          // Used for parsing command-line flags
	"net/http"      // Provides HTTP client and server implementations
	"net/url"       // Used for URL manipulation
	"path/filepath" // For working with file paths
	"time"

	"github.com/gin-gonic/gin" // Gin framework for HTTP web services
	"github.com/google/uuid"   // To generate unique identifiers
)

/*
 * Main entry point of the application
 */
func main() {
	// Flag to enable or disable debug mode
	debugMode := flag.Bool("debug", false, "Enable or disable debug mode. If true, additional logging or debugging output will be provided.")
	flag.BoolVar(debugMode, "d", false, "Enable or disable debug mode. If true, additional logging or debugging output will be provided.")

	// Flag for setting the server port
	serverPort := flag.String("port", "3000", "The port number the server will listen on. Default is 3000.")
	flag.StringVar(serverPort, "p", "3000", "The port number the server will listen on. Default is 3000.")

	// Flag for setting the server hostname or IP
	serverHost := flag.String("host", "0.0.0.0", "The hostname or IP address of the server. Default is localhost.")
	flag.StringVar(serverHost, "h", "0.0.0.0", "The hostname or IP address of the server. Default is localhost.")

	// Parse the command-line flags
	flag.Parse()

	// Build the server address from the hostname and port
	serverAddr := *serverHost + ":" + *serverPort

	// Create a new Gin engine instance
	r := gin.Default()

	// Set the Gin mode based on the debug flag
	if *debugMode {
		gin.SetMode(gin.DebugMode) // Enable debug mode
	} else {
		gin.SetMode(gin.ReleaseMode) // Use release mode
	}

	// Load HTML templates and static assets
	r.LoadHTMLGlob("./public/templates/*") // Load all HTML templates from the specified directory
	r.Static("/assets", "./public/assets") // Serve static files like CSS, JS, and images

	// Define a handler for the root endpoint
	r.GET("/", func(c *gin.Context) {
		// Determine the link scheme (http or https) based on the request
		linkScheme := "http"
		if c.Request.TLS != nil {
			linkScheme = "https"
		}

		// Build the home link
		link := linkScheme + "://" + c.Request.Host

		// Render the HTML template with the home link
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Home": link,
		})
	})

	// Define a POST endpoint for file uploads
	r.POST("/api/upload", func(c *gin.Context) {
		// Retrieve the uploaded file
		file, _ := c.FormFile("file")

		// Generate a unique file name using UUID
		id := uuid.New()
		fileName := id.String() + filepath.Ext(file.Filename)

		// Determine the link scheme (http or https)
		linkScheme := "http"
		if c.Request.TLS != nil {
			linkScheme = "https"
		}

		// Build the file download link
		link := linkScheme + "://" + c.Request.Host + "/file/" + url.QueryEscape(fileName)

		// Handle case where no file was uploaded
		if file == nil {
			c.JSON(400, gin.H{"error": "No file uploaded"})
			return
		}

		// Save the uploaded file to the 'uploads' directory
		if err := c.SaveUploadedFile(file, "./uploads/"+fileName); err != nil {
			c.JSON(500, gin.H{"error": "Failed to save file"})
			return
		}

		// Return the download link and additional file details in the response
		c.JSON(200, gin.H{
			"id":            id,
			"file_name":     fileName,
			"original_name": file.Filename,
			"file_size":     file.Size,
			"file_type":     file.Header.Get("Content-Type"),
			"upload_time":   time.Now(),
			"link":          link,
		})
	})

	// Define a GET endpoint to serve uploaded files
	r.GET("/file/:filename", func(c *gin.Context) {
		// Retrieve the file name from the URL parameter
		filename := c.Param("filename")

		// Get the file extension
		fileExtension := filepath.Ext(filename)

		// Define a map of MIME types for common file extensions
		mimeTypes := map[string]string{
			"txt":  "text/plain",
			"jpg":  "image/jpeg",
			"jpeg": "image/jpeg",
			"png":  "image/png",
			"pdf":  "application/pdf",
			"json": "application/json",
			"mp4":  "video/mp4",
			"mp3":  "audio/mpeg",
			"zip":  "application/zip",
		}

		// Set the appropriate Content-Type header if the file extension is recognized
		if mimeType, exists := mimeTypes[fileExtension]; exists {
			c.Header("Content-Type", mimeType)
		} else {
			// Default Content-Type for unknown file types
			c.Header("Content-Type", "application/octet-stream")
		}

		// Serve the file from the 'uploads' directory
		c.File("./uploads/" + filename)
	})

	// Start the server on the specified address
	r.Run(serverAddr)
}
