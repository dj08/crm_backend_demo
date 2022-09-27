## CRM Backend Demo App

This is a simple demo API created using Golang. It implements a trivial customer relationship management backend.
Each customer has an entry as follows:

    Id        uint
    Name      string
    Role      string
    Email     string
    Phone     string
    Contacted bool

They are stored serially in an in-memory database. Which is more or less a glorified map.

### Supported paths
The API runs on port 8000. The following requests are available:
- Getting a single customer through a http://localhost:8000/customers/{id} path
- Getting all customers through a the http://localhost:8000/customers path
- Creating a customer through a http://localhost:8000/customers path
- Updating a customer through a http://localhost:8000/customers/{id} path
- Deleting a customer through a http://localhost:8000/customers/{id} path

### How to Run
Contrary to the much awaited shorthand, you need to use the following command to get this thing running:
    go run .

You can also run some basic tests:
    go test
    
### Dependencies
The code is largely self-contained. It does depend on the following extenal libraries:
- github.com/gorilla/mux v1.8.0

Apart from this, it uses the following internal libraries:
- "encoding/json"
- "fmt"
- "io/ioutil"
- "net/http"
- "strconv"

### Sample Requests

Enjoy!