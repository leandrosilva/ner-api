# NER-API

This is a simple API to play with `Named Entity Recognition` and `Go`. Why? Well, why not.

## Try it out

In one terminal:

    $ go run *.go

And in a second terminal:

    $ curl -i -XPOST http://localhost:8080/train \
           -H "Content-Type: multipart/form-data" \
           -F "dataset=@./samples/superheroes.jsonl"

You should get something like:

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Sun, 11 Aug 2019 21:10:02 GMT
    Content-Length: 36

    {"Success":true,"Label":"SUPERHERO"}

Then fire a quick analysis:

    $ curl -i -XPOST http://localhost:8080/recognize \
           -H "Content-Type: application/json" \
           -d'{"Model":"SUPERHERO","content":"Batman blah blah blah before Thor blah blah. Well, of course, blah blah blah in Gotham City, and blah blah blah while The Joker and Wolverine was in Canada. Batman again. Annnd Batman blah blah again. How about Thor? I do not know."}'

Annnd boom! There you have it:

    Content-Type: application/json
    Date: Sun, 11 Aug 2019 21:10:08 GMT
    Content-Length: 165

    {
        "Success": true,
        "Model": "SUPERHERO",
        "Entities": [{
            "Entity": {
                "Text": "Batman",
                "Label": "SUPERHERO"
            },
            "Count": 3
        }, {
            "Entity": {
                "Text": "Thor",
                "Label": "SUPERHERO"
            },
            "Count": 2
        }, {
            "Entity": {
                "Text": "Joker",
                "Label": "SUPERVILLAIN"
            },
            "Count": 1
        }, {
            "Entity": {
                "Text": "Wolverine",
                "Label": "SUPERHERO"
            },
            "Count": 1
        }]
    }

## How about edge cases?

Nope. Just happy & nice paths. Get over it.
