# NER-API

This is a simple API to play with Named Entity Recognition and Go. Why? Why not.

## Try it out

In one terminal:

    $ go run *.go

And in a second terminal:

    $ curl -i -XPOST http://localhost:8080/train -H "Content-Type: multipart/form-data" -F "dataset=@./samples/superheroes.jsonl"

You should get something like:

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Sun, 11 Aug 2019 21:10:02 GMT
    Content-Length: 36

    {"Success":true,"Label":"SUPERHERO"}

Then fire a quick search:

    $ curl -i -XPOST http://localhost:8080/recognize -H "Content-Type: application/json" -d'{"Model":"SUPERHERO","content":"Batman blah blah blah Thor, blah blah in Gotham City, blah blah blah The Joker"}'

Annnd boom!

    Content-Type: application/json
    Date: Sun, 11 Aug 2019 21:10:08 GMT
    Content-Length: 165

    {"Success":true,"Model":"SUPERHERO","Entities":[{"Entity":{"Text":"Batman","Label":"SUPERHERO"},"Count":1},{"Entity":{"Text":"Thor","Label":"SUPERHERO"},"Count":1}]}

## How about edge cases?

Nope. Just happy & nice paths. Get over it.