package middleware

import (
	"bufio"
	"gofibre-project/config" // Import project's config package
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// readFileContent reads a file and returns its content as a string.
func readFileContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return strings.Join(lines, "\n"), nil
}

// matchRegexList checks if a string matches any regex pattern in a list.
func matchRegexList(s string, patterns []string) bool {
	for _, pattern := range patterns {
		matched, _ := regexp.MatchString(pattern, s) // Error ignored for simplicity in this context
		if matched {
			return true
		}
	}
	return false
}

// NewAccessControlMiddleware creates a new middleware for access control.
func NewAccessControlMiddleware(cfg config.AccessControlConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// a. Overall Check
		if !cfg.Enabled {
			return c.Next()
		}

		// b. Preview Mode
		if c.Query(cfg.PreviewParam) == "1" {
			return c.Next()
		}

		clientIP := c.IP()
		userAgent := string(c.Request().Header.UserAgent()) // More direct way to get user agent

		// c. Blacklist Mode
		if cfg.BlacklistMode.Enabled {
			if matchRegexList(clientIP, cfg.BlacklistMode.IPs) {
				htmlContent, err := readFileContent(cfg.BlacklistMode.Page)
				if err != nil {
					// Log error: log.Printf("Error reading blacklist page %s: %v", cfg.BlacklistMode.Page, err)
					return c.Status(fiber.StatusInternalServerError).SendString("Error page unavailable")
				}
				return c.Status(fiber.StatusForbidden).Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8).SendString(htmlContent)
			}
		}

		// d. Spider Mode (Whitelist for verified spiders, special page for others)
		if cfg.SpiderMode.Enabled {
			isAllowedSpider := false
			isMobileAgent := false // Determine this once

			// Check if mobile based on User-Agent
			for _, keyword := range cfg.SpiderMode.MobileKeywords {
				if strings.Contains(strings.ToLower(userAgent), strings.ToLower(keyword)) {
					isMobileAgent = true
					break
				}
			}

			for _, spiderKeyword := range cfg.SpiderMode.AllowedSpiders {
				if strings.Contains(strings.ToLower(userAgent), strings.ToLower(spiderKeyword)) {
					// User agent matches an allowed spider keyword, now verify IP
					if matchRegexList(clientIP, cfg.SpiderMode.VerifiedIPs) {
						isAllowedSpider = true
						break
					}
					if cfg.SpiderMode.VerifyByHostname {
						hostnames, err := net.LookupAddr(clientIP)
						if err == nil {
							for _, hostname := range hostnames {
								if strings.Contains(strings.ToLower(hostname), strings.ToLower(spiderKeyword)) {
									isAllowedSpider = true
									break
								}
							}
						}
						if isAllowedSpider { // Break from outer spiderKeyword loop
							break
						}
					}
				}
			}

			if !isAllowedSpider {
				var pagePath string
				if isMobileAgent {
					pagePath = cfg.SpiderMode.DefaultMobPage
				} else {
					pagePath = cfg.SpiderMode.DefaultPcPage
				}
				htmlContent, err := readFileContent(pagePath)
				if err != nil {
					// Log error: log.Printf("Error reading spider mode page %s: %v", pagePath, err)
					return c.Status(fiber.StatusInternalServerError).SendString("Error page unavailable")
				}
				// It's unusual to return 200 OK for a "blocked" spider, usually it's 403 or a special page.
				// But sticking to the requirement from the description.
				return c.Status(fiber.StatusOK).Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8).SendString(htmlContent)
			}
		}

		// e. Default Pass
		return c.Next()
	}
}
