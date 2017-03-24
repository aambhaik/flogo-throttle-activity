package throttle

var jsonMetadata = `{
  "name": "tibco-throttle",
  "version": "0.0.1",
  "description": "Simple Endpoint Throttle Activity",
  "inputs":[
    {
      "name": "endPoint",
      "type": "string",
      "required": true
    },
    {
      "name": "limitPerMinute",
      "type": "integer",
      "required": true
    },
    {
      "name": "disable",
      "type": "boolean"
    }
  ],
  "outputs": [
    {
      "name": "throttled",
      "type": "boolean"
    }
  ]
}`
