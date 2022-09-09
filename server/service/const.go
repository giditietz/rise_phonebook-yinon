package service

const (
	sqlQueryWhere             = "where"
	sqlQueryAnd               = "and"
	sqlQueryGetAll            = "getAllContact"
	sqlQueryInsertContact     = "insertContact"
	sqlQueryInsertAddress     = "insertAddress"
	sqlQueryGetContactAddress = "getContactAddress"
	sqlQueryInsertPhone       = "insertPhone"
	sqlQueryDeleteContact     = "deleteContact"
	sqlQueryEditContact       = "editContact"
	sqlQueryGetContactPhones  = "getContactPhones"
	sqlQueryEditPhone         = "editPhone"
	sqlQueryEditAddress       = "editAddress"
	sqlQueryGetNumOfContacts  = "getNumOfContacts"
	sqlSeparatorValues        = ", "
)

const (
	retrieveResultLimit = 10
)

const (
	firstNameFieldInDB   = "first_name"
	lastNameFieldInDB    = "last_name"
	contactIdFieldInDB   = "contact_id"
	phoneIdFieldInDB     = "phone_id"
	addressIdFieldInDB   = "address_id"
	descriptionFieldInDB = "description"
	phoneNumberFieldInDB = "phone_number"
	cityFieldInDB        = "city"
	streetFieldInDB      = "street"
	homeNumberFieldInDB  = "home_number"
	apartmentFieldInDB   = "apartment"
)
