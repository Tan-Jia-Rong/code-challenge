# Switcheo Code Challenge

Switcheo Code Challenge for Backend Engineer Internship Position

Applicant: Tan Jia Rong

### Problem 4: Three ways to sum to n
Three unique implementation of sun-to-n is provided in the [`sum-to-n.go`](src/problem4/sum-to-n.go) file.

Steps to run:
```shell
cd src/problem4

go run sum-to-n.go
```

### Problem 5: A Crude Chain
References: https://docs.ignite.com/guide/blog/intro

**NOTE: Unable to implement using scaffolded code provided. Hence, a blog application is built from scratch using Ignite CLI**

***Problems Faced:***
Blockchain has to be updated from v0 to v1. which is not supported by Ignite CLI v0.24. 

Upon following the migration steps outlined in the documents and updating to v0.27, there seems to be some missing dependencies and import statements. ```[see branches: migrate]```

**CRUD Commands:**
1. `create-post`: Creates a new post with a title and body `CRUD: C`
2. `show-post`: Show the details of an existing post associated with the given post id `CRUD: R`
3. `list-post`: Shows details of all existing posts `CRUD: R`
4. `update-post`: Updates an existing post `CRUD: U`
5. `delete-post`: Deletes an existing post `CRUD: D`

Steps to run to start the server:
```shell
cd src/problem5/blog

ignite chain serve
```

Steps to run after starting server:
```shell
cd /Users/{your computer name}/go/bin

# Create a blog post with title hello, body world associated to creator alice
./blogd tx blog create-post hello world --from alice

# Show blog post associated with the post id: 0
./blogd q blog show-post 0

# Create a blog post with title foo, body bar associated to creator bob
./blogd tx blog create-post foo bar --from bob

# List all blog post
./blogd q blog list-post    

# Update existing blog post
./blogd tx blog update-post hello cosmos 0 --from alice

# Verify that it has been updated
./blogd q blog show-post 0

# Delete blog post
./blogd tx blog delete-post 0 --from alice

# Verify that is has beed deleted
./blogd q blog list-post
```

### Problem 6: Cosmos Architecture

1. Diagrams can be found in the [`src/problem6/diagrams`](src/problem6/diagrams) directory.

2. Documentation can be found in the [`src/problem6/README.md`](src/problem6/README.md) file.

3. Shortcomings and potential improvements can be found in [`src/problem6/README.md`](src/problem6/README.md) under **Section 2**.