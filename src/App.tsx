import React, { useState } from 'react';
import { Terminal } from 'lucide-react';

interface Command {
  type: 'input' | 'output';
  content: string;
}

const JokeTerminal = () => {
  const [commands, setCommands] = useState<Command[]>([]);
  const [input, setInput] = useState('');
  const [loading, setLoading] = useState(false);

  const listCommands = {
    'help': 'Displays a list of available commands and their usage.',
    'clear': 'Clears the terminal screen.',
    'rndj --category -c': 'Fetches a random joke. Use `--category` or `-c` to specify a category.',
    'lsc': 'Lists all available joke categories.',
    'lsj --category -c': 'Lists all jokes. Use `--category` or `-c` to filter by a specific category.',
    'catj --id': 'Retrieves a specific joke by its ID.',
    'addc': 'Adds a new joke category.',
    'addj --category -c': 'Adds a new joke to an existing category. Use `--category` or `-c` to specify the category.',
    'addj --category --id': 'Adds an existing joke to another category. Use `--category` to specify the category and `--id` to specify the joke ID.',
    'good --id': 'Likes a joke. Use `--id` to specify the joke ID.',
    'bad --id': 'Dislikes a joke. Use `--id` to specify the joke ID.'
  };

  const handleCommand = (cmd: string) => {
    setLoading(true);
    const newCommands = [...commands];
    newCommands.push({ type: 'input', content: cmd }); // Add the input command to the history
  
    const args = cmd.split(' '); // Split the command into parts
    const command = args[0]; // The first part is the command
    const params = args.slice(1); // The rest are parameters
  
    let response = '';
  
    switch (command) {
      case 'help':
        response = 'Available commands:\n\n' +
          Object.entries(listCommands)
            .map(([key, value]) => `**${key}**: ${value}`) // TODO: Have to figure out why \t or \n doesn't not work in web, but a bit lazy right now
            .join('\n\n');
        break;
      case 'clear':
        setCommands([]);
        setInput('');
        setLoading(false);
        return; 
      case 'rndj':
        response = `Fetching a random joke${
          params.includes('--category') || params.includes('-c') ? ` from category: ${params[params.indexOf('--category') + 1] || params[params.indexOf('-c') + 1]}` : ''
        }`;
        // TODO: Call the API to fetch a random joke
        break;
      case 'lsc':
        response = 'Listing all categories';
        // TODO: Call the API to fetch all categories
        break;
      case 'lsj':
        response = `Listing all jokes${
          params.includes('--category') || params.includes('-c') ? ` from category: ${params[params.indexOf('--category') + 1] || params[params.indexOf('-c') + 1]}` : ''
        }`;
        // TODO: Call the API to fetch jokes
        break;
      case 'catj':
        response = `Fetching joke with ID: ${params[params.indexOf('--id') + 1]}`;
        // TODO: Call the API to fetch a joke by ID
        break;
      case 'addc':
        response = 'Adding a new category';
        // TODO: Call the API to add a category
        break;
      case 'addj':
        response = `Adding a joke${
          params.includes('--category') || params.includes('-c') ? ` to category: ${params[params.indexOf('--category') + 1] || params[params.indexOf('-c') + 1]}` : ''
        }`;
        // TODO: Call the API to add a joke
        break;
      case 'good':
        response = `Liking joke with ID: ${params[params.indexOf('--id') + 1]}`;
        // TODO: Call the API to like a joke
        break;
      case 'bad':
        response = `Disliking joke with ID: ${params[params.indexOf('--id') + 1]}`;
        // TODO: Call the API to dislike a joke
        break;
      default:
        response = 'Command not found. Type "help" for available commands.';
        break;
    }
  
    newCommands.push({ type: 'output', content: response }); // Add the output to the history
    setCommands(newCommands);
    setLoading(false);
    setInput('');
  };

  return (
    <div className="min-h-screen bg-black p-4">
      <div className="max-w-3xl mx-auto">
        <div className="bg-gray-800 rounded-t-lg p-2 flex items-center gap-2">
          <Terminal className="w-4 h-4 text-green-400" />
          <span className="text-green-400 text-sm font-mono">HaHaSaaS Terminal v1.0.0</span>
        </div>
        
        <div className="bg-black border border-gray-800 rounded-b-lg p-4 font-mono text-sm">
          <div className="h-96 overflow-y-auto space-y-2">
            <div className="text-green-400">
              Welcome to HaHaSaaS - Humor as a Service{'\n'}
              Type 'help' for available commands.
            </div>
            
            {commands.map((entry, idx) => (
              <div 
                key={idx} 
                className={`${
                  entry.type === 'input' ? 'text-green-400' : 'text-gray-300'
                }`}
              >
                {entry.content}
              </div>
            ))}
            
            <div className="flex items-center text-green-400">
              <span>guest@hahasaas:~$&nbsp;</span>
              <input
                type="text"
                value={input}
                onChange={(e) => setInput(e.target.value)}
                onKeyPress={(e) => {
                  if (e.key === 'Enter') {
                    handleCommand(input);
                  }
                }}
                className="flex-1 bg-transparent outline-none border-none text-green-400"
                autoFocus
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default JokeTerminal;