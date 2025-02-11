import React, { useState } from 'react';
import { Terminal } from 'lucide-react';

interface Command {
  type: 'input' | 'output';
  content: string;
}

interface Joke {
  id: number;
  content: string;
  likes: number;
  dislikes: number;
}

const JokeTerminal = () => {
  const [commands, setCommands] = useState<Command[]>([]);
  const [input, setInput] = useState('');
  const [loading, setLoading] = useState(false);

  // Update the listCommands object in your component:
  const listCommands = {
    'help': 'Displays a list of available commands and their usage.',
    'clear': 'Clears the terminal screen.',
    'rndj': 'Fetches a random joke from all jokes.',
    'rndj --category|-c CATEGORY': 'Fetches a random joke from a specific category.',
    'lsc': 'Lists all available joke categories.',
    'lsj': 'Lists all jokes.',
    'lsj --category|-c CATEGORY': 'Lists all jokes in a specific category.',
    'catj --id ID': 'Retrieves a specific joke by its ID.',
    'addc CATEGORY': 'Adds a new joke category.',
    'addj --category|-c CATEGORY JOKE': 'Adds a new joke to a category.',
    'addj --category|-c CATEGORY --id ID': 'Adds an existing joke to another category.',
    'good --id ID': 'Likes a joke by ID.',
    'bad --id ID': 'Dislikes a joke by ID.'
  };

  const getCategory = (params: string[]) => {
    if (params.includes('--category')) {
      return params[params.indexOf('--category') + 1];
    }
    if (params.includes('-c')) {
      return params[params.indexOf('-c') + 1];
    }
    return null;
  };

  // Fetch a random joke
  const fetchRandomJoke = async () => {
    const response = await fetch('/api/joke/random');
    const data = await response.json();
    return data.joke;
  };

  // Fetch a random joke by category
  const fetchRandomJokeByCategory = async (category: string) => {
    const response = await fetch(`/api/joke/random/${category}`);
    const data = await response.json();
    return data.joke;
  };

  // Fetch all categories
  const fetchCategories = async () => {
    const response = await fetch('/api/joke/categories');
    const data = await response.json();
    return data.categories;
  };

  // Fetch all jokes
  const fetchAllJokes = async () => {
    const response = await fetch('/api/joke/all');
    const data = await response.json();
    return data.jokes;
  };

  // Fetch jokes by category
  const fetchJokesByCategory = async (category: string) => {
    const response = await fetch(`/api/joke/category/${category}`);
    const data = await response.json();
    return data.jokes;
  };

  // Fetch a joke by ID
  const fetchJokeByID = async (id: string) => {
    const response = await fetch(`/api/joke/${id}`);
    const data = await response.json();
    return data.joke;
  };

  // Add a new category
  const addCategory = async (name: string) => {
    const response = await fetch('/api/joke/category', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ name }),
    });
    if (!response.ok) {
      throw new Error('Failed to add category');
    }
    return 'Category added successfully';
  };

  // Add a new joke to a category
  const addJoke = async (jokeContent: string, category: string) => {
    const response = await fetch('/api/joke', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ 
        Category: category, 
        Joke: jokeContent  
      }),
    });
    if (!response.ok) {
      throw new Error('Failed to add joke. Status: ' + response.status);
    }
    return 'Joke added successfully';
  };
  

  // Add an existing joke to a category
  const addJokeToCategory = async (jokeId: string, category: string) => {
    const response = await fetch(`/api/joke/category/${category}/${jokeId}`, {
      method: 'POST',
    });
    if (!response.ok) {
      throw new Error('Failed to add joke to category. Status: ' + response.status);
    }
    return 'Joke added to category successfully';
  };

  // Like a joke
  const likeJoke = async (id: string) => {
    const response = await fetch(`/api/joke/${id}/like`, {
      method: 'POST',
    });
    if (!response.ok) {
      throw new Error('Failed to like joke');
    }
    return 'Joke liked successfully';
  };

  // Dislike a joke
  const dislikeJoke = async (id: string) => {
    const response = await fetch(`/api/joke/${id}/dislike`, {
      method: 'POST',
    });
    if (!response.ok) {
      throw new Error('Failed to dislike joke');
    }
    return 'Joke disliked successfully';
  };

  const formatJokeOutput = (joke: Joke) => {
    return `ID: ${joke.id} | Content: ${joke.content} | Likes: ${joke.likes} | Dislikes: ${joke.dislikes}`;
  };

  const handleCommand = async (cmd: string) => {
    setLoading(true);
    const newCommands = [...commands];
    newCommands.push({ type: 'input', content: cmd });

    const args = cmd.split(' ');
    const command = args[0];
    const params = args.slice(1);

    let response = '';

    try {
      switch (command) {
        case 'help':
          response = 'Available commands:\n\n' +
            Object.entries(listCommands)
              .map(([key, value]) => `${key}: ${value}`)
              .join('\n\n');
          break;

        case 'clear':
          setCommands([]);
          setInput('');
          setLoading(false);
          return;

          case 'rndj':
            const category = getCategory(params);
            const randomJoke = category 
              ? await fetchRandomJokeByCategory(category)
              : await fetchRandomJoke();
            response = formatJokeOutput(randomJoke);
            break;

        case 'lsc':
          const categories = await fetchCategories();
          response = 'Available Categories:\n' + categories.join('\n');
          break;

        case 'lsj':
          const listCategory = getCategory(params);
          const jokes = listCategory ? await fetchJokesByCategory(listCategory) : await fetchAllJokes();
          response = 'Jokes:\n' + jokes.join('\n');
          break;

        case 'catj':
          const id = params[params.indexOf('--id') + 1];
          if (!id) {
            response = 'Please provide a joke ID using --id';
          } else {
            const joke = await fetchJokeByID(id);
            response = formatJokeOutput(joke);
          }
          break;

        case 'addc':
          const name = params.join(' ');
          if (!name) {
            response = 'Please provide a category name.';
          } else {
            response = await addCategory(name);
          }
          break;

          // Update the handleCommand's addj case:
      case 'addj':
        const jokeCategory = getCategory(params);
        const idIndex = params.indexOf('--id');
        const hasId = idIndex !== -1;
        
        if (!jokeCategory) {
          response = 'Please provide a category using --category or -c';
          break;
        }

        if (hasId) {
          const jokeId = params[idIndex + 1];
          try {
            response = await addJokeToCategory(jokeId, jokeCategory);
          } catch (error: any) {
            response = error.message;
          }
        } else {
          // Get all parameters except the category flag and its value
          const jokeContent = params
            .filter((param, index) => {
              const categoryFlagIndex = params.indexOf('--category');
              const shortCategoryFlagIndex = params.indexOf('-c');
              return !['--category', '-c'].includes(param) && 
                index !== categoryFlagIndex + 1 && 
                index !== shortCategoryFlagIndex + 1;
            })
            .join(' ');
          
          if (!jokeContent) {
            response = 'Please provide joke content.';
          } else {
            try {
              response = await addJoke(jokeContent, jokeCategory);
            } catch (error: any) {
              response = error.message;
            }
          }
        }
        break;
        case 'good':
          const likeId = params[params.indexOf('--id') + 1];
          if (!likeId) {
            response = 'Please provide a joke ID using --id';
          } else {
            response = await likeJoke(likeId);
          }
          break;

        case 'bad':
          const dislikeId = params[params.indexOf('--id') + 1];
          if (!dislikeId) {
            response = 'Please provide a joke ID using --id';
          } else {
            response = await dislikeJoke(dislikeId);
          }
          break;

        default:
          response = 'Command not found. Type "help" for available commands.';
          break;
      }
    } catch (error: any) {
      response = `Error: ${error.message}`;
    }

    newCommands.push({ type: 'output', content: response });
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
                className={`whitespace-pre-wrap ${
                  entry.type === 'input' ? 'text-green-400' : 'text-gray-300'
                }`}
              >
                {entry.type === 'input' ? '$ ' : ''}{entry.content}
              </div>
            ))}
            
            <div className="flex items-center text-green-400">
              <span>$ </span>
              <input
                type="text"
                value={input}
                onChange={(e) => setInput(e.target.value)}
                onKeyPress={(e) => {
                  if (e.key === 'Enter' && !loading) {
                    handleCommand(input);
                  }
                }}
                className="flex-1 bg-transparent outline-none border-none text-green-400"
                disabled={loading}
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