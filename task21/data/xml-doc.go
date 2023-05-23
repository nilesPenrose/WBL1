package data

type xmlData struct {
}

func (d xmlData) convertToJson() string {
	return ""
}

type XmlDataAdapter struct {
	xmlData *xmlData
}

func (adapter XmlDataAdapter) SendJsonData() {
	adapter.xmlData.convertToJson()
	println("send data in json")

}
