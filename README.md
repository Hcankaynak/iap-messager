# iap-messager



## Installation

In project directory you can run the following command to start a container

```bash
docker compose up
```
<hr>

## Usage

After starting the container, you can access the application at `http://localhost:3000`

There is also a swagger documentation available at `http://localhost:5000/swagger/index.html`

<hr>

## Workflow

![Alt text](./workflow.png)

1- Automatic message sender triggered.
2.1- There may be a more than one service running so, When service retrieve the message information save to redis to avoid other services to send the same message.
2.2- Service retrieve list of message id that will be excluded.
3- Service retrieve messages from postgresql. Exclude messages that retrieved from redis.
4- Service make a request to related webhook.
5- Service save response to redis.
