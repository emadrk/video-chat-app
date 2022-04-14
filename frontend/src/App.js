// import logo from './logo.svg';
// import './App.css';

// function App() {
//   return (
//     <div className="App">
//       <header className="App-header">
//         <img src={logo} className="App-logo" alt="logo" />
//         <p>
//           Edit <code>src/App.js</code> and save to reload.
//         </p>
//         <a
//           className="App-link"
//           href="https://reactjs.org"
//           target="_blank"
//           rel="noopener noreferrer"
//         >
//           Learn React
//         </a>
//       </header>
//     </div>
//   );
// }

// export default App;

// import { useState } from 'react';
import { BrowserRouter, Switch, Route} from "react-router-dom";
// import { BrowserRouter,Route,Switch} from "react-dom";
import CreateRoom from './components/CreateRoom';
import Room from './components/room';


function App() {
  return (
    <div className="App">
      <BrowserRouter>
      <Switch>
      {/* <Routes> */}
        <Route path="/" exact component={CreateRoom}></Route>
        <Route path="/room/:roomID" component={Room}></Route>
      </Switch>
      {/* </Routes> */}
      </BrowserRouter>
      
    </div>
  )
}

export default App;
