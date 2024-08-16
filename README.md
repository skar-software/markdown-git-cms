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

## B Classic Token (for Organization repos)

![token-classic-org1](https://github.com/user-attachments/assets/6eca3d22-b198-4d9c-bae2-b52024def856)

![token-classic-org2](https://github.com/user-attachments/assets/354a194f-f722-4078-8f00-4a992e954471)


**Note: In Github profile, public email should have an email. If it does not exist, go to email settings and update.**

![image](https://github.com/user-attachments/assets/3c949a76-3b62-4d0c-a064-efebc390167d)


# Usage (screenshots to be updated)
**1) open `localhost:3000` in your browser** 
  
![image](https://github.com/user-attachments/assets/c77ac03d-6597-47ec-88d1-810bdc24d54e)

**2) Add Github Token, and select your github repo to view the markdown (.md) files** 
  
![image](https://github.com/user-attachments/assets/31154449-bc6f-4d81-913e-9e3e94efd5d8)

**3) Open the Github repo, based on Github Token Permissions, to view and edit the markdown (files).**
   
![image](https://github.com/user-attachments/assets/2dde8197-a36c-4374-a99c-9d0e47aa23ad)

**4) After you are done editing, click update to push the changes automatically to Github repo** 
  
![updated repo](https://github.com/user-attachments/assets/fed39a7b-9029-49e8-8ae6-8473eb27bfba)
