import { useState } from 'react';
// import { BrowserRouter, Switch, Route} from "react-router-dom";
import { BrowserRouter,Route,Switch} from "react-dom";
import CreateRoom from './components/CreateRoom';
import Room from './components/room';


function App() {
  return (
    <div className="App">
      <BrowserRouter>
      <Switch>
      {/* <Routes> */}
        <Route path="/" component={CreateRoom}></Route>
        <Route path="/room/:roomID" component={Room}></Route>
      </Switch>
      {/* </Routes> */}
      </BrowserRouter>
      
    </div>
  )
}

export default App;
