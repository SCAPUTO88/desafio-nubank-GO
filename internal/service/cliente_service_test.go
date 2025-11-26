package service

import (
	"errors"
	"testing"
	"time"

	"github.com/SCAPUTO88/desafio-nubank-GO/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mocks ---

type MockClienteRepository struct {
	mock.Mock
}

func (m *MockClienteRepository) Create(cliente *domain.Cliente) error {
	args := m.Called(cliente)
	if args.Error(0) == nil {
		cliente.ID = 1
	}
	return args.Error(0)
}

func (m *MockClienteRepository) FindAllWithContatos() ([]domain.Cliente, error) {
	args := m.Called()
	return args.Get(0).([]domain.Cliente), args.Error(1)
}

func (m *MockClienteRepository) FindByID(id uint) (*domain.Cliente, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Cliente), args.Error(1)
}

// Novo Mock para o Publisher
type MockEventPublisher struct {
	mock.Mock
}

func (m *MockEventPublisher) Publish(topicID string, message interface{}) error {
	args := m.Called(topicID, message)
	return args.Error(0)
}

func (m *MockEventPublisher) Close() {
	m.Called()
}

// --- Testes ---

func TestCreateCliente_Success(t *testing.T) {
	mockRepo := new(MockClienteRepository)
	mockPublisher := new(MockEventPublisher)
	service := NewClienteService(mockRepo, mockPublisher)

	dto := domain.CreateClienteDTO{
		Nome:  "Fulano",
		Email: "emailfake@example.com",
	}

	mockRepo.On("Create", mock.AnythingOfType("*domain.Cliente")).Return(nil)
	mockPublisher.On("Publish", "new-client-created", mock.Anything).Return(nil).Maybe()

	result, err := service.Create(dto)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, "Fulano", result.Nome)

	time.Sleep(10 * time.Millisecond) 

	mockRepo.AssertExpectations(t)
}

func TestCreateCliente_Error(t *testing.T) {
	mockRepo := new(MockClienteRepository)
	mockPublisher := new(MockEventPublisher)
	service := NewClienteService(mockRepo, mockPublisher)

	dto := domain.CreateClienteDTO{Nome: "John Doe", Email: "erro@banco.com"}

	mockRepo.On("Create", mock.Anything).Return(errors.New("db_error"))

	result, err := service.Create(dto)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "db_error", err.Error())
}