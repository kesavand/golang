/*package provides the interfaces for data srialization*/
package api


type serialize interface {
	marshal(interface{}) ([]byte)
	unmarshal([]byte) ([]byte)
}
