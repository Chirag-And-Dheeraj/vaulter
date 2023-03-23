# Thinking

## Validation

### What?

- email -> valid email
- pin -> length 6
- 

## File Upload flow

- first api call --> send metadata of file from client to server
- generate hash of the file on server (!important, don't know why, but it seems important, think about this...)
- create an row in the db
- is_uploaded = false
- return the created instance to the client
- client uses the hash generated to rename the file
- second api call --> send the actual file to the server
- once file is uploaded successfully, is_uploaded = true
- respond to the client
- done

