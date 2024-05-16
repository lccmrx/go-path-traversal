# Go Path Traversal

As well known, path traversal attacks are a type of attack where an attacker
tries to acquire access to restricted routes in your webserver.

Say you created some routes as shown:
- INTERNAL&nbsp; GET /hello/
- PUBLIC  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp; GET /hello/world

And you would apply some middleware routing rules between them.

> For te purpose of this example, I'm not applying any routing rules nor middleware.
> Let's assume that the `GET /hello` route is simply unknown for the user.

An attacker could simply add a `../` to walk back onto a resource.

So interestingly enough, a `GET /hello/world/..` is a valid request.

---------------------------------------------
I've created 2 samples here:
- one using Go's default HTTP package
- one using Go's Echo package, designed by LabStack.

Both of them show the same behaviour. I've not explored Echo deep enough to know, but it seems it wraps the go default Mux.


Feel free to import it onto Postman or use the cURL request in [this file](./curl.http).

Valid Server Ports | Server
-------------------|-------
8000 | Go Http
8001 | LabStack Echo
