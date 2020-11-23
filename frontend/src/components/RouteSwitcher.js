import {
Switch,
Route
} from  "react-router-dom";

import Dashboard from './Dashboard';
import Home from './Home';
import Login from './Login';
import Signup from './Signup';
import NotFound from './NotFound';

function RouteSwitcher() {
return (
    <Switch>
        <Route exact path="/">
            <Home />
        </Route>
        <Route path="/signup">
            <Signup />
        </Route>
        <Route path="/login">
            <Login />
        </Route>
        <Route path="/dashboard">
            <Dashboard />
        </Route>
        <Route path="*">
            <NotFound />
        </Route>
    </Switch>
);
}
export default RouteSwitcher;