I need some thoughts on this...

Feature to be implemented: File upload from client, extracting and storing metadata of the files in a database table, generating a hash and saving the file either on the local server or static file storage with the filename as the hash string.

Here's my approach to this: Client side pr, take the file as input, extract metadata and send it the the backend (only metadata) --> First API Call.

On the backend, take the metadata, create a row in DB, and mark the flag is_uploaded=False.

Then generate a hash using the metadata^ as the input and saving that in the DB, and sending back a 201 response along with the hash as a response.

Then using the hash, take the file and send it to the server --> Second API call.

On the backend, rename the file as the hash, then store it locally or in the static file storage, updating is_uploaded as True and sending a 201 response.

Is this approach good? bad? seriously bad? Can I change/improve something?

Let me know. 
https://wa.me/+918451904417