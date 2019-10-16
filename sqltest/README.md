# SQL TEST

A spike to test availability of a SQL server instance over VPN/VPC

## Pre-requisites

1. AWS CLI
2. Dep

## Usage

1. fetch dependencies

    ```bash
    make vendor
    ```

2. Copy credentials from LastPass  ---> .env 

3. populate env vars

    ```bash
    export $(grep -v '^#' .env | xargs)
    ```


to test local connectivity:

    ```bash
    make testLocal
    ```


to deploy to AWS:

    ```bash
    make createfunc
    ```
