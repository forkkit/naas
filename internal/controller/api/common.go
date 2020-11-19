package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nilorg/naas/internal/model"
	"github.com/nilorg/naas/internal/pkg/contexts"
	"github.com/nilorg/naas/internal/service"
	"github.com/nilorg/sdk/convert"
)

type common struct {
}

// SelectQueryChildren 查询方法
// @Tags 		Common（通用）
// @Summary		查询角色
// @Description	org:组织
// @Description	xxxx:其他
// @Accept  json
// @Produce	json
// @Param q query string true "查询参数" Enums(org,xxxx)
// @Param pageSize query int true "页大小"
// @Success 200	{object}	Result{data=model.ResultSelect}
// @Router /common/select [GET]
// @Security OAuth2AccessCode
func (c *common) SelectQueryChildren() gin.HandlerFunc {
	return QueryChildren(map[string]gin.HandlerFunc{
		"organization":    c.SelectOrganizationList,
		"resource_server": c.SelectResourceServerList,
		"role":            c.SelectRoleList,
		"oauth2_scope":    c.SelectOAuth2ScopeList,
	})
}

// SelectOrganizationList 组织Select列表
func (*common) SelectOrganizationList(ctx *gin.Context) {
	id := ctx.Query("id")
	parentCtx := contexts.WithGinContext(ctx)
	if id != "" {
		org, err := service.Organization.GetOneByID(parentCtx, model.ConvertStringToID(id))
		if err != nil {
			writeError(ctx, err)
		} else {
			writeData(ctx, &model.ResultSelect{
				Label: org.Name,
				Value: org.ID,
			})
		}
	} else {
		name := ctx.Query("name")
		limit := convert.ToInt(ctx.Query("limit"))
		list, err := service.Organization.ListByName(parentCtx, name, limit)
		if err != nil {
			writeError(ctx, err)
		} else {
			var results []*model.ResultSelect
			for _, item := range list {
				results = append(results, &model.ResultSelect{
					Label: item.Name,
					Value: item.ID,
				})
			}
			writeData(ctx, results)
		}
	}
}

// SelectResourceServerList 资源服务器Select列表
func (*common) SelectResourceServerList(ctx *gin.Context) {
	id := ctx.Query("id")
	parentCtx := contexts.WithGinContext(ctx)
	if id != "" {
		res, err := service.Resource.GetServer(parentCtx, model.ConvertStringToID(id))
		if err != nil {
			writeError(ctx, err)
		} else {
			writeData(ctx, &model.ResultSelect{
				Label: res.Name,
				Value: res.ID,
			})
		}
	} else {
		name := ctx.Query("name")
		limit := convert.ToInt(ctx.Query("limit"))
		list, err := service.Resource.ListByName(parentCtx, name, limit)
		if err != nil {
			writeError(ctx, err)
		} else {
			var results []*model.ResultSelect
			for _, item := range list {
				results = append(results, &model.ResultSelect{
					Label: item.Name,
					Value: item.ID,
				})
			}
			writeData(ctx, results)
		}
	}
}

// SelectRoleList 角色Select列表
func (*common) SelectRoleList(ctx *gin.Context) {
	id := ctx.Query("id")
	parentCtx := contexts.WithGinContext(ctx)
	if id != "" {
		res, err := service.Role.GetOneByCode(parentCtx, model.ConvertStringToCode(id))
		if err != nil {
			writeError(ctx, err)
		} else {
			writeData(ctx, &model.ResultSelect{
				Label: res.Name,
				Value: res.Code,
			})
		}
	} else {
		name := ctx.Query("name")
		organizationID := model.ConvertStringToID(ctx.Query("organization_id"))
		limit := convert.ToInt(ctx.Query("limit"))
		var (
			list []*model.Role
			err  error
		)
		if organizationID > 0 {
			list, err = service.Role.ListByNameForOrganization(parentCtx, organizationID, name, limit)
		} else {
			list, err = service.Role.ListByName(parentCtx, name, limit)
		}
		if err != nil {
			writeError(ctx, err)
		} else {
			var results []*model.ResultSelect
			for _, item := range list {
				results = append(results, &model.ResultSelect{
					Label: item.Name,
					Value: item.Code,
				})
			}
			writeData(ctx, results)
		}
	}
}

// SelectOAuth2ScopeList OAuth2范围列表
func (*common) SelectOAuth2ScopeList(ctx *gin.Context) {
	parentCtx := contexts.WithGinContext(ctx)
	var (
		list []*model.OAuth2Scope
		err  error
	)
	list, err = service.OAuth2.AllScope(parentCtx)
	if err != nil {
		writeError(ctx, err)
	} else {
		var results []*model.ResultSelect
		for _, item := range list {
			results = append(results, &model.ResultSelect{
				Label: item.Name,
				Value: item.Code,
			})
		}
		writeData(ctx, results)
	}
}
