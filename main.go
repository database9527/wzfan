package main

import (
	"fmt"
	"log"
	"gofibre-project/config"
	"gofibre-project/middleware" // For AccessControlMiddleware
	"gofibre-project/services"
	"gofibre-project/utils" // For SyncLocalDataToRedis
	"github.com/gofiber/fiber/v2" // Fiber import
	"github.com/flosch/pongo2/v6" // Pongo2 import
	// "time" // Required for CacheService if custom TTL is used in Set
	// "github.com/go-redis/redis/v8" // For the actual redis client if used directly in main
)

func main() {
	// --- Load Configuration ---
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}
	fmt.Printf("Loaded configuration: %+v\n", cfg) // For verification

	// --- Initialize Redis Service (Conceptual) ---
	// This would ideally be uncommented and used.
	// For now, keeping it conceptual to avoid runtime errors if Redis isn't available
	// in the test environment or if go-redis isn't correctly fetched.
	/*
	redisService, err := services.NewRedisService(services.RedisConfig{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if err != nil {
		log.Fatalf("Error initializing Redis service: %s", err)
	}
	log.Println("RedisService initialized (conceptually)")

	pageCacheRedisService, err := services.NewRedisService(services.RedisConfig{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.PageCacheDB,
	})
	if err != nil {
		log.Fatalf("Error initializing Redis service for page cache: %s", err)
	}
	log.Println("PageCache RedisService initialized (conceptually)")

	cacheService := services.NewCacheService(pageCacheRedisService, services.CacheConfig{
		TTLSeconds: cfg.PageCache.TTLSeconds,
		Enabled:    cfg.PageCache.Enabled,
		// Prefix: "pagecache:", // Example prefix
	})
	log.Println("CacheService initialized (conceptually)")
	*/

	// --- Sync Local Data to Redis (Conceptual) ---
	// This section demonstrates where data synchronization would occur.
	// It assumes a Redis client (redisService) is available and configured.
	// The actual redisService initialization is commented out above.
	// If redisService were available:
	/*
	if redisService != nil && cfg != nil { // Ensure redisService and cfg are not nil
		log.Println("Attempting to sync local data to Redis...")
		// The SyncLocalDataToRedis function expects a client that can operate on a specific DB.
		// If redisService.GetClient() returns a general client, SyncLocalDataToRedis handles
		// creating a temporary, DB-specific client from its options.
		err := utils.SyncLocalDataToRedis(
			redisService.GetClient(),      // Pass the base client
			"./config/data/",              // Path to data files
			"default_site",                // Placeholder siteID
			cfg.Redis.LocalDataDB,         // Target DB index from config
		)
		if err != nil {
			log.Printf("Error syncing local data to Redis: %v", err)
		} else {
			log.Println("Local data sync attempt finished.")
		}
	} else {
		log.Println("Skipping local data sync: Redis service or config not available (conceptually).")
	}
	*/

	// --- Initialize Fibre App ---
	app := fiber.New()

	// --- Register Access Control Middleware (Conceptual) ---
	// This middleware should be registered early, before static files or routes.
	// The check for cfg.AccessControl.Enabled is now done inside the middleware itself.
	/*
	if cfg != nil { // Ensure config is loaded
		aclMiddleware := middleware.NewAccessControlMiddleware(cfg.AccessControl)
		app.Use(aclMiddleware)
		log.Println("Access Control Middleware conceptually registered.")
	} else {
		log.Println("Skipping Access Control Middleware: Config not available (conceptually).")
	}
	*/

	// Serve static files from ./public directory
	// Requests to /css/style.css will serve ./public/css/style.css
	app.Static("/", "./public")

	// Basic root route using Pongo2 template
	app.Get("/", func(c *fiber.Ctx) error {
		// Sample data for the template
		data := pongo2.Context{
			"title":            "Page Title from Go", // This is for the specific page
			"site_title":       cfg.Site.Title,
			"site_keywords":    cfg.Site.Keywords,
			"site_description": cfg.Site.Description,
			// Add other variables that tpl.html might expect
		}

		// Render the template
		htmlContent, err := services.RenderTemplate("tpl.html", data)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			// In a real app, return a proper error page
			return c.Status(fiber.StatusInternalServerError).SendString("Error rendering template")
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8) // Ensure UTF-8
		return c.SendString(htmlContent)
	})

	// --- Start Server (Commented out due to execution issues) ---
	if cfg != nil && cfg.Server.Port != "" {
		log.Println("Conceptual: Server would start on port:", cfg.Server.Port)
		// log.Fatal(app.Listen(":" + cfg.Server.Port))
	} else {
		log.Println("Conceptual: Server configuration missing, cannot start.")
		// log.Fatal(app.Listen(":3000")) // Fallback, also commented
	}

	log.Println("Conceptual main.go execution finished. No actual server started.")
}
