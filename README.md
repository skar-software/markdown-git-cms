# Installation instructions
**1) Run docker**

docker pull anonranger/pubdesk

docker run --name pubdeskmd -d -p 3000:4173 anonranger/pubdeskmd
   
4) give `Contents` permissions to the token in `Repository Permissions` section

### Back
1) add `.env.dev` file, you can take variable names from `.env.example`

### Front
1) add `fe/.env`, you can take variable names from `.env.example`

### Set up Docker
1) Run below command to compile docker image from code
`docker run --name pubdeskmd -e DOMAIN_NAME= -e GITHUB_APP_ID= -e GITHUB_CLIENT_ID= -e GITHUB_SECRET= -d -p PORT:4173 sahildhingraa/pubdeskmd`
2) replace `PORT`with the port you want to run the app on

# Usage
**1) open `localhost:3000` in your browser** 
  
![homepage](https://github.com/user-attachments/assets/a3cdb411-de43-42c3-906e-d862881a7aea)

**2) Add Github Token, and select your github repo to view the markdown (.md) files** 
  
![token](https://github.com/user-attachments/assets/1a12258b-d6c2-49ed-9354-67ba04a460ac)

![GitHub Repo](https://github.com/user-attachments/assets/4c56130d-35bc-4e60-b97f-9626bfe3557f)

**3) Open the Github repo, based on Github Token Permissions, to view and edit the markdown (files).**
   
![md file](https://github.com/user-attachments/assets/51ab08e0-df17-499f-9dba-e03ae0f8757f)

![WYSWYG](https://github.com/user-attachments/assets/fa6c9d9d-f63f-420e-ad35-33b6236ef940)

**4) After you are done editing, click update to push the changes automatically to Github repo** 
  
![updated file](https://github.com/user-attachments/assets/c744a2d9-297c-419d-b5a1-a76cd112e348)

![updated repo](https://github.com/user-attachments/assets/fed39a7b-9029-49e8-8ae6-8473eb27bfba)
