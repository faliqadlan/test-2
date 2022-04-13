# simple api

## Tech Stack

[![Go.Dev reference](https://img.shields.io/badge/-echo%20-05122A?style=flat&logo=go)](https://github.com/labstack/echo)
[![Go.Dev reference](https://img.shields.io/badge/-gorm%20-05122A?style=flat&logo=go)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/-aws--s3%20-05122A?style=flat&logo=Amazon%20Aws)](https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/go/example_code/s3)

[OPEN API](https://app.swaggerhub.com/apis/faliqadlan/simple_api/1.0.0-oas3)

[POSTMANT](https://app.getpostman.com/join-team?invite_code=6f3882e5cbe848756e6092ae1da4c01f&target_code=caacbce1308cdc5235798d61c97e631a)

<details>
<summary>HLA</summary>
<img src="challenge-2/images/hla.png">
</details>
<details>
<details>
<summary>Movie</summary>

| Feature Patient | Endpoint | Query Param                                                 | Request Body | JWT Token | Utility              |
| --------------- | -------- | ----------------------------------------------------------- | ------------ | --------- | -------------------- |
| POST            | /movie   | -                                                           | \_           | NO        | add movie            |
| DELETE          | /movie   | -                                                           | -            | NO        | delete current movie |
| PUT             | /movie   | -                                                           | -            | NO        | update current movie |
| GET             | /movie   | tittle, description, artist, genres, movie_uid, limit, page | -            | NO        | get current product  |

</details>
<p align="right">(<a href="#top">back to top</a>)</p>

# How to Use

### 1. Download

```bash
git clone https://github.com/faliqadlan/test-2.git
```

### 2.1 create file .env

```bash
touch .env
```

### 2.2 open .env

```
nano .env
```

### 2.3 write .env

```bash
PORT=8080
DB=mysql
DB_Name=<database name>
DB_PORT=3306
DB_HOST=localhot
DB_USERNAME=<username>/root
DB_PASSWORD=<password>
DB_LOC=Local
S3_REGION=<reqion AWS S3>/ap-southeast-1
S3_ID=<KEY ID AWS S3>
S3_SECRET=<SECRETKEY AWS S3>
```

# Contact

[![Gmail](https://img.shields.io/badge/gmail-faliq67%40gmail.com-red?logo=Gmail)](https://mail.google.com/)
[![Github](https://img.shields.io/badge/github-faliqadlan-lightgrey?logo=Github)](https://github.com/faliqadlan)
[![LinkedIn](https://img.shields.io/badge/linkedIn-faliqadlan-blue?logo=LinkedIn)](https://www.linkedin.com/in/faliqa/)

<p align="right">(<a href="#top">back to top</a>)</p>
