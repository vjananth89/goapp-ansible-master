# goapp-ansible-master

## The Go Application

This repo consists of a simple web application built in Golang/Postgres. The application presents an html form at its root and stores its information in a Postgres server in the backend.

 - The application uses the gorilla/mux, which is a URL router and
   dispatcher for golang.
  -  The other package used in this project is the
   `pq`, which is a GO Postgres driver for the `database/sql` package.

The project can be tested locally by downloading the contents of the repo `$GOHOME` . Once you navigate to the goapp-postgres-master folder, run the command below to execute the file and launch the application at `localhost:8080`

    go run main.go personHandler.go store.go

*Note: You will also require Postgres to be installed before you can successfully use the application. The information specific for Postgres is detailed below.*

#### Postgres Setup Overview

The Postgres server runs locally and can be accessed via port `5432` in the T2 micro EC2 instance that is referenced in the Ansible playbook and the inventory file in the repo. 

The guide for setting up the postgres instance in this project can be found in the link [here](http://postgresguide.com/setup/install.html).

    -- Table: public.person
    
    -- DROP TABLE public.person;
    
    CREATE TABLE public.person
    (
        name character(32) COLLATE pg_catalog."default" NOT NULL,
        color character varying(32) COLLATE pg_catalog."default" NOT NULL,
        pet_pref character varying(32) COLLATE pg_catalog."default" NOT NULL,
        CONSTRAINT data_pkey PRIMARY KEY (name)
    )
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

## Ansible Setup Overview

Ansible can be setup using the command below:

    sudo yum install ansible

The `playbook.yml` in the repo has all the steps necessary to download the ssh into the EC2 micro instance and deploy this golang app.

*Note: The playbook will only deploy the instance in the cloud referenced in the inventory file. The inventory file will have to be updated accordingly in order to deploy and test this application on different servers.*

The Ansible Playbook does the following:

 - It cleans up the workspace before a new deploy.
 - It downloads the artifacts from this repo - [https://github.com/vjananth89/goapp-ansible-master](https://github.com/vjananth89/goapp-ansible-master)
 - Downloads the packages necessary for this project to be launched (Golang Mux and Postgres pq)
 - Launches the app!

## EC2 Host Details

This project has been deployed to an EC2 instance in AWS at [http://ec2-34-239-126-166.compute-1.amazonaws.com:8080](http://ec2-34-239-126-166.compute-1.amazonaws.com:8080/)
