import  React, { useState } from 'react';
import ExerciseCards from './ExerciseCards';
import ExerciseAccordion from './ExerciseAccordion';

const App: React.FC = () =>  {
  const [view, setView] = useState<'cards' | 'accordion'>('cards');

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="flex justify-center mt-4">
        <button
          onClick={() => setView('cards')}
          className={`mx px-4 py-2 rounded ${view === 'cards' ? 'bg-blue-600 text-white' : 'bg-white border'
            }`} 
        >
          Card View
        </button>
        <button
         onClick={() => setView('accordion')}
         className={`mx-2 px-4 py-2 rounded ${view === 'accordion' ? 'bg-blue-600 text-white' : 'bg-white border'}`} 
        >
          Accordion View
        </button>
      </div>
      {view === 'cards' ? <ExerciseCards/> : <ExerciseAccordion/>}
    </div>
  );
};

export default App;



