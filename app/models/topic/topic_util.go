package topic

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
	"gorm.io/gorm/clause"

	"github.com/gin-gonic/gin"
)

// 我们希望返回的内容里附带 user 和 category 关联信息，需要修改下 topic.Get 方法。
func Get(idstr string) (topic Topic) {
	database.DB.Preload(clause.Associations).Where("id = ?", idstr).First(&topic)
	return
}

//In summary, the difference between these two functions lies in whether they eagerly load associated data (Preload with clause.Associations) or only retrieve the primary data (Where clause without Preload).
//func Get(idstr string) (topic Topic) {
//	database.DB.Where("id", idstr).First(&topic)
//	return
//}

func GetBy(field, value string) (topic Topic) {
	database.DB.Where("? = ?", field, value).First(&topic)
	return
}

func All() (topics []Topic) {
	database.DB.Find(&topics)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Topic{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (topics []Topic, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Topic{}),
		&topics,
		app.V1URL(database.TableName(&Topic{})),
		perPage,
	)
	return
}
