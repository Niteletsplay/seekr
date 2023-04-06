## Post Person


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"id": "2"
}'
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 0,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {},
	"gender": "",
	"hobbies": "",
	"id": "2",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"ssn": "",
	"tags": []
}
```

**Status Code:** 201


## Overwrite Person


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"id": "1"
}'
```

**Response:**

```json
{
	"message": "overwritten person"
}
```

**Status Code:** 202


## Get Person by ID


**Curl Request:**

```sh
curl -X GET http://localhost:8080/people/2
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 0,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {},
	"gender": "",
	"hobbies": "",
	"id": "2",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"ssn": "",
	"tags": []
}
```

**Status Code:** 200


## Get Person which does not exsist


**Curl Request:**

```sh
curl -X GET http://localhost:8080/people/100
```

**Response:**

```json
null
```

**Status Code:** 404


## Post person with included email


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"accounts": null,
	"address": "",
	"age": 10,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {
		"fsdfadsfasdfasdf@gmail.com": {
			"mail": "fsdfadsfasdfasdf@gmail.com"
		}
	},
	"hobbies": "",
	"id": "10",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Email test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": null,
	"relations": null,
	"religion": "",
	"sources": null,
	"ssn": "",
	"tags": null
}'
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 10,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {
		"fsdfadsfasdfasdf@gmail.com": {
			"mail": "fsdfadsfasdfasdf@gmail.com",
			"provider": "gmail",
			"services": {},
			"skipped_services": {},
			"src": "",
			"valid": true,
			"value": 0
		}
	},
	"gender": "",
	"hobbies": "",
	"id": "10",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Email test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"ssn": "",
	"tags": []
}
```

**Status Code:** 201


## Post person with included email detecting only discord as a services


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"accounts": null,
	"age": 10,
	"email": {
		"has_discord_account@gmail.com": {
			"mail": "has_discord_account@gmail.com"
		}
	},
	"id": "11",
	"name": "Email test"
}'
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 10,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {
		"has_discord_account@gmail.com": {
			"mail": "has_discord_account@gmail.com",
			"provider": "fake_mail",
			"services": {
				"Discord": {
					"icon": "./images/mail/discord.png",
					"link": "",
					"name": "Discord",
					"username": ""
				}
			},
			"skipped_services": {},
			"src": "",
			"valid": true,
			"value": 0
		}
	},
	"gender": "",
	"hobbies": "",
	"id": "11",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Email test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"ssn": "",
	"tags": []
}
```

**Status Code:** 201


## Post person with included email detecting all services


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"accounts": null,
	"age": 10,
	"email": {
		"all@gmail.com": {
			"mail": "all@gmail.com"
		}
	},
	"id": "12",
	"name": "Email test"
}'
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 10,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {
		"all@gmail.com": {
			"mail": "all@gmail.com",
			"provider": "gmail",
			"services": {
				"Discord": {
					"icon": "./images/mail/discord.png",
					"link": "",
					"name": "Discord",
					"username": ""
				},
				"Spotify": {
					"icon": "./images/mail/spotify.png",
					"link": "",
					"name": "Spotify",
					"username": ""
				},
				"Twitter": {
					"icon": "./images/mail/twitter.png",
					"link": "",
					"name": "Twitter",
					"username": ""
				},
				"Ubuntu GPG": {
					"icon": "https://ubuntu.com/favicon.ico",
					"link": "https://keyserver.ubuntu.com/pks/lookup?search=all@gmail.com\u0026op=index",
					"name": "Ubuntu GPG",
					"username": ""
				},
				"keys.gnupg.net": {
					"icon": "https://www.gnupg.org/favicon.ico",
					"link": "https://keys.gnupg.net/pks/lookup?search=all@gmail.com\u0026op=index",
					"name": "keys.gnupg.net",
					"username": ""
				}
			},
			"skipped_services": {},
			"src": "",
			"valid": true,
			"value": 0
		}
	},
	"gender": "",
	"hobbies": "",
	"id": "12",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Email test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"ssn": "",
	"tags": []
}
```

**Status Code:** 201


## Post person with included email and discord check failing


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"accounts": null,
	"age": 13,
	"email": {
		"discord_error@gmail.com": {
			"mail": "discord_error@gmail.com"
		}
	},
	"hobbies": "",
	"id": "13",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Email test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": null,
	"political": "",
	"prevoccupation": "",
	"relations": null,
	"sources": null,
	"tags": null
}'
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 13,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {
		"discord_error@gmail.com": {
			"mail": "discord_error@gmail.com",
			"provider": "fake_mail",
			"services": {},
			"skipped_services": {
				"Discord": true
			},
			"src": "",
			"valid": true,
			"value": 0
		}
	},
	"gender": "",
	"hobbies": "",
	"id": "13",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Email test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"ssn": "",
	"tags": []
}
```

**Status Code:** 201


## Post person with included email detected as a fake email by seekr


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"accounts": null,
	"age": 10,
	"email": {
		"fake_mail@gmail.com": {
			"mail": "fake_mail@gmail.com"
		}
	},
	"id": "14",
	"name": "Email test"
}'
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 10,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {
		"fake_mail@gmail.com": {
			"mail": "fake_mail@gmail.com",
			"provider": "fake_mail",
			"services": {},
			"skipped_services": {},
			"src": "",
			"valid": true,
			"value": 0
		}
	},
	"gender": "",
	"hobbies": "",
	"id": "14",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Email test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"ssn": "",
	"tags": []
}
```

**Status Code:** 201


## Post Person (civil status)


**Curl Request:**

```sh
curl -X GET http://localhost:8080/getAccounts/snapchat-exsists
```

**Response:**

```json
{
	"Snapchat-snapchat-exsists": {
		"bio": null,
		"blog": "",
		"created": "",
		"firstname": "",
		"followers": 0,
		"following": 0,
		"id": "",
		"lastname": "",
		"location": "",
		"profilePicture": null,
		"service": "Snapchat",
		"updated": "",
		"url": "",
		"username": "snapchat-exsists"
	}
}
```

**Status Code:** 200


## Post Person (civil status)
Possible values are: Single,Married,Widowed,Divorced,Separated

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"civilstatus": "Single",
	"id": "15"
}'
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 0,
	"bday": "",
	"civilstatus": "Single",
	"club": "",
	"education": "",
	"email": {},
	"gender": "",
	"hobbies": "",
	"id": "15",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"ssn": "",
	"tags": []
}
```

**Status Code:** 201


## Post Person (invalid civil status)
Possible values are: Single,Married,Widowed,Divorced,Separated

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"civilstatus": "Invalid",
	"id": "16"
}'
```

**Response:**

```json
{
	"message": "civil staus invalid"
}
```

**Status Code:** 400


## Post Person (missing id)


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{}'
```

**Response:**

```json
{
	"message": "missing id"
}
```

**Status Code:** 400


## Post Person (invalid religion)
Check [surce code](https://github.com/seekr-osint/seekr/blob/main/api/religion_type.go) for valid religions 

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"id": "17",
	"religion": "invalid"
}'
```

**Response:**

```json
{
	"message": "invalid religion"
}
```

**Status Code:** 400


## Post Person (invalid SSN)
Check [surce code](https://github.com/seekr-osint/seekr/blob/main/api/religion_type.go) for valid religions 

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"id": "18",
	"ssn": "invalid"
}'
```

**Response:**

```json
{
	"message": "invalid SSN"
}
```

**Status Code:** 400


## Post Person (invalid civil status)
Possible values are: Male,Female,Other

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"gender": "Invalid",
	"id": "19"
}'
```

**Response:**

```json
{
	"message": "invalid gender"
}
```

**Status Code:** 400


