package cmd

import (
	"github.com/spf13/cobra"
	"grpc-gateway-example/server"
	"log"
)

// 创建附加命令
// 本地标签：在本地分配一个标志，该标志仅适用于该特定命令。
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC hello-world server",
	// 运行:典型的实际工作功能。大多数命令只会实现这一点；
	// 另外还有PreRun、PreRunE、PostRun、PostRunE等等不同时期的运行命令，但比较少用，具体使用时再查看亦可
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recover error : %v", err)
			}
		}()

		server.Run()
	},
}

// 在 init() 函数中定义flags和处理配置。
func init() {
	// 我们定义了一个flag，值存储在&server.ServerPort中，长命令为--port，短命令为-p，，默认值为50052。
	// 命令的描述为server port。这一种调用方式成为Local Flags 本地标签
	serverCmd.Flags().StringVarP(&server.ServerPort, "port", "p", "50052", "server port")
	serverCmd.Flags().StringVarP(&server.CertPemPath, "cert-pem", "", "./certs/server.pem", "cert-pem path")
	serverCmd.Flags().StringVarP(&server.CertKeyPath, "cert-key", "", "./certs/server.key", "cert-key path")
	serverCmd.Flags().StringVarP(&server.CertServerName, "cert-server-name", "", "grpc-gateway", "server's hostname")
	serverCmd.Flags().StringVarP(&server.SwaggerDir, "swagger-dir", "", "proto", "path to the directory which contains swagger definitions")

	// AddCommand向这父命令（rootCmd）添加一个或多个命令
	rootCmd.AddCommand(serverCmd)
}
