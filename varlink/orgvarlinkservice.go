// Generated with varlink-generator -- https://github.com/varlink/go/cmd/varlink-generator
package varlink

func doReplyError(c *Call, name string, parameters interface{}) error {
	return c.sendMessage(&serviceReply{
		Error:      name,
		Parameters: parameters,
	})
}

func (c *Call) ReplyInterfaceNotFound(interfaceA string) error {
	var out struct {
		Interface string `json:"interface,omitempty"`
	}
	out.Interface = interfaceA
	return doReplyError(c, "org.varlink.service.InterfaceNotFound", &out)
}

func (c *Call) ReplyMethodNotFound(method string) error {
	var out struct {
		Method string `json:"method,omitempty"`
	}
	out.Method = method
	return doReplyError(c, "org.varlink.service.MethodNotFound", &out)
}

func (c *Call) ReplyMethodNotImplemented(method string) error {
	var out struct {
		Method string `json:"method,omitempty"`
	}
	out.Method = method
	return doReplyError(c, "org.varlink.service.MethodNotImplemented", &out)
}

func (c *Call) ReplyInvalidParameter(parameter string) error {
	var out struct {
		Parameter string `json:"parameter,omitempty"`
	}
	out.Parameter = parameter
	return doReplyError(c, "org.varlink.service.InvalidParameter", &out)
}

func (c *Call) replyGetInfo(vendor string, product string, version string, url string, interfaces []string) error {
	var out struct {
		Vendor     string   `json:"vendor,omitempty"`
		Product    string   `json:"product,omitempty"`
		Version    string   `json:"version,omitempty"`
		Url        string   `json:"url,omitempty"`
		Interfaces []string `json:"interfaces,omitempty"`
	}
	out.Vendor = vendor
	out.Product = product
	out.Version = version
	out.Url = url
	out.Interfaces = interfaces
	return c.Reply(&out)
}

func (c *Call) replyGetInterfaceDescription(description string) error {
	var out struct {
		Description string `json:"description,omitempty"`
	}
	out.Description = description
	return c.Reply(&out)
}

func (s *orgvarlinkserviceInterface) VarlinkDispatch(call Call, methodname string) error {
	return nil
}

func (s *orgvarlinkserviceInterface) VarlinkGetName() string {
	return `org.varlink.service`
}

func (s *orgvarlinkserviceInterface) VarlinkGetDescription() string {
	return `# The Varlink Service Interface is provided by every varlink service. It
# describes the service and the interfaces it implements.
interface org.varlink.service

# Get a list of all the interfaces a service provides and information
# about the implementation.
method GetInfo() -> (
  vendor: string,
  product: string,
  version: string,
  url: string,
  interfaces: string[]
)

# Get the description of an interface that is implemented by this service.
method GetInterfaceDescription(interface: string) -> (description: string)

# The requested interface was not found.
error InterfaceNotFound (interface: string)

# The requested method was not found
error MethodNotFound (method: string)

# The interface defines the requested method, but the service does not
# implement it.
error MethodNotImplemented (method: string)

# One of the passed parameters is invalid.
error InvalidParameter (parameter: string)`
}

type orgvarlinkserviceInterface struct{}

func orgvarlinkserviceNew() *orgvarlinkserviceInterface {
	return &orgvarlinkserviceInterface{}
}
