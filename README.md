# Rise Phonebook

This is a project to save your contact inside a web app.
The server is written in Go lang and the Client in reactJS.

prerequisites:

- docker-compose version 1.27+.
- ports 3306, 9000, and 3000 are available for docker use.

To run the project install docker-compose and from the home directory run:

```bash
docker-compose up
```

Now you can see the client at the address: http://localhost:3000
The server is running on the address: http://localhost:9000

The following are the routes to get and modify data directly in the server:
Address example:

```
http://localhost:9000/api/contact
```

## Get data from the server

#### All the URLs in this section using HTTP GET method

To get list of contacts use the following URL:

```
http://localhost:9000/api/contacts
```

this is equivalent to:

```
http://localhost:9000/api/contacts?page=0
```

this will return the 10 first contacts.  
To get the next ten contacts we should modify the page number in the query URL:

```
http://localhost:9000/api/contacts?page=1
```

To search in your contact list, we can add more params in the query
(currently supporting first name and last name).
Let's say we want to see all the contact who have y in their first name.

```
http://localhost:9000/api/contacts?first_name=y
```

If we want to see all the contacts who have y in their last name.

```
http://localhost:9000/api/contacts?last_name=y
```

To see all the contacts with y in their first name and y in their last name:

```
http://localhost:9000/api/contacts?last_name=y&first_name=y
```

If the search results contain more than 10 contacts we can add the a page param to the URL query to see the next 10 contacts in the result.

```
http://localhost:9000/api/contacts?page=1&last_name=y&first_name=y
```

## Add data to the server

#### All the URL in this section using HTTP POST method

To add a new contact we use the following URL:

```
http://localhost:9000/api/contacts
```

Together with the request we should send a payload JSON object of the new contact.

This is an example for a valid JSON object to create new contact:

```JSON
{
    "first_name": "The",
    "last_name": "New one",
    "address": [
        {
            "description": "rise",
            "city": "and",
            "street": "shine",
            "home_number": "12",
            "apartment": "12"
        }
    ],
    "phone": [
        {
            "description": "mobile",
            "phone_number": "123456789"
        }
    ]
}
```

- first_name - string **Mandatory**
- last_name - string **Mandatory**

address - array of address object.
The address object contain:

- description - string
- city - string **Mandatory**
- street - string
- apartment - string

The address fields is not mandatory.

phone - array of phone object.
The phone object contain:

- description - string
- phone_number - string **Mandatory**

A curl request to create new contact:

```bash
curl -X POST -d '{"first_name":"aaa","last_name":"dddd","address":[{"description":"ads","city":"adfs","street":"afds","home_number":"12","apartment":"12"}],"phone":[{"description":"dddd","phone_number":"123456789"}]}' http://localhost:9000/api/contacts
```

## Delete data in the server

#### All the URL in this section using HTTP DELETE method

To delete contact we use the following URL:

```
http://localhost:9000/api/contacts/:id
```

for example to delete the contact with ID 1:

```
http://localhost:9000/api/contacts/1
```

curl command example:

```bash
curl -X DELETE http://localhost:9000/api/contacts/1
```

## Update data in the server

#### All the URL in this section using HTTP PUT method

we can modify every field in the contact but first name and last name must have content

```JSON
{
    "first_name": "Updated",
    "last_name": "Man",
    "address": [
        {
            "AddressID": 3,
            "description": "Home",
            "city": "Kings landing",
            "street": "Flee button",
            "home_number": "10",
            "apartment": "9"
        },
        {
            "AddressID": 4,
            "description": "Work",
            "city": "WinterFell",
            "street": "Weirwood",
            "home_number": "21",
            "apartment": "9"
        },
        {
            "description": "College",
            "city": "NYC",
            "street": "Jump street",
            "home_number": "22",
            "apartment": "21"
        }
    ],
    "phone": [
        {
            "PhoneID": 3,
            "description": "Mobile",
            "phone_number": "052-1234567"
        },
        {
            "description": "Home",
            "phone_number": "03-9998889"
        }
    ]
}
```

- first_name - string
- last_name - string

address - array of address object.
The address object contain:
If the address does exist:

- AddressID - int **Mandatory**
- description - string
- city - string
- street - string
- apartment - string

If the address doesn't exist:

- description - string
- city - string **Mandatory**
- street - string
- apartment - string

The address fields is not mandatory.

phone - array of phone object.
The phone object contain:

If the phone does exist:

- PhoneID - int **Mandatory**
- phone_number - string
- description - string

If the phone doesn't exist:

- phone_number - string **Mandatory**
- description - string

example using curl to modify contact with ID 2:

```bash
curl -X PUT -d '{"first_name":"Update","last_name":"Man","address":[{"AddressID":3,"description":"Home","city":"Kings landig","street":"Flee button","home_number":"10","apartment":"9"},{"AddressID":4,"description":"Work","city":"Winterfell","street":"Weirwood","home_number":"21","apartment":"9"},{"description":"College","city":"NYC","street":"Jump street","home_number":"22","apartment":"21"}],"phone":[{"PhoneID":3,"description":"Mobile","phone_number":"052-1234567"},{"PhoneID":4,"description":"Home","phone_number":"052-9991111"}]}' http://localhost:9000/api/contacts/2
```
