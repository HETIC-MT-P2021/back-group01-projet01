# Image Gallery API

## Description

This repository contains all Golang API code for the image gallery project

## Usage


```git config core.hooksPath .githooks```
> Configure GitHooks


``` docker-compose up --build```
> Run the project


If you want to test the Backend API on Postman, you can use `elm_project.postman_collection.json`.

If you want to see an image after uploading it , you can see it on `http:localhost:8000/{image_id}/{image_slug).{image_extension}`


## Resources

### Images

> JSON Format

| Field           | Type                  | Description                       |
| --------------- | --------------------- | --------------------------------  |
| id              | int                   | id for the image entity           |
| name            | string                | image name                        |
| description     | string (text)         | image description (optional)      |
| created_at      | `string (y:m:d:hh:mm)`| image creation date               |
| updated_at      | `string (y:m:d:hh:mm)`| image update date                 |
| tags            | [ string ]            | image tags                        |
| category_id     | int                   | image category id                 |

> Go struct : Image

| Field           | Type                | Description                       |
| --------------- | ------------------- | --------------------------------  |
| ID              | int64               | id for the image entity           |
| Name            | string              | image name                        |
| Description     | string              | image description (optional)      |
| Slug            | string              | image slug for storage (generated)|
| Type            | string              | image type                        |
| CreatedAt       | `*time.Time`        | image creation date               |
| UpdatedAt       | `*time.Time`        | image update date                 |
| Tags            | `[]*Tags`           | image tags                        |
| CategoryID      | int64               | image category id                 |
| Category        | `*Category`         | image category                    |


### Category

> JSON Format

| Field           | Type                  | Description                       |
| --------------- | --------------------- | --------------------------------  |
| id              | int                   | id for the category entity        |
| name            | string                | category name                     |
| description     | string (text)         | category description (optional)   |
| created_at      | `string (y:m:d:hh:mm)`| category creation date            |
| updated_at      | `string (y:m:d:hh:mm)`| category update date              |

> Go struct : Category

| Field           | Type                | Description                       |
| --------------- | ------------------- | --------------------------------  |
| ID              | int64               | id for the category entity        |
| Name            | string              | category name                     |
| Description     | string              | category description (optional)   |
| CreatedAt       | `*time.Time`        | category creation date            |
| UpdatedAt       | `*time.Time`        | category update date              |


### Tags

> JSON Format

| Field           | Type                  | Description                       |
| --------------- | --------------------- | --------------------------------  |
| id              | int                   | id for the tag entity             |
| name            | string                | tag name                          |
| created_at      | `string (y:m:d:hh:mm)`| tag creation date                 |
| updated_at      | `string (y:m:d:hh:mm)`| tag update date                   |

> Go struct : Tags

| Field           | Type                | Description                       |
| --------------- | ------------------- | --------------------------------  |
| ID              | int64               | id for the tag entity             |
| Name            | string              | tag name                          |
| CreatedAt       | `*time.Time`        | tag creation date                 |
| UpdatedAt       | `*time.Time`        | tag update date                   |

## Endpoints

### LIST 

* [Get an image metadata by ID](#get-an-image-by-id)
* [Get all images metadata](#get-all-images)
* [Post an image metadata](#post-an-image-metadata)
* [Upload an image](#upload-an-image)
* [Get an image](#post-an-image)
* [Update an image](#update-an-image)
* [Delete an image](#update-an-image)
* [Get a category by ID](#get-a-category-by-id)
* [Get all categories](#get-all-categories)
* [Create a new category](#create-a-new-category) 
* [Update a category](#update-a-category)
* [Delete a category](#delete-a-category)

### Get an image by ID <a name="get-an-image-by-id"></a>

```http
GET /images/1         //where 1 is the ID of the image to be retrieved
Content-type : application/json
```
```http
HTTP/1.1 200 OK 
Content-type: application/json

{
	"id" : "1",
	"name" : "cute_cat_picture.png",
	"description" : "i are developer i make computer beep boop beep beep boop",
	"slug" : "9hjtv67dpk",
	"type" : ".png",
	"created_at" : "2020:04:05:15:53",
	"updated_at" : "2020:04:06:08:23",
	"category_id": 1,
	"tags" : ["cat","cute"]
}
```

### Get all images <a name="get-all-images"></a>

```http
GET /images     

// can be filtered by update date tag and/or category

GET /images?updated_at=asc
GET /images?updated_at=desc
GET /images?category=1
GET /images?tag=1
Content-type : application/json
```
```http
HTTP/1.1 200 OK 
Content-type: application/json

{
	"id" : 1,
	"name" : "cute_cat_picture.png",
	"description" : "i are developer i make computer beep boop beep beep boop",
	"slug" : "9hjtv67dpk",
	"type" : ".png",
	"created_at" : "2020:04:05:15:53",
	"updated_at" : "2020:04:05:15:53",
	"category_id" : 1,
	"tags" : ["cat","cute"]
},
{
	"id" : 2,
	"name" : "cute_dog_picture.png",
	"description" : "doggo",
	"slug" : "9hjtv67dpk",
	"type" : ".png",
	"created_at" : "2020:04:03:12:53",
	"updated_at" : "2020:04:03:12:53",
	"category_id" : 1,
	"tags" : ["cat","cute"]
}
```

### Post an image metadata<a name="post-an-image-metadata"></a> 

``` http
POST /images
Content-type : application/json
{
	"name" : "cute_dog_picture.png",
	"description" : "doggo",
	"category_id" : 1,
	"tags" : ["dog","cute"]
}
```

```http
HTTP/1.1 200 OK 
Content-type: application/json

{
	"id" : 2,
	"name" : "cute_dog_picture.png",
	"description" : "doggo",
	"created_at" : "2020:04:03:12:53",
	"updated_at" : "2020:04:03:12:53",
	"slug" : "9hjtv67dpk",
	"type" : ".png",
	"category_id" : 1,
	"tags" : ["dog","cute"]
}

```

### Upload an image <a name="upload-an-image"></a>
``` http
POST /upload/{image_id}
Content-type : multipart/form-data

key: "file"
```

### Get an image <a name="get-an-image"></a> 

``` http
GET /uploads/{image_id}/{image_slug}.{image_extension)
Content-type : application/json

```

### Update an image <a name="update-an-image"></a>

``` http
PUT /images/2
Content-type : application/json
{
	"name" : "doggy",
	"description" : "my dogoo",
	"tags": ["dog", "doggy"]
}
```

```http
HTTP/1.1 200 OK 
Content-type: application/json

{
	"id" : 2,
	"name" : "doggy",
	"description" : "my dogoo",
	"created_at" : "2020:04:03:12:53",
	"updated_at" : "2020:04:04:17:28",
	"slug" : "9hjtv67dpk",
	"type" : ".png",
	"category_id": 1,
	"tags" : ["dog","doggy"]
}

```

### Delete an image <a name="delete-an-image"></a>

``` http
DELETE /images/2                       // deletes image only
DELETE /images/2?delete_mode=hard      // deletes both image and image metadata
Content-type : application/json
```

```http
HTTP/1.1 204 No Content 
Content-type: application/json
```

### Get a category by ID <a name="get-a-category-by-id"></a>

```http
GET /categories/1      
Content-type : application/json
```
```http
HTTP/1.1 200 OK 
Content-type: application/json

{
	"id" : 1,
	"name" : "cars",
	"description" : "vroum",
	"created_at" : "2020:04:05:15:53",
	"updated_at" : "2020:04:06:08:23",
}
```

### Get all categories <a name="get-all-categories"></a>

```http
GET /categories                          
GET /categories?updated_at=desc         
Content-type : application/json
```
```http
HTTP/1.1 200 OK 
Content-type: application/json

{
	"id" : 1,
	"name" : "cars",
	"description" : "vroum",
	"created_at" : "2020:04:05:15:53",
	"updated_at" : "2020:04:05:15:53",
}, 
{
	"id" : 2,
	"name" : "animals",
	"description" : "A collection of animal images",
	"category_id": 1,
	"created_at" : "2020:04:05:15:53",
	"updated_at" : "2020:04:05:15:53",
}

```


//TODO : add get, get all, update, delete

### Create a new category <a name="create-a-new-category"></a>

``` http
POST /categories
Content-type : application/json
{
	"name" : "animals",
	"description" : "A collection of animal images"
}
```

```http
HTTP/1.1 200 OK 
Content-type: application/json

{
	"id" : 2,
	"name" : "animals",
	"description" : "A collection of animal images",
	"created_at" : "2020:04:05:15:53",
	"updated_at" : "2020:04:05:15:53",
}
```


### Update a category <a name="update-a-category"></a>

``` http
PUT /categories/2
Content-type : application/json
{
	"name" : "memes",
	"description" : "animal memes"
}
```

```http
HTTP/1.1 200 OK 
Content-type: application/json

{
	"id" : 2,
	"name" : "memes",
	"description" : "animal memes",
	"created_at" : "2020:04:05:15:53",
	"updated_at" : "2020:04:06:08:23",
}

```

### Delete a category <a name="delete-a-category"></a>

``` http
DELETE /categories/1
Content-type : application/json

```

```http
HTTP/1.1 204 No Content 
Content-type: application/json
```

