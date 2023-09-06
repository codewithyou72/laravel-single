//设置静态文件夹public、新增文件需要手动重启应用

package common

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/zeromicro/go-zero/rest"
	"laravel-single/app/internal/config"
	"net/http"
	"os"
	"path/filepath"
)

func PublicHandler(filePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, fmt.Sprintf("filePath=%sMethod must be GET", filePath), http.StatusMethodNotAllowed)
			return
		}
		http.ServeFile(w, r, filePath)
	}
}

// 设置public目录下面的内容为路径
func PublicStaticFileHandler(engine *rest.Server, c config.Config) {
	dir := c.PublicDir

	engine.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/",
			Handler: PublicHandler(filepath.Join(dir, "index.html")),
		},
	)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			realPath := path
			realPath = filepath.ToSlash(realPath)
			truncatedPath, err := filepath.Rel(dir, realPath)
			if err != nil {
				return err
			}
			fmt.Println(aurora.Green(fmt.Sprintf("truncatedPath=%s realPath=%s", filepath.ToSlash(truncatedPath), realPath)))

			engine.AddRoute(
				rest.Route{
					Method:  http.MethodGet,
					Path:    "/" + truncatedPath,
					Handler: PublicHandler(realPath),
				},
			)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", dir, err)
	}
}
