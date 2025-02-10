# React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react/README.md) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## Expanding the ESLint configuration

If you are developing a production application, we recommend updating the configuration to enable type aware lint rules:

- Configure the top-level `parserOptions` property like this:

```js
export default tseslint.config({
  languageOptions: {
    // other options...
    parserOptions: {
      project: ['./tsconfig.node.json', './tsconfig.app.json'],
      tsconfigRootDir: import.meta.dirname,
    },
  },
})
```

- Replace `tseslint.configs.recommended` to `tseslint.configs.recommendedTypeChecked` or `tseslint.configs.strictTypeChecked`
- Optionally add `...tseslint.configs.stylisticTypeChecked`
- Install [eslint-plugin-react](https://github.com/jsx-eslint/eslint-plugin-react) and update the config:

```js
// eslint.config.js
import react from 'eslint-plugin-react'

export default tseslint.config({
  // Set the react version
  settings: { react: { version: '18.3' } },
  plugins: {
    // Add the react plugin
    react,
  },
  rules: {
    // other rules...
    // Enable its recommended rules
    ...react.configs.recommended.rules,
    ...react.configs['jsx-runtime'].rules,
  },
})
```


# HaHaSaaS (Humor as a Service)

A terminal-themed joke API interface that brings laughter with a cyberpunk twist. This project provides a way to interact with joke APIs through a retro-terminal interface.

## Getting Started

### Prerequisites

- Node.js (v14 or higher)
- npm or yarn

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Andebugulin/hahasaas.git
   cd hahasaas

Install dependencies:
- `bashCopynpm install`

Start the development server:
- `bashCopynpm start`


### Project Structure 
```bash
Copyhahasaas/
├── src/
│   ├── components/
│   │   └── JokeTerminal.jsx
│   ├── App.js
│   └── index.js
├── public/
└── README.md
```

### Available Commands

- `help`: Show available commands
- `random`: Get a random joke
- `category <name>`: Get a joke from specific category
- `list`: Show all joke categories
- `clear`: Clear terminal screen


### Contributing
Feel free to open issues and pull requests!

### License 
MIT License

`To complete the setup:`
1. Replace the contents of `src/App.js` with:
```jsx
import JokeTerminal from './components/JokeTerminal';

function App() {
  return (
    <div className="App">
      <JokeTerminal />
    </div>
  );
}

export default App;
```
