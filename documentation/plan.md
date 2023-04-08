![[buildspace]]

- Images --> Photos    |
- Videos --> Movies    | --> Memories?
- Audio --> Songs       |
- PDFs --> Documents

- chunked uploads for—
	- images
	- videos
	- audio
	- all documents (.pdf, .docx, .xls, .pptx. etc..)
- wrangled data store
	- can be decrypted and formed back to file only on client
- self-host
	- raspberry-pi
	- organization/group/family/friends
		- one user account + pool storage (don't track who uploaded/deleted/modified) ? granular permissions? download only?
		- different accounts + pool storage (track who uploaded/deleted/modified)— regular version
		- different accounts + partitioned storage— VaulterPRO
	- options to host on—
		- GCP
		- AWS

- folder structure—
	- user_root_folder (~/)
		- file_1 (~/file_1)
		- file_2
		- ...
		- file_n
		- folder_1 (~/folder_1)
			- folder_1_file_1(~/folder_1/folder_1_file_1)
			- ...
		- folder_2
		- ...
		- folder_n

- flow
	- (/) ==> index.html ==> landing page
	- (/vault) ==> vault.html ==> actual vault
		- when visiting vault^ if session not exist, redirect (/login)
	- (/login) ==> login.html ==> login page
		- when visiting login^ if make db call, check if master account does not exist, redirect (/setup)
	- (/setup) ==> setup.html ==> register page
	- (/viewer/{file-hash}) ==> viewer.html ==> file viewing page

- pages
	- index.html
	- vault.html
	- login.html
	- setup.html
	- viewer.html
	- dashboard.html
