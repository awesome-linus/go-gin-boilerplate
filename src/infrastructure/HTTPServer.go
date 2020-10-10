package infrastructure

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/awesome-linus/go-gin-mysql-todo-api/src/config"
	"github.com/awesome-linus/go-gin-mysql-todo-api/src/domain"
	"github.com/awesome-linus/go-gin-mysql-todo-api/src/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// 注意: MySQL用ドライバは削除すると接続できなくなるので注意
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

type Server struct {
	router *gin.Engine
	DB     *gorm.DB
	Logger *zap.Logger
}

func (s *Server) Init(env string) {
	log.Printf("env: %s", env)
}

func (s *Server) Middleware() {
	s.router.Use(middleware.RecordUaAndTime)
}

func NewServerWithMySQL(db *gorm.DB, logger *zap.Logger) *Server {
	return &Server{
		router: gin.Default(),
		DB:     db,
		Logger: logger,
	}
}

func (s *Server) Router() {
	h := NewHandlerWithMySQL(s.DB)

	// HealthCheck
	s.router.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"healthCheck": "ok",
		})
	})

	// 静的ファイルのパスを指定
	s.router.Static("/views", "./views")

	// Todo
	s.router.GET("/todos", h.TodoList)
	s.router.GET("/todos/:todoId", h.ShowTodo)
	s.router.POST("/todos", h.AddTodo)
	s.router.DELETE("/todos/:todoId", h.DeleteTodo)
	s.router.PATCH("/todos/:todoId", h.ChangeTodo)

	if err := s.router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}

}

func StartHTTPServer() {
	var (
		port = flag.String("port", "8888", "addr to bind")
		env  = flag.String("env", "develop", "実行環境 (production, staging, develop)")
	)
	flag.Parse()

	logger := CreateLogger()
	defer logger.Sync()

	db, err := gorm.Open("mysql", config.GetDsn())
	if err != nil {
		log.Fatal(db, "Unable to connect to MySQL server.")
	}

	defer db.Close()

	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	db.SingularTable(true)

	// マイグレーション（テーブルが無い時は自動生成）
	// db.AutoMigrate(&entity.Product{})
	db.AutoMigrate(&domain.Todo{})

	s := NewServerWithMySQL(db, logger)
	s.Init(*env)
	s.Middleware()
	s.Router()
	log.Println("Starting app")
	_ = http.ListenAndServe(fmt.Sprint(":", *port), s.router)

	fmt.Println("test")
	fmt.Println(port, env)
	fmt.Println(config.GetDsn())
}
