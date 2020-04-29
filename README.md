**Send Message**
----
API for sending a message and send it to storage (in this case we just save it in-memory)
  
  * **URL**
  
    `/v1/message/send`
  
  * **Method:**
  
    `POST`
  
  * **Success Response:**
  
    * **Code:** 200 <br />
      **Content:** 
      ```
      {
        "success": true,
        "errorMsg": "",
        "data": [
          {
            "content": "test1"
          },
          {
            "content": "test2"
          }
        ]
      }
      ```
   
  * **Error Response:**
  
      * **Code:** 400 Bad Request <br />
      **Content:**
      ```
      {
          "success": false,
          "errorMsg": "message content cannot be empty"
      }
      ```
    
      OR
      
      * **Code:** 500 Internal Server Error <br />
          **Content:**
      ```
      {
          "success": false,
          "message": "failed to send message"
      }
      ```
  
  * **Sample cURL request:**
  
    ```
    curl --request POST \
      --url http://localhost:8080/v1/message/send \
      --header 'content-type: application/json' \
      --data '{
    	"message": {
    		"content": "test2"
    	}
    }'
    ```

----

**Get All Messages**
----
Get all previously sent messages.

* **URL**

  `/v1/message/get`

* **Method:**

  `GET`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
    ```
    {
      "success": true,
      "errorMsg": ""
    }
    ```

* **Sample cURL command:**

  ```
  curl --request POST \
    --url http://localhost:8080/v1/message/send \
    --header 'content-type: application/json' \
    --data '{
  	"message": {
  		"content": "test2"
  	}
  }'
  ```
  
  **Realtime Message**
  ----
  API for sending and displaying back message in real time using WebSocket.
    
    * **URL**
    
      `/v1/message/realtime`
    
    * **Method:**
    
      `GET (Upgraded to WebSocket)`
    
    * **Success Response:**
    
      * **Code:** 101 (Switching Protocols) <br />
        **Content:** 
        ```
        {
          "success": true,
          "errorMsg": "",
          "data": [
            {
              "content": "test1"
            },
            {
              "content": "test2"
            }
          ]
        }
        ```
     
    * **Error Response:**
    
        * **Code:** 400 Bad Request <br />
        **Content:**
        ```
        {
            "success": false,
            "errorMsg": "cannot open websocket connection"
        }
        ```
      
        OR
        
        * **Code:** 500 Internal Server Error <br />
            **Content:**
        ```
        {
            "success": false,
            "message": "<specific error message>"
        }
        ```
    
    * **Sample cURL request:**
    
      ```
      curl --include \
           --no-buffer \
           --header "Connection: Upgrade" \
           --header "Upgrade: websocket" \
           --header "Host: localhost:8080" \
           --header "Origin: http://localhost:8080" \
           --header "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" \
           --header "Sec-WebSocket-Version: 13" \
           http://localhost:8080/v1/message/realtime
      ```
  
  For testing purpose we need another tool beside cURL. Because, apart from being usable to test the initial handshake,
  curl has no support for WebSocket. It is impossible to actually exchange data using WebSocket with cURL.
  I used Simple WebSocket Client which is actually a Chrome extension for testing WebSocket.
  https://chrome.google.com/webstore/detail/simple-websocket-client/pfdhoblngboilpfeibdedpjgfnlcodoo/related?hl=en
  
  Sample request:
  ```
  {
    "content": "test2"
  }
  ```
  ----
  
  To run the code, simply run this following command:
  ```
  $ cd ./cmd/realtime-api
  $ go run app.go
  ```