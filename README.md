# tibco-throttle
This activity provides your flogo application the ability to throttle traffic to a target endpoint.


## Installation

```bash
flogo add activity github.com/TIBCOSoftware/flogo-contrib/tree/master/activity/throttle
```

## Schema
Inputs and Outputs:

```json
{
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
}
```
## Settings
| Setting     | Description    |
|:------------|:---------------|
| endPoint  | The target endpoint url |         
| limitPerMinute    | Limit number of invocations per minute |
| disable       | Flag to disable the throttle policy |
Note: if disable is set to true, limitPerMinute is ignored
## Configuration Examples
### Increment
Configure a task to increment a 'messages' counter:

```json
{
  "id": 3,
  "type": 1,
  "activityType": "tibco-throttle",
  "name": "Throttle the target API invocations",
  "attributes": [
    { "name": "endPoint", "value": "http://localhost:9980/hello" },
    { "name": "limitPerMinute", "value": 5 }
  ]
}
```