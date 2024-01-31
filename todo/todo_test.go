package todo

import "testing"

func TestCreateTodoNotAllowSleepTask(t *testing.T) {
	// Arrange
	handler := NewTodoHandler(&TestDB{})
	context := &TestContext{}

	// Act
	handler.NewTask(context)

	// Assert
	want := "not allowed"
	if want != context.v["error"] {
		t.Errorf("want %s but get %s\n", want, context.v["error"])
	}
}

type TestDB struct{}

func (TestDB) New(*Todo) error {
	return nil
}

type TestContext struct {
	v map[string]interface{}
}

func (TestContext) Bind(v interface{}) error {
	*v.(*Todo) = Todo{
		Title: "sleep",
	}
	return nil
}

func (c *TestContext) JSON(statusCode int, v interface{}) {
	c.v = v.(map[string]interface{})
}

func (TestContext) TransactionID() string {
	return "TestTransactionID"
}

func (TestContext) Audience() string {
	return "UnitTest"
}
