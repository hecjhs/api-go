package middleware

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/hecjhs/api-go/api/models"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

// Queue
type Queue struct {
	Domain   string
	Weight   int
	Priority int
}

// Compute priority
func (q *Queue) ComputePriority(domain string) (string error) {
	return nil
}

// Que declaration
var Que []string

// Repository should implement common methods
type Repository interface {
	Read() []*Queue
	ReadFromDB(domain string) models.Domain
}

func (q *Queue) Read() []*Queue {
	path, _ := filepath.Abs("")
	file, err := os.Open(path + "/api/middleware/domain.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var final []*Queue
	tmp := &Queue{}
	count := 0
	for scanner.Scan() {
		count++
		if scanner.Text() == "" {
			count = 0
			continue
		}
		switch count {
		case 1:
			tmp.Domain = scanner.Text()
		case 2:
			r := strings.Split(scanner.Text(), ":")[1]
			res, _ := strconv.Atoi(r)
			tmp.Weight = res
		case 3:
			r := strings.Split(scanner.Text(), ":")[1]
			res, _ := strconv.Atoi(r)
			tmp.Priority = res
			final = append(final, tmp)
			tmp = &Queue{}
		}
	}
	return final
}

func (q *Queue) ReadFromDB(domain string) models.Domain {
	var path, _ = filepath.Abs("")
	db, _ := gorm.Open("sqlite3", path+"/api/fixtures/data.db")
	defer db.Close()
	d := models.Domain{Domain: domain}
	db.First(&d)
	return d
}

// MockQueue should mock an Array of Queues
func MockQueue() []*Queue {
	return []*Queue{
		{
			Domain:   "alpha",
			Weight:   5,
			Priority: 5,
		},
		{
			Domain:   "omega",
			Weight:   1,
			Priority: 5,
		},
		{
			Domain:   "beta",
			Weight:   5,
			Priority: 1,
		},
	}
}

// ProxyMiddleware should queue our incoming requests
func ProxyMiddleware(c iris.Context) {
	var QPriority = make(map[string][]string)
	domain := c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}
	var repo Repository
	repo = &Queue{}
	d := repo.ReadFromDB(domain)
	if d.Priority < 5 && d.Weight < 5 {
		QPriority["high"] = append(QPriority["high"], domain)
	} else if d.Priority < 5 || d.Weight < 5 {
		QPriority["midium"] = append(QPriority["midium"], domain)

	} else {
		QPriority["low"] = append(QPriority["low"], domain)
	}
	Que = append(Que, QPriority["high"]...)
	Que = append(Que, QPriority["midium"]...)
	Que = append(Que, QPriority["low"]...)

	c.Next()
}
