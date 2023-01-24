package validate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructWithOneError(t *testing.T) {
	type TestStruct struct {
		FieldTest string `validate:"required"`
	}

	testStruct := TestStruct{}
	testStruct.FieldTest = ""

	err := Struct(testStruct)
	assert.NotNil(t, err)
	assert.Equal(t, "{\"message\":\"FieldTest is a required field\",\"field\":\"FieldTest\"}", err.Error())
}

func TestStructWithManyError(t *testing.T) {
	type TestStruct struct {
		FieldTest  string `validate:"required"`
		FieldTest2 string `validate:"email"`
	}

	testStruct := TestStruct{}
	testStruct.FieldTest = ""
	testStruct.FieldTest2 = "teste.com"

	err := Struct(testStruct)
	assert.NotNil(t, err)
	assert.Equal(t, "{\"message\":\"FieldTest is a required field\",\"field\":\"FieldTest\"}|{\"message\":\"FieldTest2 must be a valid email address\",\"field\":\"FieldTest2\"}", err.Error())
}

func TestStructWithoutError(t *testing.T) {
	type TestStruct struct {
		FieldTest  string `validate:"required"`
		FieldTest2 string `validate:"email"`
	}

	testStruct := TestStruct{}
	testStruct.FieldTest = "Teste 1"
	testStruct.FieldTest2 = "teste@teste.com"

	err := Struct(testStruct)
	assert.Nil(t, err)
}

func TestErrorMessage(t *testing.T) {
	err := ErrorMessage("FieldTest", "FieldTest is a required field")
	assert.NotNil(t, err)
	assert.Equal(t, "{\"message\":\"FieldTest is a required field\",\"field\":\"FieldTest\"}", err.Error())

}
