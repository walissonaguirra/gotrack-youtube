# GoTrack - Frontend

![Vue.js](https://img.shields.io/badge/Vue.js-3.5+-4FC08D?logo=vuedotjs&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-6.3+-646CFF?logo=vite&logoColor=white)
![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-4.1+-06B6D4?logo=tailwindcss&logoColor=white)
![Node.js](https://img.shields.io/badge/Node.js-≥18-339933?logo=nodedotjs&logoColor=white)
![License](https://img.shields.io/badge/Licença-MIT-blue)

Interface do GoTrack construída como uma **SPA (Single Page Application)** com Vue.js 3, utilizando Composition API, Single File Components (`.vue`) e Tailwind CSS v4.

O frontend é compilado pelo Vite e o output (`dist/`) é embarcado no binário Go via `embed.FS`, servido localmente por um servidor HTTP para o webview.

## Requisitos

- [Node.js](https://nodejs.org/) >= 18
- npm (incluso com o Node.js)

## Instalação

```bash
cd frontend
npm install
```

## Scripts disponíveis

| Comando          | Descrição                                                  |
| ---------------- | ---------------------------------------------------------- |
| `npm run dev`    | Inicia o servidor de desenvolvimento com HMR (Hot Reload)  |
| `npm run build`  | Gera o build de produção em `dist/`                        |

## Build para produção

O build gera os arquivos estáticos otimizados que serão embarcados no binário Go:

```bash
cd frontend
npm run build
```

Isso cria a pasta `dist/` com o HTML, CSS e JS minificados. Em seguida, compile o Go normalmente:

```bash
cd ..
go build -o gotrack
```

O `main.go` utiliza `//go:embed frontend/dist` para incluir esses arquivos no binário final.

## Stack tecnológica

| Tecnologia              | Versão  | Função                                       |
| ----------------------- | ------- | -------------------------------------------- |
| Vue.js                  | 3.5+    | Framework reativo para a interface            |
| Vite                    | 6.3+    | Bundler e dev server                          |
| Tailwind CSS            | 4.1+    | Framework CSS utility-first                   |
| @vitejs/plugin-vue      | 5.2+    | Compilação de Single File Components `.vue`   |
| @tailwindcss/vite       | 4.1+    | Integração do Tailwind com o Vite             |

## Estrutura de arquivos

```
frontend/
├── index.html                          # Entry point do Vite
├── package.json                        # Dependências e scripts
├── vite.config.js                      # Configuração do Vite
├── src/
│   ├── main.js                         # Cria e monta a app Vue
│   ├── style.css                       # Tailwind + tema customizado
│   ├── App.vue                         # Componente raiz
│   ├── components/
│   │   ├── AppSidebar.vue              # Sidebar de navegação
│   │   ├── AppDashboard.vue            # Dashboard com estatísticas
│   │   ├── ChapterView.vue            # Conteúdo de um capítulo
│   │   ├── LessonItem.vue             # Item individual de aula
│   │   ├── VideoPlayer.vue            # Player YouTube embutido
│   │   └── PomodoroTimer.vue          # Widget Pomodoro flutuante
│   └── composables/
│       └── usePomodoro.js             # Lógica reativa do timer
└── dist/                               # Output do build (gerado)
```

## Arquitetura

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

O `App.vue` mantém o estado central (`modules`, `currentChapter`, `stats`, `view`) e distribui via props. Eventos sobem via `emit`. Não há necessidade de Pinia/Vuex — a aplicação é simples o suficiente para gerenciar estado no componente raiz.

### Comunicação com o backend Go

O backend Go injeta funções globais no `window` via `webview.Bind()`. O frontend as acessa com `window.<nomeDaFunção>()`:

| Função JS                          | Descrição                                  |
| ----------------------------------- | ------------------------------------------ |
| `window.goGetModules()`             | Retorna todos os módulos com progresso      |
| `window.goToggleLesson(lessonId)`   | Alterna conclusão de uma aula               |
| `window.goGetStats()`              | Retorna estatísticas gerais                 |
| `window.goGetYouTubeURL(lessonId)`  | Retorna URL de embed do YouTube             |
| `window.goSaveTimerSession(min, ch)`| Salva sessão do Pomodoro no banco            |
| `window.goIsChapterUnlocked(ch)`    | Verifica se um capítulo está desbloqueado   |

Todas retornam Promises (o webview serializa as respostas automaticamente). Os dados trafegam como JSON strings, parseados com `JSON.parse()` no frontend.

O Go serve os arquivos de `frontend/dist` via um servidor HTTP local em porta efêmera (`127.0.0.1:0`) e abre uma janela webview apontando para esse endereço.
