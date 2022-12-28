# A Mosbot function

Was playing with the [fn project](https://fnproject.io/) locally. 
Haven't yet tried to deploy this to OCI. 

**To get up and running locally:**

 - [Install Fn](https://fnproject.io/tutorials/install/) - I have a vagrant file to do this part (I'll try to post soon).
 - Create the mosbot app `fn create app mosbot`
 - fn --verbose deploy --app mosbot --local

From there inspect to get the end point:

```bash
fn inspect function mosbot fnmosbot
```

```json
{
	"annotations": {
		"fnproject.io/fn/invokeEndpoint": "http://localhost:8080/invoke/01GNCVHNE9NG8G00GZJ000001G"
	},
	"app_id": "01GNBY0XWKNG8G00GZJ000000J",
	"created_at": "2022-12-28T17:12:43.721Z",
	"id": "01GNCVHNE9NG8G00GZJ000001G",
	"idle_timeout": 30,
	"image": "fnmosbot:0.0.21",
	"memory": 128,
	"name": "fnmosbot",
	"timeout": 30,
	"updated_at": "2022-12-28T17:12:43.721Z"
}
```

Request a url from it:

```bash
curl -X "POST" -H "Content-Type: application/json" -d '{"text":"doc 123456.1"}' http://localhost:8080/invoke/01GNCVHNE9NG8G00GZJ000001G
```

```json
{"text":"https://support.oracle.com/epmos/faces/DocumentDisplay?id=123456.1","response_type":"in_channel"}
```

Enable the UI:

It's also possible to interact with the functions from a web interface.

```bash
docker run --rm -it --link fnserver:api -p 4000:4000 -e "FN_API_URL=http://api:8080" fnproject/ui
```

Once it starts point your browser to:

http://localhost:4000

