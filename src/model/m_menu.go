package model

import (
	"context"

	"github.com/LyricTian/gin-admin/src/schema"
)

// IMenu 菜单存储接口
type IMenu interface {
	// 查询分页数据
	QueryPage(ctx context.Context, params schema.MenuPageQueryParam, pageIndex, pageSize uint) (int, []*schema.Menu, error)
	// 查询列表数据
	QueryList(ctx context.Context, params schema.MenuListQueryParam) ([]*schema.Menu, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string) (*schema.Menu, error)
	// 检查编号是否存在
	CheckCode(ctx context.Context, code string, parentID string) (bool, error)
	// 检查子级是否存在
	CheckChild(ctx context.Context, parentID string) (bool, error)
	// 创建数据
	Create(ctx context.Context, item schema.Menu) error
	// 更新数据
	Update(ctx context.Context, recordID string, item schema.Menu) error
	// 更新数据和分级码
	UpdateLevelCode(ctx context.Context, recordID, levelCode string) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
}
