package project

import (
	"context"
	"fmt"
	"net/http"

	authoritypb "github.com/ez-deploy/protobuf/authority"
	"github.com/ez-deploy/protobuf/model"
	projectpb "github.com/ez-deploy/protobuf/project"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ez-deploy/apiserver/handle/internal/errors"
	"github.com/ez-deploy/apiserver/models"
	"github.com/ez-deploy/apiserver/restapi/operations/project_ops"
)

func New(authorityClient authoritypb.AuthorityOpsClient, projectClient projectpb.ProjectOpsClient) *ProjectOpsImpl {
	return &ProjectOpsImpl{}
}

type ProjectOpsImpl struct {
	authorityClient authoritypb.AuthorityOpsClient
	projectclient   projectpb.ProjectOpsClient
}

const (
	resourceTypeProject = "project"
	projectActionManage = "manage"
)

// ProjectOpsCreateProject create a project for user.
func (h *ProjectOpsImpl) ProjectOpsCreateProject(params project_ops.ProjectOpsCreateProjectParams, principal *models.IdentityVerifyResp) middleware.Responder {
	createdProject := params.Body.Project

	createProjReq := &projectpb.CreateProjectReq{
		Project: &model.Project{
			Name:     createdProject.Name,
			Describe: createdProject.Describe,
		},
	}

	createProjResp, err := h.projectclient.CreateProject(params.HTTPRequest.Context(), createProjReq)
	if err != nil {
		return errors.NewCommonResposeWithErr(err)
	} else if createProjResp.Error != nil {
		return errors.NewCommonResponseWithErrBody(createProjResp.Error)
	}

	err = h.setProjectAuthorityForUser(params.HTTPRequest.Context(), createdProject.Name, principal.Identity.Email)
	if err != nil {
		return errors.NewCommonResposeWithErr(err)
	}

	return project_ops.NewProjectOpsCreateProjectOK().WithPayload(&models.ModelCommonResp{})
}

func (h *ProjectOpsImpl) ProjectOpsDeleteProject(params project_ops.ProjectOpsDeleteProjectParams, principal *models.IdentityVerifyResp) middleware.Responder {
	if params.ProjectName == nil {
		return middleware.Error(http.StatusBadRequest, "project name is needed")
	}
	deletedProjectName := *params.ProjectName

	ok, err := h.haveAuthorityToProject(principal, deletedProjectName)
	if err != nil {
		return errors.NewCommonResposeWithErr(err)
	} else if !ok {
		return newForbiddenRespFromMsg("can not visit given project")
	}

	if err := h.deleteProject(params.HTTPRequest.Context(), deletedProjectName); err != nil {
		return errors.NewCommonResposeWithErr(err)
	}

	return project_ops.NewProjectOpsDeleteProjectOK().WithPayload(&models.ModelCommonResp{})
}

func (h *ProjectOpsImpl) ProjectOpsList(params project_ops.ProjectOpsListParams, principal *models.IdentityVerifyResp) middleware.Responder {
	req := &model.Identity{
		Email: principal.Identity.Email,
		Name:  principal.Identity.Name,
	}

	resp, err := h.authorityClient.ListAuthoritiesByIdentity(params.HTTPRequest.Context(), req)
	if err != nil {
		return errors.NewCommonResposeWithErr(err)
	} else if resp.Error != nil {
		return errors.NewCommonResponseWithErrBody(resp.Error)
	}

	httpRespBody := &project_ops.ProjectOpsListOKBody{}
	for _, proj := range resp.Authorities.Authorities {
		httpRespBody.Projects = append(httpRespBody.Projects, proj.Resource.Name)
	}

	return project_ops.NewProjectOpsListOK().WithPayload(httpRespBody)
}

func (h *ProjectOpsImpl) setProjectAuthorityForUser(ctx context.Context, projectName string, useremail string) error {
	setAuthorityReq := &authoritypb.Authorities{Authorities: []*model.Authority{
		{
			Identity: &model.Identity{
				Email: useremail,
			},
			Resource: &model.Resource{
				Type: resourceTypeProject,
				Name: projectName,
			},
			Action: projectActionManage,
		},
	}}

	setAuthorityResp, err := h.authorityClient.SetAuthorities(ctx, setAuthorityReq)
	if err != nil {
		return err
	} else if setAuthorityResp.Error != nil {
		return fmt.Errorf("set authority error: %v", setAuthorityResp.Error.Message)
	}

	return nil
}

// deleteProject.
func (h *ProjectOpsImpl) deleteProject(ctx context.Context, projectName string) error {
	deleteProjectReq := &projectpb.DeleteProjectReq{ProjectName: projectName}
	deleteProjectResp, err := h.projectclient.DeleteProject(ctx, deleteProjectReq)
	if err != nil {
		return err
	} else if deleteProjectResp.Error != nil {
		return fmt.Errorf(deleteProjectResp.Error.Message)
	}

	// delete authorities for project.
	projectResource := &model.Resource{Type: resourceTypeProject, Name: projectName}
	projectAuthorities, err := h.authorityClient.ListAuthoritiesByResource(ctx, projectResource)
	if err != nil {
		return err
	} else if projectAuthorities.Error != nil {
		return fmt.Errorf(projectAuthorities.Error.Message)
	}

	deleteAuthorityResp, err := h.authorityClient.DeleteAuthorities(ctx, projectAuthorities.Authorities)
	if err != nil {
		return err
	}

	if len(deleteAuthorityResp.FailMessages) == 0 {
		return nil
	}

	errMsg := ""
	for _, msg := range deleteAuthorityResp.FailMessages {
		errMsg += fmt.Sprintf(
			"delete %v--%v-->%v error: %v",
			msg.Authority.Identity,
			msg.Authority.Action,
			msg.Authority.Resource,
			msg.Error.Message,
		)
	}

	return fmt.Errorf(errMsg)
}
