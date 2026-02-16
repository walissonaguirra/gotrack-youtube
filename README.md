# GoTrack

Aplicacao desktop para acompanhar o progresso no curso **Aprenda Go** do YouTube. Construida em Go com webview para a interface grafica e SQLite para persistencia local.

## Visao geral

O GoTrack permite:

- Acompanhar o progresso em 21 capitulos do curso Aprenda Go
- Marcar aulas como concluidas e visualizar o percentual de conclusao
- Assistir aos videos do YouTube diretamente na aplicacao (player embutido)
- Desbloquear exercicios ao completar os capitulos teoricos correspondentes
- Usar um timer Pomodoro integrado para sessoes de estudo focado
- Visualizar estatisticas gerais no dashboard (progresso, aulas completas, tempo estudado)

Todos os dados sao armazenados localmente em um banco SQLite no diretorio de configuracao do usuario.

## Requisitos

- **Go** >= 1.21
- **Node.js** >= 18 (para build do frontend)
- **Dependencias do webview** (bibliotecas de sistema)

### Dependencias de sistema para o webview

O pacote `webview/webview_go` requer bibliotecas nativas do sistema operacional:

**Linux (Debian/Ubuntu):**
```bash
sudo apt install libgtk-3-dev libwebkit2gtk-4.0-dev
```

**Linux (Arch/Manjaro):**
```bash
sudo pacman -S gtk3 webkit2gtk
```

**macOS:**
Nenhuma dependencia extra (usa WebKit nativo).

**Windows:**
Nenhuma dependencia extra (usa WebView2/Edge).

## Instalacao e execucao

```bash
# 1. Clonar o repositorio
git clone <url-do-repositorio>
cd gotrack-youtube-refactor

# 2. Build do frontend
cd frontend
npm install
npm run build
cd ..

# 3. Compilar e executar
go build -o gotrack
./gotrack
```

A aplicacao abre uma janela de 1280x800 com a interface do GoTrack.

## Estrutura do projeto

```
gotrack-youtube-refactor/
├── main.go                         # Ponto de entrada da aplicacao
├── go.mod                          # Modulo Go e dependencias
├── go.sum                          # Checksums das dependencias
├── internal/
│   ├── models/
│   │   └── models.go               # Structs de dados (Lesson, Chapter, Module, Stats)
│   ├── database/
│   │   ├── database.go             # Conexao, migracoes e schema SQLite
│   │   └── queries.go              # Operacoes de banco (CRUD, estatisticas)
│   ├── handlers/
│   │   └── handlers.go             # Bindings JS <-> Go para o webview
│   └── course/
│       └── course.go               # Conteudo do curso (capitulos, aulas, IDs YouTube)
└── frontend/                       # SPA Vue.js (ver frontend/README.md)
    ├── src/                        # Codigo-fonte Vue
    └── dist/                       # Build de producao (embarcado no binario)
```

## Arquitetura

### Como a aplicacao funciona

```
┌──────────────────────────────────────────────────────┐
│                    Binario Go                        │
│                                                      │
│  ┌─────────────┐  embed.FS   ┌────────────────────┐ │
│  │  Frontend    │◄───────────│  frontend/dist/     │ │
│  │  (Vue SPA)   │            │  (arquivos estaticos│ │
│  │              │            │   compilados)       │ │
│  └──────┬───────┘            └────────────────────┘ │
│         │                                            │
│         │ webview.Bind()                             │
│         │ (chamadas JS -> Go)                        │
│         ▼                                            │
│  ┌─────────────┐            ┌────────────────────┐  │
│  │  Handlers   │───────────►│    Database         │  │
│  │             │            │    (SQLite)          │  │
│  └─────────────┘            └────────────────────┘  │
│                                                      │
│  ┌─────────────┐                                     │
│  │  Course     │ Conteudo hard-coded do curso        │
│  │  (dados)    │ (capitulos, aulas, IDs YouTube)     │
│  └─────────────┘                                     │
│                                                      │
│  ┌─────────────────────────────────────┐             │
│  │  HTTP Server (127.0.0.1:porta)      │             │
│  │  Serve frontend via http.FileServer │             │
│  └─────────────────────────────────────┘             │
│                                                      │
│  ┌─────────────────────────────────────┐             │
│  │  Webview (janela nativa)            │             │
│  │  Navega para o HTTP server local    │             │
│  └─────────────────────────────────────┘             │
└──────────────────────────────────────────────────────┘
```

1. O `main.go` inicializa o banco, popula os dados do curso, embarca os arquivos do frontend e inicia um servidor HTTP local em porta efemera
2. Uma janela webview e criada e navega ate o endereco do servidor local
3. Funcoes Go sao expostas ao JavaScript via `webview.Bind()`, permitindo que o frontend chame o backend diretamente
4. O frontend Vue.js se comunica com o backend exclusivamente por essas funcoes injetadas no `window`

### Sem servidor externo

A aplicacao e completamente local. O servidor HTTP existe apenas para servir os arquivos estaticos ao webview — nenhuma porta e exposta externamente (`127.0.0.1` apenas) e a porta e efemera (escolhida automaticamente pelo SO).

## Dependencias Go

| Pacote | Funcao |
| --- | --- |
| `github.com/webview/webview_go` | Janela nativa com motor de renderizacao web (GTK+WebKit no Linux, WebView2 no Windows, WebKit no macOS) |
| `modernc.org/sqlite` | Driver SQLite puro em Go (sem CGo, sem dependencia de libsqlite3) |

O uso de `modernc.org/sqlite` em vez de `mattn/go-sqlite3` elimina a necessidade de um compilador C e permite compilacao cruzada simples.

## Pacotes internos

### `internal/models`

Define as structs de dados usadas em toda a aplicacao. Todas possuem tags `json` para serializacao automatica.

| Struct | Descricao | Campos principais |
| --- | --- | --- |
| `Lesson` | Uma aula individual do curso | `ID`, `Chapter`, `LessonNumber`, `Title`, `YouTubeID`, `IsExercise`, `Completed` |
| `Chapter` | Agrupa aulas em um capitulo numerado | `Number`, `Title`, `Lessons[]`, `IsExercise`, `RequiredBy`, `Unlocked`, `Progress` |
| `Module` | Agrupa capitulos em niveis de aprendizado | `Name`, `Chapters[]`, `Progress` |
| `Stats` | Estatisticas gerais de progresso | `TotalLessons`, `CompletedLessons`, `Progress`, `TotalMinutes` |
| `TimerSession` | Registro de sessao do Pomodoro | `ID`, `StartedAt`, `DurationMinutes`, `Chapter` |

**Formato do ID das aulas:**
```
cap{NN}_aula{NN}
```
Exemplo: `cap02_aula05` = Capitulo 2, Aula 5.

### `internal/database`

Camada de acesso ao banco SQLite. Gerencia conexao, migracoes e todas as queries.

**Localizacao do banco:**
```
~/.config/gotrack/gotrack.db     (Linux)
~/Library/Application Support/gotrack/gotrack.db  (macOS)
%AppData%\gotrack\gotrack.db     (Windows)
```

O diretorio e determinado automaticamente por `os.UserConfigDir()`.

#### Schema

**Tabela `lessons`:**

| Coluna | Tipo | Descricao |
| --- | --- | --- |
| `id` | TEXT PK | Identificador unico (`cap01_aula01`) |
| `chapter` | INTEGER | Numero do capitulo |
| `lesson_number` | INTEGER | Numero da aula dentro do capitulo |
| `title` | TEXT | Titulo da aula |
| `youtube_id` | TEXT | ID do video no YouTube |
| `is_exercise` | INTEGER | 1 se for exercicio, 0 caso contrario |
| `completed` | INTEGER | 1 se concluida, 0 caso contrario |
| `completed_at` | TEXT | Data/hora da conclusao (RFC3339) |

**Tabela `timer_sessions`:**

| Coluna | Tipo | Descricao |
| --- | --- | --- |
| `id` | INTEGER PK | Auto-incremento |
| `started_at` | TEXT | Data/hora do inicio (RFC3339) |
| `duration_minutes` | INTEGER | Duracao em minutos |
| `chapter` | INTEGER | Capitulo estudado durante a sessao |

#### Queries disponiveis

| Metodo | Descricao |
| --- | --- |
| `SeedLessons(lessons)` | Insere ou atualiza aulas no banco (upsert via `ON CONFLICT`) |
| `ToggleLesson(id)` | Alterna o status de conclusao de uma aula. Registra `completed_at` ao completar |
| `GetLessonCompletions()` | Retorna mapa `id -> bool` com status de todas as aulas |
| `IsChapterComplete(chapter)` | Verifica se todas as aulas de um capitulo estao concluidas |
| `SaveTimerSession(minutes, chapter)` | Registra uma sessao do Pomodoro |
| `GetYouTubeID(lessonID)` | Retorna o ID do YouTube para uma aula |
| `GetStats()` | Calcula estatisticas gerais: total de aulas, concluidas, progresso e tempo total estudado |

### `internal/handlers`

Camada de binding entre o JavaScript do frontend e o Go. Cada metodo e exposto ao `window` do navegador via `webview.Bind()`.

| Metodo Go | Funcao JS | Retorno | Descricao |
| --- | --- | --- | --- |
| `GetModules()` | `window.goGetModules()` | JSON string | Retorna todos os modulos com progresso calculado e status de desbloqueio |
| `ToggleLesson(id)` | `window.goToggleLesson(id)` | JSON string | Alterna conclusao de uma aula e retorna modulos atualizados |
| `IsChapterUnlocked(ch)` | `window.goIsChapterUnlocked(ch)` | bool | Verifica se um capitulo esta desbloqueado |
| `SaveTimerSession(min, ch)` | `window.goSaveTimerSession(min, ch)` | error | Salva sessao do Pomodoro |
| `GetStats()` | `window.goGetStats()` | JSON string | Retorna estatisticas gerais |
| `GetYouTubeURL(id)` | `window.goGetYouTubeURL(id)` | string | Retorna URL de embed do YouTube |
| `SeedLessons()` | (interno) | error | Popula o banco com dados do curso na inicializacao |

**Fluxo do `GetModules()`:**
1. Busca o mapa de conclusoes no banco (`GetLessonCompletions`)
2. Carrega a estrutura completa do curso (`course.AllModules`)
3. Para cada capitulo: verifica desbloqueio, marca aulas concluidas, calcula progresso
4. Para cada modulo: calcula progresso agregado
5. Serializa tudo em JSON e retorna como string

### `internal/course`

Conteudo hard-coded do curso "Aprenda Go". Nenhum dado e buscado externamente — tudo esta definido diretamente no codigo Go.

#### Estrutura do curso

| Modulo | Capitulos | Descricao |
| --- | --- | --- |
| **Iniciante** | 1-7 | Visao geral, variaveis, tipos, fundamentos, fluxo de controle + exercicios |
| **Intermediario** | 8-17 | Arrays, slices, maps, structs, funcoes, ponteiros, aplicacoes + exercicios |
| **Avancado** | 18, 20, 21 | Concorrencia, goroutines, canais + exercicios |
| **Ecossistema** | 19 | Terminal, workspace, IDEs, GitHub, compilacao cruzada, pacotes |

**Total:** 21 capitulos, ~150 aulas

#### Sistema de desbloqueio

Capitulos de exercicio ficam bloqueados ate que o capitulo teorico correspondente seja completado:

```
Exercicios #1 (Cap 3)  <- requer Cap 2 (Variaveis, Valores & Tipos)
Exercicios #2 (Cap 5)  <- requer Cap 4 (Fundamentos da Programacao)
Exercicios #3 (Cap 7)  <- requer Cap 6 (Fluxo de Controle)
Exercicios #4 (Cap 9)  <- requer Cap 8 (Agrupamentos de Dados)
Exercicios #5 (Cap 11) <- requer Cap 10 (Structs)
Exercicios #6 (Cap 13) <- requer Cap 12 (Funcoes)
Exercicios #7 (Cap 15) <- requer Cap 14 (Ponteiros)
Exercicios #8 (Cap 17) <- requer Cap 16 (Aplicacoes)
Exercicios #9 (Cap 20) <- requer Cap 18 (Concorrencia)
```

Este mapeamento esta definido em `ExerciseRequirements` no pacote `course`.

## Ponto de entrada: `main.go`

O `main.go` orquestra toda a inicializacao em sequencia:

1. **Banco de dados** — Abre (ou cria) o SQLite em `~/.config/gotrack/gotrack.db`
2. **Seed** — Popula o banco com os dados do curso (upsert, preserva progresso existente)
3. **Servidor HTTP** — Inicia um `http.FileServer` servindo `frontend/dist` em `127.0.0.1:0`
4. **Webview** — Cria janela nativa de 1280x800
5. **Bindings** — Expoe 6 funcoes Go ao JavaScript via `webview.Bind()`
6. **Navegacao** — Direciona o webview para o servidor local
7. **Execucao** — Entra no event loop do webview (`w.Run()`)

Os arquivos do frontend sao embarcados no binario usando a diretiva `//go:embed frontend/dist`, eliminando a necessidade de distribuir arquivos separados.

## Compilacao cruzada

Gracas ao `modernc.org/sqlite` (SQLite puro em Go, sem CGo), a compilacao cruzada funciona nativamente:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o gotrack-linux

# Windows
GOOS=windows GOARCH=amd64 go build -o gotrack.exe

# macOS
GOOS=darwin GOARCH=arm64 go build -o gotrack-macos
```

**Nota:** O webview requer as bibliotecas nativas do sistema de destino instaladas na maquina onde o binario sera executado.

## Armazenamento de dados

Todos os dados sao locais. Nenhuma informacao e enviada para servidores externos.

- **Banco:** `~/.config/gotrack/gotrack.db`
- **Formato:** SQLite 3
- **Backup:** Basta copiar o arquivo `.db`
- **Reset:** Deletar o arquivo `.db` reinicia todo o progresso

## Fluxo de desenvolvimento

```bash
# Terminal 1 — Dev server do frontend (com HMR)
cd frontend && npm run dev

# Terminal 2 — Compilar e executar o Go (quando precisar testar bindings)
cd frontend && npm run build && cd .. && go build -o gotrack && ./gotrack
```

O dev server do Vite (`npm run dev`) e util para iterar rapidamente no frontend, mas as funcoes `window.go*` so estao disponiveis quando executado dentro do webview via o binario Go.

---

Este projeto está licenciado sob a
[Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
[![License: CC BY-NC-SA 4.0](https://licensebuttons.net/l/by-nc-sa/4.0/88x31.png)](https://creativecommons.org/licenses/by-nc-sa/4.0/)
