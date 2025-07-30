import React, {useEffect, useState} from 'react';

type ExerciseData = {
    [category: string]: string[];
};

const ExerciseAccordion: React.FC = () => {
    const [exercisesBySport, setExercisesBySport] = useState<ExerciseData>({});
    const [expanded, setExpanded] = useState<string | null>(null);

    useEffect(() => {
        fetch("http://localhost:8080/lib/exercises", {
            method: "GET",
            credentials: "include",
            headers: {
                "Content-Type" : "application/json",
            },
        })
          .then((res) => res.json())
          .then((data) => setExercisesBySport(data.exercises))
          .catch((err) => console.error(err));
    }, []);
    
    const toggle = (category: string) => {
        setExpanded(prev => (prev === category ? null : category));
    };

    return (
        <div className="w-full max-w-2xl mx-auto p-6">
            <h1 className="text-3xl font-bold mb-6 text-center">Exercise Categories</h1>
            {Object.entries(exercisesBySport).map(([category, exercises]) => (
              <div key={category} className="mb-4 border rounded-lg">
                <button 
                  className="w-full text-left p-4 bg-blue-100 hover:bg-blue-200 font-semibold capitalize"
                  onClick={() => toggle(category)}
                  >
                    {category}
                  </button>
                  {expanded === category && (
                    <div className="bg-white p-4 border-t">
                        <ul>
                            {exercises.map((exercise, index) => (
                                <li key={index}>{exercise}</li>
                            ))}
                        </ul>
                    </div>
                  )}
              </div>
            ))}
        </div>
    );
};

export default ExerciseAccordion;


