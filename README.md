# PubdeskMD

## Markdown editor with Github integration - to edit and update files

# Installation instructions
**1) Run the code below to open the application**

Pull docker image from Docker Hub
```
docker pull anonranger/pubdeskmd
```

Start the docker server at localhost:3000
```
docker run --name pubdeskmd -d -p 3000:4173 anonranger/pubdeskmd
```

**2) Create a Personal Access Token**

A) Click on Profile -> Settings -> Developer Settings -> Personal Access Token (Fine grained token) - for **Personal** repos

OR

B) Click on Profile -> Settings -> Developer Settings -> Personal Access Token (Classic token) - for **Organization** repos

## A Fine grained Token (for Personal repos)

**Security recommendation:** When creating a Fine grained token, keep the validity window to a small timeline. Also give access only to specific repos you want to edit.

![token-perm1](https://github.com/user-attachments/assets/d3784bde-67c7-408b-9653-f226458c0e6b)

![token-perm2](https://github.com/user-attachments/assets/3efd6dd1-ae93-4990-85cb-faeab8493982)

![token-perm3](https://github.com/user-attachments/assets/294e637c-e867-4bfb-975f-b676a2499603)

![token-perm4](https://github.com/user-attachments/assets/c9d68906-10dd-47be-8ed1-3ae80faa5098)

## B Classic Token (for Organization repos)**

![token-classic-org1](https://github.com/user-attachments/assets/6eca3d22-b198-4d9c-bae2-b52024def856)

![token-classic-org2](https://github.com/user-attachments/assets/354a194f-f722-4078-8f00-4a992e954471)


**Note: In Github profile, public email should have an email. If it does not exist, go to email settings and update.**

![image](https://github.com/user-attachments/assets/3c949a76-3b62-4d0c-a064-efebc390167d)


# Usage (screenshots to be updated)
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
