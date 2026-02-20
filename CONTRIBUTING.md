# Guia de Contribuição

Obrigado por considerar contribuir com o GoTrack! Este guia descreve as convenções e o fluxo de trabalho adotados no projeto.

> **Idioma:** Este é um projeto brasileiro voltado para o curso Aprenda Go (em pt-br). Todos os commits, PRs, issues, comentários no código e documentação devem ser escritos em **português brasileiro**.
>
> **Plataforma:** O GoTrack tem suporte apenas para **Linux**.

## Branches

Todas as branches partem da `main` e seguem o padrão `tipo/descricao-em-kebab-case`:

| Prefixo | Uso | Exemplo |
|---|---|---|
| `feat/` | Nova funcionalidade | `feat/modo-escuro` |
| `fix/` | Correção de bug | `fix/fechar-player-ao-trocar-capitulo` |
| `docs/` | Documentação | `docs/reescrita-readme` |
| `style/` | Alterações visuais (sem lógica) | `style/ajuste-cores-sidebar` |
| `chore/` | Manutenção, refatoração, traduções | `chore/traducao-erros-banco` |
| `test/` | Testes | `test/cobertura-database` |

## Commits

Usamos [Conventional Commits](https://www.conventionalcommits.org/pt-br/) em **português**. O formato é:

```
tipo: descrição curta no imperativo
```

### Tipos permitidos

| Tipo | Quando usar |
|---|---|
| `feat` | Nova funcionalidade |
| `fix` | Correção de bug |
| `docs` | Apenas documentação |
| `style` | Formatação, ajustes visuais (sem mudança de lógica) |
| `refactor` | Refatoração sem alterar comportamento |
| `test` | Adição ou correção de testes |
| `chore` | Tarefas de manutenção (build, dependências, traduções) |

### Exemplos

```
feat: adiciona modo escuro com toggle na sidebar
fix: fecha player ao navegar para outro capítulo
docs: traduz comentários do código Vue para pt-br
style: destaca visualmente a aula em reprodução
chore: traduz mensagens de erro do banco para pt-br
```

### Regras

- Escreva no **imperativo**: "adiciona", não "adicionado" ou "adicionando"
- Primeira letra **minúscula** após o tipo
- **Sem ponto final** na descrição
- Se precisar de mais contexto, pule uma linha e adicione o corpo:

```
feat: adiciona modo escuro com toggle na sidebar

- Implementa alternância dark/light com persistência em localStorage
- Adiciona classes dark: em todos os componentes Vue
- Configura custom variant dark no Tailwind CSS v4
```

## Pull Requests

### Título

Descritivo e em português, **sem o prefixo do tipo**:

```
Adiciona modo escuro com toggle na sidebar
```

### Corpo

Inclua uma seção `## Resumo` com bullets explicando as alterações. Veja o template que aparece automaticamente ao abrir um PR.

### Fluxo

1. Crie uma branch a partir da `main`
2. Faça seus commits seguindo as convenções acima
3. Abra um PR para a `main`
4. Aguarde a revisão

## Versionamento

Usamos [SemVer](https://semver.org/lang/pt-BR/) e geramos o changelog com [conventional-changelog](https://github.com/conventional-changelog/conventional-changelog):

| Tipo de mudança | Bump de versão | Exemplo |
|---|---|---|
| `feat` | Minor (`0.1.0` → `0.2.0`) | Nova funcionalidade |
| `fix` | Patch (`0.2.0` → `0.2.1`) | Correção de bug |
| `BREAKING CHANGE` | Major (`0.2.0` → `1.0.0`) | Mudança incompatível |

## Ambiente de desenvolvimento

```bash
# Clonar o repositório
git clone https://github.com/walissonaguirra/gotrack-youtube.git
cd gotrack-youtube

# Instalar dependências do frontend
cd frontend && npm install && cd ..

# Executar em modo desenvolvimento
cd frontend && npm run dev

# Build e execução completa
cd frontend && npm run build && cd .. && go build -o gotrack && ./gotrack
```

### Requisitos

- **Linux** (única plataforma suportada)
- **Go** >= 1.21
- **Node.js** >= 18
- Dependências do webview (veja o [README](README.md#dependências-de-sistema-para-o-webview))
