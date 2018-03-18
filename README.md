# Packet Guardian API - Go

This is a Golang library to interact with the API exposed by a Packet Guardian instance.

## Supported Endpoints

| Endpoint                         | GET | POST | DELETE |
|----------------------------------|:---:|:----:|:------:|
| /api/device                      |     |  ?   |        |
| /api/device/user/:username       |     |      |   ?    |
| /api/device/reassign             |     |  ?   |        |
| /api/device/mac/:mac/description |     |  ?   |        |
| /api/device/mac/:mac/expiration  |     |  ?   |        |
| /api/device/:mac                 |  ?  |      |        |
| /api/blacklist/user/:username    |     |  X   |   X    |
| /api/blacklist/device            |     |  ?   |   ?    |
| /api/user                        |     |  ?   |   ?    |
| /api/user/:username              |  X  |      |        |
| /api/status                      |  X  |      |        |

* X = Supported
* ? = Not supported
* Blank = Endpoint doesn't exist
