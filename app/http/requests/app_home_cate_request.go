package requests

import (
	"github.com/gin-gonic/gin"
	requests "go_ypc/pkg/request"
	"go_ypc/pkg/response"
)

// 绑定模型
type HomeCate struct {
	Title    string `json:"title" binding:"required,checkDate"`
	Subtitle string `json:"subtitle" binding:"required"`
	// Cate     string `json:"cate" binding:"required"`
	// Sort     uint64 `json:"sort" binding:"required,gte=1,unique=Sort"`
	// Status   uint64 `json:"status" binding:"required"`
}

func ValidateHomeCate(c *gin.Context) HomeCate {

	// 自定义错误消息
	messages := map[string]map[string]string{
		"Title": {
			"required":  "标题 不能为空.",
		},
	}

	// 验证
	var form HomeCate

	if err := c.ShouldBind(&form); err != nil {
		response.UnprocessableEntity(requests.ParseCustomErrors(&HomeCate{}, err, messages), c)
		return HomeCate{}
	}

	return form
}
