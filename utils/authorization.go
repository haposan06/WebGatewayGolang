package utils

import (
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"tc-web-gateway/utils/errors"
)

type Authorization struct {
	*casbin.Enforcer
}

var authInstance *Authorization = nil

func Init() *Authorization{
	if authInstance == nil {
		policy_db := gormadapter.NewAdapter("postgres", "host=159.89.205.12 port=5432 user=merchant password=merchant dbname=web-gateway", true) // Your driver and data source
		authEnforcer, err := casbin.NewEnforcerSafe("./utils/auth_model.conf", policy_db)
		authEnforcer.EnableAutoSave(true)
		if err != nil {
			log.Fatal(err)
		} else {
			authInstance = &Authorization{authEnforcer}
			return authInstance
		}
	}
	return authInstance
}

func (a *Authorization) CheckAccessRules(role string, r *http.Request) (int, error) {
	var err error

	res, err := a.EnforceSafe(role, r.URL.Path, r.Method)

	if err != nil {
		return http.StatusInternalServerError, errors.ErrSystemInternal
	}
	if res {
		return http.StatusOK, nil
	} else {
		return http.StatusForbidden, errors.ErrUnauthorized
	}
}

func AddPolicy(roleName string, path string, method string) (bool, error) {
	authEnforcer:= Init()
	result, err := authEnforcer.AddPolicySafe(roleName, path, method)
	return result, err
}

func RemovePolicy(roleName string, path string, method string) (bool, error) {
	authEnforcer:= Init()
	result, err := authEnforcer.RemovePolicySafe(roleName, path , method)
	return result, err
}