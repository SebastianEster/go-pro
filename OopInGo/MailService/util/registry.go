package util

var registeredInterfaces = make(map[string]interface{})

func Register(interfaceName string, interfaceInstance interface{}) {
	registeredInterfaces[interfaceName] = interfaceInstance
}

func Get(interfaceName string) interface{} {
	return registeredInterfaces[interfaceName]
}
