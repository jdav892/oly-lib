import React, {useEffect, useState} from 'react';

type ExerciseData = {
    [category: string] :string[];
};

const ExerciseCards: React.FC = () => {
    const [excercisesBySport, setExercisesBySport] = useState<ExerciseData>({});
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetch("http://localhost:8080/lib/exercises", {
            method: "GET",
            credentials: "include",
            headers: {
                "Content-Type" : "application/json"
            },
        })
          .then((res) => res.json())
          .then((data) => {
            setExercisesBySport(data.exercises);
            setLoading(false);
          });
    }, []);

    if (loading) return <p className="text-center mt-8">Loading...</p>;

    return (
        <div className="p-6">
            <h1 className="text-3xl font-bold mb-6 text-center">Exercise Categories</h1>
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                {Object.entries(excercisesBySport).map(([category, exercises]) => (
                    <div key={category} className="bg-white shadow-md rounded-xl p-4 border">
                        <h2 className="text-xl font-semibold mb-4 capitalize text-blue-700">
                            {category}
                        </h2>
                        <ul className="list-disc list-inside text-gray-700 space-y-1">
                            {exercises.map((exercise, index) => (
                                <li key={index}>{exercise}</li>
                            ))}  
                        </ul>
                    </div>
                ))}
            </div>
        </div>
    )

};

export default ExerciseCards;
