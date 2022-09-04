package service

import serverutils "phonebook/server/server-utils"

func getWhereCond(fieldName string, id string) (string, error) {
	var ret string
	where, err := serverutils.GetQuery(sqlQueryWhere)
	if err != nil {
		return "", err
	}
	ret += where

	ret += serverutils.AddValuesToQuery(fieldName, id)

	return ret, nil
}
