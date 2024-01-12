import './App.css';
import { MyRoutes } from './MyRoutes.js';

function App() {
  return (
    <>
    <div className='head-div'>
      <button className='login'>Login</button>
      <button className='login'>Sign up</button>
    </div>
    <div className="App">
      <MyRoutes/>
    </div>
    </>
  );
}

export default App;
