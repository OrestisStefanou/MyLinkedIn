import React from 'react';
import ReactDom from 'react-dom';
import {BrowserRouter as Router,Route,Switch} from 'react-router-dom';
import SignUp from "./SignUp"
import SignIn from "./SignIn"
import WeclomePage from "./WelcomePage"

const App = () => {
    return(
        <Router>
            <Switch>
                <Route exact path='/'>
                    <WeclomePage/>
                </Route>
                <Route exact path='/signup'>
                    <SignUp/>
                </Route>
                <Route exact path='/signin'>
                    <SignIn/>
                </Route>
            </Switch>
        </Router>
    )
}

ReactDom.render(<App/>,document.getElementById('root'));
