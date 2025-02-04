package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"local.package/models"
)

func GetDetailTask(c *gin.Context) {
	var task models.Task

	// DBに接続
	db := DbConnection()
	// db.DB()を設定することでconnectionとして使用したDBをクローズする準備ができる
	dbForClose, _ := db.DB()

	// URL内のQueryStringのIDをc.Param("id")で受け取って検索した件数（RowsAffected）を取得
	resultCount := db.Where("id = ?", c.Param("id")).Limit(1).Find(&task).RowsAffected

	// 検索結果が0件の場合は400エラーを返却する
	if resultCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no record for get detail"})
		return
	}
	// 検索完了後(この関数の処理終了後)にAPIで使用していたDBをクローズする
	defer dbForClose.Close()

	// 取得できた場合は取得内容をJSON形式で表示する（1件）
	c.JSON(http.StatusOK, task)
}
