## User
- id
- first_name
- last_name
- email
- master_password
- profile_photo
- profile_creation_datetime
- active

## File
- id
- name
- owner (foreignkey User.id)
- size
- type
- parent_folder(foreignkey Folder.id)
- upload_time
- modified_time

## Folder
- id
- name
- owner (foreignkey User.id)
- parent_folder (foreignkey Folder.id)
- creation_time
- modified_time