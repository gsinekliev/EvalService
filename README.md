# EvalService
Service that evaluates mathematical expressions written in english

## Running locally
To run locally do `make run` from the root directory of the project.

## Run the tests
To run the unit tests do `make test` from the root directory of the project.

## Endpoints

### POST /evaluate
Endpoint accepts JSON in the format:
```
{
    "expression": <expression>
}
```
And returns JSON:
```
{
    "result": <result>
}
```
Where `<result>` is a floating point number or error and `<expression>` can be described as:
```
<expression> := What is <number> (<operator> <number>)* ?
<operator> := plus|minus|multiplied by|divided by
<number> := any integer
```
Examples:
```
What is 3?                     --> 3
What is 4 multiplied by 10?    --> 40
What is 4 divided by 2 plus 3? --> 5
```

### POST /validate
Endpoint accepts JSON in the format:
```
{
    "expression": <expression>
}
```
And returns JSON:
```
{
    "valid": <valid>
    (optional) "reason": <reason>
}
```
### GET /errors
Outputs list of the validation errors with their frequency(errors for the same expression and endpoint are considered the same)