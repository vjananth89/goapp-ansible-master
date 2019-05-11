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

## EC2 Host Details
