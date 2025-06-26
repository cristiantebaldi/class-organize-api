package main

import (
	"context"

	"github.com/cristiantebaldi/class-organize-api/internals/adapter/postgres"
	"github.com/cristiantebaldi/class-organize-api/internals/util"
	"github.com/cristiantebaldi/class-organize-api/pkg/di"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	conn := postgres.Initialize(ctx)
	validate := util.NewValidator()
	defer conn.Close()

	alocacaoService := di.ConfigAlocacaoDI(conn, validate)
	professorService := di.ConfigProfessorDI(conn, validate)
	salaService := di.ConfigSalaDI(conn, validate)
	turmaService := di.ConfigTurmaDI(conn, validate)

	router := gin.Default()

	router.POST("/alocacao", alocacaoService.Create)
	router.PUT("/alocacao/:id", alocacaoService.Update)
	router.GET("/alocacao", alocacaoService.Fetch)
	router.GET("/alocacao/:id", alocacaoService.GetByID)
	router.DELETE("/alocacao/:id", alocacaoService.Delete)
	router.POST("/alocacao/automatico", alocacaoService.CreateAlocacoesAutomaticamente)
	
	router.POST("/professor", professorService.Create)
	router.PUT("/professor/:id", professorService.Update)
	router.GET("/professor", professorService.Fetch)
	router.GET("/professor/:id", professorService.GetByID)
	router.DELETE("/professor/:id", professorService.Delete)
	
	router.POST("/sala", salaService.Create)
	router.PUT("/sala/:id", salaService.Update)
	router.GET("/sala", salaService.Fetch)
	router.GET("/sala/:id", salaService.GetByID)
	router.DELETE("/sala/:id", salaService.Delete)
	
	router.POST("/turma", turmaService.Create)
	router.PUT("/turma/:id", turmaService.Update)
	router.GET("/turma", turmaService.Fetch)
	router.GET("/turma/:id", turmaService.GetByID)
	router.DELETE("/turma/:id", turmaService.Delete)
	
	port := viper.GetString("server.http.port")

	router.Run(":" + port)
}
