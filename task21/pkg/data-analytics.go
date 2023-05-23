package pkg

type DataAnalytic interface {
	SendJsonData()
}

type JsonData struct {
}

func (d JsonData) SendJsonData() {
}
