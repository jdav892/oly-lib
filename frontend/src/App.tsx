import  React from 'react';
import './App.css';
import ViewChoice from './components/ViewChoice';
import Header from './components/Header';
import Footer from './components/Footer';

const App: React.FC = () =>  {

  return (
    <div>
      <Header/>
      <ViewChoice/>
      <Footer/>
    </div>
  );
};

export default App;
