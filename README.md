<div id="top"></div>

<br />

<h3 align="center">My Blog</h3>

  <p align="center">
    Simple Blog Restful API in Golang
    <br />
    <a href="https://github.com/its-lana/my-blog/issues">Report Bug</a>
    ·
    <a href="https://github.com/its-lana/my-blog/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
        <li><a href="#feature">Feature</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#consume-api">How to consume API</a></li>
      </ul>
    </li>
  </ol>
</details>

<br/>

## About The Project

<br/>

### Built With

-  Golang
-  MySQL
-  HttpRouter
-  Testify

<p align="right">(<a href="#top">back to top</a>)</p>

### Feature

-  Restful API
-  The use of Interfaces in the project
-  Functions are made modular so they are easy to reuse
-  Authorize Api Gateway using X-Api-Key
-  Error Handling
-  There is a model formatting for both request and response
-  Testing that directly hits the API using testify

<p align="right">(<a href="#top">back to top</a>)</p>

## Getting Started

This will give instructions on setting up your project locally. To get a local copy up and running follow these simple example steps.

### Installation

1. Clone the repo

   ```sh
   https://github.com/its-lana/my-blog.git
   ```

2. Create and Setting Database <br/>
   For a guide to creating a database and creating a table posts, you can open the file <a href="https://github.com/its-lana/my-blog/tree/main/app/init.sql">app/init.sql</a>

3. Setting Database Password <br/>
   Open file <a href="https://github.com/its-lana/my-blog/tree/main/app/database.go">app/database.go</a>, then change password with your password

4. Run project
   ```sh
   go run main.go
   ```
5. For testing, you can following this step. <br/>
   a. create database for test, eg. myblog_test <br/>
   b. create table posts <br/>
   c. open file <a href="https://github.com/its-lana/my-blog/tree/main/test/post_controller_test.go">test/post_controller_test.go</a> <br/>
   d. change password with your MySQL password <br/>
   e. enjoy with test <br/>

<p align="right">(<a href="#top">back to top</a>)</p>

<br/>

### Consume API

<!-- <h3 align="left">How to Consume API</h3> -->

You can see the guide in the following file : <a href="https://github.com/its-lana/my-blog/tree/main/test.http">Test.http</a>

<p align="right">(<a href="#top">back to top</a>)</p>
