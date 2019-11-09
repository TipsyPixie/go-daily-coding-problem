package problem016

type OrderIdLogger struct {
	stringList []string
	cursor     int
}

func NewOrderIdLogger(size int) *OrderIdLogger {
	return &OrderIdLogger{
		stringList: make([]string, size, size),
		cursor:     0,
	}
}

func (logger *OrderIdLogger) Record(orderId string) {
	logger.stringList[logger.cursor] = orderId
	logger.cursor = (logger.cursor + 1) % len(logger.stringList)
}

func (logger *OrderIdLogger) GetLast(index int) string {
	if index <= 0 {
		index = 1
	}
	logLength := len(logger.stringList)
	return logger.stringList[(logger.cursor+logLength-index)%logLength]
}
