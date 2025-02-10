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
