Okay so basically its a Logging service. Its an api you can call from the backend (prefered) to log something. the Structure of such an logobject may look something like this: 

```json
{
    "severity": "High|Critical|Security......",¨
    "message": "Some random message",
    "request": {
        "jsonObject of the request"
    },
    "userId": "UUID",
    "requestUrl": "url",
    "response": "",
    "requestKey": "",
    "customLifetime": "123"
}

```


And then this gets sent via the backend to another microservice which then stores this in a seperate Database. Then this can be displayed on a pretty webapp with a lot of filters and search options. 


Also each severity has a lifetime, which can be set by the admins of the log service. They can change it inside of their setting inside the webapp. 

This basic implementation will only manage a singular app, so you cant connect multible backends to the same log instance. each one has to get a new microservice as well as db. 

