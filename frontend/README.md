# GoTrack - Frontend

Interface do GoTrack construida como uma **SPA (Single Page Application)** com Vue.js 3, utilizando Composition API, Single File Components (`.vue`) e Tailwind CSS v4.

O frontend e compilado pelo Vite e o output (`dist/`) e embarcado no binario Go via `embed.FS`, servido localmente por um servidor HTTP para o webview.

## Requisitos

- [Node.js](https://nodejs.org/) >= 18
- npm (incluso com o Node.js)

## Instalacao

```bash
cd frontend
npm install
```

## Scripts disponiveis

| Comando          | Descricao                                                  |
| ---------------- | ---------------------------------------------------------- |
| `npm run dev`    | Inicia o servidor de desenvolvimento com HMR (Hot Reload)  |
| `npm run build`  | Gera o build de producao em `dist/`                        |

## Build para producao

O build gera os arquivos estaticos otimizados que serao embarcados no binario Go:

```bash
cd frontend
npm run build
```

Isso cria a pasta `dist/` com o HTML, CSS e JS minificados. Em seguida, compile o Go normalmente:

```bash
cd ..
go build -o gotrack
```

O `main.go` utiliza `//go:embed frontend/dist` para incluir esses arquivos no binario final.

## Stack tecnologica

| Tecnologia              | Versao  | Funcao                                       |
| ----------------------- | ------- | -------------------------------------------- |
| Vue.js                  | 3.5+    | Framework reativo para a interface            |
| Vite                    | 6.3+    | Bundler e dev server                          |
| Tailwind CSS            | 4.1+    | Framework CSS utility-first                   |
| @vitejs/plugin-vue      | 5.2+    | Compilacao de Single File Components `.vue`   |
| @tailwindcss/vite       | 4.1+    | Integracao do Tailwind com o Vite             |

## Estrutura de arquivos

```
frontend/
├── index.html                          # Entry point do Vite
├── package.json                        # Dependencias e scripts
├── vite.config.js                      # Configuracao do Vite
├── src/
│   ├── main.js                         # Cria e monta a app Vue
│   ├── style.css                       # Tailwind + tema customizado
│   ├── App.vue                         # Componente raiz
│   ├── components/
│   │   ├── AppSidebar.vue              # Sidebar de navegacao
│   │   ├── AppDashboard.vue            # Dashboard com estatisticas
│   │   ├── ChapterView.vue            # Conteudo de um capitulo
│   │   ├── LessonItem.vue             # Item individual de aula
│   │   ├── VideoPlayer.vue            # Player YouTube embutido
│   │   └── PomodoroTimer.vue          # Widget Pomodoro flutuante
│   └── composables/
│       └── usePomodoro.js             # Logica reativa do timer
└── dist/                               # Output do build (gerado)
```

## Arquitetura

### Padroes Vue utilizados

- **`<script setup>`** — sintaxe moderna e concisa do Vue 3 em todos os componentes
- **Composition API** — `ref()`, `computed()`, `onMounted()`, `onUnmounted()`
- **Single File Components** — template, script e estilos no mesmo arquivo `.vue`
- **Props + Emits** — comunicacao unidirecional entre componentes pai e filho
- **Composables** — logica reutilizavel extraida em funcoes (`usePomodoro`)

### Fluxo de dados

```
App.vue (estado global)
├── AppSidebar.vue
│   Props: modules, activeChapter
│   Emits: navigate, show-dashboard
├── AppDashboard.vue
│   Props: modules, stats
├── ChapterView.vue
│   Props: chapter, isLocked
│   Emits: toggle-lesson
│   ├── VideoPlayer.vue
│   │   Props: url
│   │   Emits: close
│   └── LessonItem.vue
│       Props: lesson
│       Emits: toggle, watch
└── PomodoroTimer.vue
    Props: getCurrentChapter
    Composable: usePomodoro()
```

O `App.vue` mantem o estado central (`modules`, `currentChapter`, `stats`, `view`) e distribui via props. Eventos sobem via `emit`. Nao ha necessidade de Pinia/Vuex — a aplicacao e simples o suficiente para gerenciar estado no componente raiz.

### Comunicacao com o backend Go

O backend Go injeta funcoes globais no `window` via `webview.Bind()`. O frontend as acessa com `window.<nomeDaFuncao>()`:

| Funcao JS                          | Descricao                                  |
| ----------------------------------- | ------------------------------------------ |
| `window.goGetModules()`             | Retorna todos os modulos com progresso      |
| `window.goToggleLesson(lessonId)`   | Alterna conclusao de uma aula               |
| `window.goGetStats()`               | Retorna estatisticas gerais                 |
| `window.goGetYouTubeURL(lessonId)`  | Retorna URL de embed do YouTube             |
| `window.goSaveTimerSession(min, ch)`| Salva sessao do Pomodoro no banco            |
| `window.goIsChapterUnlocked(ch)`    | Verifica se um capitulo esta desbloqueado   |

Todas retornam Promises (o webview serializa as respostas automaticamente). Os dados trafegam como JSON strings, parseados com `JSON.parse()` no frontend.

## Componentes

### `App.vue`

Componente raiz que orquestra toda a aplicacao:

- Gerencia o estado global reativo (`modules`, `currentChapter`, `stats`, `view`)
- Decide qual view renderizar (`dashboard` ou `chapter`) via renderizacao condicional (`v-if`/`v-else-if`)
- Inicializa a aplicacao no `onMounted`: carrega modulos, exibe dashboard, solicita permissao de notificacao

### `AppSidebar.vue`

Barra lateral fixa com navegacao entre modulos e capitulos:

- Exibe barra de progresso geral (calculada via `computed`)
- Lista modulos agrupados com seus capitulos
- Indica visualmente: capitulo ativo (highlight azul), completo (badge verde), bloqueado (icone de cadeado com opacidade reduzida)
- Emite `navigate` ao clicar em um capitulo e `show-dashboard` ao clicar no logo

### `AppDashboard.vue`

Tela inicial com visao geral do progresso:

- Tres cards de estatisticas: progresso percentual, aulas completas, tempo estudado
- Lista de modulos com barras de progresso individuais
- Dados recebidos via props `modules` e `stats`

### `ChapterView.vue`

Exibe o conteudo de um capitulo selecionado:

- **Estado bloqueado**: mostra icone de cadeado e mensagem indicando pre-requisito
- **Estado desbloqueado**: mostra titulo, barra de progresso, lista de aulas e player de video
- Gerencia o `videoUrl` localmente (estado do player pertence a esta view)
- Delega renderizacao de cada aula para `LessonItem` e do player para `VideoPlayer`

### `LessonItem.vue`

Representa uma unica aula na lista:

- Checkbox circular para marcar/desmarcar conclusao (emite `toggle`)
- Titulo e numero da aula
- Botao de play para assistir ao video (emite `watch`)
- Estilizacao condicional: borda verde e fundo claro quando completa

### `VideoPlayer.vue`

Player de video YouTube embutido:

- Renderiza um iframe responsivo (aspect ratio 16:9 via padding-top trick)
- Visivel apenas quando `url` nao esta vazio (`v-if`)
- Botao para fechar o player (emite `close`)

### `PomodoroTimer.vue`

Widget flutuante no canto inferior direito:

- Dois estados visuais: expandido (controles completos) e minimizado (apenas timer)
- Toda a logica delegada ao composable `usePomodoro`

## Composable: `usePomodoro`

Encapsula toda a logica do timer Pomodoro em uma funcao reutilizavel:

**Estado reativo:**
- `mode` — modo atual (`focus`, `short`, `long`)
- `running` — se o timer esta rodando
- `remaining` — segundos restantes
- `minimized` — se o widget esta minimizado

**Computeds:**
- `timeDisplay` — tempo formatado `MM:SS`
- `currentMode` — objeto do modo atual (label + minutos)
- `modes` — lista de modos para renderizar os botoes

**Metodos:**
- `toggle()` — inicia ou pausa o timer
- `reset()` — reinicia o timer com o tempo total do modo atual
- `setMode(mode)` — troca o modo (apenas quando pausado)
- `toggleMinimize()` — alterna entre expandido e minimizado

**Comportamento ao completar:**
- Se modo `focus`: salva a sessao no banco via `goSaveTimerSession`
- Envia notificacao do navegador
- Troca automaticamente para o proximo modo (`focus` -> `short`, `short`/`long` -> `focus`)

**Cleanup:**
- `onUnmounted` limpa o `setInterval` para evitar memory leak

## Estilizacao

O Tailwind CSS v4 e configurado em `src/style.css` usando a nova sintaxe com `@theme`:

```css
@import "tailwindcss";

@theme {
  --color-primary: #0d6efd;
  --color-secondary: #6c757d;
  --color-success: #198754;
  --color-danger: #dc3545;
  --color-warning: #ffc107;
  --color-info: #0dcaf0;
  --color-light: #f8f9fa;
  --color-dark: #212529;
  --color-border: #dee2e6;
}
```

As cores ficam disponiveis como classes utilitarias (`text-primary`, `bg-success`, `border-border`, etc.). Toda a estilizacao e feita inline via classes Tailwind nos templates — nenhum componente usa `<style>`.

## Integracao com o Go

O ciclo completo de build e execucao:

```bash
# 1. Instalar dependencias (apenas na primeira vez)
cd frontend && npm install

# 2. Build do frontend
npm run build

# 3. Compilar o binario Go (embarca frontend/dist)
cd .. && go build -o gotrack

# 4. Executar
./gotrack
```

O Go serve os arquivos de `frontend/dist` via um servidor HTTP local em porta efemera (`127.0.0.1:0`) e abre uma janela webview apontando para esse endereco.
