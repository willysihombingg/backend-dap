package user

import (
	"fmt"

	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/connector/supabase/user"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/internal/helper"
	"gitlab.com/willysihombing/task-c3/internal/presentations"
	"gitlab.com/willysihombing/task-c3/internal/ucase/contract"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/tracer"
	"golang.org/x/crypto/bcrypt"
)

type createNewUser struct {
	connectorAdmin user.Connectorer
}

func CreateUser(connectorAdmin user.Connectorer) contract.UseCase {
	return &createNewUser{
		connectorAdmin: connectorAdmin,
	}
}

func (u *createNewUser) Serve(dctx *appctx.Data) appctx.Response {
	var (
		ctx      = tracer.SpanStart(dctx.Request.Context(), "ucase.document_stamp_category")
		param    presentations.CreateUser
		response interface{}
		err      = dctx.Cast(&param)
		lf       = logger.NewFields(
			logger.EventName("CreateNewUser"),
			logger.Any("query", fmt.Sprintf("*")),
		)
	)

	authInfo := helper.AuthInfoFromContext(ctx)
	fmt.Println(authInfo.Role)

	if authInfo.Role != "admin" {
		logger.ErrorWithContext(ctx, fmt.Sprintf("user is not an admin"), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeAuthenticationFailure).WithMessage("Unauthorized!")
	}

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	param.Password = string(hashedPassword)
	fmt.Println(param.Password)

	request := u.connectorAdmin.Insert(param)
	err = u.connectorAdmin.Send(ctx, request, response)

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("error on connector all users %v", err), lf...)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError).WithMessage("Create users error!")
	}

	return *appctx.NewResponse().WithStatus("Success create Users").WithCode(consts.CodeStatusCreated).WithData("Created").WithMessage("Create Users Success")
}
