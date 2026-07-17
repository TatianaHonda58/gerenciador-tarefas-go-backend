# ⚙️ Gerenciador de Tarefas — API REST de Alta Performance (Go / Golang)

Este repositório abriga o ecossistema de Backend e inteligência de negócios do Gerenciador de Tarefas. Desenvolvida em **Go**, esta API foi projetada com foco em concorrência, baixo consumo de memória, alto throughput de requisições e testabilidade nativa.

---

## 🚀 Tecnologias e Conceitos Utilizados

A arquitetura utiliza os recursos nativos e idiomáticos da linguagem Go para microsserviços modernos:

* **Go (Golang)** (Linguagem compilada de altíssima performance e concorrência nativa)
* **Go Testing Package** (Mecanismo nativo e poderoso para asserções e testes automatizados)
* **Nett/HTTP package** (Roteamento nativo e leve, sem overhead de frameworks pesados)

## 📌 Diferenciais de Engenharia de Qualidade (QA + Dev)

* **Cultura Core de Testes Unitários:** O projeto utiliza a abordagem idiomática do Go (`*_test.go`), garantindo que handlers e lógica de negócios sejam testados isoladamente através de mocks e gravadores de resposta (`httptest.NewRecorder`).
* **Design Leve e Resiliente:** Arquitetura limpa que facilita o isolamento de funções para a aplicação de testes de integração futuros.

## 🛠️ Como Executar a API Localmente

### Pré-requisitos
Você precisará do **Go** instalado e configurado em sua máquina.

### Passo a Passo

1. **Clone o repositório:**
   ```bash
   git clone [https://github.com/TatianaHonda58/gerenciador-tarefas-go-backend.git](https://github.com/TatianaHonda58/gerenciador-tarefas-go-backend.git)
