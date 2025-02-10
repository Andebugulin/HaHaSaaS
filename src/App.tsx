import React, { useState } from 'react';
import { Terminal } from 'lucide-react';

const JokeTerminal = () => {
  const [jokes, setJokes] = useState([]);
  const [input, setInput] = useState('');
  const [loading, setLoading] = useState(false);

  const commands = {
    'help': 'Available commands: random, category <name>, list, clear',
    'random': 'Fetching random joke...',
    'clear': 'clear',
    'list': 'Available categories: programming, dad-jokes, science'
  };

  const handleCommand = (cmd) => {
    setLoading(true);
    const newJokes = [...jokes];
    
    newJokes.push({ type: 'input', content: `guest@hahasaas:~$ ${cmd}` });
    
    if (cmd === 'clear') {
      setJokes([]);
    } else {
      const response = commands[cmd] || 'Command not found. Type "help" for available commands.';
      newJokes.push({ type: 'output', content: response });
    }
    
    setJokes(newJokes);
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
            
            {jokes.map((entry, idx) => (
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