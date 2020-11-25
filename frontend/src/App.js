//import logo from './logo.svg';
import './App.css';
import RouteSwitcher from './components/RouteSwitcher'
import {CssBaseline, Container} from '@material-ui/core';
import { BrowserRouter as Router } from  "react-router-dom";

function App() {
  return (
    <Router>
      <Container maxWidth="lg">
        <CssBaseline />
        <RouteSwitcher />
      </Container>
    </Router>
  );
}

export default App;
