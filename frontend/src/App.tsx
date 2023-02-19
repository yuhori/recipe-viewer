import './App.css';
import { Routes, Route } from 'react-router-dom';
import List from './pages/List';

const homeUrl = process.env.PUBLIC_URL;


function App() {
  return (
    <div className="App">
      <Routes>
        <Route path={ homeUrl + "/" } element={<List />} />
      </Routes>
    </div>
  );
}

export default App;
