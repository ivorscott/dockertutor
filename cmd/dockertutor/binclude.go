// Code generated by binclude; DO NOT EDIT.

package main

import (
	"github.com/lu4p/binclude"
	"time"
)

var (
	_binclude0	= []byte("[\n  {\n    \"Title\": \"LESSON 1: TESTING YOUR DOCKER COMPOSE VERSION\",\n    \"Exercise\": \"Run docker-compose --version to check the version of Docker Compose installed on your machine.\",\n    \"Answer\": \"docker-compose --version\",\n    \"Explanation\": \"\",\n    \"Complete\": false,\n    \"Setup\": [],\n    \"Teardown\": []\n  },\n  {\n    \"Title\": \"LESSON 2: RUN DOCKER COMPOSE UP \",\n    \"Exercise\": \"Using Compose is basically a three-step process:\\nDefine your app’s environment with a Dockerfile so it can be reproduced anywhere.\\nDefine the services that make up your app in docker-compose.yml so they can be run together in an isolated environment.\\nRun docker-compose up and Compose starts and runs your entire app.\\n\\nIn your project directory under compose/lesson2 you have the following a Dockerfile and docker-compose.yml: \\nStart up your application by running docker-compose up\",\n    \"Example\": \"compose/lesson2\",\n    \"Answer\": \"docker-compose up\",\n    \"Explanation\": \"Compose pulls a Redis image, builds an image for your code, and starts the services you defined. In this case, the code is statically copied into the image at build time.\",\n    \"Complete\": false,\n    \"Setup\": [],\n    \"Teardown\": []\n  }\n]\n")
	_binclude1	= []byte("[\n  {\n    \"Title\": \"LESSON 1.1: TESTING YOUR DOCKER VERSION\",\n    \"Exercise\": \"Run docker version to check the version of Docker installed on your machine.\",\n    \"Answer\": \"docker version\",\n    \"Explanation\": \"\",\n    \"Complete\": false,\n    \"Setup\": [],\n    \"Teardown\": []\n  },\n  {\n    \"Title\": \"LESSON 1.2: TEST DOCKER INSTALLATION\",\n    \"Exercise\": \"Test that your installation works by running the hello-world Docker image. Run docker run hello-world.\",\n    \"Answer\": \"docker run hello-world\",\n    \"Explanation\": \"\",\n    \"Complete\": false,\n    \"Setup\": [],\n    \"Teardown\": []\n  }\n]\n")
	_binclude2	= []byte("[\n  {\n    \"Title\": \"LESSON 1.1: ENABLE SWARM MODE\",\n    \"Exercise\": \"Run docker swarm init to enable Docker swarm mode on your machine.\",\n    \"Answer\": \"docker swarm init\",\n    \"Explanation\": \"\",\n    \"Complete\": false,\n    \"Setup\": [],\n    \"Teardown\": [\"docker swarm leave --force\"]\n  },\n  {\n    \"Title\": \"LESSON 1.2: LIST THE NODES IN YOUR CLUSTER\",\n    \"Exercise\": \"Currently, you have a single node swarm cluster. List the nodes in your swarm. Run docker node ls\",\n    \"Answer\": \"docker node ls\",\n    \"Explanation\": \"\",\n    \"Complete\": false,\n    \"Setup\": [\"docker swarm init\"],\n    \"Teardown\": [\"docker swarm leave --force\"]\n  }\n]\n")
	BinFS		= &binclude.FileSystem{Files: map[string]*binclude.File{"../../lessons": {Filename: "lessons", Mode: 2147484141, ModTime: time.Unix(1593372841, 830025974), Compression: binclude.None, Content: nil}, "../../lessons/docker-compose.json": {Filename: "docker-compose.json", Mode: 420, ModTime: time.Unix(1593372841, 829754877), Compression: binclude.None, Content: _binclude0}, "../../lessons/docker.json": {Filename: "docker.json", Mode: 420, ModTime: time.Unix(1593183709, 737535923), Compression: binclude.None, Content: _binclude1}, "../../lessons/swarm.json": {Filename: "swarm.json", Mode: 420, ModTime: time.Unix(1593183709, 738394770), Compression: binclude.None, Content: _binclude2}}}
)