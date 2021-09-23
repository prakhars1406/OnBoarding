# Recommendation Rendering Service
RRS manages WRU APIs which at a high level are of 2 types,<br/>
 - APIs to fetch recommendations
 - APIs to push data like articles and events

All APIs are documented at [API-document](https://alfred.wru.ai/api/)

# How to build and run locally

**1. Prerequisites for the set up:**<br />
- 1.1) MongoDB<br />

>One can use the docker-compose.yml file to run all the data stores as containers using the command<br />
> `docker-compose up -d`<br />
>If not continue with the downloads below<br /> 

- 1.2) GoLang minimum 1.13
- 1.3) Docker Desktop

 >**Note:** Skip 2 if 1.1 is installed using Docker compose file                       [go to (3)]<br />

**2. Installing prerequisites Locally:**

- 2.1) Download and install MongoDB

**3. Creating mongo sims database collection:**<br />
  - Download and extract json file from the given link: [Json-File]()<br />

 >**Note:** Skip 3.3, if installed using Docker compose file <br />
  - 3.1) If using mongo on Docker desktop use the below command<br />
       - `docker cp Json_Directoy_Path <conatiner_id>:/var/lib/Json`<br />
  - 3.2) Open mongo shell running on docker<br/>
       - `Go to /var/lib/Json in docker container`                        
  - 3.3) If you are doing it locally<br />
       - `Go to the path which contains the json file in it.`<br />

  - Run all the following commands one after the other: <br />

    `mongoimport --db sims --collection experiments --file experiments.json --jsonArray`

    `mongoimport --db sims --collection column_mappings --file column_mappings.json --jsonArray`

    `mongoimport --db sims --collection merchant --file merchant.json --jsonArray`

    `mongoimport --db sims --collection merchant_config --file merchant_config.json --jsonArray`

    `mongoimport --db sims --collection merchant_extract_ref --file merchant_extract_ref.json --jsonArray`

    `mongoimport --db sims --collection merchant_pages --file merchant_pages.json --jsonArray`

    `mongoimport --db sims --collection merchant_platform --file merchant_platform.json --jsonArray`

    `mongoimport --db sims --collection pageType --file pageType.json --jsonArray`

    `mongoimport --db sims --collection page_recommendations --file page_recommendations.json --jsonArray`

    `mongoimport --db sims --collection poll_collection --file poll_collection.json --jsonArray`

    `mongoimport --db sims --collection user --file user.json --jsonArray`

    `mongoimport --db sims --collection sims_job_progression --file sims_job_progression.json --jsonArray`

    `mongoimport --db sims --collection customers_role --file customers_role.json --jsonArray`

    `mongoimport --db sims --collection customers --file customers.json --jsonArray`

    `mongoimport --db sims --collection fieldIndicesForSimsType --file fieldIndicesForSimsType.json --jsonArray`

    `mongoimport --db sims --collection route_description --file route_description.json --jsonArray`

    `mongoimport --db sims --collection route_info --file route_info.json --jsonArray`

> Once you are done executing all the above commands, as a result it should reflect the mongoDB with 15 new collections in “sims”.<br />

**4. Setting up DSS-RRS Locally:**
- To set up DSS-RRS you need to have access to GitLab and a GitHub account.<br />

  - 4.1) Run the below command to clone

      - `git --version`                                       
      - `git clone 
      - `go mod download`

  - 4.2) To run the DSS-RRS service execute the below command<br>(This command will be used always to run the service)<br />

      - `go run app/*.go -file=dev.json`

  - 4.4) To make sure that the application is loaded successfully open the below URL<br/>
      - [output](http://localhost:9001/metrics)
> If you get the Output status in the Browser it's becuase the application is loaded successfully.
