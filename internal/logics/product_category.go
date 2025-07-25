package logics

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/model/entity"
	"DeviceManagement/internal/service"
	"context"
	"database/sql"
	"errors"
	"log"
	"sync"
)

var (
	productCategoryOnce     sync.Once
	productCategoryInstance *productCategory
)

type productCategory struct {
}

func NewProductCategory() service.IProductCategory {
	productCategoryOnce.Do(func() {
		productCategoryInstance = &productCategory{}
	})
	return productCategoryInstance
}

func (l *productCategory) Add(ctx context.Context, req *v1.AddProductCategoryReq) (out *model.ProductCategory, err error) {
	dataInsert := map[string]interface{}{
		dao.ProductCategory.Columns().PID:   req.PID,
		dao.ProductCategory.Columns().OrgID: req.OrgID,
		dao.ProductCategory.Columns().Name:  req.Name,
	}

	id, err := dao.ProductCategory.Ctx(ctx).Data(dataInsert).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	out = &model.ProductCategory{
		ID: id,
	}
	return
}

func (l *productCategory) Delete(ctx context.Context, req *v1.DeleteProductCategoryReq) (err error) {
	err = l.deleteCategory(ctx, req.ID)
	if err != nil {
		return err
	}
	return
}

func (l *productCategory) Update(ctx context.Context, req *v1.UpdateProductCategoryReq) (err error) {
	dataUpdate := map[string]interface{}{
		dao.ProductCategory.Columns().PID:  req.PID,
		dao.ProductCategory.Columns().Name: req.Name,
	}

	_, err = dao.ProductCategory.Ctx(ctx).Where(dao.ProductCategory.Columns().ID, req.ID).Data(dataUpdate).Update()
	if err != nil {
		return err
	}

	return
}

func (l *productCategory) GetTreeByID(ctx context.Context, id int64) (out *model.ProductCategoryTree, err error) {
	// 首先验证节点是否存在
	var exist bool
	exist, err = dao.ProductCategory.Ctx(ctx).Where(dao.ProductCategory.Columns().ID, id).Exist()
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.New("分类不存在")
	}

	// 获取整个子树的所有节点（包括当前节点及其所有子节点）
	categories, err := l.getAllSubTreeNodes(ctx, id)
	if err != nil {
		return nil, err
	}

	for _, v := range categories {
		log.Printf("%+v\n", v)
	}

	// 构建树结构
	out = l.buildTree(id, categories)

	return
}

func (l *productCategory) GetTreeByOrgID(ctx context.Context, orgID string) (out model.ProductCategoryTreeList, err error) {
	// 获取顶层节点
	var result []*entity.TProductCategory
	err = dao.ProductCategory.Ctx(ctx).Fields(dao.ProductCategory.Columns().ID).Where(dao.ProductCategory.Columns().OrgID, orgID).Where(dao.ProductCategory.Columns().PID, 0).Scan(&result)
	if err != nil {
		return nil, err
	}

	for _, v := range result {
		tree, err := l.GetTreeByID(ctx, v.ID)
		if err != nil {
			return nil, err
		}
		out = append(out, tree)
	}

	return
}

// getAllSubTreeNodes 获取指定节点及其所有子节点
func (l *productCategory) getAllSubTreeNodes(ctx context.Context, rootID int64) (nodes []*model.ProductCategory, err error) {
	// 使用广度优先搜索收集所有子节点
	queue := []int64{rootID}
	nodeMap := make(map[int64]*model.ProductCategory)

	for len(queue) > 0 {
		currentID := queue[0]
		queue = queue[1:]

		// 如果当前节点已经处理过，跳过
		if _, exists := nodeMap[currentID]; exists {
			continue
		}

		// 获取当前节点信息
		var currentNode entity.TProductCategory
		if err = dao.ProductCategory.Ctx(ctx).Where(dao.ProductCategory.Columns().ID, currentID).Scan(&currentNode); err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			return nil, err
		}

		// 将当前节点添加到结果中
		nodeMap[currentID] = l.convertProductCategoryToLogic(&currentNode)

		// 查找当前节点的所有子节点
		var children []*entity.TProductCategory
		if err = dao.ProductCategory.Ctx(ctx).Fields(dao.ProductCategory.Columns().ID).Where(dao.ProductCategory.Columns().PID, currentID).Scan(&children); err != nil {
			return nil, err
		}

		ids := make([]int64, 0)
		for _, child := range children {
			ids = append(ids, child.ID)
		}
		log.Printf("当前节点ID: %d, 子节点ID: %v", currentID, ids)

		// 将子节点ID添加到队列中
		for _, child := range children {
			queue = append(queue, child.ID)
		}
	}

	// 将map转换为slice
	for _, node := range nodeMap {
		nodes = append(nodes, node)
	}

	return nodes, nil
}

// buildTree 构建分类树
// topID 根节点ID
// categories 包含根节点、及其子节点的所有节点
func (l *productCategory) buildTree(topID int64, categories []*model.ProductCategory) (root *model.ProductCategoryTree) {
	// 找到根节点
	for _, v := range categories {
		if v.ID == topID {
			root = &model.ProductCategoryTree{
				ProductCategory: v,
			}
			// TODO: Bug why
			// categories = append(categories[:i], categories[i+1:]...)
			break
		}
	}

	if root == nil {
		return nil
	}

	// 构建子节点
	for _, v := range categories {
		if v.PID == root.ID {
			root.Children = append(root.Children, l.buildTree(v.ID, categories))
		}
	}

	return
}

func (l *productCategory) deleteCategory(ctx context.Context, id int64) (err error) {
	idsToDelete := make([]int64, 0)
	queue := []int64{id}

	for len(queue) > 0 {
		currentID := queue[0]
		queue = queue[1:]
		idsToDelete = append(idsToDelete, currentID)

		var children []*entity.TProductCategory
		if err = dao.ProductCategory.Ctx(ctx).Where(dao.ProductCategory.Columns().PID, currentID).Scan(&children); err != nil {
			return err
		}

		for _, child := range children {
			queue = append(queue, child.ID)
		}
	}

	if len(idsToDelete) > 0 {
		_, err = dao.ProductCategory.Ctx(ctx).Where(dao.ProductCategory.Columns().ID, idsToDelete).Delete()
	}

	return err
}

func (l *productCategory) convertProductCategoryToLogic(in *entity.TProductCategory) (out *model.ProductCategory) {
	out = &model.ProductCategory{
		ID:        in.ID,
		PID:       in.PID,
		OrgID:     in.OrgID,
		Name:      in.Name,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}

	return
}
