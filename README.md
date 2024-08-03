# Installation instructions
1) [Register github app]('https://docs.github.com/en/apps/creating-github-apps/registering-a-github-app/registering-a-github-app')

2) [create a github personal access token]('https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens)

3) give `Contents` permissions to the token in `Repository Permissions` section

### Back
1) edit `.env.dev` file, you can take variable names from `.env.example`

### Front
1) edit `fe/.env`, you can take variable names from `.env.example`

### Set up Docker
2) add port to expose frontend in `docker-compose.yml` file (default is 3000)

1) `docker-compose up --build`

# Usage
1) open `localhost:3000` in your browser

homepage:
![homepage](https://github.com/user-attachments/assets/925a93da-77f2-4f6c-b5bc-48f5d06d6a69)

2) Add Github Token, and select your github repo to view the markdown (.md) files

github token:
![token](https://github.com/user-attachments/assets/cf5f3226-18c5-443b-aa82-99567178112c)

repositories:
![Github Repo](https://github.com/user-attachments/assets/cba04add-3fca-4e09-bdef-0a80f88ee809)


3) Open the Github repo, based on Github Token Permissions, to view and edit the markdown (files). 

md file:
![md file](https://github.com/user-attachments/assets/bd78c8be-3287-4420-b5ee-9ba804ec690d)

WYSWYG editor:
![WYSWYG](https://github.com/user-attachments/assets/eaffdd60-283f-464c-99aa-0badce192468)

After you are done editing, click update to push the changes automatically to Github repo

updated file:
![image](https://github.com/user-attachments/assets/4e0e7db6-c0ff-463a-8468-322d1582e29b)

updated repo:

![updated repo](https://github.com/user-attachments/assets/2dda36b1-6fc8-424a-b5af-4a336cad61c9)