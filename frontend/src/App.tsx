import './App.css';
import { Routes, Route } from 'react-router-dom';
import List from './pages/List';


function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/recipe-viewer/list" element={<List />} />
      </Routes>
    </div>
  );
}

export default App;
