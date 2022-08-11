import React from 'react';
import "bootstrap/dist/css/bootstrap.min.css"
import './App.css';
import Navbar from './components/Navbar';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Home from './pages/Home';
import Reports from './pages/Reports';
import Products from './pages/Products';
import Register from './pages/Register';
import Bitly from './pages/Bitly';
import Login from './pages/Login';

function App() {
  return (
    <>
      <Router>
        <Navbar />
        <Switch>
          <Route path='/' exact component={Home} />
          <Route path='/reports' component={Reports} />
          <Route path='/products' component={Products} />
          <Route path='/login' component={Login} />
          <Route path='/register' component={Register} />
          <Route path='/bitly/:id' component={Bitly}/>
        </Switch>
      </Router>
    </>
  );
}

export default App;
