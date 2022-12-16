// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"cherry-disk/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: CoreHandler(serverCtx),
			},
		},
	)

	//用户
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/detail",
				Handler: UserDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register/send/code",
				Handler: MailCodeSendRegisterHandler(serverCtx),
			},
			//分享文件详细
			{
				Method:  http.MethodGet,
				Path:    "/share/basic/detail",
				Handler: ShareBasicDetailHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/neko",
					Handler: CoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/upload",
					Handler: FileUploadHandler(serverCtx),
				},
				{
					// TODO
					// no tested
					Method:  http.MethodPost,
					Path:    "/user/repository/save",
					Handler: UserRepositorySaveHandler(serverCtx),
				},
				{
					// TODO
					//no tested
					Method:  http.MethodPost,
					Path:    "/user/file/list",
					Handler: UserFileListHandler(serverCtx),
				},
			}...,
		),
	)
}
